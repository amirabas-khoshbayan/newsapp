package scheduler

import (
	"github.com/go-co-op/gocron"
	"github.com/labstack/gommon/log"
	"newsapp/service/publishservice"
	"sync"
)

type Config struct {
	PublishWaitedNewsInSeconds int `yaml:"publish_waited_news_in_seconds"`
}

type Scheduler struct {
	config     Config
	schedule   gocron.Scheduler
	publishSvc publishservice.Publish
}

func New(cfg Config, publishSvc publishservice.Publish) Scheduler {
	return Scheduler{
		config:     cfg,
		publishSvc: publishSvc,
	}
}

func (s *Scheduler) Start(done <-chan bool, wg *sync.WaitGroup) {
	defer wg.Done()

	s.schedule.Every(s.config.PublishWaitedNewsInSeconds).Second().Do(s.PublishWaitedNewsList)

	s.schedule.StartAsync()

	// waited for gracefully shutdown
	<-done

	log.EnableColor()
	log.Info("stop scheduler ...")
	s.schedule.Stop()

}

func (s *Scheduler) PublishWaitedNewsList() {

}
