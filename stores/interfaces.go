package stores

import "github.com/zopping/mock-test/models"

type Finder interface {
	Find(id int) (*models.User, error)
}
