package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncodeDecodeBookingID(t *testing.T) {
	id := int64(123456789)

	// Test encode
	hash, err := EncodeBookingID(id)
	assert.NoError(t, err)
	assert.NotEmpty(t, hash)

	// Test decode
	decodedID, err := DecodeBookingID(hash)
	assert.NoError(t, err)
	assert.Equal(t, id, decodedID)
}

func TestDecodeBookingID_InvalidHash(t *testing.T) {
	invalidHash := "invalidhash"
	id, err := DecodeBookingID(invalidHash)
	assert.Error(t, err)
	assert.Equal(t, int64(0), id)
}
