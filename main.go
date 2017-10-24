package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {

	// (try to) get the port from heroku config vars
	port := getENV("PORT")
	//schHour := getENV("SCHEDULE_HOUR")
	//schMinute := getENV("SCHEDULE_MINUTE")
	//schSecond := getENV("SCHEDULE_SECOND")

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
		handler.monitor.Update()
		handler.notifyAll()

		/*
			dur := durUntilTime(schHour, schMinute, schSecond)
			time.Sleep(dur)
		*/
		break // TODO
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

// calculate duration until next HH:MM:SS
func durUntilClock(hour, minute, second int) time.Duration {
	t := time.Now()

	// the time this HH:MM:SS is happening
	when := time.Date(t.Year(), t.Month(), t.Day(), hour,
		minute, second, 0, t.Location())

	// d is the time until next such time
	d := when.Sub(t)

	// if duration is negative, add a day
	if d < 0 {
		when = when.Add(24 * time.Hour)
		d = when.Sub(t)
	}

	return d
}

// calculate duration until time is when
func durUntilTime(when time.Time) time.Duration {
	return when.Sub(time.Now())
}
