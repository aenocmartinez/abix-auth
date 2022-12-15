package controller

import (
	"abix360/src/usecase"
	formrequest "abix360/src/view/form-request"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var req formrequest.RegisterFormRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	registerUseCase := usecase.RegisterUseCase{}
	code, err := registerUseCase.Execute(req.Name, req.Email, req.Password)
	if err != nil {
		c.JSON(code, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Usuario registrado exitosamente"})
}
