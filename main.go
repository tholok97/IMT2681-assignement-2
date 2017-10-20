package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	// get the port heroku assignened for us
	port := os.Getenv("PORT")

	if port == "" { // ....if heroku didn't give us a port
		panic("No port specified")
	}

	// set up handlers
	http.HandleFunc("/", handleSubscriberRequest)
	http.HandleFunc("/latest/", handleLatest)
	http.HandleFunc("/average/", handleAverage)

	// start listening on port
	fmt.Println("Listening on port " + port + "...")
	err := http.ListenAndServe(":"+port, nil)

	// if error, panic
	if err != nil {
		panic(err)
	}
}

// handle subscriber requests
func handleSubscriberRequest(res http.ResponseWriter, req *http.Request) {
	status := http.StatusNotImplemented
	http.Error(res, http.StatusText(status), status)
}

// handle requests about latests data
func handleLatest(res http.ResponseWriter, req *http.Request) {
	status := http.StatusNotImplemented
	http.Error(res, http.StatusText(status), status)
}

// handle requests about average data
func handleAverage(res http.ResponseWriter, req *http.Request) {
	status := http.StatusNotImplemented
	http.Error(res, http.StatusText(status), status)
}
