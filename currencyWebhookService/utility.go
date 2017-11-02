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

// FixerIOPayload contains response from FixerIO
type FixerIOPayload struct {
	Base  string             `json:"base"`
	Date  string             `json:"date"`
	Rates map[string]float32 `json:"rates"`
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
func DurUntilClock(hour, minute, second int) time.Duration {
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

// DurUntilTime calculate duration until time is when
func DurUntilTime(when time.Time) time.Duration {
	return when.Sub(time.Now())
}
