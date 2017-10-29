package main

// StubCurrencyMonitor fakes the behavior expected by a CurrencyMonitor. Used
//	for testing. The fields are used to determine the outcome of the methods.
type StubCurrencyMonitor struct {
	err        error
	nextVal    float32
	validCurrs []string
}

// generate new currency monitor stub
func StubCurrencyMonitorFactory(err error, nextVal float32) StubCurrencyMonitor {
	monitor := StubCurrencyMonitor{err: err, nextVal: nextVal}
	monitor.validCurrs = []string{
		"AUD", "BGN", "BRL", "CAD", "CHF", "CNY", "CZK", "DKK", "GBP", "HKD",
		"HRK", "HUF", "IDR", "ILS", "INR", "JPY", "KRW", "MXN", "MYR", "NOK",
		"NZD", "PHP", "PLN", "RON", "RUB", "SEK", "SGD", "THB", "TRY", "USD",
		"ZAR", "EUR",
	}
	return monitor
}

// Update the dataset (faked)
func (monitor *StubCurrencyMonitor) Update(currencyAPIURL string) error {
	return monitor.err
}

// Latest currency (faked)
func (monitor *StubCurrencyMonitor) Latest(curr1, curr2 string) (float32,
	error) {

	if !isIn(curr1, monitor.validCurrs) || !isIn(curr2, monitor.validCurrs) {
		return 923, errInvalidCurrency
	}
	return monitor.nextVal, nil
}

// Average currency (faked)
func (monitor *StubCurrencyMonitor) Average(curr1, curr2 string) (float32, error) {

	if !isIn(curr1, monitor.validCurrs) || !isIn(curr2, monitor.validCurrs) {
		return 923, errInvalidCurrency
	}
	return monitor.nextVal, nil
}

// is str in slice ?
func isIn(str string, slice []string) bool {
	for _, v := range slice {
		if str == v {
			return true
		}
	}
	return false
}
