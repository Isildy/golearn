package main

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"UsersIndex",
		"GET",
		"/users",
		UsersIndex,
	},
	Route{
		"UsersShow",
		"GET",
		"/users/{id}",
		UsersShow,
	},
	Route{
		"UsersCreate",
		"POST",
		"/users",
		UsersCreate,
	},
	Route{
		"UsersDelete",
		"GET",
		"/users/{id}",
		UsersDelete,
	},

}