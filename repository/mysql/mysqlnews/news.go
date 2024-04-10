package mysqlnews

import (
	"context"
	"newsapp/entity"
)

func (d *DB) InsertNews(news entity.News) (entity.News, error) {
	result, err := d.newsConn.Conn().Exec(`INSERT INTO user(title, short_description, description,image_file_name,creator_user_id) VALUES(?,?,?,?,?) `,
		news.Title, news.ShortDescription, news.Description, news.ImageFileName, news.CreatorUserID)
	if err != nil {
		return entity.News{}, err
	}

	insertId, _ := result.LastInsertId()

	news.ID = uint(insertId)

	return news, nil
}

func (d *DB) GetNewsByID(ctx context.Context, newsID int) (entity.News, error) {
	row := d.newsConn.Conn().QueryRowContext(ctx, ` SELECT * FROM news WHERE id = ?`, newsID)

	news, err := scanNews(row)
	if err != nil {
		return entity.News{}, err
	}

	return news, nil
}

func (d *DB) GetNewsByTitle(ctx context.Context, title string) (entity.News, error) {
	//TODO implement me
	panic("implement me")
}

func (d *DB) GetNewsList(ctx context.Context) ([]entity.News, error) {
	newsList := make([]entity.News, 0)
	rows, err := d.newsConn.Conn().QueryContext(ctx, ` SELECT * FROM news`)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		news, err := scanNews(rows)
		if err != nil {
			return nil, err
		}
		newsList = append(newsList, news)
	}

	return newsList, nil
}

func (d *DB) UpdateNewsByModel(ctx context.Context, news entity.News) (entity.News, error) {
	//TODO implement me
	panic("implement me")
}

func (d *DB) DeleteNews(ctx context.Context, newsID string) error {
	//TODO implement me
	panic("implement me")
}
