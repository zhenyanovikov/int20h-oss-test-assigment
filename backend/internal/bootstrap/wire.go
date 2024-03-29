//go:build wireinject
// +build wireinject

package bootstrap

import (
	"github.com/google/wire"
	"oss-backend/internal/config"
	"oss-backend/internal/persistence"
	"oss-backend/internal/persistence/postgres"
	"oss-backend/internal/service"
	"oss-backend/internal/service/auction"
	"oss-backend/internal/service/auth"
	"oss-backend/internal/service/httpserver"
	"oss-backend/internal/service/media"
	"oss-backend/internal/service/user"
)

func Up() (*Dependencies, error) {
	wire.Build(
		wire.Bind(new(service.User), new(*user.Service)),
		wire.Bind(new(service.Auth), new(*auth.Service)),
		wire.Bind(new(service.Auction), new(*auction.Service)),
		wire.Bind(new(service.Media), new(*media.Service)),
		wire.Bind(new(persistence.Auth), new(*postgres.Postgres)),
		wire.Bind(new(persistence.User), new(*postgres.Postgres)),
		wire.Bind(new(persistence.Auction), new(*postgres.Postgres)),
		//wire.Bind(new(persistence.Cache), new(*redis.Redis)),

		config.New,
		httpserver.New,
		getPostgresConfig,

		postgres.New,
		auth.New,
		user.New,
		auction.New,
		media.New,
		//redis.New,
		NewDependencies,
	)
	return &Dependencies{}, nil
}

func getPostgresConfig(cfg config.Config) config.Postgres {
	return cfg.Postgres
}
