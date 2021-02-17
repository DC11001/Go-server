package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

//Defining the type route and routes as an array of route
type Route struct {
	Name       string
	MethodType string
	Path       string
	Handler    http.HandlerFunc
}

type Routes []Route

//This func return the mux.Router used in the http.ListenAndServe
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	//This loop trough the array setting the routes to the mux router
	for _, route := range routes {
		router.
			Name(route.Name).
			Methods(route.MethodType).
			Path(route.Path).
			Handler(route.Handler)
	}
	return router
}

//Examples routes with basic http methods. The functions are called from the actions file.
var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"PostRoute",
		"POST",
		"/postroute",
		PostFunction,
	},
	Route{
		"GetRouteWithParams",
		"GET",
		"/id/{id}",
		GetByID,
	}, Route{
		"DeleteRoute",
		"DELETE",
		"/delete",
		Delete,
	},
}
