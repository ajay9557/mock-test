package stores

import (
	"context"
	"database/sql"

	"github.com/zopping/mock-test/models"
)

type user struct {
	db *sql.DB
}

func New(db *sql.DB) *user {
	return &user{
		db: db,
	}
}
func (u *user) Find(id int) (*models.User, error) {
	user := &models.User{}
	query := "select id, name from users where id =? "
	rows, err := u.db.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		user = &models.User{}
		rows.Scan(&user.Id, &user.Name)
	}
	return user, nil
}

func (u *user) Create(id int, name string) (int, error) {
	query := "insert into users(id, name) values (?, ?)"
	resp, err := u.db.ExecContext(context.TODO(), query, id, name)
	if err != nil {
		return 0, err
	}
	userId, err := resp.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(userId), nil
}

func (u *user) Update(id int, name string) error {
	query := "update users set name = ? where id = ? "
	resp, err := u.db.ExecContext(context.TODO(), query, name, id)
	if err != nil {
		return err
	}
	_, err = resp.LastInsertId()
	if err != nil {
		return err
	}
	return nil
}
