package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	// (try to) get the port from heroku config vars
	port := getENV("PORT")
	currencyAPI := "testing"

	// set up handler (TODO will use real db and monitor eventually)
	db := VolatileSubscriberDBFactory()
	monitor := StubCurrencyMonitorFactory(nil, 1.6)
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

	// update monitor -> notify -> sleep
	for {
		handler.monitor.Update(currencyAPI)
		handler.notifyAll()

		// calculate time until next update/notify and sleep
		break // DEBUG
	}
}

// get environment variable. If something goes wrong: PANIC
func getENV(name string) string {
	ret := os.Getenv(name)
	if ret == "" {
		panic("Missing env variable: " + ret)
	}
	fmt.Println("Read env ", name, " = ", ret)
	return ret
}
