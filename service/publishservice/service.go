package publishservice

import (
	"context"
	"newsapp/contract/broker"
	"newsapp/entity"
	"newsapp/param/newsparam"
	"time"
)

type Config struct {
	WaitingTimeout time.Duration `yaml:"waiting_timeout"`
}

type Publish interface {
	Publish(event entity.Event, payload string)
}

type Repository interface {
	AddNewsToWaitingList(ctx context.Context, newsID uint, category entity.Category) error
	GetWaitingListNewsByCategory(ctx context.Context, category entity.Category) ([]entity.WaitingNews, error)
	RemoveNewsFromWaitingList(category entity.Category, newsIDs []uint)
}

type Service struct {
	config    Config
	repo      Repository
	publisher broker.Publisher
}

func New(cfg Config, repo Repository, pub broker.Publisher) Service {
	return Service{
		config:    cfg,
		repo:      repo,
		publisher: pub,
	}
}
func (s Service) AddNewsToWaitingList(ctx context.Context, req newsparam.AddToWaitingListRequest) (newsparam.AddToWaitingListResponse, error) {
	err := s.repo.AddNewsToWaitingList(ctx, req.UserID, req.Category)
	if err != nil {
		return newsparam.AddToWaitingListResponse{}, err
	}

	return newsparam.AddToWaitingListResponse{Timeout: s.config.WaitingTimeout}, nil
}
