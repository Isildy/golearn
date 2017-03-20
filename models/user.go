package models

import (
	"golern/config"
	_ "github.com/mattn/go-sqlite3"
	"fmt"
)

type User struct {
	Id		int64	 `json:"id"`
	First_name  	string	 `json:"name"`
}

type Users []User

func UsersAll() ([]User, error) {
	users := []User{}
	err := config.DB.Select(&users, "SELECT * FROM user")
	if err != nil {
		fmt.Println(err)
		return users, err
	} else {
		return users, nil
	}
}

func UserFindById(id string) (*User, error) {
	var user User
	err := config.DB.Get(&user, "SELECT * FROM user WHERE id=$1", id)
	if err != nil {
		return &User{}, err
	} else {
		return &user, nil
	}
}

func (u User) Save() (User, error) {
	res, err := config.DB.Exec("INSERT INTO user(id, first_name) values(?, ?)", nil, u.First_name)
	id, err := res.LastInsertId()
	u.Id = id
	if err != nil {
		return User{}, err
	} else {
		return u, nil
	}
}

func (u User) Delete() (User, error){
	_, err := config.DB.Exec("delete from user where id=?", u.Id)
	if err != nil {
		return User{}, err
	} else {
		return u, nil
	}
}

func (u User) String() string {
	return fmt.Sprintf("{id:%d, first_name:%s}", u.Id, u.First_name)
}


