package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
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

func (handler *SubscriberHandler) handleSubscriberRequest_POST(res http.ResponseWriter, req *http.Request) {
	// attempt to decode the POST json
	var s Subscriber
	err := json.NewDecoder(req.Body).Decode(&s)

	// if couldn't decode -> bad req
	// (SHOULD ALSO FAIL FOR NON-COMPLIANT JSON)
	if err != nil {
		status := http.StatusBadRequest
		http.Error(res, http.StatusText(status), status)
		return
	}

	// (try to) add the student
	id, addErr := handler.db.Add(s)

	// if couldn't add -> internal server error
	//  (client's responsability to retry)
	if addErr != nil {
		status := http.StatusInternalServerError
		http.Error(res, http.StatusText(status), status)
	}

	// respond with id given by db
	fmt.Fprint(res, id)
}

func (handler *SubscriberHandler) handleSubscriberRequest_GET(res http.ResponseWriter, req *http.Request) {

	// try to pick out the id from the url
	parts := strings.Split(req.URL.String(), "/")
	if len(parts) < 2 {
		status := http.StatusBadRequest
		http.Error(res, http.StatusText(status), status)
		return
	}

	// convert (string) id to int
	id, err := strconv.Atoi(parts[1])
	if err != nil {
		status := http.StatusBadRequest
		http.Error(res, http.StatusText(status), status)
		return
	}

	// attempt to fetch subscriber with given id
	sub, err := handler.db.Get(id)
	if err != nil {
		status := http.StatusNotFound
		http.Error(res, http.StatusText(status), status)
		return
	}

	// TODO hva om sending av json feiler? content-type vil være json likevel?
	http.Header.Add(res.Header(), "content-type", "application/json")

	// decode and send the sub
	err = json.NewEncoder(res).Encode(sub)
	if err != nil {
		status := http.StatusInternalServerError
		http.Error(res, http.StatusText(status), status)
		return
	}

}

// handle subscriber requests
func (handler *SubscriberHandler) handleSubscriberRequest(res http.ResponseWriter, req *http.Request) {

	// switch on the method of the request
	switch req.Method {
	case "POST":
		handler.handleSubscriberRequest_POST(res, req)
	case "GET":
		handler.handleSubscriberRequest_GET(res, req)
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
