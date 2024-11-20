package logic

import (
	"app/globals"
	"database/sql"

	// "encoding/json"
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"
	_ "modernc.org/sqlite"
	// "strconv"
)

// openDB открывает соединение с базой данных SQLite
func openDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite", "db.sqlite")
	if err != nil {
		return nil, errors.New("error opening database: " + err.Error())
	}
	return db, nil
}

// CreateUser создает нового пользователя
func CreateUser(name, login, password string) (sql.Result, error) {
	// Хеширование пароля перед сохранением
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("error hashing password: " + err.Error())
	}

	db, err := openDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	result, err := db.Exec("INSERT INTO users (name, login, password) VALUES (?, ?, ?)", name, login, hashedPassword)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("error executing SQL command: " + err.Error())
	}
	return result, nil
}

// GetUserByLogin получает пользователя по его логину
func GetUserByLogin(login string) (globals.User, error) {
	db, err := openDB()
	if err != nil {
		return globals.User{}, err
	}
	defer db.Close()

	row := db.QueryRow("SELECT id, name, login, password FROM users WHERE login = ?", login)
	u := globals.User{}

	err = row.Scan(&u.Id, &u.Name, &u.Login, &u.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return globals.User{}, errors.New("user not found")
		}
		fmt.Println(err)
		return globals.User{}, errors.New("error retrieving account: " + err.Error())
	}

	return u, nil
}

// VerifyPassword проверяет соответствие хешированного пароля
func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
