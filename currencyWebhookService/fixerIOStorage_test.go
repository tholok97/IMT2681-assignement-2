package currencyWebhookService

import (
	"errors"
	"testing"

	mgo "gopkg.in/mgo.v2"
)

func TestFixerIOUpdate(t *testing.T) {

	payload := FixerIOPayload{
		Base: "EUR", Date: "2014-03-04",
		Rates: map[string]float32{"NOK": 9.4, "USD": 1.4},
	}

	faker := faker{payload: payload, err: errors.New("test error")}

	// prepare db
	fios := FixerIOStorage{
		DatabaseURL:    "jlskdjfdksj::::",
		DatabaseName:   "test",
		CollectionName: "fiosTest",
		FixerIOURL:     "",
	}

	var err error

	err = fios.Update(faker.fakeGetter)
	if err == nil {
		t.Error("expected error with bad DatabaseURL")
		return
	}

	fios.DatabaseURL = "mongodb://localhost"
	err = fios.Update(faker.fakeGetter)
	if err == nil {
		t.Error("expected error with no fixerio URL")
		return
	}

	faker.err = nil
	err = fios.Update(faker.fakeGetter)
	if err != nil {
		t.Error("didn't expect error. err: ", err.Error())
		return
	}
}

func TestFixerIOGetLatest(t *testing.T) {

	// prepare db
	fios := FixerIOStorage{
		DatabaseURL:    "mongodb://localhost",
		DatabaseName:   "test",
		CollectionName: "currencytest",
		FixerIOURL:     "",
	}

	// prepare rates
	rates := map[string]float32{
		"EUR": 1,
		"NOK": 2,
	}

	mRates := MongoRate{Name: "latest", Rates: rates}

	// open session
	session, err := mgo.Dial(fios.DatabaseURL)
	if err != nil {
		t.Error(err.Error())
		return
	}
	defer session.Close()

	// insert rates
	session, err = mgo.Dial(fios.DatabaseURL)
	if err != nil {
		t.Error(err.Error())
		return
	}
	defer session.Close()

	err = session.DB(fios.DatabaseName).C(fios.CollectionName).Insert(mRates)
	if err != nil {
		t.Error(err.Error())
		return
	}

	// assert latest
	rate, err := fios.Latest("EUR", "NOK")
	if err != nil {
		t.Error(err.Error())
		return
	}

	// assert correct rate returned
	if rate != 2 {
		t.Errorf("wrong rate returned, got %v, expected %v", rate, 2)
		return
	}
}

func TestFixerIOBadDB(t *testing.T) {

	// prepare db
	fios := FixerIOStorage{
		DatabaseURL:    "",
		DatabaseName:   "",
		CollectionName: "",
		FixerIOURL:     "",
	}

	var err error

	// test latest
	_, err = fios.Latest("", "")
	if err == nil {
		t.Error("expected latest to fail because empty dbname, it didn't")
		return
	}

	// test average
	_, err = fios.Average("", "")
	if err == nil {
		t.Error("expected average to fail because empty dbname, it didn't")
		return
	}

}
