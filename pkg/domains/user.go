package domains

import "github.com/hmdrzaa11/example-go-api/pkg/dtos"

type User struct {
	Username string
	Email    string
	Password string
}

func (u *User) ToDto() *dtos.UserResponse {
	return &dtos.UserResponse{
		Username: u.Username,
		Email:    u.Email,
	}
}

type UserRepository interface {
	GetById(id int) (*User, error)
	GetByEmail(email string) (*User, error)
}
