// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// MaintenanceModeVerify maintenance mode verify
//
// swagger:model MaintenanceModeVerify
type MaintenanceModeVerify struct {

	// changed
	Changed *bool `json:"changed,omitempty"`

	// reason
	Reason *string `json:"reason,omitempty"`
}

// Validate validates this maintenance mode verify
func (m *MaintenanceModeVerify) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this maintenance mode verify based on context it is used
func (m *MaintenanceModeVerify) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MaintenanceModeVerify) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MaintenanceModeVerify) UnmarshalBinary(b []byte) error {
	var res MaintenanceModeVerify
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
