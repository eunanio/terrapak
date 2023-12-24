package token

import (
	"terrapak/internal/api/auth/roles"
	"terrapak/internal/db/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func CreateDefaultPAT(db_client *gorm.DB) (client *entity.ApiKeys, editor *entity.ApiKeys) {
	clientKey := uuid.New()
	editorKey := uuid.New()

	clientApikey := entity.ApiKeys{}
	editorApikey := entity.ApiKeys{}

	clientApikey.Name = "Default Terraform Client"
	clientApikey.Token = clientKey.String()
	clientApikey.Role = int(roles.Reader)

	editorApikey.Name = "Github Actions Editor"
	editorApikey.Token = editorKey.String()
	editorApikey.Role = int(roles.Editor)
	existingKeys := clientApikey.ReadAll(db_client)

	if len(existingKeys) == 0 {
		clientApikey.Create(db_client, &clientApikey)
		editorApikey.Create(db_client, &editorApikey)
		return &clientApikey, &editorApikey
	}

	return nil, nil

}