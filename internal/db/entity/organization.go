package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Organization struct {
	ModelBase
	Name string `json:"name"`
}

func (Organization) TableName() string {
	return "organizations"
}

func (o *Organization) Up(client *gorm.DB) {
	err := client.AutoMigrate(&Organization{})
	if err != nil {
		panic("error migrating organizations table")
	}
}

func (o *Organization) Create(client *gorm.DB) {
	 client.Create(o)
}

func (o *Organization) Read(client *gorm.DB, id uuid.UUID) (organization *Organization) {
	client.Where("id = ?", id).First(&organization)
	return organization
	
}

func (o *Organization) Update(client *gorm.DB) {
	client.Save(o)
}