package usecase

import (
	"abix360/src/dao/mysql"
	"abix360/src/domain"
	abixjwt "abix360/src/infraestructure/abix-jwt"
	"errors"
)

type ResetPasswordUseCase struct{}

func (useCase *ResetPasswordUseCase) Execute(id int64, password string) (int, error) {
	var repository domain.UserRepository = mysql.ConnectDBAuth()
	user := domain.FindUserById(id, repository)
	if !user.Exists() {
		return 202, errors.New("el usuario no existe")
	}

	user.WithRepository(repository).WithPassword(abixjwt.HashAndSalt([]byte(password)))
	err := user.Update()
	if err != nil {
		return 500, err
	}

	return 200, nil
}
