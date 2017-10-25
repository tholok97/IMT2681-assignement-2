package main

import (
	"fmt"

	mgo "gopkg.in/mgo.v2"
)

type SubscriberMongoDB struct {
	URL                      string
	Name                     string
	SubscriberCollectionName string
}

// SubscriberMongoDB returns a fresh SubscriberMongoDB
func SubscriberMongoDBFactory(url, name, collectionName string) (SubscriberMongoDB, error) {

	db := SubscriberMongoDB{URL: url, Name: name, SubscriberCollectionName: collectionName}

	session, err := mgo.Dial(db.URL)
	if err != nil {
		return SubscriberMongoDB{}, nil
	}
	defer session.Close()

	index := mgo.Index{
		Key:        []string{"webhookurl"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}

	err = session.DB(db.Name).C(db.SubscriberCollectionName).EnsureIndex(index)
	if err != nil {
		return SubscriberMongoDB{}, err
	}

	return db, nil
}

// Add adds a subscriber to the db
func (db *SubscriberMongoDB) Add(s Subscriber) (int, error) {
	session, err := mgo.Dial(db.URL)
	if err != nil {
		return -1, err
	}
	defer session.Close()

	err = session.DB(db.Name).C(db.SubscriberCollectionName).Insert(s)
	if err != nil {
		return -1, err
	}
	return 923, nil
}

// Remove subscriber with id. Err if not found
func (db *SubscriberMongoDB) Remove(id int) error {
	session, err := mgo.Dial(db.URL)
	if err != nil {
		return err
	}
	defer session.Close()

	return nil
}

// Count returns the number of subscribers in the db
func (db *SubscriberMongoDB) Count() (int, error) {
	return 923, nil
}

// Get gets subscriber with id
func (db *SubscriberMongoDB) Get(id int) (Subscriber, error) {

	session, err := mgo.Dial(db.URL)
	if err != nil {
		return Subscriber{}, err
	}
	defer session.Close()

	var sub Subscriber
	err = session.DB(db.Name).C(db.SubscriberCollectionName).FindId(id).One(&sub)
	if err != nil {
		fmt.Println(err.Error())
		return Subscriber{}, err
	}
	return sub, nil
}

// GetAll gets all subscribers as slice
func (db *SubscriberMongoDB) GetAll() ([]Subscriber, error) {
	return nil, nil
}
