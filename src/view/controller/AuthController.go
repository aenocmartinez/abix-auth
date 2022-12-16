package controller

import (
	"abix360/src/usecase"
	"abix360/src/view/dto"
	formrequest "abix360/src/view/form-request"
	"net/http"
	"strconv"

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

func Login(c *gin.Context) {
	var req formrequest.LoginFormRequest
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	useCase := usecase.LoginUseCase{}
	dataLogin, err := useCase.Execute(req.Email, req.Password)
	if err != nil {
		c.JSON(202, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": dataLogin.Token,
		"_id":   dataLogin.Id,
	})
}

func Logout(c *gin.Context) {
	useCase := usecase.LogoutUseCase{}
	code, err := useCase.Execute(c)
	if err != nil {
		c.JSON(code, gin.H{"error": err.Error()})
		return
	}
	c.JSON(code, gin.H{"message": "su sesión ha finalizado con éxito"})
}

func ResetPassword(c *gin.Context) {
	var req formrequest.ResetPasswordFormRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	useCase := usecase.ResetPasswordUseCase{}
	code, err := useCase.Execute(req.Id, req.Password)
	if err != nil {
		c.JSON(code, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "La contraseña se ha actualizado con éxito"})
}

func UpdateInfoPersonal(c *gin.Context) {
	var req formrequest.InfoPersonalFormRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	useCase := usecase.UpdateInfoPersonalUseCase{}
	code, err := useCase.Execute(dto.InfoPersonalDTO{
		Name:  req.Name,
		Email: req.Email,
		Id:    req.Id,
	})

	if err != nil {
		c.JSON(code, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "La información se ha actualizado con éxito"})
}

func FindUser(c *gin.Context) {
	var strId string = c.Param("id")
	if len(strId) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "parámetro no válido"})
		return
	}

	id, err := strconv.Atoi(strId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "parámetro no válido"})
		return
	}

	useCase := usecase.FindUserUseCase{}
	user, err := useCase.Execute(int64(id))
	if err != nil {
		c.JSON(202, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}

func UnsuscribeUser(c *gin.Context) {
	var strId string = c.Param("id")
	if len(strId) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "parámetro no válido"})
		return
	}

	id, err := strconv.Atoi(strId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "parámetro no válido"})
		return
	}

	useCase := usecase.UnsuscribeUseCase{}
	code, err := useCase.Execute(int64(id))
	if err != nil {
		c.JSON(code, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "el usuario ha sido dado de baja del sistema"})
}
