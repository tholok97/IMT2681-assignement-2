package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	// (try to) get the port from heroku config vars
	port := getENV("PORT")
	schHour := getIntENV("SCHEDULE_HOUR")
	schMinute := getIntENV("SCHEDULE_MINUTE")
	schSecond := getIntENV("SCHEDULE_SECOND")
	fixerIOURL := getENV("FIXER_IO_URL")
	mongoDBURL := getENV("MONGO_DB_URL")
	mongoDBDatabaseName := getENV("MONGO_DB_CURRENCY_DATABASE_NAME")

	// set up handler (TODO will use real db and monitor eventually)
	db, err := SubscriberMongoDBFactory("localhost", "assignement_2", "subscribers")
	if err != nil {
		panic(err.Error())
	}
	monitor := FixerIOStorage{
		DatabaseURL:    mongoDBURL,
		DatabaseName:   mongoDBDatabaseName,
		CollectionName: "currencies",
		FixerIOURL:     fixerIOURL,
	}
	err = monitor.Update()
	if err != nil {
		panic("couldn't first-time-update monitor: " + err.Error())
	}

	// set up handler
	handler := SubscriberHandlerFactory(&db, &monitor)

	// set up handlerfuncs
	http.HandleFunc("/", handler.handleSubscriberRequest)
	http.HandleFunc("/latest", handler.handleLatest)
	http.HandleFunc("/average", handler.handleAverage)
	http.HandleFunc("/evaluationtrigger", handler.handleEvaluationTrigger)

	// start listening on port
	fmt.Println("Listening on port " + port + "...")
	err = http.ListenAndServe(":"+port, nil)

	// if we couldn't set up the server, give up
	if err != nil {
		panic(err)
	}

	// update monitor -> notify -> sleep
	for {
		handler.monitor.Update()
		handler.notifyAll()

		dur := durUntilClock(schHour, schMinute, schSecond)
		time.Sleep(dur)
	}
}
