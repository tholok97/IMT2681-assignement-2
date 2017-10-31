package currencyWebhookService

import "strconv"

// VolatileSubscriberDB is a subscriber database that uses volatile memory.
// used for testing
type VolatileSubscriberDB struct {
	subscribers map[string]Subscriber
	nextID      int
}

// VolatileSubscriberDBFactory returns a fresh VolatileSubscriberDB
func VolatileSubscriberDBFactory() VolatileSubscriberDB {
	db := VolatileSubscriberDB{subscribers: make(map[string]Subscriber)}
	return db
}

// Add adds a subscriber to the db
func (db *VolatileSubscriberDB) Add(s Subscriber) (string, error) {
	db.subscribers[strconv.Itoa(db.nextID)] = s
	db.nextID++
	return strconv.Itoa(db.nextID - 1), nil
}

// Remove subscriber with id. Err if not found
func (db *VolatileSubscriberDB) Remove(id string) error {
	_, ok := db.subscribers[id]
	if ok {
		delete(db.subscribers, id)
		return nil
	}
	return errNotFound
}

// Count returns the number of subscribers in the db
func (db *VolatileSubscriberDB) Count() (int, error) {
	return len(db.subscribers), nil
}

// Get gets subscriber with id
func (db *VolatileSubscriberDB) Get(id string) (Subscriber, error) {
	var err error
	s, ok := db.subscribers[id]
	if !ok {
		err = errNotFound
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
