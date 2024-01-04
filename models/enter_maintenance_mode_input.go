// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// EnterMaintenanceModeInput enter maintenance mode input
//
// swagger:model EnterMaintenanceModeInput
type EnterMaintenanceModeInput struct {

	// shutdown vms
	ShutdownVms []string `json:"shutdown_vms,omitempty"`
}

// Validate validates this enter maintenance mode input
func (m *EnterMaintenanceModeInput) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this enter maintenance mode input based on context it is used
func (m *EnterMaintenanceModeInput) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *EnterMaintenanceModeInput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EnterMaintenanceModeInput) UnmarshalBinary(b []byte) error {
	var res EnterMaintenanceModeInput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
