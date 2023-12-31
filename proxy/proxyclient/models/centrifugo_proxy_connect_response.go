// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CentrifugoProxyConnectResponse centrifugo proxy connect response
//
// swagger:model CentrifugoProxyConnectResponse
type CentrifugoProxyConnectResponse struct {

	// disconnect
	Disconnect *CentrifugoProxyDisconnect `json:"disconnect,omitempty"`

	// error
	Error *CentrifugoProxyError `json:"error,omitempty"`

	// result
	Result *CentrifugoProxyConnectResult `json:"result,omitempty"`
}

// Validate validates this centrifugo proxy connect response
func (m *CentrifugoProxyConnectResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDisconnect(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateError(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResult(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CentrifugoProxyConnectResponse) validateDisconnect(formats strfmt.Registry) error {
	if swag.IsZero(m.Disconnect) { // not required
		return nil
	}

	if m.Disconnect != nil {
		if err := m.Disconnect.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("disconnect")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("disconnect")
			}
			return err
		}
	}

	return nil
}

func (m *CentrifugoProxyConnectResponse) validateError(formats strfmt.Registry) error {
	if swag.IsZero(m.Error) { // not required
		return nil
	}

	if m.Error != nil {
		if err := m.Error.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("error")
			}
			return err
		}
	}

	return nil
}

func (m *CentrifugoProxyConnectResponse) validateResult(formats strfmt.Registry) error {
	if swag.IsZero(m.Result) { // not required
		return nil
	}

	if m.Result != nil {
		if err := m.Result.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("result")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("result")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this centrifugo proxy connect response based on the context it is used
func (m *CentrifugoProxyConnectResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDisconnect(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateError(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateResult(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CentrifugoProxyConnectResponse) contextValidateDisconnect(ctx context.Context, formats strfmt.Registry) error {

	if m.Disconnect != nil {

		if swag.IsZero(m.Disconnect) { // not required
			return nil
		}

		if err := m.Disconnect.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("disconnect")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("disconnect")
			}
			return err
		}
	}

	return nil
}

func (m *CentrifugoProxyConnectResponse) contextValidateError(ctx context.Context, formats strfmt.Registry) error {

	if m.Error != nil {

		if swag.IsZero(m.Error) { // not required
			return nil
		}

		if err := m.Error.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("error")
			}
			return err
		}
	}

	return nil
}

func (m *CentrifugoProxyConnectResponse) contextValidateResult(ctx context.Context, formats strfmt.Registry) error {

	if m.Result != nil {

		if swag.IsZero(m.Result) { // not required
			return nil
		}

		if err := m.Result.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("result")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("result")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CentrifugoProxyConnectResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CentrifugoProxyConnectResponse) UnmarshalBinary(b []byte) error {
	var res CentrifugoProxyConnectResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
