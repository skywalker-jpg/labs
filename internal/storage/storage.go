package storage

import (
	"Labs2/internal/models"
	"context"
	"database/sql"
)

type JokeStorage interface {
	CreateJoke(ctx context.Context, wh models.JokeDTO) error
	GetJokes(ctx context.Context) ([]models.Joke, error)
}

type CatStorage interface {
	CreateCat(ctx context.Context, wh models.CatDTO) error
	GetCats(ctx context.Context) ([]models.Cat, error)
}

type NumberStorage interface {
	CreateNumber(ctx context.Context, wh models.NumberDTO) error
	GetNumbers(ctx context.Context) ([]models.Number, error)
}

type Storage struct {
	JokeStorage
	CatStorage
	NumberStorage
}

func NewStorage(db *sql.DB) *Storage {
	return &Storage{
		JokeStorage:   NewJokeRepo(db),
		CatStorage:    NewCatRepo(db),
		NumberStorage: NewNumberRepo(db),
	}
}
