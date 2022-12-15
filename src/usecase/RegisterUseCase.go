package usecase

import (
	"abix360/src/domain"
	"abix360/src/infraestructure/datasource/mysql"
	"errors"
)

type RegisterUseCase struct{}

func (useCase *RegisterUseCase) Execute(name, email, password string) (int, error) {
	var repository domain.UserRepository = mysql.ConnectDBAuth()
	user := domain.FindUserByEmail(email, repository)
	if user.Exists() {
		return 404, errors.New("el usuario ya se encuentra registrado")
	}

	user = *domain.NewUser(name, email).WithPassword(password).WithRepository(repository)
	return 200, user.Create()
}
