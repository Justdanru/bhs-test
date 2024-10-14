package handler

import (
	"context"
	"encoding/json"
	"github.com/Justdanru/bhs-test/internal/controller/http/v1/models"
	"github.com/Justdanru/bhs-test/internal/usecase/service"
	ctxlogger "github.com/Justdanru/bhs-test/pkg/context/logger"
	"github.com/gorilla/mux"
	"log/slog"
	"net/http"
	"strconv"
	"time"
)

type UserHandler struct {
	*ErrorsHandler
	userService service.UserService
}

func NewUserHandler(
	errorsHandler *ErrorsHandler,
	userService service.UserService,
) *UserHandler {
	return &UserHandler{
		ErrorsHandler: errorsHandler,
		userService:   userService,
	}
}

type UserResponse struct {
	User *models.User `json:"user,required"`
}

func (h *UserHandler) Handle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	logger, err := ctxlogger.FromContext(r.Context())
	if err != nil {
		h.responseError(w, err)
		return
	}

	if _, ok := vars["user_id"]; !ok {
		logger.Error("couldn't fetch user_id from URL", "error", models.ErrUserIdNotPassed)
		h.responseError(w, models.ErrUserIdNotPassed)
		return
	}

	logger = logger.With(slog.Group(
		"query_parameters",
		slog.String("user_id", vars["user_id"]),
	))

	userId, err := strconv.ParseUint(vars["user_id"], 10, 64)
	if err != nil {
		logger.Error("couldn't parse user_id", "error", err)
		h.responseError(w, models.ErrWrongUserIdFormat)
		return
	}

	ctx, done := context.WithTimeout(context.Background(), 1*time.Second)
	ctx = ctxlogger.ContextWithLogger(ctx, logger)

	defer done()

	user, err := h.userService.User(ctx, service.UserFilter{
		Id: userId,
	})
	if err != nil {
		h.responseError(w, err)
		return
	}

	apiUser := models.NewUserFromModel(user)

	response := &UserResponse{
		User: apiUser,
	}

	responseBytes, err := json.Marshal(response)
	if err != nil {
		logger.Error("couldn't marshal response", "error", err)
		h.responseError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write(responseBytes)
	if err != nil {
		logger.Error("couldn't write response", "error", err)
		h.responseError(w, err)
		return
	}

	return
}
