package formrequest

type ResetPasswordFormRequest struct {
	Id       int64  `json:"id" binding:"required"`
	Password string `json:"password" binding:"required"`
}
