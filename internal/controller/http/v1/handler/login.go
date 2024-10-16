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

type LoginHandler struct {
	*ErrorsHandler
	authService service.AuthService
	userService service.UserService
}

func NewLoginHandler(
	authService service.AuthService,
	userService service.UserService,
	errorsHandler *ErrorsHandler,
) *LoginHandler {
	return &LoginHandler{
		ErrorsHandler: errorsHandler,
		authService:   authService,
		userService:   userService,
	}
}

type LoginRequest struct {
	UserCredentials *models.UserCredentials `json:"credentials,required"`
}

type LoginResponse struct {
	Ok   bool         `json:"ok,required"`
	Auth *models.Auth `json:"auth,required"`
}

func (h *LoginHandler) Handle(w http.ResponseWriter, r *http.Request) {
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

	user, err := h.userService.User(ctx, service.UserFilter{Username: request.UserCredentials.Username})
	if err != nil {
		h.responseError(w, err)
		return
	}

	ok, err := user.CheckPassword(request.UserCredentials.Password)
	if err != nil {
		logger.Error("error comparing password", "error", err)
		h.responseError(w, err)
		return
	}

	var response LoginResponse

	response.Ok = ok

	if ok {
		token, err := h.authService.NewToken(ctx, user.Id())
		if err != nil {
			h.responseError(w, err)
			return
		}

		response.Auth = &models.Auth{AccessToken: token}
	}

	if bytes, err = json.Marshal(&response); err != nil {
		logger.Error("couldn't marshal response", "error", err)
		h.responseError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write(bytes)
	if err != nil {
		logger.Error("couldn't write response", "error", err)
		h.responseError(w, err)
		return
	}
}
