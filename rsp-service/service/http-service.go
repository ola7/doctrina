package service

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

// StartHTTPService starts an HTTP service on a given port
func StartHTTPService(port string) {

	r := NewRouter()    // see mux-router.go
	http.Handle("/", r) // injects our router to root path of http listener

	logrus.Println("Starting HTTP listener at: " + port)
	err := http.ListenAndServe(":"+port, nil) // blocking wait
	if err != nil {
		logrus.Println("An error occured starting HTTP listener at port: " + port)
		logrus.Fatalln("Error: " + err.Error())
	}
}
