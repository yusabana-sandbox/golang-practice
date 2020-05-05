package chapter8

import (
	"testing"
)

func TestDoEven(t *testing.T) {
	cases := []struct {
		name     string
		input    int
		expected bool
	}{
		{"isEven_1", 1, false},
		{"isEven_2", 2, true},
		{"isEven_5", 5, false},
		{"isEven_0", 0, true},
		{"isEven_-1", -1, false},
		{"isEven_-2", -2, true},
		{"isEven_-5", -5, false},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if actual := DoEven(c.input); c.expected != actual {
				t.Errorf("want IsOdd(%d) = %v, got %v", c.input, c.expected, actual)
			}
		})
	}
}
