package utils

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContextWithUserId_And_UserIdFromContext(t *testing.T) {
	ctx := context.Background()
	userID := int64(12345)

	// Store userID in context
	ctxWithUser := ContextWithUserId(ctx, userID)

	// Lấy userID từ context
	got := UserIdFromContext(ctxWithUser)
	assert.Equal(t, userID, got)

	// Case when no userID is in context
	gotEmpty := UserIdFromContext(ctx)
	assert.Equal(t, int64(0), gotEmpty)
}
