package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ModelBase
	Name         string       `json:"name"`
	Email        string       `json:"email"`
	AuthorityID  string       `json:"authority_id"`
	OrganizationID uuid.UUID  
	Organization  Organization `json:"organization"`
}

func (User) TableName() string {
	return "users"
}

func (u *User) Up(client *gorm.DB) {
	err := client.AutoMigrate(&User{})
	if err != nil {
		panic("error migrating users table")
	}
}