package services

import (
	"log/slog"
	"terrapak/internal/db/client"
	"terrapak/internal/db/entity"

	"github.com/google/uuid"
)

type OrganizationService struct{}

func (os *OrganizationService) Create(organization entity.Organization) {
	client := client.Default()
	organization.Create(client)
}

func (os *OrganizationService) UpdateName(id uuid.UUID, name string) {

	client := client.Default()
	organization := entity.Organization{}
	organization = *organization.Read(client, id)
	if organization.ID != uuid.Nil {
		organization.Name = name
		organization.Update(client)
	}else {
		slog.Error("Organization not found when attempting to update name")
	}
}

func (os *OrganizationService) FindByName(name string) (organization entity.Organization) {
	client := client.Default()
	client.Raw("SELECT * FROM organizations WHERE name = ?",name).Scan(&organization)
	return organization
}