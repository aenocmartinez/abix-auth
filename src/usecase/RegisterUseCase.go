package usecase

import (
	"abix360/shared"
	"abix360/src/dao/mysql"
	"abix360/src/domain"
	"errors"
)

type RegisterUseCase struct{}

func (useCase *RegisterUseCase) Execute(name, email, password string) (int, error) {
	var repository domain.UserRepository = mysql.ConnectDBAuth()

	user := domain.FindUserByEmail(email, repository)
	if user.Exists() {
		return 202, errors.New("el usuario ya se encuentra registrado")
	}

	user = *domain.NewUser(name, email).WithRepository(repository).WithPassword(shared.HashAndSalt([]byte(password)))
	return 200, user.Create()
}
