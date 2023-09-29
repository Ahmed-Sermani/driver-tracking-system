package grpc

import (
	"context"
	"time"

	"github.com/Ahmed-Sermani/driver-tracking-system/emulation/emulationpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/durationpb"
)

type EmulationRepository struct {
	client emulationpb.EmulationServiceClient
}

func NewEmulationRepository(conn *grpc.ClientConn) *EmulationRepository {
	return &EmulationRepository{client: emulationpb.NewEmulationServiceClient(conn)}
}

func (r *EmulationRepository) StartEmulation(ctx context.Context, channelId string) error {
	md, _ := metadata.FromIncomingContext(ctx)
	newCtx := metadata.NewOutgoingContext(ctx, md)

	_, err := r.client.StartEmulation(
		newCtx,
		&emulationpb.StartEmulationRequest{
			Channel:  channelId,
			Freq:     durationpb.New(10 * time.Second),
			Duration: durationpb.New(3 * time.Hour),
		},
	)
	return err
}
