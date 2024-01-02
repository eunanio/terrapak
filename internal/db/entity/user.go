package entity

import (
	"terrapak/internal/api/auth/roles"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ModelBase
	Name         		string       `json:"name"`
	Email        		string       `json:"email"`
	AuthorityID  		string       `json:"authority_id"`
	OrganizationID 	   uuid.UUID  
	Organization  	   Organization `json:"organization"`
	PasswordHash       string
	Role 	   		   roles.UserRoles `json:"role"`
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

func (u *User) Create(client *gorm.DB) {
	client.Create(u)
}

func (u *User) Read(client *gorm.DB, email string) (user *User) {
	client.Where("token = ?", email).First(&user)
	return user
}

func (u *User) ReadAll(client *gorm.DB) (list []User) {
	client.Raw("SELECT * FROM users").Scan(&list)
	return list
}

func (u *User) ReadByExternalID(client *gorm.DB, id string) (user *User) {
	client.Raw("SELECT * FROM users WHERE authority_id = ?", id).Scan(&user)
	return user
}

func (u *User) Update(client *gorm.DB, user *User) {
	client.Save(&user)
}

func (u *User) Delete(client *gorm.DB, user *User) {
	client.Delete(&user)
}

