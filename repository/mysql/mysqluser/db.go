package mysqluser

import (
	"newsapp/entity"
	"newsapp/repository/mysql"
	"time"
)

type DB struct {
	userConn *mysql.MySQLDB
}

func New(conn *mysql.MySQLDB) *DB {
	return &DB{userConn: conn}
}

func scanUser(scanner mysql.Scanner) (entity.User, error) {
	var registerDate time.Time
	var user entity.User

	err := scanner.Scan(&user.ID, &user.PhoneNumber, &user.Password, &user.Email, &user.LastName, &user.FirstName, user.Role, &registerDate)
	if err != nil {
		return entity.User{}, err
	}
	return user, nil

}
