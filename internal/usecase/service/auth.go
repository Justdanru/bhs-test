package service

import "context"

type AuthService interface {
	NewToken(ctx context.Context, userId uint64) (string, error)
	VerifyToken(ctx context.Context, token string) (bool, error)
}
