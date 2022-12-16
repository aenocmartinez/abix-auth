package usecase

import (
	"abix360/src/dao/mysql"
	"abix360/src/domain"
	"errors"
)

type UnsuscribeUseCase struct{}

func (useCase *UnsuscribeUseCase) Execute(id int64) (int, error) {
	var repository domain.UserRepository = mysql.NewUserDao()
	user := domain.FindUserById(id, repository)
	if !user.Exists() {
		return 202, errors.New("el usuario no existe")
	}

	user.WithState(false).WithRepository(repository)
	err := user.Update()
	if err != nil {
		return 500, err
	}

	return 200, nil
}
