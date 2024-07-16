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

// VirtualPrivateCloudFloatingIP virtual private cloud floating Ip
//
// swagger:model VirtualPrivateCloudFloatingIp
type VirtualPrivateCloudFloatingIP struct {

	// entity async status
	EntityAsyncStatus *EntityAsyncStatus `json:"entityAsyncStatus,omitempty"`

	// external ip
	ExternalIP *string `json:"external_ip,omitempty"`

	// external subnet
	// Required: true
	ExternalSubnet *NestedVirtualPrivateCloudExternalSubnet `json:"external_subnet"`

	// id
	// Required: true
	ID *string `json:"id"`

	// local id
	// Required: true
	LocalID *string `json:"local_id"`

	// vpc
	// Required: true
	Vpc *NestedVirtualPrivateCloud `json:"vpc"`
}

// Validate validates this virtual private cloud floating Ip
func (m *VirtualPrivateCloudFloatingIP) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEntityAsyncStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExternalSubnet(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocalID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVpc(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VirtualPrivateCloudFloatingIP) validateEntityAsyncStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.EntityAsyncStatus) { // not required
		return nil
	}

	if m.EntityAsyncStatus != nil {
		if err := m.EntityAsyncStatus.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("entityAsyncStatus")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("entityAsyncStatus")
			}
			return err
		}
	}

	return nil
}

func (m *VirtualPrivateCloudFloatingIP) validateExternalSubnet(formats strfmt.Registry) error {

	if err := validate.Required("external_subnet", "body", m.ExternalSubnet); err != nil {
		return err
	}

	if m.ExternalSubnet != nil {
		if err := m.ExternalSubnet.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("external_subnet")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("external_subnet")
			}
			return err
		}
	}

	return nil
}

func (m *VirtualPrivateCloudFloatingIP) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *VirtualPrivateCloudFloatingIP) validateLocalID(formats strfmt.Registry) error {

	if err := validate.Required("local_id", "body", m.LocalID); err != nil {
		return err
	}

	return nil
}

func (m *VirtualPrivateCloudFloatingIP) validateVpc(formats strfmt.Registry) error {

	if err := validate.Required("vpc", "body", m.Vpc); err != nil {
		return err
	}

	if m.Vpc != nil {
		if err := m.Vpc.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vpc")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vpc")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this virtual private cloud floating Ip based on the context it is used
func (m *VirtualPrivateCloudFloatingIP) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateEntityAsyncStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateExternalSubnet(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVpc(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VirtualPrivateCloudFloatingIP) contextValidateEntityAsyncStatus(ctx context.Context, formats strfmt.Registry) error {

	if m.EntityAsyncStatus != nil {
		if err := m.EntityAsyncStatus.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("entityAsyncStatus")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("entityAsyncStatus")
			}
			return err
		}
	}

	return nil
}

func (m *VirtualPrivateCloudFloatingIP) contextValidateExternalSubnet(ctx context.Context, formats strfmt.Registry) error {

	if m.ExternalSubnet != nil {
		if err := m.ExternalSubnet.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("external_subnet")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("external_subnet")
			}
			return err
		}
	}

	return nil
}

func (m *VirtualPrivateCloudFloatingIP) contextValidateVpc(ctx context.Context, formats strfmt.Registry) error {

	if m.Vpc != nil {
		if err := m.Vpc.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vpc")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vpc")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *VirtualPrivateCloudFloatingIP) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VirtualPrivateCloudFloatingIP) UnmarshalBinary(b []byte) error {
	var res VirtualPrivateCloudFloatingIP
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}