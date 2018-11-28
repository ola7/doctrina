package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"strconv"

	"../dbclient"
	"../model"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

// DBClient is the DBC instance
var DBClient dbclient.IBoltClient

// allows health state to be overriden for testability
var isHealthy = true

var client = &http.Client{}

// this method is always called first and only once
// we're doing a bypass of keepalives here
func init() {
	var transport http.RoundTripper = &http.Transport{
		DisableKeepAlives: true,
	}
	client.Transport = transport
}

// GetUser handles requests on the /users route
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

	// assign local ip
	user.ServedBy = getIP()

	// invoke the quotes-service
	q, err := getQuote()
	if err == nil {
		user.Quote = q
	}

	data, _ := json.Marshal(user)
	writeJSONResponse(w, http.StatusOK, data)
}

// HealthCheck handles request on the /health route
func HealthCheck(w http.ResponseWriter, r *http.Request) {

	// since we're here http service is known to be up - just check db connection
	dbUp := DBClient.CheckStatus()
	if dbUp && isHealthy {
		data, _ := json.Marshal(healthCheckResponse{Status: "up"})
		writeJSONResponse(w, http.StatusOK, data) // returns 200 if ok
	} else {
		data, _ := json.Marshal(healthCheckResponse{Status: "database unaccessible"})
		writeJSONResponse(w, http.StatusServiceUnavailable, data) // returns 503 if bad
	}
}

// SetHealthState allows a caller to override health check for testing purposes
func SetHealthState(w http.ResponseWriter, r *http.Request) {

	var state, err = strconv.ParseBool(mux.Vars(r)["state"])

	if err != nil {
		logrus.Println("health state must be true or false")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	isHealthy = state
	w.WriteHeader(http.StatusOK)
}

func getQuote() (model.Quote, error) {
	req, _ := http.NewRequest("GET", "http://quotes-service:8080/api/quote?strength=4", nil)
	resp, err := client.Do(req)

	if err == nil && resp.StatusCode == 200 {
		quote := model.Quote{}
		bytes, _ := ioutil.ReadAll(resp.Body)
		json.Unmarshal(bytes, &quote)
		return quote, nil
	} else {
		return model.Quote{}, fmt.Errorf("Some error")
	}
}

func getIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		logrus.Println("Error occurred getting IP:", err)
		return "n/a"
	}
	for _, address := range addrs {
		// check the address type and if it is not a loopback the display it
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	panic("unable to determine non-loopback ip address - exiting")
}

type healthCheckResponse struct {
	Status string `json:"status"`
}

func writeJSONResponse(w http.ResponseWriter, status int, data []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Length", strconv.Itoa(len(data)))
	w.WriteHeader(status)
	w.Write(data)
}
