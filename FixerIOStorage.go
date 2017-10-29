package main

import (
	"fmt"
	"time"

	mgo "gopkg.in/mgo.v2"
)

type FixerIOStorage struct {
	DatabaseURL    string
	DatabaseName   string
	CollectionName string
}

func (fios *FixerIOStorage) Update(url string) error {

	// delete old entry
	session, err := mgo.Dial(fios.DatabaseURL)
	defer session.Close()
	if err != nil {
		return err
	}

	err = session.DB(fios.DatabaseName).DropDatabase()
	if err != nil {
		return err
	}

	// UPDATE LATEST

	// get payload
	payload, err := fetchFixerIO(url+"/latest?base=EUR", getJSON)
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
		fmt.Printf("error in Insert(): %v", err.Error())
		return err
	}

	// UPDATE AVERAGE
	average, err := generateAverage(url)
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
		fmt.Printf("error in Insert(): %v", err.Error())
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

type MongoRate struct {
	Name  string             `name`
	Rates map[string]float32 `rates`
}
