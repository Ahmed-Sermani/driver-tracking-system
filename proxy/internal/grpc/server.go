package grpc

import (
	"context"
	"fmt"

	"github.com/Ahmed-Sermani/driver-tracking-system/proxy/proxypb"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type centrifugoProxyServer struct {
	proxypb.UnimplementedCentrifugoProxyServer
	log           zerolog.Logger
	emulationRepo *EmulationRepository
}

func RegisterServer(emulationRepo *EmulationRepository, log zerolog.Logger, registrar grpc.ServiceRegistrar) error {
	proxypb.RegisterCentrifugoProxyServer(registrar, &centrifugoProxyServer{log: log, emulationRepo: emulationRepo})
	return nil
}

// Connect handles connect events from Centrifugo
func (s *centrifugoProxyServer) Connect(ctx context.Context, req *proxypb.ConnectRequest) (*proxypb.ConnectResponse, error) {
	s.log.Printf("connect event: %+v", req)
	// TODO: implement your own logic to validate and process the connect event
	// for example, you can check the client credentials, return custom data,
	// subscribe the client to some channels, attach some meta information, etc.
	res := &proxypb.ConnectResponse{
		Result: &proxypb.ConnectResult{
			User: ctx.Value("id").(string),
		},
	}
	return res, nil
}

// Refresh handles refresh events from Centrifugo
func (s *centrifugoProxyServer) Refresh(ctx context.Context, req *proxypb.RefreshRequest) (*proxypb.RefreshResponse, error) {
	s.log.Printf("refresh event: %+v", req)
	// TODO: implement your own logic to validate and process the refresh event
	// for example, you can check if the client session is still valid,
	// prolong it or expire it, return some custom data, etc.
	res := &proxypb.RefreshResponse{}
	return res, nil
}

// Subscribe handles subscribe events from Centrifugo
func (s *centrifugoProxyServer) Subscribe(ctx context.Context, req *proxypb.SubscribeRequest) (*proxypb.SubscribeResponse, error) {
	s.log.Printf("subscribe event: %+v", req)
	// TODO: implement your own logic to validate and process the subscribe event
	// for example, you can check if the client has permission to subscribe to the channel,
	// return some custom data or channel information, etc.
	res := &proxypb.SubscribeResponse{}
	if req.GetChannel() != fmt.Sprintf("deliveries:%s", ctx.Value("id").(string)) {
		s.log.Info().Str("channel", req.GetChannel()).Str("otherchannel", fmt.Sprintf("deliveries:%s", ctx.Value("id").(string))).Msg("error matching channels")
		return nil, status.Errorf(codes.Unauthenticated, "unauthorized to subscribe to this channel")
	}
	if err := s.emulationRepo.StartEmulation(ctx, req.GetChannel()); err != nil {
		s.log.Info().Str("channel", req.GetChannel()).Err(err).Msg("error starting emulation")
		return nil, status.Errorf(codes.Internal, "internal error")
	}
	return res, nil
}

// Publish handles publish events from Centrifugo
func (s *centrifugoProxyServer) Publish(ctx context.Context, req *proxypb.PublishRequest) (*proxypb.PublishResponse, error) {
	s.log.Printf("publish event: %+v", req)
	// TODO: implement your own logic to validate and process the publish event
	// for example, you can check if the client has permission to publish to the channel,
	// modify or reject the publication data, etc.
	res := &proxypb.PublishResponse{}
	return res, nil
}

// SubRefresh handles sub_refresh events from Centrifugo
func (s *centrifugoProxyServer) SubRefresh(ctx context.Context, req *proxypb.SubRefreshRequest) (*proxypb.SubRefreshResponse, error) {
	s.log.Printf("sub_refresh event: %+v", req)
	// TODO: implement your own logic to validate and process the sub_refresh event
	// for example, you can check if the client subscription is still valid,
	// prolong it or expire it, return some custom data or channel information, etc.
	res := &proxypb.SubRefreshResponse{}
	return res, nil
}

// RPC handles rpc events from Centrifugo
func (s *centrifugoProxyServer) RPC(ctx context.Context, req *proxypb.RPCRequest) (*proxypb.RPCResponse, error) {
	s.log.Printf("rpc event: %+v", req)
	// TODO: implement your own logic to validate and process the rpc event
	// for example, you can execute some custom logic based on the rpc method and params,
	// return some custom data or an error, etc.
	res := &proxypb.RPCResponse{}
	return res, nil
}
