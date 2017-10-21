package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
)

// make request with given method+body, assert statuscode and return pointer
//	to response (nil if failure)
func reqTest(t *testing.T, ts *httptest.Server, method string, body io.Reader, expectedCode int) *http.Response {

	// instantiate test client
	client := &http.Client{}

	// create a request to our mock HTTP server
	req, err := http.NewRequest(method, ts.URL, body)
	if err != nil {
		t.Errorf("error constructing valid request")
	}

	// do request
	resp, err := client.Do(req)
	if err != nil {
		t.Errorf("error doing valid request")
	}

	// reach to the response from the request
	if resp.StatusCode != expectedCode {
		t.Errorf("expected statuscode %d, received %d", expectedCode, resp.StatusCode)
	}

	return resp
}

func TestSubscriberHandler_handleSubscriberRequest_POST(t *testing.T) {

	// instantiate test handler using volatile db (shouldn't fail)
	db := VolatileSubscriberDBFactory()
	handler := SubscriberHandlerFactory(&db)

	// instantiate mock HTTP server
	ts := httptest.NewServer(http.HandlerFunc(handler.handleSubscriberRequest))
	defer ts.Close()

	// fully valid
	validBody := strings.NewReader(`{
		"webhookURL": "http://remoteUrl:8080/randomWebhookPath",
		"baseCurrency": "EUR",
		"targetCurrency": "NOK",
		"minTriggerValue": 1.50, 
		"maxTriggerValue": 2.55
		}`)

	// json correct, but missing one field: invalid (TODO: doesn't work)
	invalidBody := strings.NewReader(`{
		"webhookURL": "http://remoteUrl:8080/randomWebhookPath",
		"baseCurrency": "EUR",
		"targetCurrency": "NOK",
		"maxTriggerValue": 2.55
		}`)

	// json incorrect, invalid
	veryInvalidBody := strings.NewReader(`{
		"webhookURL": "http://remoteUrl:8080/randomWebhookPath"",
		}`)

	// asssert that correct error codes are returned (store valid response)
	reqTest(t, ts, http.MethodPost, invalidBody, http.StatusBadRequest)
	reqTest(t, ts, http.MethodPost, veryInvalidBody, http.StatusBadRequest)
	resp := reqTest(t, ts, http.MethodPost, validBody, http.StatusOK)

	// test valid response:

	if resp == nil {
		t.Error("erroring in getting response from server")
	}

	// assert that body is a single int
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Error("error converting body in response to []byte")
	}

	str := string(b)
	_, err = strconv.Atoi(str)
	if err != nil {
		t.Error("body in response to subscriber request isn't an id (int)")
	}
}

func TestSubscriberHandler_handleSubscriberRequest_GET(t *testing.T) {

	// instantiate test handler using volatile db (shouldn't fail)
	db := VolatileSubscriberDBFactory()
	handler := SubscriberHandlerFactory(&db)

	// instantiate mock HTTP server
	ts := httptest.NewServer(http.HandlerFunc(handler.handleSubscriberRequest))
	defer ts.Close()

	// test ids
	validID := 1
	invalidID := 2

	// sneak stuff into the db
	testSub := Subscriber{WebhookURL: "testing"}
	db.subscribers[validID] = testSub

	// assert that request for valid id returns OK
	validIDBody := strings.NewReader(strconv.Itoa(validID))
	resp := reqTest(t, ts, http.MethodGet, validIDBody, http.StatusOK)

	// assert that request for invalid id doesn't succeed
	invalidIDBody := strings.NewReader(strconv.Itoa(invalidID))
	reqTest(t, ts, http.MethodGet, invalidIDBody, http.StatusNotFound)

	// assert that malformed request returns bad request
	malformedIDBody := strings.NewReader("THIS IS NOT AN ID xD")
	reqTest(t, ts, http.MethodGet, malformedIDBody, http.StatusBadRequest)

	// test body of response from valid request:

	if resp == nil {
		t.Error("error getting response from server")
	}

	// attempt to unmarshall
	var s Subscriber
	err := json.NewDecoder(resp.Body).Decode(&s)
	if err != nil {
		t.Error("error while unmarshalling response:", err.Error())
	}

	// assert that it contains our test data
	if s.WebhookURL != testSub.WebhookURL {
		t.Error("returned wrong subscriber from get request")
	}

}
