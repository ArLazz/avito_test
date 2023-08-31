package service

import (
	"context"
	"database/sql"
	_ "github.com/lib/pq"
	"strings"
	"time"
)

func (slugService) AddUser(ctx context.Context, addSlugs string, deleteSlugs string, userId string) (string, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return "DB opening error", err
	}
	defer db.Close()

	var exists bool = false
	if addSlugs != "" {
		for _, slug := range strings.Split(addSlugs, ",") {
			exists = false
			err = db.QueryRow("SELECT EXISTS (SELECT 1 FROM segments WHERE segment_name = $1)", slug).Scan(&exists)
			if !exists {
				return "The segment with the name: " + slug + " does not exist.", err
			}

			slugId := 0
			_ = db.QueryRow("SELECT segment_id FROM segments WHERE segment_name = $1", slug).Scan(&slugId)

			exists = false
			err = db.QueryRow("SELECT EXISTS (SELECT 1 FROM users WHERE user_id = $1 AND segment = $2)", userId, slugId).Scan(&exists)
			if exists {
				return "The user " + userId + " is already in the segment " + slug, err
			}
		}
	}

	if deleteSlugs != "" {
		for _, slug := range strings.Split(deleteSlugs, ",") {
			exists = false
			err = db.QueryRow("SELECT EXISTS (SELECT 1 FROM segments WHERE segment_name = $1)", slug).Scan(&exists)
			if !exists {
				return "The segment with the name: " + slug + " does not exist.", err
			}

			slugId := 0
			_ = db.QueryRow("SELECT segment_id FROM segments WHERE segment_name = $1", slug).Scan(&slugId)
			exists = false
			err = db.QueryRow("SELECT EXISTS (SELECT 1 FROM users WHERE user_id = $1 AND segment = $2)", userId, slugId).Scan(&exists)
			if !exists {
				return "The user " + userId + " is not a member of the segment " + slug + ".", err
			}
		}
	}

	res := "To the user " + userId
	if addSlugs != "" {
		for _, slug := range strings.Split(addSlugs, ",") {
			slugId := 0
			_ = db.QueryRow("SELECT segment_id FROM segments WHERE segment_name = $1", slug).Scan(&slugId)
			_, err = db.Exec("INSERT INTO users (user_id, segment) values ($1, $2)", userId, slugId)
			if err != nil {
				return "Error adding a segment.", err
			}

			_, err = db.Exec("INSERT INTO usershistory (user_id, segment_name, operation, time) values ($1, $2, $3, $4)",
				userId, slug, "add", time.Now().Format("2006-01-02 15:04:05"))
			if err != nil {
				return "Error adding to history.", err
			}

		}
		res += " segment/s added: " + addSlugs
	}
	if deleteSlugs != "" {
		for _, slug := range strings.Split(deleteSlugs, ",") {
			slugId := 0
			_ = db.QueryRow("SELECT segment_id FROM segments WHERE segment_name = $1", slug).Scan(&slugId)
			_, err = db.Exec("DELETE FROM users WHERE user_id = $1 AND segment = $2",
				userId, slugId)
			if err != nil {
				return "Segment Deletion error.", err
			}

			_, err = db.Exec("INSERT INTO usershistory (user_id, segment_name, operation, time) values ($1, $2, $3, $4)",
				userId, slug, "delete", time.Now().Format("2006-01-02 15:04:05"))
			if err != nil {
				return "Error adding to history.", err
			}
		}
		if addSlugs != "" {
			res += " and"
		}
		res += " segment/s deleted: " + deleteSlugs + "."
	}

	return res, nil
}
