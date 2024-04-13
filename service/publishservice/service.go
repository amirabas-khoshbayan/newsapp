package publishservice

import (
	"context"
	"newsapp/contract/broker"
	"newsapp/entity"
	"newsapp/param/newsparam"
	"newsapp/pkg/protobufencoder"
	"sync"
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
	RemoveNewsFromWaitingList(ctx context.Context, category entity.Category, newsIDs []uint)
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
func (s Service) PublishWaitedNewsList(ctx context.Context) error {

	var wg sync.WaitGroup
	for _, cat := range entity.CategoryList() {
		wg.Add(1)
		go s.PublishNews(ctx, cat, &wg)
	}

	wg.Wait()

	return nil
}
func (s Service) PublishNews(ctx context.Context, category entity.Category, wg *sync.WaitGroup) {

	defer wg.Done()

	waitingNews, err := s.repo.GetWaitingListNewsByCategory(ctx, category)
	if err != nil {
		return
	}

	publishedNewsToBeRemoved := make([]uint, 0)
	for _, news := range waitingNews {
		protoBufNews := entity.PublishedNews{Category: news.Category, NewsIDs: []uint{news.NewsID}}

		go s.publisher.Publish(entity.PublishingNewsPublishedEvent, protobufencoder.EncodePublishingNewsPublishedEvent(protoBufNews))

		publishedNewsToBeRemoved = append(publishedNewsToBeRemoved, news.NewsID)
	}

	go s.repo.RemoveNewsFromWaitingList(ctx, category, publishedNewsToBeRemoved)
}
