package domain

type UserRepository interface {
	Create(user User) error
	FindByEmail(email string) User
	UpdateToken(id int64, token string) error
	FindByToken(token string) User
	FindById(id int64) User
	Update(user User) error
}
