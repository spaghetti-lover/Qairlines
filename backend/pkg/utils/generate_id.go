package utils

import (
	"sync"

	"github.com/speps/go-hashids"
)

// HashIDGenerator is used for decoding and encoding ID
type HashIDGenerator struct {
	hasher *hashids.HashID
	mu     sync.Mutex
}

// Singleton instance
var (
	hashIDGenerator *HashIDGenerator
	once            sync.Once
)

// GetHashIDGenerator return instance of HashIDGenerator
func GetHashIDGenerator() *HashIDGenerator {
	once.Do(func() {
		hd := hashids.NewData()
		hd.Salt = "ducanhdeptrai" // Lưu ý: Nên đưa salt vào config/env
		hd.MinLength = 6
		h, _ := hashids.NewWithData(hd)

		hashIDGenerator = &HashIDGenerator{
			hasher: h,
		}
	})
	return hashIDGenerator
}

// EncodeBookingID encode ID from int64 to hashid
func (g *HashIDGenerator) EncodeBookingID(id int64) (string, error) {
	g.mu.Lock()
	defer g.mu.Unlock()

	encoded, err := g.hasher.EncodeInt64([]int64{id})
	if err != nil {
		return "", err
	}
	return encoded, nil
}

// DecodeBookingID giải mã chuỗi hashid thành int64
func (g *HashIDGenerator) DecodeBookingID(hash string) (int64, error) {
	g.mu.Lock()
	defer g.mu.Unlock()

	decoded, err := g.hasher.DecodeInt64WithError(hash)
	if err != nil || len(decoded) == 0 {
		return 0, err
	}
	return decoded[0], nil
}

// Hàm wrapper tiện lợi nếu muốn dùng trực tiếp
func EncodeBookingID(id int64) (string, error) {
	return GetHashIDGenerator().EncodeBookingID(id)
}

// Hàm wrapper tiện lợi nếu muốn dùng trực tiếp
func DecodeBookingID(hash string) (int64, error) {
	return GetHashIDGenerator().DecodeBookingID(hash)
}
