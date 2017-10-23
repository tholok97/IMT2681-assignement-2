package main

import (
	"errors"
)

// error(s) used for this type:
var errDoesNotExist = errors.New("subscriber does not exist")

// VolatileSubscriberDB is a subscriber database that uses volatile memory.
// used for testing
type VolatileSubscriberDB struct {
	subscribers map[int]Subscriber
	nextID      int
}

// VolatileSubscriberDBFactory returns a fresh VolatileSubscriberDB
func VolatileSubscriberDBFactory() VolatileSubscriberDB {
	db := VolatileSubscriberDB{subscribers: make(map[int]Subscriber)}
	return db
}

// Add adds a subscriber to the db
func (db *VolatileSubscriberDB) Add(s Subscriber) (int, error) {
	db.subscribers[db.nextID] = s
	db.nextID++
	return db.nextID - 1, nil
}

// Add adds a subscriber to the db
func (db *VolatileSubscriberDB) Remove(id int) error {
	return errors.New("NOT IMPLEMENTED")
}

// Count returns the number of subscribers in the db
func (db *VolatileSubscriberDB) Count() (int, error) {
	return len(db.subscribers), nil
}

// Get gets subscriber with id
func (db *VolatileSubscriberDB) Get(id int) (Subscriber, error) {
	var err error
	s, ok := db.subscribers[id]
	if !ok {
		err = errDoesNotExist
	}
	return s, err
}

// GetAll gets all subscribers as slice
func (db *VolatileSubscriberDB) GetAll() ([]Subscriber, error) {
	all := make([]Subscriber, 0, len(db.subscribers))
	for _, s := range db.subscribers {
		all = append(all, s)
	}
	return all, nil
}
