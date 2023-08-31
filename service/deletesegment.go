package service

import (
	"context"
	"database/sql"
	_ "github.com/lib/pq"
)

func (slugService) DeleteSegment(ctx context.Context, slug string) (string, error) {
	if slug == "" {
		return "The segment name is set incorrectly.", nil
	}

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return "DB opening error", err
	}
	defer db.Close()

	var exists bool = false
	err = db.QueryRow("SELECT EXISTS (SELECT 1 FROM segments WHERE segment_name = $1)", slug).Scan(&exists)
	if !exists {
		return "The segment with the name: " + slug + " does not exist.", err
	}
	_, err = db.Exec("DELETE FROM segments WHERE segment_name = $1", slug)
	if err != nil {
		return "Error deleting a segment with the name:" + slug, err
	}

	return "The segment with the name: " + slug+ " has been deleted.", nil
}