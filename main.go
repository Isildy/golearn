package main

import (
	"golern/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"strconv"
)

func main() {
	//users := models.AllUsers()
	//fmt.Println(users)
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/users", UsersIndex)
	router.HandleFunc("/users/{id}", UsersShow)
	log.Fatal(http.ListenAndServe(":8080", router))
}

func UsersIndex(w http.ResponseWriter, r *http.Request) {
	users, err := models.AllUsers()
	if err != nil {
		fmt.Println(err)
	}
	json.NewEncoder(w).Encode(users)
}

func UsersShow(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.ParseInt(params["id"], 0, 64)
	user := models.User{Id: id, First_name: "testuser"}
	fmt.Fprintln(user.Save())
}