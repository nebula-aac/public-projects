//go:build wireinject
// +build wireinject

package user

import (
	"database/sql"

	"github.com/google/wire"
)

func Wire(db *sql.DB) *handler {
	wire.Build(ProviderSet)
	return &handler{}
}
