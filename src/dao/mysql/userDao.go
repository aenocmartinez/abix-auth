package mysql

import (
	"abix360/database"
	"abix360/src/domain"

	"bytes"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type UserDao struct {
	db *database.ConnectDB
}

func ConnectDBAuth() *UserDao {
	return &UserDao{
		db: database.Instance(),
	}
}

func (u *UserDao) Create(user domain.User) error {
	var query bytes.Buffer
	query.WriteString("INSERT INTO users(name, email, password) VALUES (?,?,?)")

	stmt, err := u.db.Source().Conn().Prepare(query.String())

	if err != nil {
		log.Println("abix-auth / UserDao / Create / conn.Prepare: ", err.Error())
	}

	_, err = stmt.Exec(user.Name(), user.Email(), user.Password())
	if err != nil {
		log.Println("abix-auth / UserDao / Create / stmt.Exec: ", err.Error())
	}

	return err
}

func (u *UserDao) FindByEmail(email string) domain.User {
	var user domain.User
	var cad bytes.Buffer

	cad.WriteString("SELECT id, name, email, password FROM users WHERE email = ?")
	row := u.db.Source().Conn().QueryRow(cad.String(), email)

	var name, password string
	var id int64

	row.Scan(&id, &name, &email, &password)

	user = *domain.NewUser(name, email).WithId(id).WithPassword(password)

	return user
}
