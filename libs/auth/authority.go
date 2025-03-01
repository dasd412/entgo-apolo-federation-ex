package auth

type Role int

const (
	Admin = 1 << iota
	User
)

func ConvertRole(role string) Role {
	switch role {
	case "admin":
		return Admin
	default:
		return User
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
