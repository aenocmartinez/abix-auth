package dto

type InfoPersonalDTO struct {
	Email     string `json:"email"`
	Name      string `json:"name"`
	Id        int64  `json:"id"`
	State     bool   `json:"state"`
	CreatedAt string `json:"createdAt"`
}
