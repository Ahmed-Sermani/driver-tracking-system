// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CentrifugoProxyPublishRequest centrifugo proxy publish request
//
// swagger:model CentrifugoProxyPublishRequest
type CentrifugoProxyPublishRequest struct {

	// b64data
	B64data string `json:"b64data,omitempty"`

	// channel
	Channel string `json:"channel,omitempty"`

	// client
	Client string `json:"client,omitempty"`

	// data
	// Format: byte
	Data strfmt.Base64 `json:"data,omitempty"`

	// encoding
	Encoding string `json:"encoding,omitempty"`

	// meta
	// Format: byte
	Meta strfmt.Base64 `json:"meta,omitempty"`

	// protocol
	Protocol string `json:"protocol,omitempty"`

	// transport
	Transport string `json:"transport,omitempty"`

	// user
	User string `json:"user,omitempty"`
}

// Validate validates this centrifugo proxy publish request
func (m *CentrifugoProxyPublishRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this centrifugo proxy publish request based on context it is used
func (m *CentrifugoProxyPublishRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CentrifugoProxyPublishRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CentrifugoProxyPublishRequest) UnmarshalBinary(b []byte) error {
	var res CentrifugoProxyPublishRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
