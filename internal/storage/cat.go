package storage

import (
	"Labs2/internal/models"
	"context"
	"database/sql"
)

type CatRepo struct {
	db *sql.DB
}

func NewCatRepo(db *sql.DB) *CatRepo {
	return &CatRepo{
		db: db,
	}
}

func (s *CatRepo) CreateCat(ctx context.Context, cat models.CatDTO) error {
	query := `INSERT INTO cats (cat_id, url, created_at) VALUES ($1, $2, NOW())`
	_, err := s.db.ExecContext(ctx, query, cat.CatID, cat.URL)
	return err
}

func (s *CatRepo) GetCats(ctx context.Context) ([]models.Cat, error) {
	query := `SELECT id, cat_id, url, created_at FROM cats`
	rows, err := s.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cats []models.Cat
	for rows.Next() {
		var cat models.Cat
		if err := rows.Scan(&cat.ID, &cat.CatID, &cat.URL, &cat.CreatedAt); err != nil {
			return nil, err
		}
		cats = append(cats, cat)
	}

	return cats, nil
}
