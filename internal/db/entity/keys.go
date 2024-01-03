package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ApiKeys struct {
	ModelBase
	Name  string `json:"name"`
	Token string `json:"key"`
	Role  int    `json:"role"`
	UserID uuid.UUID `json:"user_id"`
}

func (ApiKeys) TableName() string {
	return "apikeys"
}

func (a *ApiKeys) Up(client *gorm.DB) {
	err := client.AutoMigrate(&ApiKeys{})
	if err != nil {
		panic("error migrating apikeys table")
	}
}

func (a *ApiKeys) Create(client *gorm.DB) {
	client.Create(a)
}

func (a *ApiKeys) Read(client *gorm.DB,token string) (apikey *ApiKeys) {
	client.Where("token = ?", token).First(&apikey)
	return apikey
}

func (a *ApiKeys) ReadAll(client *gorm.DB) (list []ApiKeys) {
	client.Raw("SELECT * FROM apikeys").Scan(&list)
	return list
}

func (a *ApiKeys) Delete(client *gorm.DB, apikey *ApiKeys) {
	client.Delete(&apikey)
}

func (a *ApiKeys) DeleteByUser(client *gorm.DB, user_id uuid.UUID) {
	client.Where("user_id = ?", user_id).Delete(&a)
}