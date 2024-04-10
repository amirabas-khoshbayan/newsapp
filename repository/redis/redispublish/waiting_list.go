package redispublish

import (
	"context"
	"newsapp/entity"
)

func (d DB) AddNewsToWaitingList(ctx context.Context, newsID uint, category entity.Category) error {
	//TODO implement me
	panic("implement me")
}

func (d DB) GetWaitingListNewsByCategory(ctx context.Context, category entity.Category) ([]entity.WaitingNews, error) {
	//TODO implement me
	panic("implement me")
}

func (d DB) RemoveNewsFromWaitingList(category entity.Category, newsIDs []uint) {
	//TODO implement me
	panic("implement me")
}
