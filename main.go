package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	fios := FixerIOStorage{
		DatabaseURL:    "localhost",
		DatabaseName:   "assignement_2",
		CollectionName: "currencies",
	}

	supererr := fios.Update("http://api.fixer.io")
	if supererr != nil {
		panic(supererr.Error())
	}

	// (try to) get the port from heroku config vars
	port := os.Getenv("PORT")
	if port == "" {
		panic("No port specified")
	}

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
}
