package middlewares

import (
	"app/globals"
	"app/logic"
	"database/sql"
	"errors"
)

// Функция получения пользователя
func GetUser(login string) (globals.User, error) {
	result, err := logic.GetToId_sql(login)
	if err != nil {
		if err == errors.New("error account") {
			return globals.User{}, errors.New("пользователь не найден")
		}
		return globals.User{}, errors.New("error sql")
	}
	return result, nil
}

// Функция создания пользователя
func CreateUser(name, login, password string) (sql.Result, error) {
	result, err := logic.Create_sql(name, login, password)
	if err != nil {
		return nil, errors.New("error sql")
	}
	return result, nil
}
