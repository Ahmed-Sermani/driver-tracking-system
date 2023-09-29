package grpc

import (
	"context"

	"github.com/Ahmed-Sermani/driver-tracking-system/users/internal/application"
	"github.com/Ahmed-Sermani/driver-tracking-system/users/internal/domain"
	"github.com/Ahmed-Sermani/driver-tracking-system/users/userspb"
	"github.com/google/uuid"
	"google.golang.org/grpc"
)

type server struct {
	app application.App
	userspb.UnimplementedUsersServiceServer
}

var _ userspb.UsersServiceServer = (*server)(nil)

func RegisterServer(app application.App, registrar grpc.ServiceRegistrar) error {
	userspb.RegisterUsersServiceServer(registrar, server{app: app})
	return nil
}

func (s server) RegisterUser(ctx context.Context, request *userspb.RegisterUserRequest) (*userspb.RegisterUserResponse, error) {
	id := uuid.New().String()
	err := s.app.RegisterUser(ctx, application.RegisterUser{
		ID:    id,
		Name:  request.GetName(),
		Phone: request.GetPhone(),
	})
	return &userspb.RegisterUserResponse{Id: id}, err
}

func (s server) Login(ctx context.Context, request *userspb.LoginRequest) (*userspb.LoginResponse, error) {
	token, err := s.app.LoginUser(ctx, application.LoginUser{Phone: request.Phone})
	if err != nil {
		return nil, err
	}
	return &userspb.LoginResponse{Token: token}, nil
}

func (s server) GetUser(ctx context.Context, request *userspb.GetUserRequest) (*userspb.GetUserResponse, error) {
	customer, err := s.app.GetUser(ctx, application.GetUser{
		ID: request.GetId(),
	})
	if err != nil {
		return nil, err
	}

	return &userspb.GetUserResponse{
		User: s.userFromDomain(customer),
	}, nil
}

func (s server) userFromDomain(customer *domain.User) *userspb.User {
	return &userspb.User{
		Id:      customer.ID,
		Name:    customer.Name,
		Enabled: customer.Enabled,
	}
}
