package auth

import (
	"context"
	"errors"
	"strconv"
)

type userIdContextKey string

const userIDKey userIdContextKey = "userID"

func WithUserId(ctx context.Context, userId string) context.Context {
	return context.WithValue(ctx, userIDKey, userId)
}

// UserIdFromContext MSA인 경우, userId를 ctx에서 받아와야 함
func UserIdFromContext(ctx context.Context) (int, error) {
	userIdString, ok := ctx.Value(userIDKey).(string)

	if !ok {
		return -1, errors.New("user ID not found in context")
	}

	userId, err := strconv.Atoi(userIdString)

	if err != nil {
		return -1, err
	}

	return userId, nil
}

type authorityContextKey string

const userAuthorityKey authorityContextKey = "userAuthority"

func WithUserAuthority(ctx context.Context, v Authority) context.Context {
	return context.WithValue(ctx, userAuthorityKey, v)
}

func UserAuthorityFromContext(ctx context.Context) Authority {
	v, _ := ctx.Value(userAuthorityKey).(Authority)
	return v
}

type apiOperationNameContextKey string

const apiOperationNameKey apiOperationNameContextKey = "apiOperationName"

func WithApiOperationName(ctx context.Context, name string) context.Context {
	return context.WithValue(ctx, apiOperationNameKey, name)
}

func ApiOperationNameFromContext(ctx context.Context) (string, error) {
	operationName, ok := ctx.Value(apiOperationNameKey).(string)
	if !ok {
		return "", errors.New("operationName not found in context")
	}

	return operationName, nil
}
