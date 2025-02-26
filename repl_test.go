package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "    hello  world   ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "   HellO  WoRlD  ",
			expected: []string{"hello", "world"},
		},
	}
	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("incorrect slice length correct slice length is:%v your slice length is:%v", len(c.expected), len(actual))
			break
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("word at index %v does not match. Your word is %v but expected word is %v", i, word, expectedWord)
			}
		}
	}
}
