package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/fed-605/url-shortener-go/internal/storage"
	"github.com/jackc/pgx/v5/pgconn"
)

// save url and alias to the db
func (s *PostgresStore) SaveUrl(url string, alias string) error {
	const op = "storage.postgres.SaveUrl" // helper

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)

	defer cancel()

	query := `INSERT INTO urls(url,alias)
	VALUES($1,$2);`

	if _, err := s.db.Exec(ctx, query, url, alias); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return fmt.Errorf("%s: %w", op, storage.ErrURLExists)
		}
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil

}

// get url from db by alias
func (s *PostgresStore) GetUrl(alias string) (string, error) {
	const op = "storage.postgres.GetUrl"

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)

	defer cancel()

	var url string

	query := `SELECT url FROM urls WHERE alias = $1;`

	if err := s.db.QueryRow(ctx, query, alias).Scan(&url); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", fmt.Errorf("%s: %w", op, storage.ErrURLNotFound)
		}
		return "", fmt.Errorf("%s: %w", op, err)
	}
	return url, nil
}

// delete url and alias from the db
func (s *PostgresStore) DeleteUrl(alias string) error {
	const op = "storage.postgres.DeleteUrl"

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)

	defer cancel()

	query := `DELETE FROM urls WHERE alias = $1;`

	res, err := s.db.Exec(ctx, query, alias)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	rows := res.RowsAffected()

	switch {
	case rows == 0:
		return fmt.Errorf("%s: %w", op, storage.ErrURLNotFound)
	case rows > 1:
		return fmt.Errorf("%s: %w", op, storage.ErrUnexpectedRows)
	}
	return nil
}

// get all records
func (s *PostgresStore) GetAllRecords() ([]storage.Url, error) {
	const op = "storage.postgres.GetAllRecords"

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)

	defer cancel()

	query := `SELECT id,url,alias FROM urls`

	rows, err := s.db.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	defer rows.Close()

	var urls []storage.Url

	for rows.Next() {
		var url storage.Url

		if err := rows.Scan(&url.Id, &url.Url, &url.Alias); err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}

		urls = append(urls, url)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return urls, nil

}
