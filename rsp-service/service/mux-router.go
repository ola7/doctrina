package service

import (
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

// NewRouter returns pointer to Gorilla mux.Router we can use as a handler.
func NewRouter() *mux.Router {

	// create Gorilla router
	router := mux.NewRouter().StrictSlash(true)

	// iterate over routes in http-routes.go and attach them to the router instance
	// and attach them to Gorilla instance
	for _, route := range routes {
		router.Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
		logrus.Println("Attaching route:", route.Pattern)
	}

	return router
}
