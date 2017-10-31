package currencyWebhookService

import "testing"
import "encoding/json"
import "reflect"

type faker struct {
	payload FixerIOPayload
	err     error
}

func (f *faker) fakeGetter(url string) ([]byte, error) {

	b, err := json.Marshal(f.payload)

	if err != nil {
		panic("internal test fail: " + err.Error())
	}

	return b, f.err
}

// Test the fetcher (veeeery ad-hoc code. Just needed to make sure it works in best-case scenario)
// TODO add more tests
func TestFetchFixerIO(t *testing.T) {

	payload := FixerIOPayload{Base: "EUR", Date: "2014-03-04", Rates: map[string]float32{"NOK": 9.4, "URD": 1.4}}
	t.Log(payload.Rates["NOK"])
	faker := faker{payload: payload, err: nil}

	data, err := fetchFixerIO("http://api.fixer.io/latest?base=EUR", faker.fakeGetter)

	if err != nil {
		t.Error("couldn't fetch fixerstuff: ", err.Error())
		return
	}

	if !reflect.DeepEqual(payload, data) {
		t.Error("fetch returned wrong payload", data)
	}

}
