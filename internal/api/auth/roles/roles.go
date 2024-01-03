package roles

import (
	"fmt"
	"strings"
)

type UserRoles int

const (
	Default UserRoles = iota
	Reader
	Editor
	Owner
)

const (
	default_name = ""
	owner_name   = "owner"
	editor_name  = "editor"
	reader_name  = "reader"
)

func (role UserRoles) String() string {
	switch role {
	case Owner:
		return owner_name
	case Editor:
		return editor_name
	case Reader:
		return reader_name
	default:
		return default_name

	}
}

func Parse(name string) UserRoles {
	switch name {
	case owner_name:
		return Owner
	case editor_name:
		return Editor
	case reader_name:
		return Reader
	default:
		return Default
	}
}

func ParseEmailRoles(role_by_email string) (email string, role UserRoles, err error) {
	if (!strings.Contains(role_by_email,"@")){
		return "",-1,fmt.Errorf("Roles must contain a valid email")
	}

	if (!strings.Contains(role_by_email,":")){
		return "",-1,fmt.Errorf("malformed role, Missing \":\"")
	}

	items := strings.Split(":",role_by_email)
	if len(items) == 2 {
		email = items[0]
		role = Parse(items[1])
		return email, role, nil
	}

	return "", -1, fmt.Errorf("Error Occured parsing email role")
}