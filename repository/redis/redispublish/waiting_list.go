package redispublish

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"newsapp/entity"
	"time"
)

func nowTime() int64 {
	return time.Now().UnixMicro()
}

func (d DB) AddNewsToWaitingList(ctx context.Context, newsID uint, category entity.Category) error {

	_, err := d.adapter.Client().ZAdd(ctx, fmt.Sprintf("%s:%s", "waitinglist", category),
		redis.Z{
			Score: float64(nowTime()), Member: fmt.Sprintf("%d", newsID)}).Result()
	if err != nil {
		return err
	}

	return nil
}

func (d DB) GetWaitingListNewsByCategory(ctx context.Context, category entity.Category) ([]entity.WaitingNews, error) {
	//TODO implement me
	panic("implement me")
}

func (d DB) RemoveNewsFromWaitingList(category entity.Category, newsIDs []uint) {
	//TODO implement me
	panic("implement me")
}
