package main

import(
	"golern/models"
	"encoding/json"
	"fmt"
	//"log"
	"net/http"
	"github.com/gorilla/mux"
)

type appHandler func(w http.ResponseWriter, r *http.Request) error

func (h appHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := h(w, r); err != nil {
		// add diferent errors hednling
		http.Error(w, http.StatusText(500), 500)
	}
}

func UsersIndex(w http.ResponseWriter, r *http.Request) error {
	users, err := models.UsersAll()
	if err != nil {
		return err
	}
	json.NewEncoder(w).Encode(users)
	return nil
}

func UsersCreate(w http.ResponseWriter, r *http.Request) error {
	name := r.URL.Query().Get("name")
	user := models.User{First_name: name}
	user, err := user.Save()
	if err != nil {
		return err
	}
	json.NewEncoder(w).Encode(user)
	return nil
}

func UsersDelete(w http.ResponseWriter, r *http.Request) error {
	// find http errors handling
	params := mux.Vars(r)
	user, err := models.UserFindById(params["id"])
	if err != nil {
		fmt.Fprintln(w, err)
	}
	user.Delete()
	fmt.Fprintln(w, "user was deleted")
	return nil

}

func UsersShow(w http.ResponseWriter, r *http.Request) error {
	params := mux.Vars(r)
	user, err := models.UserFindById(params["id"])
	if err != nil  {
		return err
		//handleError(w, r, err)
	}
	json.NewEncoder(w).Encode(user)
	return nil
}
