package hexrgb

import "testing"

func TestHexToRGB(t *testing.T) {
	type testColor struct {
		input    string
		expected string
	}

	shouldWorkColors := []testColor{
		{input: "#123412", expected: "18,52,18"},
		{input: "123412", expected: "18,52,18"},
		{input: "#123412", expected: "18,52,18"},
		{input: "#000000", expected: "0,0,0"},
		{input: "FFFFFF", expected: "255,255,255"},
	}

	shouldCauseError := []string{
		"12#1212",
		"121212#",
		"##121212",
		"##121212",
		"12121212",
		"1212",
		"12",
		"Asdfalkjs",
	}

	for _, color := range shouldWorkColors {
		output, err := ToRGB(color.input)
		if err != nil {
			t.Fatal(err)
		}
		if output != color.expected {
			t.Fatalf("Expected: %s Got: %s", color.expected, output)
		}
	}

	for _, input := range shouldCauseError {
		_, err := ToRGB(input)
		if err == nil {
			t.Fatalf("Input: %s should have caused an error", input)
		}
	}
}

func TestRGBToHex(t *testing.T) {
	type testColor struct {
		input    string
		expected string
	}

	shouldWorkColors := []testColor{
		{input: "18,52,18", expected: "#123412"},
		{input: "0,0,0", expected: "#000000"},
		{input: "255,255,255", expected: "#FFFFFF"},
	}

	shouldCauseError := []string{
		"12,,12 12",
		"12 12 12",
		"12",
		",12,12,12",
		"12,12,12,",
		"1212",
		"12",
		"#121212",
		"Asdfalkjs",
	}

	for _, color := range shouldWorkColors {
		output, err := ToHex(color.input)
		if err != nil {
			t.Fatal(err)
		}
		if output != color.expected {
			t.Fatalf("Expected: %s Got: %s", color.expected, output)
		}
	}

	for _, input := range shouldCauseError {
		_, err := ToHex(input)
		if err == nil {
			t.Fatalf("Input: %s should have caused an error", input)
		}
	}
}
