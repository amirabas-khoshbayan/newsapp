package mysqlnews

import (
	"newsapp/entity"
	"newsapp/repository/mysql"
	"time"
)

type DB struct {
	newsConn *mysql.MySQLDB
}

func New(conn *mysql.MySQLDB) *DB {
	return &DB{newsConn: conn}
}

func scanUser(scanner mysql.Scanner) (entity.News, error) {
	var createdAt time.Time
	var news entity.News

	err := scanner.Scan(&news.ID, &news.Title, &news.ShortDescription, &news.Description, &news.ImageFileName, &news.CreatorUserID, &news.VisitCount, &news.LikeCount, &createdAt)
	if err != nil {
		return entity.News{}, err
	}

	news.CreatedAt = createdAt

	return news, nil

}
