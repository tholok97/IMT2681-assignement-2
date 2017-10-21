package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// SubscriberHandler handles requests from clients. It has info on subscribers
// and currency
type SubscriberHandler struct {
	db SubscriberDB
	// currency monitor
}

// SubscriberHandlerFactory returns a fresh handler
func SubscriberHandlerFactory(db SubscriberDB) SubscriberHandler {
	handler := SubscriberHandler{db: db}
	return handler
}

// handle subscriber requests
func (handler *SubscriberHandler) handleSubscriberRequest(res http.ResponseWriter, req *http.Request) {

	// switch on the method of the request
	switch req.Method {
	case "POST":

		// attempt to decode the POST json
		var s Subscriber
		err := json.NewDecoder(req.Body).Decode(&s)

		// if couldn't decode -> bad req
		// (SHOULD ALSO FAIL FOR NON-COMPLIANT JSON)
		if err != nil {
			http.Error(res, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		// (try to) add the student
		id, addErr := handler.db.Add(s)

		// if couldn't add -> internal server error
		//  (client's responsability to retry)
		if addErr != nil {
			http.Error(res, addErr.Error(), http.StatusInternalServerError)
		}

		// respond with id given by db
		fmt.Fprint(res, id)
	default:
		status := http.StatusNotImplemented
		http.Error(res, http.StatusText(status), status)
	}
}

// handle requests about latests data
func (handler *SubscriberHandler) handleLatest(res http.ResponseWriter, req *http.Request) {
	status := http.StatusNotImplemented
	http.Error(res, http.StatusText(status), status)
}

// handle requests about average data
func (handler *SubscriberHandler) handleAverage(res http.ResponseWriter, req *http.Request) {
	status := http.StatusNotImplemented
	http.Error(res, http.StatusText(status), status)
}
