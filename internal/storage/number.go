package storage

import (
	"Labs2/internal/models"
	"context"
	"database/sql"
)

type NumberRepo struct {
	db *sql.DB
}

func NewNumberRepo(db *sql.DB) *NumberRepo {
	return &NumberRepo{
		db: db,
	}
}

func (s *NumberRepo) GetNumbers(ctx context.Context) ([]models.Number, error) {
	rows, err := s.db.Query("SELECT id, num, fact, created_at FROM number_facts")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var facts []models.Number
	for rows.Next() {
		var f models.Number
		if err := rows.Scan(&f.ID, &f.Number, &f.Fact, &f.CreatedAt); err != nil {
			return nil, err
		}
		facts = append(facts, f)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return facts, nil
}

func (s *NumberRepo) CreateNumber(ctx context.Context, number models.NumberDTO) error {
	query := `
		INSERT INTO number_facts (num, fact) VALUES ($1, $2)
    `
	_, err := s.db.ExecContext(ctx, query,
		number.Number, number.Fact)
	return err
}
