package domain

import (
	"context"
)

type UserRepository interface {
	Save(ctx context.Context, customer *User) error
	Find(ctx context.Context, customerID string) (*User, error)
	FindByPhone(ctx context.Context, phone string) (*User, error)
	Update(ctx context.Context, customer *User) error
}
