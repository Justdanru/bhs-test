package service

import (
	"context"
	"github.com/Justdanru/bhs-test/internal/models"
)

type UserFilter struct {
	Id       uint64
	Username string
	Limit    uint
	Offset   uint
}

type UserService interface {
	User(ctx context.Context, filter UserFilter) (*models.User, error)
}
