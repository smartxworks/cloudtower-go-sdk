// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NestedVirtualPrivateCloudService nested virtual private cloud service
//
// swagger:model NestedVirtualPrivateCloudService
type NestedVirtualPrivateCloudService struct {

	// id
	// Required: true
	ID *string `json:"id"`

	// internal cidr
	// Required: true
	InternalCidr *string `json:"internal_cidr"`

	// tep ip pools
	// Required: true
	TepIPPools []*NestedVirtualPrivateCloudServiceTepIPPool `json:"tep_ip_pools"`

	MarshalOpts *NestedVirtualPrivateCloudServiceMarshalOpts `json:"-"`
}

type NestedVirtualPrivateCloudServiceMarshalOpts struct {
	ID_Explicit_Null_When_Empty bool

	InternalCidr_Explicit_Null_When_Empty bool
}

func (m NestedVirtualPrivateCloudService) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	b.WriteString("{")

	first := true

	// handle nullable field id
	if m.ID != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"id\":")
		bytes, err := swag.WriteJSON(m.ID)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.ID_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"id\":null")
		first = false
	}

	// handle nullable field internal_cidr
	if m.InternalCidr != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"internal_cidr\":")
		bytes, err := swag.WriteJSON(m.InternalCidr)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.InternalCidr_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"internal_cidr\":null")
		first = false
	}

	// handle non nullable field tep_ip_pools without omitempty
	if !first {
		b.WriteString(",")
	}
	b.WriteString("\"tep_ip_pools\":")
	bytes, err := swag.WriteJSON(m.TepIPPools)
	if err != nil {
		return nil, err
	}
	b.Write(bytes)
	first = false

	b.WriteString("}")
	return b.Bytes(), nil
}

// Validate validates this nested virtual private cloud service
func (m *NestedVirtualPrivateCloudService) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInternalCidr(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTepIPPools(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NestedVirtualPrivateCloudService) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *NestedVirtualPrivateCloudService) validateInternalCidr(formats strfmt.Registry) error {

	if err := validate.Required("internal_cidr", "body", m.InternalCidr); err != nil {
		return err
	}

	return nil
}

func (m *NestedVirtualPrivateCloudService) validateTepIPPools(formats strfmt.Registry) error {

	if err := validate.Required("tep_ip_pools", "body", m.TepIPPools); err != nil {
		return err
	}

	for i := 0; i < len(m.TepIPPools); i++ {
		if swag.IsZero(m.TepIPPools[i]) { // not required
			continue
		}

		if m.TepIPPools[i] != nil {
			if err := m.TepIPPools[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tep_ip_pools" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("tep_ip_pools" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this nested virtual private cloud service based on the context it is used
func (m *NestedVirtualPrivateCloudService) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateTepIPPools(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NestedVirtualPrivateCloudService) contextValidateTepIPPools(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.TepIPPools); i++ {

		if m.TepIPPools[i] != nil {
			if err := m.TepIPPools[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tep_ip_pools" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("tep_ip_pools" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *NestedVirtualPrivateCloudService) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NestedVirtualPrivateCloudService) UnmarshalBinary(b []byte) error {
	var res NestedVirtualPrivateCloudService
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
