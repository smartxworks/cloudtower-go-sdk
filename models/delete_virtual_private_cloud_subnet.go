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

// DeleteVirtualPrivateCloudSubnet delete virtual private cloud subnet
//
// swagger:model DeleteVirtualPrivateCloudSubnet
type DeleteVirtualPrivateCloudSubnet struct {

	// id
	// Required: true
	ID *string `json:"id"`
}

// Validate validates this delete virtual private cloud subnet
func (m *DeleteVirtualPrivateCloudSubnet) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeleteVirtualPrivateCloudSubnet) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this delete virtual private cloud subnet based on context it is used
func (m *DeleteVirtualPrivateCloudSubnet) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DeleteVirtualPrivateCloudSubnet) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeleteVirtualPrivateCloudSubnet) UnmarshalBinary(b []byte) error {
	var res DeleteVirtualPrivateCloudSubnet
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
