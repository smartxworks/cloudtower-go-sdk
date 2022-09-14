// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ApplicationVMSpecStorage application Vm spec storage
//
// swagger:model ApplicationVmSpecStorage
type ApplicationVMSpecStorage struct {

	// size
	Size *int64 `json:"size,omitempty"`
}

// Validate validates this application Vm spec storage
func (m *ApplicationVMSpecStorage) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this application Vm spec storage based on context it is used
func (m *ApplicationVMSpecStorage) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ApplicationVMSpecStorage) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ApplicationVMSpecStorage) UnmarshalBinary(b []byte) error {
	var res ApplicationVMSpecStorage
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}