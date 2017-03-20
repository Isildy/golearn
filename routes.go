package main

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.Handler
}

type Routes []Route

var routes = Routes{
	Route{
		"UsersIndex",
		"GET",
		"/users",
		appHandler(UsersIndex),
	},
	Route{
		"UsersShow",
		"GET",
		"/users/{id}",
		appHandler(UsersShow),
	},
	Route{
		"UsersCreate",
		"POST",
		"/users",
		appHandler(UsersCreate),
	},
	Route{
		"UsersDelete",
		"GET",
		"/users/{id}",
		appHandler(UsersDelete),
	},

}