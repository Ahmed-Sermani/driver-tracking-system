// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CentrifugoProxySubRefreshResult centrifugo proxy sub refresh result
//
// swagger:model CentrifugoProxySubRefreshResult
type CentrifugoProxySubRefreshResult struct {

	// b64info
	B64info string `json:"b64info,omitempty"`

	// expire at
	ExpireAt string `json:"expireAt,omitempty"`

	// expired
	Expired bool `json:"expired,omitempty"`

	// info
	// Format: byte
	Info strfmt.Base64 `json:"info,omitempty"`
}

// Validate validates this centrifugo proxy sub refresh result
func (m *CentrifugoProxySubRefreshResult) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this centrifugo proxy sub refresh result based on context it is used
func (m *CentrifugoProxySubRefreshResult) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CentrifugoProxySubRefreshResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CentrifugoProxySubRefreshResult) UnmarshalBinary(b []byte) error {
	var res CentrifugoProxySubRefreshResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}