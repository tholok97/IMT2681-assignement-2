package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func fetchFixerIO(url string, jsonGetter func(url string) ([]byte, error)) (FixerIOPayload, error) {

	// get the json as bytes
	b, err := jsonGetter(url)
	if err != nil {
		return FixerIOPayload{}, err
	}

	// convert to struct
	var payload FixerIOPayload
	err = json.Unmarshal(b, &payload)

	return payload, err
}

// FixerIOPayload contains response from FixerIO
type FixerIOPayload struct {
	Base  string             `base`
	Date  string             `date`
	Rates map[string]float32 `rates`
}

// gets json as []byte from a given url
func getJSON(url string) ([]byte, error) {

	// (try to) get response from url
	resp, getErr := http.Get(url)
	if getErr != nil || resp.StatusCode != http.StatusOK {
		return nil, getErr
	}

	// we need to close the body when we're done. defer it
	defer resp.Body.Close()

	// (try to) read []byte from response
	bodyBytes, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		return nil, readErr
	}

	return bodyBytes, nil
}
