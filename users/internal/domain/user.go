package domain

import (
	"github.com/stackus/errors"
)

type User struct {
	ID      string
	Name    string
	Phone   string
	Enabled bool
}

var (
	ErrNameCannotBeBlank   = errors.Wrap(errors.ErrBadRequest, "the user name cannot be blank")
	ErrUserIDCannotBeBlank = errors.Wrap(errors.ErrBadRequest, "the user id cannot be blank")
	ErrPhoneCannotBeBlank  = errors.Wrap(errors.ErrBadRequest, "the user phone be blank")
	ErrUserAlreadyEnabled  = errors.Wrap(errors.ErrBadRequest, "the user is already enabled")
	ErrUserAlreadyDisabled = errors.Wrap(errors.ErrBadRequest, "the user is already disabled")
)

func RegisterUser(id, name, phone string) (*User, error) {
	if id == "" {
		return nil, ErrUserIDCannotBeBlank
	}

	if name == "" {
		return nil, ErrNameCannotBeBlank
	}

	if phone == "" {
		return nil, ErrPhoneCannotBeBlank
	}

	return &User{
		ID:      id,
		Name:    name,
		Phone:   phone,
		Enabled: true,
	}, nil
}

func (c *User) Enable() error {
	if c.Enabled {
		return ErrUserAlreadyEnabled
	}

	c.Enabled = true

	return nil
}

func (c *User) Disable() error {
	if !c.Enabled {
		return ErrUserAlreadyDisabled
	}

	c.Enabled = false

	return nil
}
