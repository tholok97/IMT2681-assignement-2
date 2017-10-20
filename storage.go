package main

import (
	"errors"
)

type SubscriberDB interface {
	Add(s Subscriber) (int, error)
	Count() (int, error)
	Get(id int) (Subscriber, error)
	GetAll() ([]Subscriber, error)
}

type VolitileSubscriberDB struct {
	subscribers map[int]Subscriber
	nextID      int
}

func VolitileSubscriberDBFactory() VolitileSubscriberDB {
	db := VolitileSubscriberDB{subscribers: make(map[int]Subscriber)}
	return db
}

func (db *VolitileSubscriberDB) Add(s Subscriber) (int, error) {
	db.subscribers[db.nextID] = s
	db.nextID++
	return db.nextID - 1, nil
}

func (db *VolitileSubscriberDB) Count() (int, error) {
	return len(db.subscribers), nil
}

func (db *VolitileSubscriberDB) Get(id int) (Subscriber, error) {
	var err error
	s, ok := db.subscribers[id]
	if !ok {
		err = errors.New("Subscriber does not exist!")
	}
	return s, err
}

func (db *VolitileSubscriberDB) GetAll() ([]Subscriber, error) {
	if dbCount, countErr := db.Count(); countErr != nil {
		all := make([]Subscriber, 0, dbCount)
		for _, s := range db.subscribers {
			all = append(all, s)
		}
		return all, nil
	} else {
		return nil, countErr
	}
}
