package user

import (
	"context"
	"database/sql"
	"errors"
	"github.com/Justdanru/bhs-test/internal/models"
	"github.com/Justdanru/bhs-test/internal/usecase/repository"
	"github.com/Masterminds/squirrel"
	"github.com/lib/pq"
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

	err := query.RunWith(r.sql).QueryRowContext(ctx).Scan(&id, &username, &passwordHash)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, repository.ErrUserNotFound
		}

		return nil, err
	}

	return models.BuildUser(id, username, passwordHash), nil
}

func isDuplicateEntryError(err error) bool {
	var pqErr *pq.Error
	if errors.As(err, &pqErr) {
		return pqErr.Code == errCodeDuplicateEntry
	} else {
		return false
	}
}
