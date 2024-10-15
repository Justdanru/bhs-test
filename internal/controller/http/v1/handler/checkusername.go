package handler

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/Justdanru/bhs-test/internal/usecase/repository"
	"github.com/Justdanru/bhs-test/internal/usecase/service"
	ctxlogger "github.com/Justdanru/bhs-test/pkg/context/logger"
	"io"
	"log/slog"
	"net/http"
	"time"
)

type CheckUsernameHandler struct {
	*ErrorsHandler
	userService service.UserService
}

func NewCheckUsernameHandler(
	errorsHandler *ErrorsHandler,
	userService service.UserService,
) *CheckUsernameHandler {
	return &CheckUsernameHandler{
		ErrorsHandler: errorsHandler,
		userService:   userService,
	}
}

type CheckUsernameRequest struct {
	Username string `json:"username,required"`
}

type CheckUsernameResponse struct {
	IsExists bool `json:"is_exists,required"`
}

func (h *CheckUsernameHandler) Handle(w http.ResponseWriter, r *http.Request) {
	logger, err := ctxlogger.FromContext(r.Context())
	if err != nil {
		h.responseError(w, err)
		return
	}

	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		logger.Error("error while reading request body", "error", err)
		h.responseError(w, err)
	}

	defer r.Body.Close()

	logger = logger.With(slog.String("request_body", string(bytes)))

	var request CheckUsernameRequest

	if err := json.Unmarshal(bytes, &request); err != nil {
		logger.Error("couldn't unmarshal request body", "error", err)
		h.responseError(w, err)
		return
	}

	ctx, done := context.WithTimeout(context.Background(), 1*time.Second)
	ctx = ctxlogger.ContextWithLogger(ctx, logger)

	defer done()

	response := CheckUsernameResponse{IsExists: true}

	_, err = h.userService.User(ctx, service.UserFilter{
		Username: request.Username,
		Offset:   0,
		Limit:    1,
	})
	if err != nil {
		if !errors.Is(err, repository.ErrUserNotFound) {
			h.responseError(w, err)
			return
		}

		response.IsExists = false
	}

	if bytes, err = json.Marshal(&response); err != nil {
		logger.Error("couldn't marshal response", "error", err)
		h.responseError(w, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	_, err = w.Write(bytes)
	if err != nil {
		logger.Error("couldn't write response", "error", err)
		h.responseError(w, err)
		return
	}

	return
}
