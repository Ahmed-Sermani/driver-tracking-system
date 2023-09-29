package grpc

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/Ahmed-Sermani/driver-tracking-system/emulation/emulationpb"
	"github.com/centrifugal/gocent/v3"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
)

type emulationServer struct {
	emulationpb.UnimplementedEmulationServiceServer
	log  zerolog.Logger
	cent *gocent.Client
}

func RegisterServer(centClient *gocent.Client, log zerolog.Logger, registrar grpc.ServiceRegistrar) error {
	emulationpb.RegisterEmulationServiceServer(registrar, &emulationServer{log: log, cent: centClient})
	return nil
}

func (s *emulationServer) StartEmulation(ctx context.Context, req *emulationpb.StartEmulationRequest) (*emulationpb.StartEmulationResponse, error) {
	go s.sendMessageEvery(req.GetChannel(), req.GetFreq().AsDuration(), req.GetDuration().AsDuration())
	return &emulationpb.StartEmulationResponse{}, nil
}

// sendMessageEvery  sends a message to the given channel every freq duration until the duration expires.
func (s *emulationServer) sendMessageEvery(channelId string, freq time.Duration, duration time.Duration) {
	ticker := time.NewTicker(freq)
	defer ticker.Stop()

	endTime := time.Now().Add(duration)
	for {
		select {
		case <-ticker.C:
			if time.Now().After(endTime) {
				return
			}
			lat := randomFloat(-90, 90)
			long := randomFloat(-180, 180)
			speed := rand.Intn(121)

			message := fmt.Sprintf(`{"lat": %f, "long": %f, "speed": "%dKM"}`, lat, long, speed)
			if _, err := s.cent.Publish(context.Background(), channelId, []byte(message)); err != nil {
				s.log.Error().Str("channel", channelId).Err(err).Msg("error while publishing")
			}
		}
	}
}

func randomFloat(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}
