package currencyWebhookService

import (
	"testing"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// NOTE: Database test-code heavily inspired by https://github.com/marni/imt2681_cloud/blob/master/studentdb/database_test.go

func setupDB(t *testing.T) *SubscriberMongoDB {

	db, err := SubscriberMongoDBFactory("mongodb://localhost", "subscriberTestDB")
	if err != nil {
		t.Error("couldn't setup db: " + err.Error())
	}

	return &db
}

func tearDownDB(t *testing.T, db *SubscriberMongoDB) {
	session, err := mgo.Dial(db.URL)
	if err != nil {
		t.Error(err)
	}
	defer session.Close()

	err = session.DB(db.Name).DropDatabase()
	if err != nil {
		t.Error(err)
	}
}

func TestMongoFactory(t *testing.T) {
	db, err := SubscriberMongoDBFactory("mongodb://localhost", "subscriberTestDB")
	if err != nil {
		t.Error("factory failed: " + err.Error())
	}

	count, err := db.Count()
	if err != nil {
		t.Error("db error: " + err.Error())
	}
	if count != 0 {
		t.Error("factory fresh db not empty")
	}
}

func TestMongoAdd(t *testing.T) {
	db := setupDB(t)
	defer tearDownDB(t, db)

	// test valid subscriber
	url := "notarealurl"
	s := Subscriber{WebhookURL: &url}
	id, err := db.Add(s)
	if err != nil {
		t.Error("db error: " + err.Error())
		return
	}

	if !bson.IsObjectIdHex(id) {
		t.Error("returned id is not hex")
		return
	}

	count, err := db.Count()
	if err != nil {
		t.Error("db error: ", err.Error())
		return
	}
	if count != 1 {
		t.Error("count after one add sin't 1")
		return
	}

}

func TestMongoGet(t *testing.T) {
	db := setupDB(t)
	defer tearDownDB(t, db)

	// test invalid id
	_, err := db.Get("notanid")
	if err == nil {
		t.Error("db failed to error for invalid id in get")
		return
	}

	// test valid id

	// (insert subscriber)
	url := "notarealurl"
	s := Subscriber{WebhookURL: &url}
	id, err := db.Add(s)
	if err != nil {
		t.Error("db error: " + err.Error())
		return
	}

	// (assert get)
	sub, err := db.Get(id)
	if err != nil {
		t.Error("error returned from valid get: " + err.Error())
	}

	if *sub.WebhookURL != url {
		t.Error("get returned wrong subscriber")
	}
}

func TestMongoRemove(t *testing.T) {
	db := setupDB(t)
	defer tearDownDB(t, db)

	// test invalid id
	err := db.Remove("notanid")
	if err == nil {
		t.Error("db failed to error for invalid id in remove")
		return
	}

	// test valid id

	// (insert subscriber)
	url := "notarealurl"
	s := Subscriber{WebhookURL: &url}
	id, err := db.Add(s)
	if err != nil {
		t.Error("db error: " + err.Error())
		return
	}

	// (assert remove)
	err = db.Remove(id)
	if err != nil {
		t.Error("error returned from valid remove: " + err.Error())
	}
}
