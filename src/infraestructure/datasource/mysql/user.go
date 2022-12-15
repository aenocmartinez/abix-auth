package mysql

import (
	"abix360/src/domain"
)

type MySQL struct {
}

func ConnectDBAuth() *MySQL {
	return &MySQL{}
}

func (mysql *MySQL) Create(user domain.User) error {
	return nil
}

func (mysql *MySQL) FindByEmail(email string) domain.User {
	return domain.User{}
}
