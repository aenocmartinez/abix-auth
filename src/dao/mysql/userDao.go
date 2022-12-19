package mysql

import (
	"abix360/database"
	"abix360/src/domain"
	"database/sql"

	"bytes"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type UserDao struct {
	db *database.ConnectDB
}

func NewUserDao() *UserDao {
	return &UserDao{
		db: database.Instance(),
	}
}

func (u *UserDao) Create(user domain.User) error {
	var query bytes.Buffer
	query.WriteString("INSERT INTO users(name, email, password, state) VALUES (?,?,?,?)")

	stmt, err := u.db.Source().Conn().Prepare(query.String())
	if err != nil {
		log.Println("abix-auth / UserDao / Create / conn.Prepare: ", err.Error())
	}

	_, err = stmt.Exec(user.Name(), user.Email(), user.Password(), user.State())
	if err != nil {
		log.Println("abix-auth / UserDao / Create / stmt.Exec: ", err.Error())
	}

	return err
}

func (u *UserDao) FindByEmail(email string) domain.User {
	var user domain.User
	var cad bytes.Buffer

	cad.WriteString("SELECT id, name, email, password, token, state FROM users WHERE email = ?")
	row := u.db.Source().Conn().QueryRow(cad.String(), email)

	var token sql.NullString
	var name, password string
	var state bool
	var id int64

	row.Scan(&id, &name, &email, &password, &token, &state)
	user = *domain.NewUser(name, email).WithId(id).WithPassword(password).WithState(state)

	if token.Valid {
		user.WithToken(token.String)
	}

	return user
}

func (u *UserDao) UpdateToken(id int64, token string) error {
	var query bytes.Buffer

	query.WriteString("UPDATE users SET token = ? WHERE id = ?")
	stmt, err := u.db.Source().Conn().Prepare(query.String())
	if err != nil {
		log.Println("abix-auth / UserDao / UpdateToken / conn.Prepare: ", err.Error())
	}

	_, err = stmt.Exec(token, id)
	if err != nil {
		log.Println("abix-auth / UserDao / UpdateToken / stmt.Exec: ", err.Error())
	}
	return err
}

func (u *UserDao) FindByToken(token string) domain.User {
	var user domain.User
	var cad bytes.Buffer

	cad.WriteString("SELECT id, name, email, token FROM users WHERE token = ?")
	row := u.db.Source().Conn().QueryRow(cad.String(), token)

	var tokenResult sql.NullString
	var name, email string
	var id int64

	row.Scan(&id, &name, &email, &tokenResult)
	user = *domain.NewUser(name, email).WithId(id)
	if tokenResult.Valid {
		user.WithToken(tokenResult.String)
	}

	return user
}

func (u *UserDao) FindById(id int64) domain.User {
	var user domain.User
	var cad bytes.Buffer

	cad.WriteString("SELECT id, name, email, password, token, state FROM users WHERE id = ?")
	row := u.db.Source().Conn().QueryRow(cad.String(), id)

	var token sql.NullString
	var name, email, password string
	var state bool

	row.Scan(&id, &name, &email, &password, &token, &state)
	user = *domain.NewUser(name, email).WithId(id).WithPassword(password).WithState(state)

	if token.Valid {
		user.WithToken(token.String)
	}

	return user
}

func (u *UserDao) Update(user domain.User) error {
	var query bytes.Buffer

	query.WriteString("UPDATE users SET name= ?, email=?, state=?, password=?, updated_at=NOW() WHERE id = ?")

	stmt, err := u.db.Source().Conn().Prepare(query.String())
	if err != nil {
		log.Println("abix-auth / UserDao / Update / conn.Prepare: ", err.Error())
	}

	_, err = stmt.Exec(user.Name(), user.Email(), user.State(), user.Password(), user.Id())
	if err != nil {
		log.Println("abix-auth / UserDao / Update / stmt.Exec: ", err.Error())
	}
	return err
}

func (u *UserDao) AllUsers() []domain.User {
	var users []domain.User
	var strQuery bytes.Buffer

	strQuery.WriteString("SELECT id, name, email, state, created_at FROM users order by name")
	rows, err := u.db.Source().Conn().Query(strQuery.String())
	if err != nil {
		log.Println("abix-auth / UserDao / AllUsers / s.db.Source().Conn().Query: ", err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var name, email, createdAt string
		var state bool
		var id int64
		rows.Scan(&id, &name, &email, &state, &createdAt)
		user := domain.NewUser(name, email)
		user.WithId(id).WithState(state).WithCreatedAt(createdAt)
		users = append(users, *user)
	}

	return users
}
