package usecase

import (
	"abix360/src/dao/mysql"
	"abix360/src/domain"
	abixjwt "abix360/src/infraestructure/abix-jwt"
	"errors"

	"github.com/gin-gonic/gin"
)

type LogoutUseCase struct{}

func (useCase *LogoutUseCase) Execute(c *gin.Context) (int, error) {
	var token string = abixjwt.GetTokenRequest(c)
	isValid := abixjwt.VerifyToken(token)
	if !isValid {
		return 202, errors.New("token no válido")
	}

	var repository domain.UserRepository = mysql.NewUserDao()
	user := domain.FindUserByToken(token, repository)
	if !user.Exists() {
		return 202, errors.New("su sesión ha caducado")
	}

	user.WithRepository(repository).WithToken("").UpdateToken()

	// if token != user.Token() {
	// 	return 202, errors.New("su sesión ha caducado 2")
	// }

	return 200, nil
}
