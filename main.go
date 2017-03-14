package main

import (
	"golern/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	//"strconv"
)

func main() {
	//users := models.AllUsers()
	//fmt.Println(users)
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/users", UsersIndex).Methods("GET")
	router.HandleFunc("/users/{id}", UsersShow).Methods("GET")
	router.HandleFunc("/users", UsersCreate).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func UsersIndex(w http.ResponseWriter, r *http.Request) {
	users, err := models.AllUsers()
	if err != nil {
		fmt.Println(err)
	}
	json.NewEncoder(w).Encode(users)
}

func UsersCreate(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	user := models.User{First_name: (params["name"])}
	fmt.Fprintln(w,  user.Save())
}

func UsersShow(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	//id, _ := strconv.ParseInt(params["id"], 0, 64)
	user, err := models.GetUser(params["id"])
	if err != nil {
		fmt.Println(err)
	}
	json.NewEncoder(w).Encode(user)
}