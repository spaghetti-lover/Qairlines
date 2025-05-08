package utils

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

// Random number generator that can be recovered
var r = rand.New(rand.NewSource(time.Now().UnixNano()))

// RandomInt generates a random integer between min and max
func RandomInt(min, max int) int {
	return min + r.Intn(max-min+1)
}

// RandomString generates a random string of length n
func RandomString(n int) string {
	var sb strings.Builder
	for i := 0; i < n; i++ {
		c := alphabet[r.Intn(len(alphabet))]
		sb.WriteByte(c)
	}
	return sb.String()
}

// RandomName generates a random name
func RandomName() string {
	return RandomString(6)
}

// RandomStringNum generates a random registration number like "AB1234"
func RandomStringNum() string {
	letters := RandomString(2)       // 2 random letters
	numbers := RandomInt(1000, 9999) // 4-digit random number
	return strings.ToUpper(letters) + fmt.Sprintf("-%04d", numbers)
}
