package currencyWebhookService

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"
)

// returns a struct representing the payload returned from fixerIO with the given url
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

// GetENV gets environment variable. If something goes wrong: PANIC
func GetENV(name string) string {
	ret := os.Getenv(name)
	if ret == "" {
		panic("Missing env variable: " + name)
	}
	fmt.Println("Read env ", name, " = ", ret)
	return ret
}

// GetIntENV gets environment variable as int. If something goes wrong: PANIC
func GetIntENV(name string) int {
	ret := GetENV(name)
	num, err := strconv.Atoi(ret)
	if err != nil {
		panic("Error converting env to int: " + err.Error())
	}
	return num
}

// DurUntilClock calculates duration until next HH:MM:SS
func DurUntilClock(now time.Time, hour, minute, second int) time.Duration {

	// the time this HH:MM:SS is happening
	when := time.Date(now.Year(), now.Month(), now.Day(), hour,
		minute, second, 0, now.Location())

	// d is the time until next such time
	d := when.Sub(now)

	// if duration is negative, add a day
	if d < 0 {
		when = when.Add(24 * time.Hour)
		d = when.Sub(now)
	}

	return d
}

// DurUntilTime calculate duration until time is when
func DurUntilTime(now, when time.Time) time.Duration {
	return when.Sub(now)
}
