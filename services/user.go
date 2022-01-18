package services

import (
	"errors"

	"github.com/zopping/mock-test/models"
	"github.com/zopping/mock-test/stores"
)

type user struct {
	store stores.Finder
}

func New(userStore stores.Finder) *user {
	return &user{
		store: userStore,
	}
}
func (u *user) Find(id int) (*models.User, error) {
	if id < 1 {
		return nil, errors.New("id not passed")
	}
	user, err := u.store.Find(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}
