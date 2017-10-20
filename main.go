package main

// TODO: handle erors properly

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	// get the port heroku assignened for us
	port := os.Getenv("PORT")

	if port == "" { // ....if heroku didn't give us a port (DEBUG)
		port = "8080"
	}

	// set up default path
	http.HandleFunc("/", handleNotImplemented)

	// start listening on port
	fmt.Println("Listening on port " + port + "...")
	err := http.ListenAndServe(":"+port, nil)

	// if error, panic
	if err != nil {
		panic(err)
	}
}

func handleNotImplemented(res http.ResponseWriter, req *http.Request) {
	status := http.StatusNotImplemented
	http.Error(res, http.StatusText(status), status)
}
