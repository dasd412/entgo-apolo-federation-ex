package auth

type Role int

const (
	UNKNOWN = 1 << iota
	Admin
	User
)

func ConvertRole(role string) Role {
	switch role {
	case "admin":
		return Admin
	case "user":
		return User
	default:
		return UNKNOWN
	}
}

type Authority interface {
	IsAdmin() bool
	IsUser() bool
	HasRole(role Role) bool
}

type UserAuthority struct {
	role Role
}

func NewAuthority(role Role) UserAuthority {
	return UserAuthority{role}
}

func (u UserAuthority) IsAdmin() bool {
	return u.role == Admin
}

func (u UserAuthority) IsUser() bool {
	return u.role == User
}

func (u UserAuthority) HasRole(role Role) bool {
	return u.role&role != 0
}
