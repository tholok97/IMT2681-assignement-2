package main

import (
	"fmt"
	"time"

	service "github.com/tholok97/IMT2681-assignement-2/currencyWebhookService"
)

func main() {

	// (try to) get the port from heroku config vars
	schHour := service.GetIntENV("SCHEDULE_HOUR")
	schMinute := service.GetIntENV("SCHEDULE_MINUTE")
	schSecond := service.GetIntENV("SCHEDULE_SECOND")
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

	// update monitor -> notify -> sleep
	for {
		fmt.Println("Updating monitor...")
		err := monitor.Update()
		if err != nil {
			fmt.Println("!!! Failed to update monitor (", err.Error(), ")")
		}

		fmt.Println("Notifying all subscribers...")
		err = handler.NotifyAll()
		if err != nil {
			fmt.Println("!!! Failed to notify all subscribers (", err.Error(), ")")
		}

		dur := service.DurUntilClock(schHour, schMinute, schSecond)
		fmt.Println("Sleeping ", dur, "...")
		time.Sleep(dur)
		fmt.Println("Done sleeping!")
	}
}
