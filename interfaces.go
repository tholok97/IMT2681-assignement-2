package main

import "errors"

// SubscriberDB defines how a db in the system behaves
type SubscriberDB interface {
	Add(s Subscriber) (int, error)
	Remove(id int) error
	Count() (int, error)
	Get(id int) (Subscriber, error)
	GetAll() ([]Subscriber, error)
}

// CurrencyMontitor defines how something that monitors currency behaves
type CurrencyMonitor interface {
	Update(currencyAPIURL string) error
	Latest(curr1, curr2 string) (float32, error)
	Average(curr1, curr2 string) (float32, error)
}

// error variables (so users of the interfaces can react to them)
var errInvalidCurrency = errors.New("currency used is not valid")
var errNotFound = errors.New("subscriber not found")
