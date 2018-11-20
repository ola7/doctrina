package service

import (
	"log"
	"net/http"
)

// StartHTTPService starts an HTTP service on a given port
func StartHTTPService(port string) {

	r := NewRouter()    // see mux-router.go
	http.Handle("/", r) // injects our router to root path of http listener

	log.Println("Starting HTTP listener at: " + port)
	err := http.ListenAndServe(":"+port, nil) // blocking wait
	if err != nil {
		log.Fatal("An error occured starting HTTP listener at port: " + port)
		log.Fatal("Error: " + err.Error())
	}
}
