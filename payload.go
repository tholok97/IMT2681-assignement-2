package main

// Subscriber holds info about a subscriber to currency-info in the system
type Subscriber struct {
	WebhookURL      *string  `json:"webhookURL"`
	BaseCurrency    *string  `json:"baseCurrency"`
	TargetCurrency  *string  `json:"targetCurrency"`
	MinTriggerValue *float32 `json:"minTriggerValue"`
	MaxTriggerValue *float32 `json:"maxTriggerValue"`
}

// CurrencyPayload holds info sent to subscriber
type CurrencyPayload struct {
	BaseCurrency    string  `json:"baseCurrency"`
	TargetCurrency  string  `json:"targetCurrency"`
	CurrentRate     float32 `json:"currentRate"`
	MinTriggerValue float32 `json:"minTriggerValue"`
	MaxTriggerValue float32 `json:"maxTriggerValue"`
}

// CurrencyRequest holds info received from the user during basic currency
// requests
type CurrencyRequest struct {
	BaseCurrency   *string `json:"baseCurrency"`
	TargetCurrency *string `json:"targetCurrency"`
}

// a subscriber is valid if no fields are set to nil
func validateSubscriber(s Subscriber) bool {
	if s.BaseCurrency == nil ||
		s.MaxTriggerValue == nil ||
		s.MinTriggerValue == nil ||
		s.TargetCurrency == nil ||
		s.WebhookURL == nil {

		return false
	}
	return true
}
