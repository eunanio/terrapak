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

func (u *User) Create(client *gorm.DB, user *User) {
	client.Create(user)
}

func (u *User) Read(client *gorm.DB, email string) (user *User) {
	client.Where("token = ?", email).First(&user)
	return user
}

func (u *User) ReadAll(client *gorm.DB) (list []User) {
	client.Raw("SELECT * FROM users").Scan(&list)
	return list
}

func (u *User) Update(client *gorm.DB, user *User) {
	client.Save(&user)
}

func (u *User) Delete(client *gorm.DB, user *User) {
	client.Delete(&user)
}

