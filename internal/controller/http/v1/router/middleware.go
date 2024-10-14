package router

import (
	ctxlogger "github.com/Justdanru/bhs-test/pkg/context/logger"
	"log/slog"
	"net/http"
	"os"
)

func (r *Router) initMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if req.Header.Get("Content-Type") != "application/json" {
			http.Error(w, "Content-Type unacceptable", http.StatusBadRequest)
			return
		}

		id := r.requestId.Add(1)

		logger := slog.New(slog.NewJSONHandler(os.Stdout, nil)).With(slog.Group(
			"request",
			slog.String("url", req.RequestURI),
			slog.String("method", req.Method),
			slog.Uint64("id", id),
		))

		req = req.WithContext(ctxlogger.ContextWithLogger(req.Context(), logger))

		next.ServeHTTP(w, req)
	})
}
