package utils

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContextWithUserId_And_UserIdFromContext(t *testing.T) {
	ctx := context.Background()
	userID := int64(12345)

	// Lưu userID vào context
	ctxWithUser := ContextWithUserId(ctx, userID)

	// Lấy userID từ context
	got := UserIdFromContext(ctxWithUser)
	assert.Equal(t, userID, got)

	// Trường hợp không có userID trong context
	gotEmpty := UserIdFromContext(ctx)
	assert.Equal(t, int64(0), gotEmpty)
}
