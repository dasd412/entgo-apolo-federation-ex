package auth

type Role int

const (
	UNKNOWN = 1 << iota
	Admin
	Author
	Guest
)

func ConvertRole(role string) Role {
	switch role {
	case "admin":
		return Admin
	case "author":
		return Author
	case "user":
		return Guest
	default:
		return UNKNOWN
	}
}

type Authority interface {
	IsAdmin() bool
	IsAuthor() bool
	IsGuest() bool
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

func (u UserAuthority) IsAuthor() bool {
	return u.role == Author
}

func (u UserAuthority) IsGuest() bool {
	return u.role == Guest
}

func (u UserAuthority) HasRole(role Role) bool {
	return u.role&role != 0
}
