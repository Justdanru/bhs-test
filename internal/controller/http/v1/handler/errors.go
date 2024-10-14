package handler

import (
	"errors"
	"fmt"
	"github.com/Justdanru/bhs-test/internal/usecase/repository"
	"net/http"
	"schneider.vip/problem"

	"github.com/Justdanru/bhs-test/internal/models"

	apimodels "github.com/Justdanru/bhs-test/internal/controller/http/v1/models"
)

type ErrorsHandler struct {
}

func NewErrorsHandler() *ErrorsHandler {
	return &ErrorsHandler{}
}

func (h *ErrorsHandler) responseError(w http.ResponseWriter, err error) {
	switch {
	case errors.Is(err, models.ErrPasswordTooShort):
		h.responseProblemJSON(
			w,
			"Password too short",
			"errors://password-too-short",
			http.StatusBadRequest,
		)
	case errors.Is(err, models.ErrPasswordTooLong):
		h.responseProblemJSON(
			w,
			"Password too long",
			"errors://password-too-long",
			http.StatusBadRequest,
		)
	case errors.Is(err, models.ErrUsernameTooShort):
		h.responseProblemJSON(
			w,
			"Username too short",
			"errors://username-too-short",
			http.StatusBadRequest,
		)
	case errors.Is(err, models.ErrUsernameTooLong):
		h.responseProblemJSON(
			w,
			"Username too long",
			"errors://username-too-long",
			http.StatusBadRequest,
		)
	case errors.Is(err, repository.ErrUserNotFound):
		h.responseProblemJSON(
			w,
			"User not found",
			"errors://user-not-found",
			http.StatusNotFound,
		)
	case errors.Is(err, apimodels.ErrUserIdNotPassed):
		h.responseProblemJSON(
			w,
			"User id not passed in URL",
			"errors://user-id-not-passed-in-url",
			http.StatusBadRequest,
		)
	default:
		h.responseProblemJSON(
			w,
			"Oops... Some error occurred",
			"errors://internal-server-error",
			http.StatusInternalServerError,
		)
	}
}

func (h *ErrorsHandler) responseProblemJSON(
	w http.ResponseWriter,
	problemTitle string,
	problemType string,
	problemStatus int,
) {
	_, err := problem.New(
		problem.Title(problemTitle),
		problem.Type(problemType),
		problem.Status(problemStatus),
	).WriteTo(w)

	if err != nil {
		fmt.Printf("Error while response problem+json: %s\n", err.Error())
	}
}
