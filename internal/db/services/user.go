package services

import (
	"terrapak/internal/config"
	"terrapak/internal/db/client"
	"terrapak/internal/db/entity"
)

type UserService struct{}

func (us *UserService) Create(user entity.User) {
	client := client.Default()
	org_service := OrganizationService{}
	gc := config.GetDefault()
	org := org_service.FindByName(gc.Organization)
	user.OrganizationID = org.ID
	user.Create(client)
}

func (us *UserService) Find(email string) *entity.User {
	model := entity.User{}
	client := client.Default()

	return model.Read(client, email)
}

func (us *UserService) FindByExternalID(id string) *entity.User {
	model := entity.User{}
	client := client.Default()

	return model.ReadByExternalID(client, id)
}

// func (us *UserService) MemberCount() int64 {
// 	model := entity.User{}
// 	client := client.Default()

// 	return model.Count(client)
// }