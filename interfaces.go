package main

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
	Latest(curr1, curr2 string) (float32 error)
	Average(curr1, curr2 string, days int) (float32 error)
}
