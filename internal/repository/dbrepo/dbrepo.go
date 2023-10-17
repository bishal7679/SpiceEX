package dbrepo

import (
	"database/sql"

	"github.com/bishal7679/SpiceEx/internal/config"
	"github.com/bishal7679/SpiceEx/internal/repository"
)

type postgresDBRepo struct {
	App *config.AppConfig
	DB *sql.DB
}

func NewpostgresRepo(conn *sql.DB, a *config.AppConfig) repository.DatabaseRepo {
	return &postgresDBRepo{
		App: a,
		DB: conn,
	}
}