package handler

import (
	"context"
	"encoding/json"
	"github.com/Justdanru/bhs-test/internal/controller/http/v1/models"
	"github.com/Justdanru/bhs-test/internal/usecase/service"
	ctxlogger "github.com/Justdanru/bhs-test/pkg/context/logger"
	"io"
	"log/slog"
	"net/http"
	"time"
)

type RegisterHandler struct {
	*ErrorsHandler
	userService service.UserService
}

func NewRegisterHandler(
	errorsHandler *ErrorsHandler,
	userService service.UserService,
) *RegisterHandler {
	return &RegisterHandler{
		ErrorsHandler: errorsHandler,
		userService:   userService,
	}
}

type RegisterRequest struct {
	UserCredentials *models.UserCredentials `json:"credentials,required"`
}

type RegisterResponse struct {
	User *models.User `json:"user,required"`
}

func (h *RegisterHandler) Handle(w http.ResponseWriter, r *http.Request) {
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

	var request RegisterRequest

	if err := json.Unmarshal(bytes, &request); err != nil {
		logger.Error("couldn't unmarshal request body", "error", err)
		h.responseError(w, err)
		return
	}

	ctx, done := context.WithTimeout(context.Background(), 1*time.Second)
	ctx = ctxlogger.ContextWithLogger(ctx, logger)

	defer done()

	user, err := h.userService.Register(ctx, request.UserCredentials.Username, request.UserCredentials.Password)
	if err != nil {
		h.responseError(w, err)
		return
	}

	apiUser := models.NewUserFromModel(user)

	response := RegisterResponse{
		User: apiUser,
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
