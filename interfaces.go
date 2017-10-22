package main

// SubscriberDB defines how a db in the system behaves
type SubscriberDB interface {
	Add(s Subscriber) (int, error)
	Count() (int, error)
	Get(id int) (Subscriber, error)
	GetAll() ([]Subscriber, error)
}
