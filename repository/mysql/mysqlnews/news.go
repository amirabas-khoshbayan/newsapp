package mysqlnews

import (
	"context"
	"newsapp/entity"
)

func (D *DB) InsertNews(news entity.News) (entity.News, error) {
	//TODO implement me
	panic("implement me")
}

func (D *DB) GetNewsByID(ctx context.Context, newsID int) (entity.News, error) {
	//TODO implement me
	panic("implement me")
}

func (D *DB) GetNewsByTitle(ctx context.Context, title string) (entity.News, error) {
	//TODO implement me
	panic("implement me")
}

func (D *DB) GetNewses(ctx context.Context) ([]entity.News, error) {
	//TODO implement me
	panic("implement me")
}

func (D *DB) UpdateNewsByModel(ctx context.Context, news entity.News) (entity.News, error) {
	//TODO implement me
	panic("implement me")
}

func (D *DB) DeleteNews(ctx context.Context, newsID string) error {
	//TODO implement me
	panic("implement me")
}
