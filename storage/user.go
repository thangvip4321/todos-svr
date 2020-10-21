package storage

import (
	"database/sql"
	"errors"

	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
)

var cost = 7

//should put this secret in config file?
type User struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	ID       int64  `json:"id"`
}

type UserHandler struct {
	Db *sql.DB
}

func (user *User) GenerateJwtKey() (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{"user_info": user})
	tokenString, err := token.SignedString([]byte(viper.GetString("secret"))) // jwt lib only support bytes
	if err != nil {
		return "", err
	}
	return tokenString, nil
} // should add a time-to-live to token or not???????
func (handler *UserHandler) CreateUser(user *User) (error, *User) {
	stmt, err := handler.Db.Prepare("INSERT INTO users (Username,Passwords) value (?,?)")
	if err != nil {
		return err, nil // duplicated entries still fallthrough, not managed yet
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
	rows, err := handler.Db.Query("SELECT * FROM users WHERE ID = ?", id)
	rows.Next()
	newUser := User{}
	var pass []byte
	rows.Scan(&newUser.ID, &newUser.Name, &pass)
	return nil, &newUser
}

func (handler *UserHandler) DeleteUser(id int) error {
	stmt, err := handler.Db.Prepare("DELETE FROM users WHERE ID=?")
	if err != nil {
		return err
	}
	res, err := stmt.Exec(id)
	if num, _ := res.RowsAffected(); num == 0 {
		return errors.New("no row affected!")
	}
	return nil

}
