package main

import (
	"fmt"

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

	return map[string]float32{"EUR": 923}, nil
}

type MongoRate struct {
	Name  string             `name`
	Rates map[string]float32 `rates`
}
