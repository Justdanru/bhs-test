package auth

import (
	"context"
	"github.com/Justdanru/bhs-test/config"
	ctxlogger "github.com/Justdanru/bhs-test/pkg/context/logger"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type Service struct {
	secretKey []byte
}

func NewService(cfg *config.Config) *Service {
	return &Service{
		secretKey: []byte(cfg.Auth.SecretKey),
	}
}

func (s *Service) NewToken(ctx context.Context, userId uint64) (string, error) {
	logger, err := ctxlogger.FromContext(ctx)
	if err != nil {
		return "", err
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userId,
		"iss": "bhs-test",
		"exp": time.Now().Add(24 * time.Hour).Unix(),
		"iat": time.Now().Unix(),
	})

	token, err := claims.SignedString(s.secretKey)
	if err != nil {
		logger.Error("couldn't create jwt token string", "error", err)

		return "", err
	}

	return token, nil
}

func (s *Service) VerifyToken(ctx context.Context, tokenString string) (bool, error) {
	logger, err := ctxlogger.FromContext(ctx)
	if err != nil {
		return false, err
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return s.secretKey, nil
	})
	if err != nil {
		logger.Error("error while parsing jwt token string", "error", err)
		return false, err
	}

	if !token.Valid {
		return false, nil
	}

	return true, nil
}
