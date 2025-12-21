package main

import "testing"

func TestCleanInput(t *testing.T) {

	theOne := []string{"hello", "world"}

	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "hello World",
			expected: theOne,
		},
		{
			input:    "hello World",
			expected: theOne,
		},
		{
			input:    "",
			expected: []string{},
		},
		{
			input:    "      hello	  world    ",
			expected: theOne,
		},
		{
			input:    "HellO WorlD",
			expected: theOne,
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)

		if len(actual) != len(c.expected) {
			t.Errorf(
				"lengths do not match for input %q: expected %d got %d",
				c.input, len(c.expected), len(actual),
			)
			continue
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			if word != expectedWord {
				t.Errorf(
					"cleanInput(%q)[%d] = %q; expected %q",
					c.input, i, word, expectedWord,
				)
			}
		}
	}
}
