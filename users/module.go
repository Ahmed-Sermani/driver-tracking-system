package users

import (
	"context"

	"github.com/Ahmed-Sermani/driver-tracking-system/internal/auth"
	"github.com/Ahmed-Sermani/driver-tracking-system/internal/monolith"
	"github.com/Ahmed-Sermani/driver-tracking-system/users/internal/application"
	"github.com/Ahmed-Sermani/driver-tracking-system/users/internal/grpc"
	"github.com/Ahmed-Sermani/driver-tracking-system/users/internal/logging"
	"github.com/Ahmed-Sermani/driver-tracking-system/users/internal/postgres"
	"github.com/Ahmed-Sermani/driver-tracking-system/users/internal/rest"
)

type Module struct{}

func (m Module) Startup(ctx context.Context, mono monolith.Monolith) error {
	users := postgres.NewUserRepository("users.users", mono.DB())

	var app application.App
	app = application.New(users, auth.NewTokenManager(mono.Config().JWT.Secret, mono.Config().JWT.Expiry))
	app = logging.LogApplicationAccess(app, mono.Logger())

	if err := grpc.RegisterServer(app, mono.RPC()); err != nil {
		return err
	}
	if err := rest.RegisterGateway(ctx, mono.Mux(), mono.Config().Rpc.Address()); err != nil {
		return err
	}
	if err := rest.RegisterSwagger(mono.Mux()); err != nil {
		return err
	}

	return nil
}
