package usecase

import (
	"abix360/src/dao/mysql"
	"abix360/src/domain"
	"abix360/src/view/dto"
	"errors"
)

type FindUserUseCase struct{}

func (useCase *FindUserUseCase) Execute(id int64) (dto.InfoPersonalDTO, error) {
	var userDto dto.InfoPersonalDTO
	var repository domain.UserRepository = mysql.NewUserDao()
	user := domain.FindUserById(id, repository)
	if !user.Exists() {
		return userDto, errors.New("el usuario no existe")
	}

	userDto.Email = user.Email()
	userDto.Name = user.Name()
	userDto.Id = user.Id()
	userDto.State = user.State()

	return userDto, nil
}
