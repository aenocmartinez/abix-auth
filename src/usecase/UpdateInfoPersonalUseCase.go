package usecase

import (
	"abix360/src/dao/mysql"
	"abix360/src/domain"
	"abix360/src/view/dto"
	"errors"
)

type UpdateInfoPersonalUseCase struct{}

func (useCase *UpdateInfoPersonalUseCase) Execute(infoPersonal dto.InfoPersonalDTO) (int, error) {
	var repository domain.UserRepository = mysql.NewUserDao()
	user := domain.FindUserById(infoPersonal.Id, repository)
	if !user.Exists() {
		return 202, errors.New("el usuario no existe")
	}

	user.WithEmail(infoPersonal.Email).WithName(infoPersonal.Name).WithRepository(repository)
	err := user.Update()
	if err != nil {
		return 202, err
	}

	return 200, nil
}
