package formrequest

type InfoPersonalFormRequest struct {
	Email string `json:"email" binding:"required,email"`
	Name  string `json:"name" binding:"required"`
	Id    int64  `json:"id" binding:"required"`
}
