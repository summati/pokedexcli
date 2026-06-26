package main

import "testing"

func TestClearInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input: "Hello World",
			expected: []string{
				"hello",
				"world",
			},
		},
		{
			input: "Hello WORLD",
			expected: []string{
				"hello",
				"world",
			},
		},
	}

	for _, cs := range cases {
		actual := clearInput(cs.input)
		if len(actual) != len(cs.expected) {
			t.Errorf("Expected length %d, but got %d", len(cs.expected), len(actual))
			continue
		}

		for i := range actual {
			if actual[i] != cs.expected[i] {
				t.Errorf("Expected %s, but got %s", cs.expected[i], actual[i])
			}
		}

	}
}
