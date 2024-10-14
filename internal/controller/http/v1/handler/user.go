package handler

import (
	"context"
	"encoding/json"
	"github.com/Justdanru/bhs-test/internal/controller/http/v1/models"
	"github.com/Justdanru/bhs-test/internal/usecase/service"
	"github.com/gorilla/mux"
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

	if _, ok := vars["user_id"]; !ok {
		h.responseError(w, models.ErrUserIdNotPassed)
		return
	}

	userId, err := strconv.ParseUint(vars["user_id"], 10, 64)
	if err != nil {
		h.responseError(w, models.ErrWrongUserIdFormat)
		return
	}

	ctx, done := context.WithTimeout(context.Background(), 1*time.Second)

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
		h.responseError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write(responseBytes)
	if err != nil {
		h.responseError(w, err)
		return
	}

	return
}
