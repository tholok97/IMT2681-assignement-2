package currencyWebhookService

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

// SubscriberHandler handles requests from clients. It has info on subscribers
// and currency
type SubscriberHandler struct {
	db      SubscriberDB
	Monitor CurrencyMonitor
}

// SubscriberHandlerFactory returns a fresh handler
func SubscriberHandlerFactory(db SubscriberDB, monitor CurrencyMonitor) SubscriberHandler {
	handler := SubscriberHandler{db: db, Monitor: monitor}

	return handler
}

func (handler *SubscriberHandler) HandleSubscriberRequestPOST(res http.ResponseWriter, req *http.Request) {

	// attempt to decode the POST json
	var s Subscriber
	err := json.NewDecoder(req.Body).Decode(&s)

	// if couldn't decode -> bad req
	// (SHOULD ALSO FAIL FOR NON-COMPLIANT JSON)
	if err != nil {
		respWithCode(&res, http.StatusBadRequest)
		return
	}

	// check validity of posted json
	if !validateSubscriber(s) {
		respWithCode(&res, http.StatusBadRequest)
		return
	}

	// check validity of URL in posted json
	_, err = url.ParseRequestURI(*s.WebhookURL)
	if err != nil {
		respWithCode(&res, http.StatusBadRequest)
		return
	}

	// (try to) add the student
	id, addErr := handler.db.Add(s)

	// if couldn't add -> internal server error
	//  (client's responsability to retry)
	if addErr != nil {
		respWithCode(&res, http.StatusInternalServerError)
		return
	}

	// respond with id given by db
	fmt.Fprint(res, id)
}

func (handler *SubscriberHandler) HandleSubscriberRequestGET(res http.ResponseWriter, req *http.Request) {

	// try to pick out the id from the url
	parts := strings.Split(req.URL.String(), "/")
	if len(parts) < 2 {
		respWithCode(&res, http.StatusBadRequest)
		return
	}

	// attempt to fetch subscriber with given id
	sub, err := handler.db.Get(parts[1])
	if err != nil {
		respWithCode(&res, http.StatusNotFound)
		return
	}

	http.Header.Add(res.Header(), "content-type", "application/json")

	// decode and send the sub
	err = json.NewEncoder(res).Encode(sub)
	if err != nil {
		respWithCode(&res, http.StatusInternalServerError)
		return
	}

}

func (handler *SubscriberHandler) HandleSubscriberRequestDELETE(res http.ResponseWriter, req *http.Request) {

	// try to pick out the id from the url
	parts := strings.Split(req.URL.String(), "/")
	if len(parts) < 2 {
		respWithCode(&res, http.StatusBadRequest)
		return
	}

	// attempt to delete the subscriber with id
	err := handler.db.Remove(parts[1])
	if err != nil {
		respWithCode(&res, http.StatusNotFound)
		return
	}

	// deletion succeeded. yay!
	respWithCode(&res, http.StatusOK)
}

// handle subscriber requests
func (handler *SubscriberHandler) HandleSubscriberRequest(res http.ResponseWriter, req *http.Request) {

	// switch on the method of the request
	switch req.Method {
	case "POST":
		handler.HandleSubscriberRequestPOST(res, req)
	case "GET":
		handler.HandleSubscriberRequestGET(res, req)
	case "DELETE":
		handler.HandleSubscriberRequestDELETE(res, req)
	default:
		respWithCode(&res, http.StatusNotImplemented)
	}
}

// handle requests about latests data
func (handler *SubscriberHandler) HandleLatest(res http.ResponseWriter, req *http.Request) {

	// ..only supports POST method
	if req.Method != http.MethodPost {
		respWithCode(&res, http.StatusNotImplemented)
	}

	// attempt to decode the POST json
	var currReq CurrencyRequest
	err := json.NewDecoder(req.Body).Decode(&currReq)

	// if couldn't decode -> bad req
	if err != nil {
		respWithCode(&res, http.StatusBadRequest)
		return
	}

	// check validity of posted json
	if !validateCurrencyRequest(currReq) {
		respWithCode(&res, http.StatusBadRequest)
		return
	}

	// (try to) get the latest currency info
	rate, rateErr := handler.Monitor.Latest(*currReq.BaseCurrency, *currReq.TargetCurrency)

	// if couldn't get latest -> either not found or internal error
	//  (client's responsability to retry)
	if rateErr == errInvalidCurrency {
		respWithCode(&res, http.StatusBadRequest)
		return
	} else if rateErr != nil {
		respWithCode(&res, http.StatusInternalServerError)
		return
	}

	// respond with id given by db
	fmt.Fprint(res, rate)
}

// handle requests about average data
func (handler *SubscriberHandler) HandleAverage(res http.ResponseWriter, req *http.Request) {

	// ..only supports POST method
	if req.Method != http.MethodPost {
		respWithCode(&res, http.StatusNotImplemented)
	}

	// attempt to decode the POST json
	var currReq CurrencyRequest
	err := json.NewDecoder(req.Body).Decode(&currReq)

	// if couldn't decode -> bad req
	if err != nil {
		respWithCode(&res, http.StatusBadRequest)
		return
	}

	// check validity of posted json
	if !validateCurrencyRequest(currReq) {
		respWithCode(&res, http.StatusBadRequest)
		return
	}

	// (try to) get the average currency info for the last 7 days
	rate, rateErr := handler.Monitor.Average(*currReq.BaseCurrency, *currReq.TargetCurrency)

	// if couldn't get average -> either not found or internal error
	//  (client's responsability to retry)
	if rateErr == errInvalidCurrency {
		respWithCode(&res, http.StatusBadRequest)
		return
	} else if rateErr != nil {
		respWithCode(&res, http.StatusInternalServerError)
		return
	}

	// respond with id given by db
	fmt.Fprint(res, rate)
}

// handler (for testing and debug mostly) that forces all subscribers to be notfied
func (handler *SubscriberHandler) HandleEvaluationTrigger(res http.ResponseWriter, req *http.Request) {

	// only GET supported
	if req.Method != http.MethodGet {
		respWithCode(&res, http.StatusNotImplemented)
		return
	}

	// notify all subscribers
	err := handler.NotifyAll()

	if err != nil {
		respWithCode(&res, http.StatusInternalServerError)
		return
	}
}

// notify all subscribers
func (handler *SubscriberHandler) NotifyAll() error {
	subs, err := handler.db.GetAll()
	if err != nil {
		return err
	}
	for _, s := range subs {
		handler.NotifySubscriber(s)
	}
	return nil
}

// notify single subscriber
func (handler *SubscriberHandler) NotifySubscriber(s Subscriber) {
	// TODO implement notifications
	fmt.Println("Notifying ", *s.WebhookURL)
}

// utility function for responding with a simple statuscode
func respWithCode(res *http.ResponseWriter, status int) {
	http.Error(*res, http.StatusText(status), status)
}
