package main

import (
	"strconv"
	"testing"
)

func generateTestSubscriber(url string) Subscriber {
	testText := "TEST"
	triggerVal := float32(923)
	return Subscriber{
		WebhookURL:      &url,
		BaseCurrency:    &testText,
		TargetCurrency:  &testText,
		MinTriggerValue: &triggerVal,
		MaxTriggerValue: &triggerVal,
	}
}

func TestVolatileSubscriberDBFactory(t *testing.T) {

	// make a new db
	db := VolatileSubscriberDBFactory()

	// assert that fresh db is empty
	if len(db.subscribers) != 0 {
		t.Error("factory-fresh db not empty")
	}
}

func TestVolatileSubscriberDB_Add(t *testing.T) {

	// make a new db
	db := VolatileSubscriberDBFactory()

	// test subscribers
	s1 := generateTestSubscriber("1")
	s2 := generateTestSubscriber("2")

	// add test subscribers
	id1, err1 := db.Add(s1)
	id2, err2 := db.Add(s2)

	if err1 != nil {
		t.Error(err1)
	} else if err2 != nil {
		t.Error(err2)
	}

	// assert db size
	if len(db.subscribers) != 2 {
		t.Error("db-size isn't 2 after adding two things")
	}

	// assert that the correct elements were inserted
	if db.subscribers[id1] != s1 || db.subscribers[id2] != s2 {
		t.Error("id returned from add not correct")
		t.Error("id returned from add not correct")
	}
}

func TestVolatileSubscriberDB_Remove(t *testing.T) {

	// make a new db
	db := VolatileSubscriberDBFactory()

	// test subscriber
	s1 := generateTestSubscriber("1")

	// test id's
	id1 := "2" // VALID
	id2 := "3" // INVALID

	// add test subscriber to db
	db.subscribers[id1] = s1

	// assert that deletion of valid id is possible
	err := db.Remove(id1)
	if err != nil {
		t.Error("removal of valid subscriber not possible. error: ", err.Error())
		return
	}

	// asssert that subscriber was actually deleted
	_, ok := db.subscribers[id1]
	if ok {
		t.Error("remove returned no error not subscriber was not actually removed")
		return
	}

	// assert that removal with invalid id fails (properly)
	err = db.Remove(id2)
	if err != errNotFound {
		t.Errorf("removal with invalid id didn't return the correct error. Got: %s, wanted: %s", err.Error(), errNotFound.Error())
	} else if err == nil {
		t.Errorf("removal with invalid id didn't return an error")
	}
}

func TestVolatileSubscriberDB_Count(t *testing.T) {

	// make a new db
	db := VolatileSubscriberDBFactory()

	var count int
	var err error

	// assert that empty db returns Count() = 0
	count, err = db.Count()
	if err != nil {
		t.Error(err)
	} else if count != 0 {
		t.Error("factory-fresh db not returning correct count (0)")
	}

	// add a few subscribers
	db.subscribers["1"] = generateTestSubscriber("NOT RELEVANT")
	db.subscribers["2"] = generateTestSubscriber("NOT RELEVANT")

	// assert that Count() is now 2
	count, err = db.Count()
	if err != nil {
		t.Error(err)
	} else if count != 2 {
		t.Error("count not 2 after adding 2 subscribers")
	}
}

func TestVolatileSubscriberDB_Get(t *testing.T) {

	var s Subscriber
	var err error

	// make a new db
	db := VolatileSubscriberDBFactory()

	// test subscriber
	s1 := generateTestSubscriber("1")

	// test id's
	id1 := "2" // VALID
	id2 := "3" // INVALID

	// add test subscriber to db
	db.subscribers[id1] = s1

	// assert that correct subscriber is returned from valid id
	s, err = db.Get(id1)
	if err != nil {
		t.Error(err)
	} else if s != s1 {
		t.Error("wrong subscriber returned from get")
	}

	// assert that error is returned for invalid id
	_, err = db.Get(id2)
	if err != errNotFound {
		t.Errorf("wrong kind of error returned for wrong id. "+
			"Got: %s, wanted: %s", err.Error(), errNotFound.Error())
	} else if err == nil {
		t.Error("error not returned for invalid id")
	}
}

func TestVolatileSubscriberDB_GetAll(t *testing.T) {

	// make a new db
	db := VolatileSubscriberDBFactory()

	// make test subscribers
	testSubs := []Subscriber{
		generateTestSubscriber("1"),
		generateTestSubscriber("2"),
	}

	// add test subscribers to db
	for i, s := range testSubs {
		db.subscribers[strconv.Itoa(i)] = s
	}

	// getall, and check for errors
	all, err := db.GetAll()
	if err != nil {
		t.Error(err)
	}

	// assert that the returned slice is equal to the test slice
	if !(testSubs[0] == all[0] && testSubs[1] == all[1]) &&
		!(testSubs[0] == all[1] && testSubs[1] == all[0]) {
		t.Error("subs returned from getAll() not equal to test subs")
	}
}
