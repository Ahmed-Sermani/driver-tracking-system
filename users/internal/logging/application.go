package logging

import (
	"context"

	"github.com/Ahmed-Sermani/driver-tracking-system/users/internal/application"
	"github.com/Ahmed-Sermani/driver-tracking-system/users/internal/domain"
	"github.com/rs/zerolog"
)

type Application struct {
	application.App
	logger zerolog.Logger
}

var _ application.App = (*Application)(nil)

func LogApplicationAccess(application application.App, logger zerolog.Logger) Application {
	return Application{
		App:    application,
		logger: logger,
	}
}

// LoginUser(ctx context.Context, login LoginUser) (string, error)

func (a Application) RegisterUser(ctx context.Context, register application.RegisterUser) (err error) {
	a.logger.Info().Msg("--> Users.RegisterUser")
	defer func() { a.logger.Info().Err(err).Msg("<-- Users.RegisterUser") }()
	return a.App.RegisterUser(ctx, register)
}

func (a Application) LoginUser(ctx context.Context, login application.LoginUser) (token string, err error) {
	a.logger.Info().Msg("--> Users.LoginUser")
	defer func() { a.logger.Info().Err(err).Msg("<-- Users.LoginUser") }()
	return a.App.LoginUser(ctx, login)
}

func (a Application) GetUser(ctx context.Context, get application.GetUser) (customer *domain.User,
	err error,
) {
	a.logger.Info().Msg("--> Users.GetUser")
	defer func() { a.logger.Info().Err(err).Msg("<-- Users.GetUser") }()
	return a.App.GetUser(ctx, get)
}

func (a Application) EnableUser(ctx context.Context, enable application.EnableUser) (err error) {
	a.logger.Info().Msg("--> Users.EnableUser")
	defer func() { a.logger.Info().Err(err).Msg("<-- Users.EnableUser") }()
	return a.App.EnableUser(ctx, enable)
}

func (a Application) DisableUser(ctx context.Context, disable application.DisableUser) (err error) {
	a.logger.Info().Msg("--> Users.DisableUser")
	defer func() { a.logger.Info().Err(err).Msg("<-- Users.DisableUser") }()
	return a.App.DisableUser(ctx, disable)
}
