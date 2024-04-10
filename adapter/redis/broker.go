package redis

import (
	"context"
	"github.com/labstack/gommon/log"
	"newsapp/entity"
	"time"
)

func (a Adapter) Publish(event entity.Event, payload string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
	defer cancel()

	if err := a.redisClient.Publish(ctx, string(event), payload).Err(); err != nil {
		log.Error("Adapter.redisClient.Publish error : ", err.Error())
	}
}
