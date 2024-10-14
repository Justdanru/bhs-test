package user

import (
	"context"
	"database/sql"
	"errors"
	"github.com/Justdanru/bhs-test/internal/models"
	"github.com/Justdanru/bhs-test/internal/usecase/repository"
	ctxlogger "github.com/Justdanru/bhs-test/pkg/context/logger"
	"github.com/Masterminds/squirrel"
	"github.com/lib/pq"
	"log/slog"
)

var (
	errCodeDuplicateEntry = pq.ErrorCode("23505")
)

type RepositoryPostgreSQL struct {
	sql *sql.DB
}

func NewRepositoryPostgreSQL(sql *sql.DB) *RepositoryPostgreSQL {
	return &RepositoryPostgreSQL{
		sql: sql,
	}
}

func (r *RepositoryPostgreSQL) Get(ctx context.Context, filter repository.GetFilter) (*models.User, error) {
	logger, err := ctxlogger.FromContext(ctx)
	if err != nil {
		return nil, err
	}

	logger = logger.With(
		slog.Group("filter",
			slog.Uint64("id", filter.Id),
			slog.String("username", filter.Username),
			slog.Uint64("limit", uint64(filter.Limit)),
			slog.Uint64("offset", uint64(filter.Offset)),
		),
	)

	query := squirrel.Select("id, username, password_hash").
		From("users").
		PlaceholderFormat(squirrel.Dollar)

	if filter.Id != 0 {
		query.Where(squirrel.Eq{"id": filter.Id})
	}

	if filter.Username != "" {
		query.Where(squirrel.Eq{"username": filter.Username})
	}

	if filter.Limit != 0 {
		query.Limit(uint64(filter.Limit))
	}

	if filter.Offset != 0 {
		query.Offset(uint64(filter.Offset))
	}

	var (
		id                     uint64
		username, passwordHash string
	)

	err = query.RunWith(r.sql).QueryRowContext(ctx).Scan(&id, &username, &passwordHash)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, repository.ErrUserNotFound
		}

		logger.Error("sql query execution failed", "error", err)
		return nil, err
	}

	return models.BuildUser(id, username, passwordHash), nil
}

func (r *RepositoryPostgreSQL) Add(ctx context.Context, user *models.User) (*models.User, error) {
	logger, err := ctxlogger.FromContext(ctx)
	if err != nil {
		return nil, err
	}

	logger = logger.With(slog.Group(
		"added_user",
		slog.String("username", user.Username()),
	))

	var id uint64

	err = squirrel.Insert("users").
		Columns("username", "password_hash").
		Values(user.Username(), user.PasswordHash()).
		PlaceholderFormat(squirrel.Dollar).
		Suffix("RETURNING id").
		RunWith(r.sql).
		QueryRowContext(ctx).
		Scan(&id)
	if err != nil {
		logger.Error("couldn't add new user", "error", err)

		if isDuplicateEntryError(err) {
			return nil, repository.ErrUsernameAlreadyTaken
		}

		return nil, err
	}

	user.SetId(id)

	return user, nil
}

func isDuplicateEntryError(err error) bool {
	var pqErr *pq.Error
	if errors.As(err, &pqErr) {
		return pqErr.Code == errCodeDuplicateEntry
	} else {
		return false
	}
}
