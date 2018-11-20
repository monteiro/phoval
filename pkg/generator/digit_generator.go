package generator

import (
	"crypto/rand"
	"fmt"
	"io"
)

// maximum of digits when generating a code
const MaxDigits = 6

// GenerateRandomDigits generates random digits to be used as codes to validate phone numbers
func GenerateRandomDigits() (string, error) {
	var table = [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}
	b := make([]byte, MaxDigits)
	n, err := io.ReadAtLeast(rand.Reader, b, MaxDigits)
	if n != MaxDigits || err != nil {
		return "", fmt.Errorf("Error generating random digits")
	}
	for i := 0; i < len(b); i++ {
		b[i] = table[int(b[i])%len(table)]
	}
	return string(b), nil
}
