package main

import "testing"
// import "fmt"


func TestCleanInput(t *testing.T) {

	cases := []struct {
		input string
		expected []string
	}{
		{
		input: " hello world ",
		expected: []string{"hello", "world"},
		},

		{
		input: "",
		expected: []string{},
		},

		{
		input: "Glurak breathe fire",
		expected: []string{"glurak", "breathe", "fire"},
		},

		{
		input: "Pikachu PIKA pika!!",
		expected: []string{"pikachu", "pika", "pika!!"},
		},

		{
		input: "D1g1m0n   is   not   P0k3m0n",
		expected: []string{"d1g1m0n", "is", "not", "p0k3m0n"},
		},

	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		

		if len(actual) != len(c.expected) {
			t.Errorf("Error: Length actual-expected not matched: '%v' vs '%v'", actual, c.expected)
			continue
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			
			if word != expectedWord {
				t.Errorf("Error: Actual words '%v' are different than expected '%v'", actual, c.expected)
				break
			}
		}
		
	}

}