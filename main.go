package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Id		int	 `json:"id"`
	First_name  	string	 `json:"name"`
}

type Users []User

func UsersInit() []User {
	todos := Users{
		User{Id: 1, First_name: "Write presentation"},
		User{Id: 2, First_name: "Host meetup"},
	}
	return todos
}

func main() {
	users := UsersInit
	user := &User{Id: 1, First_name: "Write presentation"}
	userout, _ := json.Marshal(user)
	fmt.Println(string(userout))
	fmt.Println("%v", users)
}