package factory

import (
	"database/sql"
	"fmt"
	"github.com/Justdanru/bhs-test/config"
	"github.com/Justdanru/bhs-test/internal/infrastructure/repository/user"
	"github.com/Justdanru/bhs-test/internal/usecase/repository"
	"github.com/google/wire"

	_ "github.com/lib/pq"
)

var repositoriesSet = wire.NewSet(
	providePostgreSQLConnection,
	user.NewRepositoryPostgreSQL,

	wire.Bind(new(repository.UserRepository), new(*user.RepositoryPostgreSQL)),
)

func providePostgreSQLConnection(cfg *config.Config) (*sql.DB, func(), error) {
	dsn := fmt.Sprintf(
		"postgresql://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.PG.User,
		cfg.PG.Password,
		cfg.PG.Host,
		cfg.PG.Port,
		cfg.PG.DB,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, nil, fmt.Errorf("error open postgresql connection. %w", err)
	}

	return db, func() {
		_ = db.Close()
	}, nil
}
