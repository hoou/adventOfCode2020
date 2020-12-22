package main

import (
	"reflect"
	"testing"
)

func TestCalculate(t *testing.T) {
	var tests = []struct {
		input    string
		expected int
	}{
		{"1 + 2 * 3 + 4 * 5 + 6", 231},
		{"1 + (2 * 3) + (4 * (5 + 6))", 51},
		{"2 * 3 + (4 * 5)", 46},
		{"5 + (8 * 3 + 9 + 3 * 4 * 3)", 1445},
		{"5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))", 669060},
		{"((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2", 23340},
	}

	for _, test := range tests {
		if output := Calculate(test.input); output != test.expected {
			t.Error("Test Failed: {} inputted, {} expected, received: {}", test.input, test.expected, output)
		}
	}
}

func TestGetToken(t *testing.T) {
	var tests = []struct {
		input         string
		expectedToken *Token
	}{
		{"1 + 2 * 3 + 4 * 5 + 6", &Token{"operand", 1}},
		{"", nil},
	}

	for _, test := range tests {
		token, _ := getFirstToken(test.input)
		if !reflect.DeepEqual(token, test.expectedToken) {
			t.Error("Token expected to be {}.", test.expectedToken)
		}
	}
}
