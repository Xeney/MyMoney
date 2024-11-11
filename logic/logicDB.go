package logic

import (
	"app/globals"
	"database/sql"

	// "encoding/json"
	"errors"
	"fmt"

	_ "modernc.org/sqlite"
	// "strconv"
)

// Создать пользователя
func Create_sql(name, login, password string) (sql.Result, error) {
	db, err := sql.Open("sqlite", "db.sqlite")
	if err != nil {
		return nil, errors.New("error connection")
	}
	defer db.Close()

	result, err := db.Exec("insert into users (name, login, password) values ($1, $2, $3)", name, login, password)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("error sql_exec")
	}
	return result, nil
}

// Получить пользователя по Login
func GetToId_sql(Login string) (globals.User, error) {
	db, err := sql.Open("sqlite", "db.sqlite")
	if err != nil {
		return globals.User{}, errors.New("error connection")
	}
	defer db.Close()
	row := db.QueryRow("select * from users where login = $1", Login)
	u := globals.User{}
	// var Rewards []byte
	err = row.Scan(&u.Id, &u.Name, &u.Login, &u.Password)
	if err != nil {
		fmt.Println(err)
		return globals.User{}, errors.New("error account")
	}

	return u, nil
}
