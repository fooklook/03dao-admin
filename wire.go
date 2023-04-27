//go:build wireinject
// +build wireinject

package main

import (
	"03dao-admin/apps/user"
	"03dao-admin/apps/user/view"
	"github.com/google/wire"
)

type application struct {
	userView *view.UserView
}

func InitApplication() (application, error) {
	wire.Build(
		user.UserSet,
		wire.Struct(new(application), "*"),
	)
	return application{}, nil
}
