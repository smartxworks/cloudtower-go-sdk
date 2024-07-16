// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// VirtualPrivateCloudCreationParams virtual private cloud creation params
//
// swagger:model VirtualPrivateCloudCreationParams
type VirtualPrivateCloudCreationParams struct {

	// description
	Description *string `json:"description,omitempty"`

	// mtu
	Mtu *int32 `json:"mtu,omitempty"`

	// name
	// Required: true
	Name *string `json:"name"`

	// vpc service id
	// Required: true
	VpcServiceID *string `json:"vpc_service_id"`
}

// Validate validates this virtual private cloud creation params
func (m *VirtualPrivateCloudCreationParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVpcServiceID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VirtualPrivateCloudCreationParams) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *VirtualPrivateCloudCreationParams) validateVpcServiceID(formats strfmt.Registry) error {

	if err := validate.Required("vpc_service_id", "body", m.VpcServiceID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this virtual private cloud creation params based on context it is used
func (m *VirtualPrivateCloudCreationParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *VirtualPrivateCloudCreationParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VirtualPrivateCloudCreationParams) UnmarshalBinary(b []byte) error {
	var res VirtualPrivateCloudCreationParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
