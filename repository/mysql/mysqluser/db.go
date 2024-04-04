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

	err := scanner.Scan(&user.ID, &user.FirstName, &user.LastName, &user.PhoneNumber, &user.Email, &registerDate, &user.Password, &user.Role)
	if err != nil {
		return entity.User{}, err
	}

	user.RegisterDate = registerDate

	return user, nil

}
