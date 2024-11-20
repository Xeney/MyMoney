package middlewares

import (
	"app/globals"
	"app/logic"
	"database/sql"
	"errors"
	"fmt"
)

// GetUser получает пользователя по логину
func GetUser(login string) (globals.User, error) {
	result, err := logic.GetUserByLogin(login)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows): // Обработка отсутствия пользователя
			return globals.User{}, errors.New("пользователь не найден")
		default:
			return globals.User{}, fmt.Errorf("при извлечении пользователя произошла ошибка: %v", err)
		}
	}
	return result, nil
}

// CreateUser создает нового пользователя с указанными данными
func CreateUser(name, login, password string) (sql.Result, error) {
	if name == "" || login == "" || password == "" {
		return nil, errors.New("имя, логин и пароль не должны быть пустыми")
	}

	result, err := logic.CreateUser(name, login, password)
	if err != nil {
		return nil, fmt.Errorf("ошибка при создании пользователя: %v", err)
	}
	return result, nil
}
