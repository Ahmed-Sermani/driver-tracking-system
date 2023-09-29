package application

import (
	"context"

	"github.com/Ahmed-Sermani/driver-tracking-system/internal/auth"
	"github.com/Ahmed-Sermani/driver-tracking-system/users/internal/domain"
)

type (
	RegisterUser struct {
		ID    string
		Name  string
		Phone string
	}

	GetUser struct {
		ID string
	}

	EnableUser struct {
		ID string
	}

	DisableUser struct {
		ID string
	}

	LoginUser struct {
		Phone string
	}

	App interface {
		RegisterUser(ctx context.Context, register RegisterUser) error
		GetUser(ctx context.Context, get GetUser) (*domain.User, error)
		LoginUser(ctx context.Context, login LoginUser) (string, error)
		EnableUser(ctx context.Context, enable EnableUser) error
		DisableUser(ctx context.Context, disable DisableUser) error
	}

	Application struct {
		users        domain.UserRepository
		tokenManager *auth.TokenManager
	}
)

var _ App = (*Application)(nil)

func New(users domain.UserRepository, tm *auth.TokenManager) *Application {
	return &Application{
		users:        users,
		tokenManager: tm,
	}
}

func (a Application) RegisterUser(ctx context.Context, register RegisterUser) error {
	user, err := domain.RegisterUser(register.ID, register.Name, register.Phone)
	if err != nil {
		return err
	}

	return a.users.Save(ctx, user)
}

func (a Application) LoginUser(ctx context.Context, login LoginUser) (string, error) {
	user, err := a.users.FindByPhone(ctx, login.Phone)
	if err != nil {
		return "", err
	}
	return a.tokenManager.NewTokenWithStandardClaims(user.ID, user.Name)
}

func (a Application) GetUser(ctx context.Context, get GetUser) (*domain.User, error) {
	return a.users.Find(ctx, get.ID)
}

func (a Application) EnableUser(ctx context.Context, enable EnableUser) error {
	user, err := a.users.Find(ctx, enable.ID)
	if err != nil {
		return err
	}

	err = user.Enable()
	if err != nil {
		return err
	}

	return a.users.Update(ctx, user)
}

func (a Application) DisableUser(ctx context.Context, disable DisableUser) error {
	user, err := a.users.Find(ctx, disable.ID)
	if err != nil {
		return err
	}

	err = user.Disable()
	if err != nil {
		return err
	}

	return a.users.Update(ctx, user)
}
