package service

import (
	"context"
	"database/sql"
	_ "github.com/lib/pq"

)

func (slugService) UserSegments(ctx context.Context, userId string) (string, error) {

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return "DB opening error", err
	}
	defer db.Close()

	var exists bool = false
	err = db.QueryRow("SELECT EXISTS (SELECT 1 FROM users WHERE user_id = $1)", userId).Scan(&exists)
	if !exists {
		return "The user with id " + userId + " does not exist.", err
	}

	rows, err := db.Query("SELECT segment_name FROM users LEFT JOIN segments ON users.segment = segments.segment_id WHERE users.user_id = $1", userId)
	if err != nil {
		return "Error", err
	}
	defer rows.Close()
	segments := ""

	for rows.Next() {
		str := []byte{}
		err := rows.Scan(&str)
		if err != nil {
			return "Error", err
		}
		segments += string(str) + "; "
	}
	return "The user " + userId + " has segments: " + segments, nil
}