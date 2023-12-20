package entity

import "gorm.io/gorm"

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