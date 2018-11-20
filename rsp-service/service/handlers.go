package service

import (
	"encoding/json"
	"net/http"
	"strconv"

	"../dbclient"
	"github.com/gorilla/mux"
)

// DBClient is the DBC instance
var DBClient dbclient.IBoltClient

// GetUser handles requests on the GetUser route
func GetUser(w http.ResponseWriter, r *http.Request) {

	// get the 'userID' path parameter from req (via mux)
	var userID = mux.Vars(r)["userID"]

	// query the user from db
	user, err := DBClient.QueryUser(userID)

	// return 404 if error
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	data, _ := json.Marshal(user)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Length", strconv.Itoa(len(data)))
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
