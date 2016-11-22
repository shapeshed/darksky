package darksky

import "testing"

func TestLatLongToString(t *testing.T) {
	s := latLongToString(52.847875, -0.664398)
	expected := "52.847875,-0.664398"
	if s != expected {
		t.Errorf("expected string %q; got %q", expected, s)
	}

}
