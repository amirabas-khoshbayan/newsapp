package scheduler

import (
	"context"
	"fmt"
	"github.com/go-co-op/gocron"
	"github.com/labstack/gommon/log"
	"newsapp/service/publishservice"
	"sync"
	"time"
)

type Config struct {
	PublishWaitedNewsInSeconds int `yaml:"publish_waited_news_in_seconds"`
}

type Scheduler struct {
	config     Config
	schedule   gocron.Scheduler
	publishSvc publishservice.Service
}

func New(cfg Config, publishSvc publishservice.Service) Scheduler {
	return Scheduler{
		config:     cfg,
		publishSvc: publishSvc,
	}
}

func (s *Scheduler) Start(done <-chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("here", s.config.PublishWaitedNewsInSeconds)

	s.schedule.Every(s.config.PublishWaitedNewsInSeconds).Second().Do(s.PublishWaitedNewsList)

	s.schedule.StartAsync()

	// waited for gracefully shutdown
	<-done

	log.EnableColor()
	log.Info("stop scheduler ...")
	s.schedule.Stop()

}

func (s *Scheduler) PublishWaitedNewsList() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()

	err := s.publishSvc.PublishWaitedNewsList(ctx)
	if err != nil {
		log.Error("Scheduler.publishSvc.PublishWaitedNewsList() error : ", err.Error())
	}
}
