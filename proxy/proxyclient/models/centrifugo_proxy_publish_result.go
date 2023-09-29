// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CentrifugoProxyPublishResult centrifugo proxy publish result
//
// swagger:model CentrifugoProxyPublishResult
type CentrifugoProxyPublishResult struct {

	// b64data
	B64data string `json:"b64data,omitempty"`

	// data
	// Format: byte
	Data strfmt.Base64 `json:"data,omitempty"`

	// skip history
	SkipHistory bool `json:"skipHistory,omitempty"`
}

// Validate validates this centrifugo proxy publish result
func (m *CentrifugoProxyPublishResult) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this centrifugo proxy publish result based on context it is used
func (m *CentrifugoProxyPublishResult) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CentrifugoProxyPublishResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CentrifugoProxyPublishResult) UnmarshalBinary(b []byte) error {
	var res CentrifugoProxyPublishResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
