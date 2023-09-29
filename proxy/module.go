package proxy

import (
	"context"

	"github.com/Ahmed-Sermani/driver-tracking-system/internal/monolith"
	"github.com/Ahmed-Sermani/driver-tracking-system/proxy/internal/grpc"
	"github.com/Ahmed-Sermani/driver-tracking-system/proxy/internal/rest"
)

type Module struct{}

func (Module) Startup(ctx context.Context, mono monolith.Monolith) error {
	conn, err := grpc.Dial(ctx, mono.Config().Rpc.Address())
	if err != nil {
		return err
	}
	if err := grpc.RegisterServer(grpc.NewEmulationRepository(conn), mono.Logger(), mono.RPC()); err != nil {
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
