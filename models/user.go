package models

import (
	_ "github.com/mattn/go-sqlite3"
	"database/sql"
	"github.com/jmoiron/sqlx"
	"fmt"
)

type User struct {
	Id		int64	 `json:"id"`
	First_name  	string	 `json:"name"`
}

type Users []User

func AllUsers() ([]User, error) {
	db, err := sqlx.Open("sqlite3", "./golern.db")
	checkErr(err)
	defer db.Close()
	users := []User{}
	err = db.Select(&users, "SELECT * FROM user")
	if err != nil {
		return users, err
	} else {
		return users, nil
	}
}

func GetUser(id string) (*User, error) {
	var user User
	db, err := sqlx.Open("sqlite3", "./golern.db")
	defer db.Close()
	err = db.Get(&user, "SELECT * FROM user WHERE id=$1", id)
	if err != nil {
		return &User{}, err
	} else {
		return &user, nil
	}
}

func (u User) Save() User {
	db, err := sql.Open("sqlite3", "./golern.db")
	checkErr(err)
	defer db.Close()

	res, err := db.Exec("INSERT INTO user(id, first_name) values(?, ?)", nil, u.First_name)
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)
	u.Id = id
	return u
}

func (u User) String() string {
	return fmt.Sprintf("{id:%d, first_name:%s}", u.Id, u.First_name)
}

func checkErr(err error) {
	if err != nil {
		panic("sql err: " + err.Error())
	}
}