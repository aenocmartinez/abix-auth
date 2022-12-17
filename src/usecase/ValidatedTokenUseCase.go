package usecase

import (
	"abix360/src/dao/mysql"
	"abix360/src/domain"
	abixjwt "abix360/src/infraestructure/abix-jwt"

	"github.com/gin-gonic/gin"
)

type ValidatedTokenUseCase struct {
}

func (v *ValidatedTokenUseCase) Execute(c *gin.Context) bool {
	token := abixjwt.GetTokenRequest(c)

	if !abixjwt.VerifyToken(token) {
		return false
	}

	var repository domain.UserRepository = mysql.NewUserDao()
	user := domain.FindUserByToken(token, repository)

	return user.Exists()
}
