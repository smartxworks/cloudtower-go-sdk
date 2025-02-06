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

// VirtualPrivateCloudSecurityGroup virtual private cloud security group
//
// swagger:model VirtualPrivateCloudSecurityGroup
type VirtualPrivateCloudSecurityGroup struct {

	// default for vpc
	DefaultForVpc *bool `json:"default_for_vpc,omitempty"`

	// description
	Description *string `json:"description,omitempty"`

	// entity async status
	EntityAsyncStatus *EntityAsyncStatus `json:"entityAsyncStatus,omitempty"`

	// id
	// Required: true
	ID *string `json:"id"`

	// label groups
	LabelGroups []*NestedVirtualPrivateCloudLabelGroup `json:"label_groups,omitempty"`

	// local id
	// Required: true
	LocalID *string `json:"local_id"`

	// name
	// Required: true
	Name *string `json:"name"`

	// vms
	Vms []*NestedVM `json:"vms,omitempty"`

	// vpc
	// Required: true
	Vpc *NestedVirtualPrivateCloud `json:"vpc"`

	MarshalOpts *VirtualPrivateCloudSecurityGroupMarshalOpts `json:"-"`
}

type VirtualPrivateCloudSecurityGroupMarshalOpts struct {
	DefaultForVpc_Explicit_Null_When_Empty bool

	Description_Explicit_Null_When_Empty bool

	EntityAsyncStatus_Explicit_Null_When_Empty bool

	ID_Explicit_Null_When_Empty bool

	LocalID_Explicit_Null_When_Empty bool

	Name_Explicit_Null_When_Empty bool

	Vpc_Explicit_Null_When_Empty bool
}

func (m VirtualPrivateCloudSecurityGroup) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	b.WriteString("{")

	first := true

	// handle nullable field default_for_vpc
	if m.DefaultForVpc != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"default_for_vpc\":")
		bytes, err := swag.WriteJSON(m.DefaultForVpc)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.DefaultForVpc_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"default_for_vpc\":null")
		first = false
	}

	// handle nullable field description
	if m.Description != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"description\":")
		bytes, err := swag.WriteJSON(m.Description)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.Description_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"description\":null")
		first = false
	}

	// handle nullable field entityAsyncStatus
	if m.EntityAsyncStatus != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"entityAsyncStatus\":")
		bytes, err := swag.WriteJSON(m.EntityAsyncStatus)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.EntityAsyncStatus_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"entityAsyncStatus\":null")
		first = false
	}

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

	// handle non nullable field label_groups with omitempty
	if swag.IsZero(m.LabelGroups) {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"label_groups\":")
		bytes, err := swag.WriteJSON(m.LabelGroups)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	}

	// handle nullable field local_id
	if m.LocalID != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"local_id\":")
		bytes, err := swag.WriteJSON(m.LocalID)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.LocalID_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"local_id\":null")
		first = false
	}

	// handle nullable field name
	if m.Name != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"name\":")
		bytes, err := swag.WriteJSON(m.Name)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.Name_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"name\":null")
		first = false
	}

	// handle non nullable field vms with omitempty
	if swag.IsZero(m.Vms) {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"vms\":")
		bytes, err := swag.WriteJSON(m.Vms)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	}

	// handle nullable field vpc
	if m.Vpc != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"vpc\":")
		bytes, err := swag.WriteJSON(m.Vpc)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.Vpc_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"vpc\":null")
		first = false
	}

	b.WriteString("}")
	return b.Bytes(), nil
}

// Validate validates this virtual private cloud security group
func (m *VirtualPrivateCloudSecurityGroup) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEntityAsyncStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLabelGroups(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocalID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVms(formats); err != nil {
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

func (m *VirtualPrivateCloudSecurityGroup) validateEntityAsyncStatus(formats strfmt.Registry) error {
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

func (m *VirtualPrivateCloudSecurityGroup) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *VirtualPrivateCloudSecurityGroup) validateLabelGroups(formats strfmt.Registry) error {
	if swag.IsZero(m.LabelGroups) { // not required
		return nil
	}

	for i := 0; i < len(m.LabelGroups); i++ {
		if swag.IsZero(m.LabelGroups[i]) { // not required
			continue
		}

		if m.LabelGroups[i] != nil {
			if err := m.LabelGroups[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("label_groups" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("label_groups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *VirtualPrivateCloudSecurityGroup) validateLocalID(formats strfmt.Registry) error {

	if err := validate.Required("local_id", "body", m.LocalID); err != nil {
		return err
	}

	return nil
}

func (m *VirtualPrivateCloudSecurityGroup) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *VirtualPrivateCloudSecurityGroup) validateVms(formats strfmt.Registry) error {
	if swag.IsZero(m.Vms) { // not required
		return nil
	}

	for i := 0; i < len(m.Vms); i++ {
		if swag.IsZero(m.Vms[i]) { // not required
			continue
		}

		if m.Vms[i] != nil {
			if err := m.Vms[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("vms" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("vms" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *VirtualPrivateCloudSecurityGroup) validateVpc(formats strfmt.Registry) error {

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

// ContextValidate validate this virtual private cloud security group based on the context it is used
func (m *VirtualPrivateCloudSecurityGroup) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateEntityAsyncStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLabelGroups(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVms(ctx, formats); err != nil {
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

func (m *VirtualPrivateCloudSecurityGroup) contextValidateEntityAsyncStatus(ctx context.Context, formats strfmt.Registry) error {

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

func (m *VirtualPrivateCloudSecurityGroup) contextValidateLabelGroups(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.LabelGroups); i++ {

		if m.LabelGroups[i] != nil {
			if err := m.LabelGroups[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("label_groups" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("label_groups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *VirtualPrivateCloudSecurityGroup) contextValidateVms(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Vms); i++ {

		if m.Vms[i] != nil {
			if err := m.Vms[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("vms" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("vms" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *VirtualPrivateCloudSecurityGroup) contextValidateVpc(ctx context.Context, formats strfmt.Registry) error {

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
func (m *VirtualPrivateCloudSecurityGroup) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VirtualPrivateCloudSecurityGroup) UnmarshalBinary(b []byte) error {
	var res VirtualPrivateCloudSecurityGroup
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
