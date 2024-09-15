package storage

import (
	"Labs2/internal/models"
	"context"
	"database/sql"
	"fmt"
)

type JokeRepo struct {
	db *sql.DB
}

func NewJokeRepo(db *sql.DB) *JokeRepo {
	return &JokeRepo{
		db: db,
	}
}

func (s *JokeRepo) CreateJoke(ctx context.Context, joke models.JokeDTO) error {
	query := `
        INSERT INTO jokes (joke_id, category, joke)
        VALUES ($1, $2, $3)
    `
	_, err := s.db.ExecContext(ctx, query, joke.ID, joke.Category, joke.Joke)
	if err != nil {
		return fmt.Errorf("failed to execute query: %w", err)
	}
	return nil
}

func (s *JokeRepo) GetJokes(ctx context.Context) ([]models.Joke, error) {
	var jokes []models.Joke
	rows, err := s.db.QueryContext(ctx, "SELECT id, joke_id, category, joke, created_at FROM jokes")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var joke models.Joke
		if err := rows.Scan(&joke.ID, &joke.JokeID, &joke.Category, &joke.Joke, &joke.CreatedAt); err != nil {
			return nil, err
		}
		jokes = append(jokes, joke)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return jokes, nil
}
