package entity

type User struct {
	Name         string       `json:"name"`
	Email        string       `json:"email"`
	AuthorityID  string       `json:"authority_id"`
	Organization Organization `gorm:"references:ID" json:"id"`
}