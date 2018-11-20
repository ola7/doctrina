package service

import (
	"net/http"
)

// Route defines a single route
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes wraps a slice of routes
type Routes []Route

// initialize routes - refer to handler funcs in handlers.go
var routes = Routes{
	Route{
		"GetUser",         // name
		"GET",             // http method
		"/users/{userID}", // route pattern
		GetUser,           // see handlers.go
	},
}
