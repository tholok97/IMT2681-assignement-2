package currencyWebhookService

import "gopkg.in/mgo.v2/bson"

// Subscriber holds info about a subscriber to currency-info in the system
// (fields are pointers to make validation of incomming requests trivial)
type Subscriber struct {
	ID              bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	WebhookURL      *string       `json:"webhookURL"`
	BaseCurrency    *string       `json:"baseCurrency"`
	TargetCurrency  *string       `json:"targetCurrency"`
	MinTriggerValue *float32      `json:"minTriggerValue"`
	MaxTriggerValue *float32      `json:"maxTriggerValue"`
}

// DialogResponse contains the response to dialogflow
type DialogResponse struct {
	DisplayText string `json:"displayText"`
	Speech      string `json:"speech"`
}

// DialogRequest contains the request from dialogflow
type DialogRequest struct {
	Results ResultField `json:"result"`
}

// ResultField contains the ResultField from the dialogflow request
type ResultField struct {
	Parameters ParameterField `json:"parameters"`
}

// ParameterField contains the ParameterField in the ResultField from the dialogflow request
type ParameterField struct {
	BaseCurrency   string `json:"baseCurrency"`
	TargetCurrency string `json:"targetCurrency"`
}

// CurrencyPayload holds info sent to subscriber
type CurrencyPayload struct {
	BaseCurrency    string  `json:"baseCurrency"`
	TargetCurrency  string  `json:"targetCurrency"`
	CurrentRate     float32 `json:"currentRate"`
	MinTriggerValue float32 `json:"minTriggerValue"`
	MaxTriggerValue float32 `json:"maxTriggerValue"`
}

// FixerIOPayload contains response from FixerIO
type FixerIOPayload struct {
	Base  string             `json:"base"`
	Date  string             `json:"date"`
	Rates map[string]float32 `json:"rates"`
}

// CurrencyRequest holds info received from the user during basic currency
// requests
// (fields are pointers to make validation of incomming requests trivial)
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

// a currency request is valid if no fields are set to nil
func validateCurrencyRequest(cr CurrencyRequest) bool {
	if cr.BaseCurrency == nil ||
		cr.TargetCurrency == nil {

		return false
	}
	return true
}
