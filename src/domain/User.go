package domain

type User struct {
	repository UserRepository
	id         int64
	name       string
	email      string
	password   string
	state      bool
	createdAt  string
	updatedAt  string
}

func NewUser(name, email string) *User {
	return &User{
		name:  name,
		email: email,
	}
}

func (u *User) Id() int64 {
	return u.id
}

func (u *User) Name() string {
	return u.name
}

func (u *User) Email() string {
	return u.email
}

func (u *User) Password() string {
	return u.password
}

func (u *User) State() bool {
	return u.state
}

func (u *User) WithId(id int64) *User {
	u.id = id
	return u
}

func (u *User) WithName(name string) *User {
	u.name = name
	return u
}

func (u *User) WithEmail(email string) *User {
	u.email = email
	return u
}

func (u *User) WithPassword(password string) *User {
	u.password = password
	return u
}

func (u *User) WithState(state bool) *User {
	u.state = state
	return u
}

func (u *User) WithRepository(repository UserRepository) *User {
	u.repository = repository
	return u
}

func (u *User) WithCreatedAt(createdAt string) *User {
	u.createdAt = createdAt
	return u
}

func (u *User) WithUpdatedAt(updatedAt string) *User {
	u.updatedAt = updatedAt
	return u
}

func (u *User) CreatedAt() string {
	return u.createdAt
}

func (u *User) UpdatedAt() string {
	return u.updatedAt
}

func (u *User) Exists() bool {
	return u.id > 0
}

func (u *User) Create() error {
	return u.repository.Create(*u)
}

func FindUserByEmail(email string, repository UserRepository) User {
	return repository.FindByEmail(email)
}
