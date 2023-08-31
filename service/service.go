package service

import (
	"context"
	_ "github.com/lib/pq"
)

type Service interface {
	CreateSegment(ctx context.Context, slug string) (string, error)
	DeleteSegment(ctx context.Context, slug string) (string, error)
	AddUser(ctx context.Context, addSlug string, deleteSlug string, userId string) (string, error)
	UserSegments(ctx context.Context, userId string) (string, error)
	UsersHistory(ctx context.Context, year string, month string) (string, error)
}

const connStr = "user=postgres password=4581 dbname=segments_db sslmode=disable"

type slugService struct{}

func NewService() Service {
	return slugService{}
}






