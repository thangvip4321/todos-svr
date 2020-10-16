package storage

import (
	"database/sql"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

var cost = 7

type User struct {
	Name     string
	Password string
	ID       int64
}

type UserHandler struct {
	db *sql.DB
}

func (handler *UserHandler) CreateUser(user *User) (error, *User) {
	stmt, err := handler.db.Prepare("INSERT INTO users (Username,Passwords) value (?,?)")
	if err != nil {
		return err, nil
	}
	defer stmt.Close()
	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), cost)
	if err != nil {
		return err, nil
	}
	result, err := stmt.Exec(user.Name, encryptedPassword)
	if err != nil {
		return err, nil
	}
	if num, _ := result.RowsAffected(); num == 0 {
		return errors.New("fake account ?"), nil
	}
	id, _ := result.LastInsertId()
	rows, err := handler.db.Query("SELECT * FROM users WHERE ID = ?", id)
	rows.Next()
	newUser := User{}
	var pass []byte
	rows.Scan(&newUser.Name, &pass, &newUser.ID)
	return nil, &newUser
}

func (handler *UserHandler) DeleteUser(id int) error {
	stmt, err := handler.db.Prepare("DELETE FROM users WHERE ID=?")
	if err != nil {
		return err
	}
	res, err := stmt.Exec(id)
	if num, _ := res.RowsAffected(); num == 0 {
		return errors.New("no row affected!")
	}
	return nil

}
