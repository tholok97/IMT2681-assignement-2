package currencyWebhookService

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
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

// HandleSubscriberRequestPOST handles requests from incomming POSTs about subscribers
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

// HandleSubscriberRequestGET handle GETs from subscriber
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

// HandleSubscriberRequestDELETE handles incomming DELETEs on subscribers
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

// HandleSubscriberRequest handles incomming subscriber requests
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

// HandleLatest handle requests about latest info
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

// HandleAverage handle incomming requests about average info
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

// HandleEvaluationTrigger (for testing and debug mostly) that forces all subscribers to be notfied
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

// HandleDialogFlow takes request from dialogflow and returns a json with data
func (handler *SubscriberHandler) HandleDialogFlow(res http.ResponseWriter, req *http.Request) {

	fmt.Println("Handling dialogflow request")

	// only POST supported
	if req.Method != http.MethodPost {
		respWithCode(&res, http.StatusNotImplemented)
		return
	}

	// decode incomming json
	var dialogRequest DialogRequest
	err := json.NewDecoder(req.Body).Decode(&dialogRequest)

	// if couldn't decode -> bad req
	if err != nil {
		respWithCode(&res, http.StatusBadRequest)
		return
	}

	// (try and) get requested rate
	rate, rateErr := handler.Monitor.Latest(
		dialogRequest.Results.Parameters.BaseCurrency,
		dialogRequest.Results.Parameters.TargetCurrency)

	// handle errors
	if rateErr == errInvalidCurrency {
		respWithCode(&res, http.StatusBadRequest)
		return
	} else if rateErr != nil {
		respWithCode(&res, http.StatusInternalServerError)
		return
	}

	// convert rate to string (to be added in response to dialogFlow)
	rateStr := strconv.FormatFloat(float64(rate), 'f', 2, 32)

	// build response string
	respString := ""
	respString += "The exchange rate between "
	respString += dialogRequest.Results.Parameters.BaseCurrency
	respString += " and "
	respString += dialogRequest.Results.Parameters.TargetCurrency
	respString += " is: "
	respString += rateStr

	// prepare response payload
	var dialogResponse DialogResponse
	dialogResponse.DisplayText = respString
	dialogResponse.Speech = respString

	// set JSON in response header
	http.Header.Add(res.Header(), "content-type", "application/json")

	// send payload as reponse
	err = json.NewEncoder(res).Encode(dialogResponse)

	// if anything went wrong -> internal server error (our fault)
	if err != nil {
		respWithCode(&res, http.StatusInternalServerError)
		return
	}
}

// NotifyAll notifies all
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

// NotifySubscriber notifies single subscriber
func (handler *SubscriberHandler) NotifySubscriber(s Subscriber) {

	// log
	fmt.Println("Notifying ", *s.WebhookURL)

	// calc rate
	rate, err := handler.Monitor.Latest(*s.BaseCurrency, *s.TargetCurrency)
	if err != nil {
		fmt.Println("\tERROR: failed to get latest between ", *s.BaseCurrency,
			" and ", *s.MinTriggerValue)
	}

	// log
	fmt.Println("\tRate: ", rate)
	fmt.Println("\tMinTriggerValue: ", *s.MinTriggerValue)
	fmt.Println("\tMaxTriggerValue: ", *s.MaxTriggerValue)

	// should notify?
	if rate >= *s.MinTriggerValue && rate <= *s.MaxTriggerValue {

		// prepare payload
		payload := CurrencyPayload{
			BaseCurrency:    *s.BaseCurrency,
			TargetCurrency:  *s.TargetCurrency,
			CurrentRate:     rate,
			MinTriggerValue: *s.MinTriggerValue,
			MaxTriggerValue: *s.MaxTriggerValue,
		}

		// try 3 times to send notification
		for i := 0; i < 3; i++ {

			// try sending notification
			err = sendNotification(*s.WebhookURL, payload)
			if err != nil {

				// failure... sleep to try again if not end of loop
				fmt.Println("\tdidn't manage to notify: ", err.Error())
				if i < 3 {
					time.Sleep(time.Second * 2)
				}
			} else {

				// success! break
				fmt.Println("\tDid notify!")
				break
			}
		}
	}

	fmt.Println("\tdone trying")

}

func sendNotification(url string, payload CurrencyPayload) error {

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(payload)
	_, err := http.Post(url, "application/json; charset=utf-8", b)
	if err != nil {
		return err
	}

	return nil
}

// utility function for responding with a simple statuscode
func respWithCode(res *http.ResponseWriter, status int) {
	http.Error(*res, http.StatusText(status), status)
}
