package newsservice

import (
	"context"
	"newsapp/entity"
	"newsapp/param/newsparam"
	"time"
)

type Repository interface {
	InsertNews(news entity.News) (entity.News, error)
	GetNewsByID(ctx context.Context, newsID int) (entity.News, error)
	GetNewsByTitle(ctx context.Context, title string) (entity.News, error)
	GetNewsList(ctx context.Context) ([]entity.News, error)
	UpdateNewsByModel(ctx context.Context, news entity.News) (entity.News, error)
	DeleteNews(ctx context.Context, newsID string) error
}

type Service struct {
	repo Repository
}

func New(repo Repository) Service {
	return Service{repo: repo}
}

func (s Service) CreateNewNews(req newsparam.CreateNewsRequest) (entity.News, error) {

	news := entity.News{
		Title:            req.Title,
		Description:      req.Description,
		ShortDescription: req.ShortDescription,
		ImageFileName:    req.ImageFileName,
		Categories:       req.Categories,
		CreatedAt:        time.Now(),
		CreatorUserID:    req.CreatorUserID,
	}

	newsRes, err := s.repo.InsertNews(news)
	if err != nil {
		return entity.News{}, err
	}

	return newsRes, nil
}
