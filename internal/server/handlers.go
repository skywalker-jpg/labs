package server

import (
	"Labs2/internal/models"
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/labstack/echo"
	"io/ioutil"
	"log/slog"
	"net/http"
)

func (s *Server) RegisterHandlers() {
	app := s.app

	apiGroup := app.Group("/api/v1")
	apiGroup.POST("/joke", s.ReserveJokeHandler)
	apiGroup.POST("/jokes", s.GetAllJokesHandler)
	apiGroup.POST("/cat", s.ReserveCatHandler)
	apiGroup.POST("/cats", s.GetAllCatsHandler)
	apiGroup.POST("/fact", s.ReserveNumberFactHandler)
	apiGroup.POST("/facts", s.GetAllNumbersFactsHandler)
	apiGroup.POST("/news", s.GetHackerNewsHandler)
	apiGroup.POST("/quotes", s.GetQuotesHandler)
	//
	//app.GET("/*", s.NotFound)
}

func (s *Server) ReserveJokeHandler(c echo.Context) error {
	requestID := c.Get("requestID").(string)

	url := "https://v2.jokeapi.dev/joke/Any?type=single"
	resp, err := http.Get(url)
	if err != nil {
		s.logger.Error("ReserveJokeHandler", slog.String("requestID", requestID),
			slog.String("error", fmt.Sprintf("Unable to fetch joke from API: %v", err)))
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch joke from API"})
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		s.logger.Error("ReserveJokeHandler", slog.String("requestID", requestID),
			slog.String("error", fmt.Sprintf("API returned non-200 status: %v", resp.StatusCode)))
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch joke from API"})
	}

	var jokeResponse models.JokeResponse
	if err := json.NewDecoder(resp.Body).Decode(&jokeResponse); err != nil {
		s.logger.Error("ReserveJokeHandler", slog.String("requestID", requestID),
			slog.String("error", fmt.Sprintf("Unable to parse joke API response: %v", err)))
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to parse joke response"})
	}

	if jokeResponse.Error {
		s.logger.Info("ReserveJokeHandler", slog.String("requestID", requestID),
			slog.String("error", "Joke API returned an error"))
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Joke API returned an error"})
	}

	jokeDTO := models.JokeDTO{
		ID:       jokeResponse.ID,
		Category: jokeResponse.Category,
		Joke:     jokeResponse.Joke,
	}

	ctx := c.Request().Context()

	err = s.Storage.CreateJoke(ctx, jokeDTO)
	if err != nil {
		s.logger.Error("ReserveJokeHandler", slog.String("requestID", requestID),
			slog.String("error", fmt.Sprintf("Unable to save joke to the database: %v", err)))
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to save joke to the database"})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Joke successfully reserved and saved to the database!",
		"joke":    jokeResponse.Joke,
		"id":      jokeResponse.ID,
	})
}

func (s *Server) GetAllJokesHandler(c echo.Context) error {
	requestID := c.Get("requestID").(string)

	jokes, err := s.Storage.GetJokes(c.Request().Context())
	if err != nil {
		s.logger.Error("Server", slog.String("requestID", requestID),
			slog.String("error", "Unable to retrieve jokes: "+err.Error()))
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Unable to retrieve jokes"})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"jokes": jokes})
}

type CatResponse struct {
	ID  string `json:"id"`
	URL string `json:"url"`
}

func (s *Server) ReserveCatHandler(c echo.Context) error {
	apiURL := fmt.Sprintf("https://api.thecatapi.com/v1/images/search?api_key=%s", "live_LqRQo2v2wg44W3YxwWTnPTAqLSc32hYbzekVTYMGbWqINxkw5L4ocBVMhJyeYUzj")

	resp, err := http.Get(apiURL)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Unable to fetch cat from API"})
	}
	defer resp.Body.Close()

	var catData []CatResponse
	if err := json.NewDecoder(resp.Body).Decode(&catData); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Unable to decode cat data"})
	}

	if len(catData) == 0 {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "No cat data available"})
	}

	cat := catData[0]

	catDTO := models.CatDTO{
		CatID: cat.ID,
		URL:   cat.URL,
	}

	ctx := c.Request().Context()

	if err := s.Storage.CreateCat(ctx, catDTO); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Unable to save cat to the database"})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"cat_id": cat.ID,
		"url":    cat.URL,
	})
}

func (s *Server) GetAllCatsHandler(c echo.Context) error {
	requestID := c.Get("requestID").(string)

	cats, err := s.Storage.GetCats(c.Request().Context())
	if err != nil {
		s.logger.Error("Server", slog.String("requestID", requestID),
			slog.String("error", "Unable to retrieve cats: "+err.Error()))
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Unable to retrieve cats"})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"cats": cats})
}

func (s *Server) GetAllNumbersFactsHandler(c echo.Context) error {
	ctx := c.Request().Context()
	facts, err := s.Storage.GetNumbers(ctx)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Unable to fetch number facts from database"})
	}
	return c.JSON(http.StatusOK, facts)
}

func (s *Server) ReserveNumberFactHandler(c echo.Context) error {
	number := c.QueryParam("number")

	apiURL := fmt.Sprintf("http://numbersapi.com/%s?json", number)
	s.logger.Info(number)
	resp, err := http.Get(apiURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("API request failed with status %d", resp.StatusCode)
	}

	var apiResponse models.NumberDTO
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &apiResponse)
	if err != nil {
		return err
	}

	err = s.Storage.CreateNumber(c.Request().Context(), apiResponse)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Unable to save number fact to database"})
	}

	return c.JSON(http.StatusOK, apiResponse.Fact)
}

func (s *Server) GetHackerNewsHandler(c echo.Context) error {
	url := "https://news.ycombinator.com/"

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("failed to fetch Hacker News: %s", resp.Status)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return err
	}

	var titles []string

	doc.Find(".title a").Each(func(i int, s *goquery.Selection) {
		title := s.Text()
		titles = append(titles, title)
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch news from Hacker News"})
	}

	return c.JSON(http.StatusOK, titles)
}

func (s *Server) GetQuotesHandler(c echo.Context) error {
	url := "http://quotes.toscrape.com/"

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("failed to fetch quotes: %s", resp.Status)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return err
	}

	var quotes []models.Quote

	doc.Find(".quote").Each(func(i int, s *goquery.Selection) {
		text := s.Find(".text").Text()
		author := s.Find(".author").Text()
		quotes = append(quotes, models.Quote{Text: text, Author: author})
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch quotes from the website"})
	}

	return c.JSON(http.StatusOK, quotes)
}
