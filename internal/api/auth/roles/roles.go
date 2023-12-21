package roles

type UserRoles int

const (
	Default UserRoles = iota
	Owner
	Editor
	Reader
)

func (role UserRoles) String() string {
	switch role {
	case Owner:
		return "OWNER"
	case Editor:
		return "EDITOR"
	case Reader:
		return "READER"
	case Default:
		return "DEFAULT"
	default:
		return "Invalid Role"

	}
}