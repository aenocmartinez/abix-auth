package usecase

import (
	"abix360/src/dao/mysql"
	"abix360/src/domain"
	abixjwt "abix360/src/infraestructure/abix-jwt"
	"errors"
	"log"
)

type LoginUseCase struct{}

func (useCase *LoginUseCase) Execute(email, password string) (abixjwt.ResponseLogin, error) {
	var responseLogin abixjwt.ResponseLogin
	var repository domain.UserRepository = mysql.ConnectDBAuth()

	var user domain.User = domain.FindUserByEmail(email, repository)
	if !user.Exists() {
		return responseLogin, errors.New("el usuario no existe")
	}

	if !user.IsActive() {
		return responseLogin, errors.New("el usuario está inactivo")
	}

	if !abixjwt.CheckPasswordHash(user.Password(), []byte(password)) {
		return responseLogin, errors.New("contraseña incorrecta")
	}

	tokendValid, err := abixjwt.GenerateJWT(email, "Admin")
	if err != nil {
		log.Println("LoginUseCase / GenerateJWT: ", err.Error())
	}

	user.WithToken(tokendValid).WithRepository(repository).UpdateToken()

	responseLogin.Email = email
	responseLogin.Token = tokendValid
	responseLogin.Id = user.Id()

	return responseLogin, nil
}
