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
	http.HandleFunc("/", handleRequest)
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

func handleRequest(res http.ResponseWriter, req *http.Request) {
	status := http.StatusNotImplemented
	http.Error(res, http.StatusText(status), status)
}

func handleLatest(res http.ResponseWriter, req *http.Request) {
	status := http.StatusNotImplemented
	http.Error(res, http.StatusText(status), status)
}

func handleAverage(res http.ResponseWriter, req *http.Request) {
	status := http.StatusNotImplemented
	http.Error(res, http.StatusText(status), status)
}
