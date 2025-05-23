package utils

import (
	"context"
)

type contextKey string

const userIdKey contextKey = "userId"

// ContextWithUserID lưu UserID vào context
func ContextWithUserId(ctx context.Context, userId int64) context.Context {
	return context.WithValue(ctx, userIdKey, userId)
}

// UserIDFromContext lấy UserID từ context
func UserIdFromContext(ctx context.Context) int64 {
	userId, ok := ctx.Value(userIdKey).(int64)
	if !ok {
		return 0
	}
	return userId
}
