package services

import (
	"errors"

	"github.com/chentihe/mongodb-api/models"
)

type UserService struct {
}

var users = []models.User{
	{
		Id:       "1",
		Name:     "Foo",
		Password: "$2a$10$kaUYL69SIghwwcsAsxRYT.EkFht264VD/DFut8aFHCESwl1WcHUoS"},
	{
		Id:       "2",
		Name:     "Bar",
		Password: "$2a$10$kaUYL69SIghwwcsAsxRYT.EkFht264VD/DFut8aFHCESwl1WcHUoS"},
}

func NewUserService() *UserService {
	return &UserService{}
}

func (s *UserService) GetUserByName(name string) (res *models.User, err error) {
	for _, user := range users {
		if user.Name == name {
			res = &user
		}
	}

	if res == nil {
		return nil, errors.New("Invalid user name")
	}

	return res, nil
}
