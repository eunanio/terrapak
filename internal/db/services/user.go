package services

import (
	"terrapak/internal/config"
	"terrapak/internal/db/client"
	"terrapak/internal/db/entity"

	"github.com/google/uuid"
)

type UserService struct{}

func (us *UserService) Create(user entity.User) *entity.User {
	client := client.Default()
	org_service := OrganizationService{}
	gc := config.GetDefault()
	org := org_service.FindByName(gc.Organization)
	user.OrganizationID = org.ID
	user.Create(client)
	return &user
}

func (us *UserService) Find(id uuid.UUID) *entity.User {
	model := entity.User{}
	client := client.Default()

	return model.Read(client, id)
}

func (us *UserService) FindByExternalID(id string) *entity.User {
	model := entity.User{}
	client := client.Default()

	return model.ReadByExternalID(client, id)
}

func (us *UserService) CreateApiKey(key entity.ApiKeys) *entity.ApiKeys {
	client := client.Default()
	key.Create(client)
	return &key
}

func (us *UserService) RemoveApiKeys(user_id uuid.UUID) {
	client := client.Default()
	key := entity.ApiKeys{}
	key.DeleteByUser(client, user_id)
}

// func (us *UserService) MemberCount() int64 {
// 	model := entity.User{}
// 	client := client.Default()

// 	return model.Count(client)
// }