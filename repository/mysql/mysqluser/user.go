package mysqluser

import (
	"context"
	"newsapp/entity"
)

func (d *DB) GetUserByID(ctx context.Context, userID string) (entity.User, error) {
	rows, err := d.userConn.Conn().QueryContext(ctx, ` SELECT * FROM user WHERE id = ?`, userID)
	if err != nil {
		return entity.User{}, err
	}

	user, err := scanUser(rows)
	if err != nil {
		return entity.User{}, err
	}

	return user, nil
}
func (d *DB) InsertUser(user entity.User) (entity.User, error) {
	result, err := d.userConn.Conn().Exec(`INSERT INTO user(first_name, last_name, phone_number,password,role, email) VALUES(?,?,?,?,?,?) `,
		user.FirstName, user.LastName, user.PhoneNumber, user.Password, user.Role, user.Email)
	if err != nil {

		return entity.User{}, err
	}
	insertId, _ := result.LastInsertId()

	user.ID = uint(insertId)

	return user, nil
}

func (d *DB) GetUserByPhoneNumber(ctx context.Context, phoneNumber string) (entity.User, error) {
	rows, err := d.userConn.Conn().QueryContext(ctx, ` SELECT * FROM user WHERE phone_number = ?`, phoneNumber)
	if err != nil {
		return entity.User{}, err
	}

	user, err := scanUser(rows)
	if err != nil {
		return entity.User{}, err
	}

	return user, nil
}

func (d *DB) GetUsers(ctx context.Context) ([]entity.User, error) {
	//TODO implement me
	panic("implement me")
}

func (d *DB) UpdateUserByModel(ctx context.Context, user entity.User) error {
	//TODO implement me
	panic("implement me")
}
