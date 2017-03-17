package main

import(
	"golern/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func UsersIndex(w http.ResponseWriter, r *http.Request) {
	users, err := models.UsersAll()
	if err != nil {
		handleError(w, r, err)
	}
	json.NewEncoder(w).Encode(users)
}

func UsersCreate(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	user := models.User{First_name: name}
	u, err := user.Save()
	if err != nil {
		handleError(w, r, err)
	}else {
		fmt.Fprintln(w, u)
	}
}

func UsersDelete(w http.ResponseWriter, r *http.Request) {
	// find http errors handling
	params := mux.Vars(r)
	user, err := models.UserFindById(params["id"])
	if err != nil {
		fmt.Fprintln(w, err)
	}else {
		user.Delete()
		fmt.Fprintln(w, "user was deleted")
	}
}

func UsersShow(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	user, err := models.UserFindById(params["id"])
	if err != nil  {
		handleError(w, r, err)
	}
	json.NewEncoder(w).Encode(user)
}

func handleError(w http.ResponseWriter, r *http.Request, err error) (int, error){
	log.Println(err)
	return fmt.Fprintln(w, err)
}
