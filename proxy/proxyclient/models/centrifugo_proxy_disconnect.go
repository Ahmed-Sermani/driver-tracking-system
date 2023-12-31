// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CentrifugoProxyDisconnect centrifugo proxy disconnect
//
// swagger:model CentrifugoProxyDisconnect
type CentrifugoProxyDisconnect struct {

	// code
	Code int64 `json:"code,omitempty"`

	// reason
	Reason string `json:"reason,omitempty"`
}

// Validate validates this centrifugo proxy disconnect
func (m *CentrifugoProxyDisconnect) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this centrifugo proxy disconnect based on context it is used
func (m *CentrifugoProxyDisconnect) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CentrifugoProxyDisconnect) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CentrifugoProxyDisconnect) UnmarshalBinary(b []byte) error {
	var res CentrifugoProxyDisconnect
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
