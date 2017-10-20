package main

// Subscriber holds info about a subscriber to currency-info in the system
type Subscriber struct {
	WebhookURL      string `json:"webhookURL"`
	BaseCurrency    string `json:"baseCurrency"`
	TargetCurrency  string `json:"targetCurrency"`
	MinTriggerValue int    `json:"minTriggerValue"`
	MaxTriggerValue int    `json:"maxTriggerValue"`
}

// CurrencyPayload holds info sent to subscriber
type CurrencyPayload struct {
	BaseCurrency    string `json:"baseCurrency"`
	TargetCurrency  string `json:"targetCurrency"`
	CurrentRate     int    `json:"currentRate"`
	MinTriggerValue int    `json:"minTriggerValue"`
	MaxTriggerValue int    `json:"maxTriggerValue"`
}

// CurrencyRequest holds info received from the user during basic currency
// requests
type CurrencyRequest struct {
	BaseCurrency   string `json:"baseCurrency"`
	TargetCurrency string `json:"targetCurrency"`
}
