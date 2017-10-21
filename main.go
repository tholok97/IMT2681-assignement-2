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

	// set up handler
	db := VolatileSubscriberDBFactory()
	handler := SubscriberHandlerFactory(&db)

	// set up handlerfuncs
	http.HandleFunc("/", handler.handleSubscriberRequest)
	http.HandleFunc("/latest/", handler.handleLatest)
	http.HandleFunc("/average/", handler.handleAverage)

	// start listening on port
	fmt.Println("Listening on port " + port + "...")
	err := http.ListenAndServe(":"+port, nil)

	// if error, panic
	if err != nil {
		panic(err)
	}
}
