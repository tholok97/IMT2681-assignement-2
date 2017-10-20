package main

type Subscriber struct {
	webhookURL      string `json:"webhookURL"`
	baseCurrency    string `json:"baseCurrency"`
	targetCurrency  string `json:"targetCurrency"`
	minTriggerValue int    `json:"minTriggerValue"`
	maxTriggerValue int    `json:"maxTriggerValue"`
}

type CurrencyPayload struct {
	baseCurrency    string `json:"baseCurrency"`
	targetCurrency  string `json:"targetCurrency"`
	currentRate     int    `json:"currentRate"`
	minTriggerValue int    `json:"minTriggerValue"`
	maxTriggerValue int    `json:"maxTriggerValue"`
}

type CurrencyRequest struct {
	baseCurrency   string `json:"baseCurrency"`
	targetCurrency string `json:"targetCurrency"`
}
