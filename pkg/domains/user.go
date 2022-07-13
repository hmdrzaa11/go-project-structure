package domains

type User struct {
	Username string
	Email    string
	Password string
}

type UserRepository interface {
	GetById(id int) (*User, error)
	GetByEmail(email string) (*User, error)
}
