package main

import "testing"

const checkMark = "\u2713"
const ballotX = "\u2717"

func TestCommaFloatSign(t *testing.T) {
	var integerStrings = []struct {
		input    string
		expected string
	}{
		{"", ""}, {"1", "1"}, {"12", "12"}, {"123", "123"},
		{"1234", "1,234"}, {"12345", "12,345"}, {"123456", "123,456"}, {"1234567", "1,234,567"},
		{"12345678", "12,345,678"}, {"123456789", "123,456,789"}, {"1234567890", "1,234,567,890"}, {"12345678901", "12,345,678,901"},
	}

	t.Log("Given the inputs to test formatting different integer strings")
	{
		for _, item := range integerStrings {
			t.Logf("\twhen formatting %s for output %s", item.input, item.expected)
			{
				formattedStr := comma(item.input)
				if formattedStr != item.expected {
					t.Fatalf("\t\tShould be formatted to %s, %s", item.expected, ballotX)
				}
				t.Logf("Should be formatted to %s, %s", item.expected, checkMark)
			}
		}
	}
}
