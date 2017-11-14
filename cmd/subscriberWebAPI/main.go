package main

import (
	"fmt"
	"net/http"

	service "github.com/tholok97/IMT2681-assignement-2/currencyWebhookService"
)

func main() {

	// (try to) get the port from heroku config vars
	port := service.GetENV("PORT")
	fixerIOURL := service.GetENV("FIXER_IO_URL")
	mongoDBURL := service.GetENV("MONGO_DB_URL")
	mongoDBDatabaseName := service.GetENV("MONGO_DB_DATABASE_NAME")

	// set up db
	db, err := service.SubscriberMongoDBFactory(mongoDBURL, mongoDBDatabaseName)
	if err != nil {
		panic("couldn't set up db" + err.Error())
	}

	// set up monitor
	monitor := service.FixerIOStorage{
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
	handler := service.SubscriberHandlerFactory(&db, &monitor)

	// set up handlerfuncs
	http.HandleFunc("/", handler.HandleSubscriberRequest)
	http.HandleFunc("/latest", handler.HandleLatest)
	http.HandleFunc("/average", handler.HandleAverage)
	http.HandleFunc("/evaluationtrigger", handler.HandleEvaluationTrigger)
	http.HandleFunc("/dialogFlow", handler.HandleDialogFlow)

	// start listening on port
	fmt.Println("Listening on port " + port + "...")
	err = http.ListenAndServe(":"+port, nil)

	// if we couldn't set up the server, give up
	if err != nil {
		panic(err)
	}
}
