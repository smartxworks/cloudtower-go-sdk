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

// DeleteVirtualPrivateCloudFloatingIP delete virtual private cloud floating Ip
//
// swagger:model DeleteVirtualPrivateCloudFloatingIp
type DeleteVirtualPrivateCloudFloatingIP struct {

	// id
	// Required: true
	ID *string `json:"id"`
}

// Validate validates this delete virtual private cloud floating Ip
func (m *DeleteVirtualPrivateCloudFloatingIP) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeleteVirtualPrivateCloudFloatingIP) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this delete virtual private cloud floating Ip based on context it is used
func (m *DeleteVirtualPrivateCloudFloatingIP) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DeleteVirtualPrivateCloudFloatingIP) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeleteVirtualPrivateCloudFloatingIP) UnmarshalBinary(b []byte) error {
	var res DeleteVirtualPrivateCloudFloatingIP
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
