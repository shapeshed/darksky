package darksky

import "testing"

func TestLatLongToString(t *testing.T) {
	s := latLongToString(52.847875, -0.664398)
	expected := "52.847875,-0.664398"
	if s != expected {
		t.Errorf("expected %q; got %q", expected, s)
	}

}

func TestMakeURL(t *testing.T) {
	params := RequestParams{
		Key:       "17b1e8cae7b654290659b438557def7e",
		Latitude:  52.847875,
		Longitude: -0.664397,
		Units:     "si",
	}
	s, err := makeURL(&params)
	if err != nil {
		t.Errorf("got err: %v", err)
	}
	expected := "https://api.darksky.net/forecast/17b1e8cae7b654290659b438557def7e/52.847875,-0.664397?units=si"
	if s != expected {
		t.Errorf("expected %q; got %q", expected, s)
	}

}
