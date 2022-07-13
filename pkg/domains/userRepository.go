package domains

import "database/sql"

type UserRepositoryDB struct {
	db *sql.DB
}

// NewUserRepository return a new userRepository
func NewUserRepository(db *sql.DB) *UserRepositoryDB {
	return &UserRepositoryDB{db}
}

func (ur *UserRepositoryDB) GetById(id int) (*User, error) {
	return nil, nil
}

func (ur *UserRepositoryDB) GetByEmail(email string) (*User, error) {
	return nil, nil
}
