package mysqluser

import (
	"context"
	"newsapp/entity"
)

func (d *DB) GetUserByID(ctx context.Context, userID string) (entity.User, error) {
	rows := d.userConn.Conn().QueryRowContext(ctx, ` SELECT * FROM user WHERE id = ?`, userID)

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
	rows := d.userConn.Conn().QueryRowContext(ctx, ` SELECT * FROM user WHERE phone_number = ?`, phoneNumber)

	user, err := scanUser(rows)
	if err != nil {
		return entity.User{}, err
	}

	return user, nil
}

func (d *DB) GetUsers(ctx context.Context) ([]entity.User, error) {
	users := make([]entity.User, 0)
	rows, err := d.userConn.Conn().QueryContext(ctx, ` SELECT * FROM user`)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		user, err := scanUser(rows)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (d *DB) UpdateUserByModel(ctx context.Context, user entity.User) error {
	row := d.userConn.Conn().QueryRowContext(ctx, `UPDATE user SET first_name = ?, last_name = ?, phone_number = ?,role = ?, email = ?  WHERE id = ? `,
		user.FirstName, user.LastName, user.PhoneNumber, user.Role, user.Email, user.ID)

	if row.Err() != nil {
		return row.Err()
	}

	return nil

}
