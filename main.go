package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	// (try to) get the envs
	port := os.Getenv("PORT")
	if port == "" {
		panic("No port specified")
	}
	fixerIOURL := os.Getenv("FIXER_IO_URL")
	if fixerIOURL == "" {
		panic("No fixerIOURL specified")
	}

	// set up handler (TODO will use real db and monitor eventually)
	db := VolatileSubscriberDBFactory()
	monitor := FixerIOStorage{
		DatabaseURL:    "localhost",
		DatabaseName:   "assignement_2",
		CollectionName: "currencies",
		FixerIOURL:     fixerIOURL,
	}

	monitor.Update()

	handler := SubscriberHandlerFactory(&db, &monitor)

	// set up handlerfuncs
	http.HandleFunc("/", handler.handleSubscriberRequest)
	http.HandleFunc("/latest", handler.handleLatest)
	http.HandleFunc("/average", handler.handleAverage)

	// start listening on port
	fmt.Println("Listening on port " + port + "...")
	err := http.ListenAndServe(":"+port, nil)

	// if we couldn't set up the server, give up
	if err != nil {
		panic(err)
	}
}
