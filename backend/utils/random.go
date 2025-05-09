package utils

import (
	"fmt"
	"math/big"
	"math/rand"
	"strings"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

// Random number generator that can be recovered
var r = rand.New(rand.NewSource(time.Now().UnixNano()))

// RandomInt generates a random integer between min and max
func RandomInt(min, max int64) int64 {
	return min + r.Int63n(max-min+1)
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

// RandomTime generates a random time like
func RandomTime() pgtype.Timestamp {
	// Tạo thời gian ngẫu nhiên trong khoảng từ giờ đến 7 ngày sau
	randomDuration := time.Duration(rand.Intn(7*24)) * time.Hour
	t := time.Now().Add(randomDuration)

	return pgtype.Timestamp{
		Time:  t,
		Valid: true,
	}
}

// RandomPrice generates a random price
func RandomPrice() pgtype.Numeric {
	price := big.NewInt(rand.Int63n(1_000_000_000)) // ví dụ giá từ 0 đến 1 tỷ

	return pgtype.Numeric{
		Int:   price,
		Exp:   -2, // 2 chữ số thập phân => vd: 12345678 => 123456.78
		Valid: true,
	}
}

// RandomEmail generate a random email
func RandomEmail() string {
	emailDomains := []string{
		"example.com",
		"testmail.com",
		"mailinator.com",
		"gmail.com",
		"yahoo.com",
	}
	username := RandomString(8)
	domain := emailDomains[rand.Intn(len(emailDomains))]
	return fmt.Sprintf("%s@%s", strings.ToLower(username), domain)
}
