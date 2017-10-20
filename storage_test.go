package main

import "testing"

func generateTestSubscriber(url string) Subscriber {
	return Subscriber{
		webhookURL:      url,
		baseCurrency:    "TEST",
		targetCurrency:  "TEST",
		minTriggerValue: 923,
		maxTriggerValue: 923,
	}
}

// tests the add functionality of a factory-new implementor of subscriberDB
func DBAddTest(t *testing.T, db SubscriberDB) {

	var count int // holds count in db
	var err error // holds potential error

	// generate the test-subscribers
	s1 := generateTestSubscriber("1")
	s2 := generateTestSubscriber("2")

	// Assert that the db is empty
	count, err = db.Count()
	if err != nil {
		t.Error(err)
	}
	if count != 0 {
		t.Error("Factory-new db is not empty")
	}

	// insert into db
	_, err = db.Add(s1)
	_, err = db.Add(s2)

	// assert nothing went wrong
	if err != nil {
		t.Error(err)
	}

	// assert that count after insert is 2
	count, err = db.Count()
	if err != nil {
		t.Error(err)
	}
	if count != 2 {
		t.Error("db count is not 2 after adding 2 subscribers")
	}
}

// test the count functionality of a factory-new implementor of SubscriberDB
func DBCountTest(t *testing.T, db SubscriberDB) {

}

func DBGetTest(t *testing.T, db SubscriberDB) {

}

func DBGetAllTest(t *testing.T, db SubscriberDB) {

}

// test all the features of a TestVolitileSubscriberDB
func TestVolitileSubscriberDB(t *testing.T) {

	// make a new db (This function should always work -> not tested)
	db := VolitileSubscriberDBFactory()

	// Test each of the functionalities
	DBAddTest(t, &db)
	DBCountTest(t, &db)
	DBGetTest(t, &db)
	DBGetAllTest(t, &db)
}
