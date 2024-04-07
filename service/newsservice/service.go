package newsservice

import (
	"context"
	"newsapp/entity"
)

type Repository interface {
	InsertNews(news entity.News) (entity.News, error)
	GetNewsByID(ctx context.Context, newsID int) (entity.News, error)
	GetNewsByTitle(ctx context.Context, title string) (entity.News, error)
	GetNewses(ctx context.Context) ([]entity.News, error)
	UpdateNewsByModel(ctx context.Context, news entity.News) (entity.News, error)
	DeleteNews(ctx context.Context, newsID string) error
}

type Service struct {
	repo Repository
}

func New(repo Repository) Service {
	return Service{repo: repo}
}
