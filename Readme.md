# API Backend Documentation

## Overview

This project is a backend server designed to handle multiple API requests related to jokes, cats, number facts, news, and quotes. The server is built using the Echo framework in Go and integrates with several external APIs to fetch and store data in a database.

## Endpoints

### 1. Jokes API

- **POST /api/v1/joke**: Fetch a random joke from JokeAPI and store it in the database.
- **POST /api/v1/jokes**: Retrieve all stored jokes from the database.

#### Example Request:
```http
POST http://localhost:8080/api/v1/joke
```

#### Example response:
```json
{
    "message": "Joke successfully reserved and saved to the database!",
    "joke": "Why did the chicken cross the road?",
    "id": 1234
}
```

### 2. **Cats API Documentation (cats_api.md)**


# Cats API

## Endpoints

- **POST /api/v1/cat**: Fetch a random cat image from TheCatAPI and store it in the database.
- **POST /api/v1/cats**: Retrieve all stored cat images from the database.

### Example request:
### Example response:
```json
{
    "cat_id": "abc123",
    "url": "https://cdn2.thecatapi.com/images/abc123.jpg"
}
```

### 3. **Number Facts API Documentation (number_facts_api.md)**

# Number Facts API

## Endpoints

- **POST /api/v1/fact?number=42**: Fetch a fact about a specific number from NumbersAPI and store it in the database.
- **POST /api/v1/facts**: Retrieve all stored number facts from the database.

### Example request:
POST http://localhost:8080/api/v1/fact?number=42
### Example response:
```json
{
    "fact": "42 is the answer to life, the universe, and everything."
}
```

### 4. **News API Documentation (news_api.md)**

# News API

## Endpoints

- **POST /api/v1/news**: Fetch the latest news titles from Hacker News.

### Example request:
POST http://localhost:8080/api/v1/news
### Example response:
```json
[
    "Latest Technology News",
    "Go 1.18 is released",
    "New advancements in AI"
]
```

### 5. **Quotes API Documentation (quotes_api.md)**

# Quotes API

## Endpoints

- **POST /api/v1/quotes**: Fetch random quotes from the website "Quotes to Scrape".

### Example request:
POST http://localhost:8080/api/v1/quotes
### Example response:
```json
[
    {
        "text": "The only limit to our realization of tomorrow is our doubts of today.",
        "author": "Franklin D. Roosevelt"
    },
    {
        "text": "Do not dwell in the past, do not dream of the future, concentrate the mind on the present moment.",
        "author": "Buddha"
    }
]