package main

import (
	"flag"
	"net/http"
	"os"
)

func main() {

	// gets flag from app execution (e.g. -port=80)
	port := flag.String("port", "80", "port on localhost to check")
	flag.Parse()

	resp, err := http.Get("http://127.0.0.1:" + *port + "/health")

	if err != nil || resp.StatusCode != 200 {
		os.Exit(1)
	}

	os.Exit(0)
}
