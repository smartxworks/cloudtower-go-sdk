// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HostAuthInfo host auth info
//
// swagger:model HostAuthInfo
type HostAuthInfo struct {

	// default user password
	DefaultUserPassword *string `json:"default_user_password,omitempty"`

	// root user password
	RootUserPassword *string `json:"root_user_password,omitempty"`
}

// Validate validates this host auth info
func (m *HostAuthInfo) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this host auth info based on context it is used
func (m *HostAuthInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HostAuthInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HostAuthInfo) UnmarshalBinary(b []byte) error {
	var res HostAuthInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}