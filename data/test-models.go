package data

import (
	"database/sql"
	"time"
)

type PostgresRepositoryTest struct {
	Conn *sql.DB
}

func NewPostgresRepositoryTest(db *sql.DB) *PostgresRepositoryTest {
	return &PostgresRepositoryTest{
		Conn: db,
	}
}

func (u *PostgresRepositoryTest) GetAll() ([]*User, error) {
	users := []*User{}
	return users, nil
}

func (u *PostgresRepositoryTest) GetByEmail(email string) (*User, error) {

	user := User{
		ID:        1,
		FirstName: "Agent",
		LastName:  "One",
		Email:     "agent@one.com",
		Password:  "",
		Active:    1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	return &user, nil
}

func (u *PostgresRepositoryTest) GetOne(id int) (*User, error) {
	user := User{
		ID:        1,
		FirstName: "Agent",
		LastName:  "One",
		Email:     "agent@one.com",
		Password:  "",
		Active:    1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return &user, nil
}

func (u *PostgresRepositoryTest) Update(user User) error {

	return nil
}

func (u *PostgresRepositoryTest) DeleteByID(id int) error {

	return nil
}

func (u *PostgresRepositoryTest) Insert(user User) (int, error) {

	return 2, nil
}

func (u *PostgresRepositoryTest) ResetPassword(password string, user User) error {
	return nil
}

func (u *PostgresRepositoryTest) PasswordMatches(plainText string, user User) (bool, error) {

	return true, nil
}
