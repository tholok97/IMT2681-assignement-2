package currencyWebhookService

import (
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// SubscriberMongoDB implements SubscriberDB using mongodb (persistant)
type SubscriberMongoDB struct {
	URL                      string
	Name                     string
	SubscriberCollectionName string
}

// SubscriberMongoDBFactory returns a fresh SubscriberMongoDB
func SubscriberMongoDBFactory(url, name string) (SubscriberMongoDB, error) {

	db := SubscriberMongoDB{URL: url, Name: name, SubscriberCollectionName: "subscribers"}

	// assert that session is possible
	session, err := mgo.Dial(db.URL)
	if err != nil {
		return SubscriberMongoDB{}, nil
	}
	defer session.Close()

	return db, nil
}

// Add adds a subscriber to the db
func (db *SubscriberMongoDB) Add(s Subscriber) (string, error) {

	session, err := mgo.Dial(db.URL)
	if err != nil {
		return "", err
	}
	defer session.Close()

	// generate new ID and insert
	s.ID = bson.NewObjectId()
	err = session.DB(db.Name).C(db.SubscriberCollectionName).Insert(s)
	if err != nil {
		return "", err
	}

	return s.ID.Hex(), nil
}

// Remove subscriber with id. Err if not found
func (db *SubscriberMongoDB) Remove(id string) error {

	session, err := mgo.Dial(db.URL)
	if err != nil {
		return err
	}
	defer session.Close()

	// if no id, fail
	if id == "" {
		return errNoId
	}

	// (id in hex -> convert)
	err = session.DB(db.Name).C(db.SubscriberCollectionName).RemoveId(bson.ObjectIdHex(id))
	if err != nil {
		return err
	}

	return nil
}

// Count returns the number of subscribers in the db
func (db *SubscriberMongoDB) Count() (int, error) {

	session, err := mgo.Dial(db.URL)
	if err != nil {
		return 923, err
	}
	defer session.Close()

	count, err := session.DB(db.Name).C(db.SubscriberCollectionName).Count()
	if err != nil {
		return 923, err
	}

	return count, nil
}

// Get gets subscriber with id
func (db *SubscriberMongoDB) Get(id string) (Subscriber, error) {

	session, err := mgo.Dial(db.URL)
	if err != nil {
		return Subscriber{}, err
	}
	defer session.Close()

	// if no id, fail
	if id == "" {
		return Subscriber{}, errNoId
	}

	// (id is in hex -> convert)
	var sub Subscriber
	err = session.DB(db.Name).C(db.SubscriberCollectionName).FindId(bson.ObjectIdHex(id)).One(&sub)
	if err != nil {
		return Subscriber{}, err
	}

	return sub, nil
}

// GetAll gets all subscribers as slice
func (db *SubscriberMongoDB) GetAll() ([]Subscriber, error) {

	session, err := mgo.Dial(db.URL)
	if err != nil {
		return nil, err
	}
	defer session.Close()

	subs := make([]Subscriber, 0)
	err = session.DB(db.Name).C(db.SubscriberCollectionName).Find(nil).All(&subs)
	if err != nil {
		return nil, err
	}

	return subs, nil
}
