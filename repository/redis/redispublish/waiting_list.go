package redispublish

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"newsapp/entity"
	"strconv"
	"time"
)

func getCategoryKey(category entity.Category) string {
	return fmt.Sprintf("%s:%s", "waitinglist", category)
}

func nowTime() int64 {
	return time.Now().UnixMicro()
}

func (d DB) AddNewsToWaitingList(ctx context.Context, newsID uint, category entity.Category) error {

	_, err := d.adapter.Client().ZAdd(ctx, getCategoryKey(category),
		redis.Z{
			Score: float64(nowTime()), Member: fmt.Sprintf("%d", newsID)}).Result()
	if err != nil {
		return err
	}

	return nil
}

func (d DB) GetWaitingListNewsByCategory(ctx context.Context, category entity.Category) ([]entity.WaitingNews, error) {

	min := fmt.Sprintf("%d", time.Now().Add(-2*time.Hour).UnixMicro())
	max := strconv.Itoa(int(nowTime()))
	list, err := d.adapter.Client().ZRangeByScoreWithScores(ctx, getCategoryKey(category), &redis.ZRangeBy{
		Min:    min,
		Max:    max,
		Offset: 0,
		Count:  0,
	}).Result()
	if err != nil {
		return nil, err
	}

	var reult = make([]entity.WaitingNews, 0)

	for _, l := range list {
		newsID, _ := strconv.Atoi(l.Member.(string))

		reult = append(reult, entity.WaitingNews{
			NewsID:    uint(newsID),
			Timestamp: int64(l.Score),
			Category:  category,
		})
	}

	return reult, nil
}

func (d DB) RemoveNewsFromWaitingList(category entity.Category, newsIDs []uint) {
	//TODO implement me
	panic("implement me")
}
