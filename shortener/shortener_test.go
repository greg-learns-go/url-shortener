package shortener

import "testing"

func TestFoo(t *testing.T) {
	testCases := []struct {
		Input    string
		Expected string
	}{
		{
			Input:    "a",
			Expected: "bba",
		},
		{
			Input:    "aa",
			Expected: "cca",
		},
		{
			Input:    "b",
			Expected: "bab",
		},
		{
			Input:    "aaaaaaaaaaaaaaaaaaaaaaaaaa",
			Expected: "AAa",
		},
	}

	for _, tC := range testCases {
		got := Shorten(tC.Input)
		if got != tC.Expected {
			t.Errorf("for %s expected %s got %s", tC.Input, tC.Expected, got)
		}
	}
}
