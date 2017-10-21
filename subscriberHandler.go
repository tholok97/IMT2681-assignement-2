package main

import "net/http"

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
func (db *SubscriberHandler) handleSubscriberRequest(res http.ResponseWriter, req *http.Request) {
	status := http.StatusNotImplemented
	http.Error(res, http.StatusText(status), status)
}

// handle requests about latests data
func (db *SubscriberHandler) handleLatest(res http.ResponseWriter, req *http.Request) {
	status := http.StatusNotImplemented
	http.Error(res, http.StatusText(status), status)
}

// handle requests about average data
func (db *SubscriberHandler) handleAverage(res http.ResponseWriter, req *http.Request) {
	status := http.StatusNotImplemented
	http.Error(res, http.StatusText(status), status)
}
