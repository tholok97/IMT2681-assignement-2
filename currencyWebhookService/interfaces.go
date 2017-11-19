package currencyWebhookService

import "errors"

// SubscriberDB defines how a db in the system behaves
type SubscriberDB interface {
	Add(s Subscriber) (string, error)
	Remove(id string) error
	Count() (int, error)
	Get(id string) (Subscriber, error)
	GetAll() ([]Subscriber, error)
}

// CurrencyMonitor defines how something that monitors currency behaves
type CurrencyMonitor interface {
	Update(func(string) ([]byte, error)) error
	Latest(curr1, curr2 string) (float32, error)
	Average(curr1, curr2 string) (float32, error)
}

// error variables (so users of the interfaces can react to them)
var errInvalidCurrency = errors.New("currency used is not valid")
var errNotFound = errors.New("subscriber not found")
var errInvalidID = errors.New("no id given to search in mongodb")
