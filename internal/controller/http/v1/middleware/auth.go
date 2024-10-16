package middleware

import (
	"github.com/Justdanru/bhs-test/internal/usecase/service"
	"net/http"
	"strings"
)

type AuthMiddleware struct {
	authService service.AuthService
}

func NewAuthMiddleware(authService service.AuthService) *AuthMiddleware {
	return &AuthMiddleware{
		authService: authService,
	}
}

func (m *AuthMiddleware) Handle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		authString := req.Header.Get("Authorization")

		authStringParts := strings.Split(authString, " ")

		if len(authStringParts) < 2 {
			http.Error(w, "Wrong 'Authorization' header format", http.StatusUnauthorized)
			return
		}

		token := authStringParts[1]

		isValid, err := m.authService.VerifyToken(req.Context(), token)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		if !isValid {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, req)
	})
}
