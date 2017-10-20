package main

type SubscriberStorage interface {
	Add(s Subscriber)
	Count() int
	Get(id int) (Subscriber, error)
	GetAll() []Subscriber
}
