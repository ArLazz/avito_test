package service

import (
	"context"
	"database/sql"
	_ "github.com/lib/pq"
)

func (slugService) CreateSegment(ctx context.Context, slug string) (string, error) {
	if slug == "" {
		return "The segment name is set incorrectly.", nil
	}

	connStr := "user=postgres password=4581 dbname=segments_db sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return "DB opening error", err
	}
	defer db.Close()

	var exists bool = true
	err = db.QueryRow("SELECT EXISTS (SELECT 1 FROM segments WHERE segment_name = $1)", slug).Scan(&exists)
	if exists {
		return "There is already a segment with the name: " + slug + ".", err
	}

	_, err = db.Exec("INSERT INTO segments (segment_name) values ($1)",
		slug)
	if err != nil {
		return "Error adding a segment.", err
	}

	return "A segment with the name: " + slug + " has been created.", nil
}