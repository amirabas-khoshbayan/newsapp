package publishservice

import (
	"context"
	"newsapp/contract/broker"
	"newsapp/entity"
	"time"
)

type Publish interface {
	Publish(event entity.Event, payload string)
}

type Repository interface {
	AddNewsToWaitingList(ctx context.Context, newsID uint, category entity.Category) error
	GetWaitingListNewsByCategory(ctx context.Context, category entity.Category) ([]entity.WaitingNews, error)
	RemoveNewsFromWaitingList(category entity.Category, newsIDs []uint)
}

type Config struct {
	WaitingTimeout time.Duration `yaml:"waiting_timeout"`
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
