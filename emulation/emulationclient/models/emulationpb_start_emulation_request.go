// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// EmulationpbStartEmulationRequest emulationpb start emulation request
//
// swagger:model emulationpbStartEmulationRequest
type EmulationpbStartEmulationRequest struct {

	// channel
	Channel string `json:"channel,omitempty"`

	// duration
	Duration string `json:"duration,omitempty"`

	// freq
	Freq string `json:"freq,omitempty"`
}

// Validate validates this emulationpb start emulation request
func (m *EmulationpbStartEmulationRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this emulationpb start emulation request based on context it is used
func (m *EmulationpbStartEmulationRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *EmulationpbStartEmulationRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EmulationpbStartEmulationRequest) UnmarshalBinary(b []byte) error {
	var res EmulationpbStartEmulationRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
