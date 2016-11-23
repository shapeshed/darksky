package darksky

import "testing"

func TestLatLongToString(t *testing.T) {
	s := latLongToString(52.847875, -0.664398)
	expected := "52.847875,-0.664398"
	if s != expected {
		t.Errorf("expected %q; got %q", expected, s)
	}

}

type makeURLTest struct {
	params   RequestParams
	expected string
}

var makeURLTests = []makeURLTest{
	{
		RequestParams{
			Key:       "17b1e8cae7b654290659b438557def7e",
			Latitude:  52.847875,
			Longitude: -0.664397,
			Exclude:   "currently,minutely",
		},
		"https://api.darksky.net/forecast/17b1e8cae7b654290659b438557def7e/52.847875,-0.664397?exclude=currently%2Cminutely",
	},
	{
		RequestParams{
			Key:       "17b1e8cae7b654290659b438557def7e",
			Latitude:  52.847875,
			Longitude: -0.664397,
			Extend:    "hourly",
		},
		"https://api.darksky.net/forecast/17b1e8cae7b654290659b438557def7e/52.847875,-0.664397?extend=hourly",
	},
	{
		RequestParams{
			Key:       "17b1e8cae7b654290659b438557def7e",
			Latitude:  52.847875,
			Longitude: -0.664397,
			Lang:      "fr",
		},
		"https://api.darksky.net/forecast/17b1e8cae7b654290659b438557def7e/52.847875,-0.664397?lang=fr",
	},
	{
		RequestParams{
			Key:       "17b1e8cae7b654290659b438557def7e",
			Latitude:  52.847875,
			Longitude: -0.664397,
			Units:     "si",
		},
		"https://api.darksky.net/forecast/17b1e8cae7b654290659b438557def7e/52.847875,-0.664397?units=si",
	},
	{
		RequestParams{
			Key:       "17b1e8cae7b654290659b438557def7e",
			Latitude:  52.847875,
			Longitude: -0.664397,
			Exclude:   "currently,minutely",
			Extend:    "hourly",
			Lang:      "fr",
			Units:     "si",
		},
		"https://api.darksky.net/forecast/17b1e8cae7b654290659b438557def7e/52.847875,-0.664397?exclude=currently%2Cminutely&extend=hourly&lang=fr&units=si",
	},
}

func TestMakeURL(t *testing.T) {

	for _, tt := range makeURLTests {
		actual, err := makeURL(&tt.params)
		if err != nil {
			t.Errorf("got err: %v", err)
		}
		if actual != tt.expected {
			t.Errorf("\nexpected %v\n actual %v", tt.expected, actual)
		}
	}

}
