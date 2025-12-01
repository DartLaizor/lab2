package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"

	"server/internal/storage"
)

type Storage struct {
	db *pgxpool.Pool
}

// NewStorage инициализирует базу данных и проверяет соединение
func NewStorage(storagePath string) (*Storage, error) {
	const op = "storage.postgres.NewStorage"

	db, err := pgxpool.New(context.Background(), storagePath)
	if err != nil {
		return nil, fmt.Errorf("%s : %w", op, err)
	}

	if err := db.Ping(context.Background()); err != nil {
		db.Close()
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	_, err = db.Exec(context.Background(),
		`		CREATE TABLE IF NOT EXISTS reviews (
			id SERIAL PRIMARY KEY,
			name TEXT NOT NULL,
			date TEXT NOT NULL,
			phone TEXT,
			email TEXT,
			technologies TEXT[],
			rating INT CHECK (rating BETWEEN 0 AND 9),
			comment TEXT
		);
	`)

	if err != nil {
		db.Close()
		return nil, fmt.Errorf("%s: failed to create table: %w", op, err)
	}
	return &Storage{db: db}, nil
}

// SaveRewiev сохраняет отзыв в БД
func (s *Storage) SaveRewiev(r storage.Review) error {
	const op = "storage.postgres.SaveReview"
	_, err := s.db.Exec(context.Background(),
		`INSERT INTO reviews (name, date, phone, email, technologies, rating, comment)
		 VALUES ($1, $2, $3, $4, $5, $6, $7)`,
		r.Name,
		r.Date,
		r.Phone,
		r.Email,
		r.Technologies,
		r.Rating,
		r.Comment,
	)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

// GetReviews возвращает все отзывы из базы данных
func (s *Storage) GetReviews() ([]storage.Review, error) {
	const op = "storage.postgres.GetReviews"

	rows, err := s.db.Query(context.Background(), `
		SELECT name, date, phone, email, technologies, rating, comment
		FROM reviews
		ORDER BY id DESC
	`)

	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	defer rows.Close()

	var reviews []storage.Review

	for rows.Next() {
		var r storage.Review
		err := rows.Scan(
			&r.Name,
			&r.Date,
			&r.Phone,
			&r.Email,
			&r.Technologies,
			&r.Rating,
			&r.Comment,
		)
		if err != nil {
			return nil, fmt.Errorf("%s: failed to scan row: %w", op, err)
		}
		reviews = append(reviews, r)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("%s: rows iteration error: %w", op, err)
	}

	return reviews, nil
}

func (s *Storage) CloseStorage() {
	if s.db != nil {
		s.db.Close()
	}

}
