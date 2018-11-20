package generator_test

import (
	"phoval/pkg/generator"
	"strconv"
	"testing"
)

func TestDigitGenerator(t *testing.T) {
	code, err := generator.GenerateRandomDigits()
	if err != nil {
		t.Errorf("Error while generating random digits")
	}
	if len(code) != generator.MaxDigits {
		t.Errorf("The code doesnt have the correct length: Expected: %d, Got: %d", len(code), 6)
	}
	if _, err := strconv.Atoi(code); err != nil {
		t.Errorf("Verification code '%s' is not a number.", code)
	}
}
