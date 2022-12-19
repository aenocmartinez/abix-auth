package usecase

import (
	"abix360/src/dao/mysql"
	"abix360/src/domain"
	"abix360/src/view/dto"
)

type ListUsersUseCase struct{}

func (useCase *ListUsersUseCase) Execute() []dto.UserDto {
	var usersDto []dto.UserDto
	var repository domain.UserRepository = mysql.NewUserDao()

	users := domain.AllUsers(repository)

	for _, user := range users {
		usersDto = append(usersDto, dto.UserDto{
			Id:        user.Id(),
			Name:      user.Name(),
			Email:     user.Email(),
			State:     user.State(),
			CreatedAt: user.CreatedAt(),
		})
	}

	return usersDto
}
