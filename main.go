package main

import (
	"golern/config"
	"net/http"
	"github.com/gorilla/mux"
	"log"
)

func main() {
	config.InitDB();

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)

	}
	log.Fatal(http.ListenAndServe(":8080", router))
}

