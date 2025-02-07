// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// VirtualPrivateCloudRouteParams virtual private cloud route params
//
// swagger:model VirtualPrivateCloudRouteParams
type VirtualPrivateCloudRouteParams struct {

	// destination
	// Required: true
	Destination *string `json:"destination"`

	// next hop local id
	// Required: true
	NextHopLocalID *string `json:"next_hop_local_id"`

	// next hop type
	// Required: true
	NextHopType *VirtualPrivateCloudRouteNextHopType `json:"next_hop_type"`

	MarshalOpts *VirtualPrivateCloudRouteParamsMarshalOpts `json:"-"`
}

type VirtualPrivateCloudRouteParamsMarshalOpts struct {
	Destination_Explicit_Null_When_Empty bool

	NextHopLocalID_Explicit_Null_When_Empty bool

	NextHopType_Explicit_Null_When_Empty bool
}

func (m VirtualPrivateCloudRouteParams) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	b.WriteString("{")

	first := true

	// handle nullable field destination
	if m.Destination != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"destination\":")
		bytes, err := swag.WriteJSON(m.Destination)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.Destination_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"destination\":null")
		first = false
	}

	// handle nullable field next_hop_local_id
	if m.NextHopLocalID != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"next_hop_local_id\":")
		bytes, err := swag.WriteJSON(m.NextHopLocalID)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.NextHopLocalID_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"next_hop_local_id\":null")
		first = false
	}

	// handle nullable field next_hop_type
	if m.NextHopType != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"next_hop_type\":")
		bytes, err := swag.WriteJSON(m.NextHopType)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.NextHopType_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"next_hop_type\":null")
		first = false
	}

	b.WriteString("}")
	return b.Bytes(), nil
}

// Validate validates this virtual private cloud route params
func (m *VirtualPrivateCloudRouteParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDestination(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNextHopLocalID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNextHopType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VirtualPrivateCloudRouteParams) validateDestination(formats strfmt.Registry) error {

	if err := validate.Required("destination", "body", m.Destination); err != nil {
		return err
	}

	return nil
}

func (m *VirtualPrivateCloudRouteParams) validateNextHopLocalID(formats strfmt.Registry) error {

	if err := validate.Required("next_hop_local_id", "body", m.NextHopLocalID); err != nil {
		return err
	}

	return nil
}

func (m *VirtualPrivateCloudRouteParams) validateNextHopType(formats strfmt.Registry) error {

	if err := validate.Required("next_hop_type", "body", m.NextHopType); err != nil {
		return err
	}

	if err := validate.Required("next_hop_type", "body", m.NextHopType); err != nil {
		return err
	}

	if m.NextHopType != nil {
		if err := m.NextHopType.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("next_hop_type")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("next_hop_type")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this virtual private cloud route params based on the context it is used
func (m *VirtualPrivateCloudRouteParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateNextHopType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VirtualPrivateCloudRouteParams) contextValidateNextHopType(ctx context.Context, formats strfmt.Registry) error {

	if m.NextHopType != nil {
		if err := m.NextHopType.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("next_hop_type")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("next_hop_type")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *VirtualPrivateCloudRouteParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VirtualPrivateCloudRouteParams) UnmarshalBinary(b []byte) error {
	var res VirtualPrivateCloudRouteParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
