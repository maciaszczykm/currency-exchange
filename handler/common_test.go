package handler

import (
	"reflect"
	"testing"
)

func TestRound(t *testing.T) {
	cases := []struct {
		value    float64
		places   int
		expected float64
	}{
		{
			3.1415,
			2,
			3.14,
		},
		{
			2.72,
			1,
			2.7,
		},
		{
			7.98,
			1,
			8,
		},
		{
			0.737,
			2,
			0.74,
		},
	}

	for _, c := range cases {
		actual := Round(c.value, c.places)
		if !reflect.DeepEqual(actual, c.expected) {
			t.Errorf("Round() returns %#v, expected %#v", actual, c.expected)
		}
	}
}
