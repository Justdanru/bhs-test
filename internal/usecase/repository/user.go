package repository

import (
	"context"
	"github.com/Justdanru/bhs-test/internal/models"
)

type GetFilter struct {
	Id       uint64
	Username string
	Limit    uint
	Offset   uint
}

type UserRepository interface {
	Get(ctx context.Context, filter GetFilter) (*models.User, error)
}
