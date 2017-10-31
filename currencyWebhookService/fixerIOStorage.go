package currencyWebhookService

/*
 * WARNING: This file is pretty ad-hoc, as the deadline is approaching fast,
	and I need this feature for the project to work
*/

import (
	"time"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// FixerIOStorage is an implementation of CurrencyMonitor that stores it's
// currencies in mongodb and gets them from fixer.io
type FixerIOStorage struct {
	DatabaseURL    string
	DatabaseName   string
	CollectionName string
	FixerIOURL     string
}

// Update the mongodb database by fetching data from url
func (fios *FixerIOStorage) Update() error {

	// delete old entry
	session, err := mgo.Dial(fios.DatabaseURL)
	if err != nil {
		return err
	}
	defer session.Close()

	session.DB(fios.DatabaseName).C(fios.CollectionName).DropCollection()

	// UPDATE LATEST

	// get payload
	payload, err := fetchFixerIO(fios.FixerIOURL+"/latest?base=EUR", getJSON)
	if err != nil {
		return err
	}

	// fetch info we want
	rates := payload.Rates
	rates["EUR"] = 1

	mRates := MongoRate{Name: "latest", Rates: rates}

	// store it
	session, err = mgo.Dial(fios.DatabaseURL)
	if err != nil {
		return err
	}
	defer session.Close()

	err = session.DB(fios.DatabaseName).C(fios.CollectionName).Insert(mRates)
	if err != nil {
		return err
	}

	// UPDATE AVERAGE
	average, err := generateAverage(fios.FixerIOURL)
	if err != nil {
		return err
	}

	mRates = MongoRate{Name: "average", Rates: average}

	// store it
	session, err = mgo.Dial(fios.DatabaseURL)
	if err != nil {
		return err
	}
	defer session.Close()

	err = session.DB(fios.DatabaseName).C(fios.CollectionName).Insert(mRates)
	if err != nil {
		return err
	}

	return nil
}

func generateAverage(url string) (map[string]float32, error) {

	payload, err := fetchFixerIO(url+"/latest?base=EUR", getJSON)
	if err != nil {
		return nil, err
	}
	temp := payload.Rates

	t := time.Now()

	// get for last 7 days
	for i := 0; i < 6; i++ {
		t = t.AddDate(0, 0, -1)
		date := t.Format("2006-01-02")

		payload, err := fetchFixerIO(url+"/"+date+"?base=EUR", getJSON)
		if err != nil {
			return nil, err
		}

		rateForDay := payload.Rates
		for k, v := range rateForDay {
			temp[k] += v
		}
	}

	// do average

	for k := range temp {
		temp[k] /= 7
	}

	// ..add EUR (=1)
	temp["EUR"] = 1

	return temp, nil
}

func (fios *FixerIOStorage) getRate(curr1, curr2, name string) (float32, error) {

	// fetch rates from mongodb
	session, err := mgo.Dial(fios.DatabaseURL)
	if err != nil {
		return 923, err
	}
	defer session.Close()

	var mrate MongoRate
	err = session.DB(fios.DatabaseName).C(fios.CollectionName).Find(bson.M{"name": name}).One(&mrate)
	if err != nil {
		return 923, err
	}

	// calculate rate
	rate1, ok1 := mrate.Rates[curr1]
	rate2, ok2 := mrate.Rates[curr2]
	if !ok1 || !ok2 {
		return 923, errInvalidCurrency
	}

	return rate2 / rate1, nil
}

// Latest rate beteween curr1 and curr2
func (fios *FixerIOStorage) Latest(curr1, curr2 string) (float32, error) {

	rate, err := fios.getRate(curr1, curr2, "latest")
	if err != nil {
		return 923, err
	}

	return rate, nil
}

// Average rate between curr1 and curr2
func (fios *FixerIOStorage) Average(curr1, curr2 string) (float32, error) {

	rate, err := fios.getRate(curr1, curr2, "average")
	if err != nil {
		return 923, err
	}

	return rate, nil
}

// MongoRate represents how how rates are stored in mongodb
type MongoRate struct {
	Name  string             `json:"name"`
	Rates map[string]float32 `json:"rates"`
}
