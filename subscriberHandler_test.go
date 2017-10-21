package main

import (
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
)

// do request with given body. 'expected' tells function if we're expecting it
//  to fail or not
func doSubscriberRequest(t *testing.T, ts *httptest.Server, body io.Reader, expected bool) {

	// instantiate test client
	client := &http.Client{}

	// create a request to our mock HTTP server
	req, err := http.NewRequest(http.MethodPost, ts.URL, body)
	if err != nil {
		t.Errorf("error constructing valid request")
	}

	// do request
	resp, err := client.Do(req)
	if err != nil {
		t.Errorf("error doing valid request")
	}

	// react to the response from the request
	switch expected {
	case true:

		// assert that request returned OK
		if resp.StatusCode != http.StatusOK {
			t.Errorf("expected statuscode ok, received %d", resp.StatusCode)
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

	case false:

		// assert that request didn't return OK
		if resp.StatusCode == http.StatusOK {
			t.Errorf("expected to fail, received %d", resp.StatusCode)
		}
	}
}

func TestSubscriberHandler_handleSubscriberRequest(t *testing.T) {

	// instantiate test handler using volatile db (shouldn't fail)
	db := VolatileSubscriberDBFactory()
	handler := SubscriberHandlerFactory(&db)

	// instantiate mock HTTP server
	ts := httptest.NewServer(http.HandlerFunc(handler.handleSubscriberRequest))
	defer ts.Close()

	// INSTANTIATE REQUEST BODIES

	// fully valid
	validBody := strings.NewReader(`{
		"webhookURL": "http://remoteUrl:8080/randomWebhookPath",
		"baseCurrency": "EUR",
		"targetCurrency": "NOK",
		"minTriggerValue": 1.50, 
		"maxTriggerValue": 2.55
		}`)

	/* THIS ONE DOESN'T WORK YET. TODO [1]
	// json correct, but missing one field: invalid
	invalidBody := strings.NewReader(`{
		"webhookURL": "http://remoteUrl:8080/randomWebhookPath",
		"baseCurrency": "EUR",
		"targetCurrency": "NOK",
		"maxTriggerValue": 2.55
		}`)
	*/

	// json incorrect, invalid
	veryInvalidBody := strings.NewReader(`{
		"webhookURL": "http://remoteUrl:8080/randomWebhookPath"",
		}`)

	doSubscriberRequest(t, ts, validBody, true)
	//doSubscriberRequest(t, ts, invalidBody, false) TODO: DOESN'T WORK YET [1]
	doSubscriberRequest(t, ts, veryInvalidBody, false)
}
