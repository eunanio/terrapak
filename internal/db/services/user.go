package services

import (
	"terrapak/internal/db/client"
	"terrapak/internal/db/entity"
)

type UserService struct{}

func (us *UserService) Create(user entity.User) {
	client := client.Default()
	user.Create(client, &user)
}

func (us *UserService) Find(email string) *entity.User {
	model := entity.User{}
	client := client.Default()

	return model.Read(client, email)
}