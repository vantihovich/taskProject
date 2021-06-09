package postgres

import (
	log "github.com/sirupsen/logrus"
)

type UsersProvider struct {
	db DB
}

type User struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func NewUsersProvider(db *DB) *UsersProvider {
	return &UsersProvider{db: *db}
}

func (u *UsersProvider) FindUserByEmailAndPassword(email, password string) (*User, error) {
	us := User{Email: email, Password: password}
	row := u.db.QueryRow(`Select id from users where username=$1 and password=$2;`, email, password)
	if err := row.Scan(&us.ID); err != nil {
		return nil, err
	}
	log.WithFields(log.Fields{}).Info("The query executed successfully")

	return &us, nil
}
