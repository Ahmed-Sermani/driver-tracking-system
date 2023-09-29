package emulation

import (
	"context"

	"github.com/Ahmed-Sermani/driver-tracking-system/emulation/internal/grpc"
	"github.com/Ahmed-Sermani/driver-tracking-system/emulation/internal/rest"
	"github.com/Ahmed-Sermani/driver-tracking-system/internal/monolith"
	"github.com/centrifugal/gocent/v3"
)

type Module struct{}

func (Module) Startup(ctx context.Context, mono monolith.Monolith) error {
	cent := gocent.New(gocent.Config{
		Addr: mono.Config().Cent.Conn,
		Key:  mono.Config().Cent.Key,
	})
	if err := grpc.RegisterServer(cent, mono.Logger(), mono.RPC()); err != nil {
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
