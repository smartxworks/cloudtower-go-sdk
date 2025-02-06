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

// VMDiskOperate Vm disk operate
//
// swagger:model VmDiskOperate
type VMDiskOperate struct {

	// modify disks
	ModifyDisks []*DiskOperateModifyDisk `json:"modify_disks,omitempty"`

	// new disks
	NewDisks *VMDiskParams `json:"new_disks,omitempty"`

	// remove disks
	RemoveDisks *VMDiskOperateRemoveDisks `json:"remove_disks,omitempty"`

	MarshalOpts *VMDiskOperateMarshalOpts `json:"-"`
}

type VMDiskOperateMarshalOpts struct {
	ModifyDisks_Explicit_Null_When_Empty bool

	NewDisks_Explicit_Null_When_Empty bool

	RemoveDisks_Explicit_Null_When_Empty bool
}

func (m VMDiskOperate) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	b.WriteString("{")

	first := true

	// handle non nullable field modify_disks with omitempty
	if swag.IsZero(m.ModifyDisks) {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"modify_disks\":")
		bytes, err := swag.WriteJSON(m.ModifyDisks)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	}

	// handle nullable field new_disks
	if m.NewDisks != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"new_disks\":")
		bytes, err := swag.WriteJSON(m.NewDisks)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.NewDisks_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"new_disks\":null")
		first = false
	}

	// handle nullable field remove_disks
	if m.RemoveDisks != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"remove_disks\":")
		bytes, err := swag.WriteJSON(m.RemoveDisks)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.RemoveDisks_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"remove_disks\":null")
		first = false
	}

	b.WriteString("}")
	return b.Bytes(), nil
}

// Validate validates this Vm disk operate
func (m *VMDiskOperate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateModifyDisks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNewDisks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRemoveDisks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VMDiskOperate) validateModifyDisks(formats strfmt.Registry) error {
	if swag.IsZero(m.ModifyDisks) { // not required
		return nil
	}

	for i := 0; i < len(m.ModifyDisks); i++ {
		if swag.IsZero(m.ModifyDisks[i]) { // not required
			continue
		}

		if m.ModifyDisks[i] != nil {
			if err := m.ModifyDisks[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("modify_disks" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("modify_disks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *VMDiskOperate) validateNewDisks(formats strfmt.Registry) error {
	if swag.IsZero(m.NewDisks) { // not required
		return nil
	}

	if m.NewDisks != nil {
		if err := m.NewDisks.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("new_disks")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("new_disks")
			}
			return err
		}
	}

	return nil
}

func (m *VMDiskOperate) validateRemoveDisks(formats strfmt.Registry) error {
	if swag.IsZero(m.RemoveDisks) { // not required
		return nil
	}

	if m.RemoveDisks != nil {
		if err := m.RemoveDisks.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("remove_disks")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("remove_disks")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this Vm disk operate based on the context it is used
func (m *VMDiskOperate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateModifyDisks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNewDisks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRemoveDisks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VMDiskOperate) contextValidateModifyDisks(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ModifyDisks); i++ {

		if m.ModifyDisks[i] != nil {
			if err := m.ModifyDisks[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("modify_disks" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("modify_disks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *VMDiskOperate) contextValidateNewDisks(ctx context.Context, formats strfmt.Registry) error {

	if m.NewDisks != nil {
		if err := m.NewDisks.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("new_disks")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("new_disks")
			}
			return err
		}
	}

	return nil
}

func (m *VMDiskOperate) contextValidateRemoveDisks(ctx context.Context, formats strfmt.Registry) error {

	if m.RemoveDisks != nil {
		if err := m.RemoveDisks.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("remove_disks")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("remove_disks")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *VMDiskOperate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VMDiskOperate) UnmarshalBinary(b []byte) error {
	var res VMDiskOperate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// VMDiskOperateRemoveDisks VM disk operate remove disks
//
// swagger:model VMDiskOperateRemoveDisks
type VMDiskOperateRemoveDisks struct {

	// disk index
	// Required: true
	DiskIndex []int32 `json:"disk_index"`

	MarshalOpts *VMDiskOperateRemoveDisksMarshalOpts `json:"-"`
}

type VMDiskOperateRemoveDisksMarshalOpts struct {
	DiskIndex_Explicit_Null_When_Empty bool
}

func (m VMDiskOperateRemoveDisks) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	b.WriteString("{")

	first := true

	// handle non nullable field disk_index without omitempty
	if !first {
		b.WriteString(",")
	}
	b.WriteString("\"disk_index\":")
	{
		bytes, err := swag.WriteJSON(m.DiskIndex)
		if err != nil {
			return nil, err
		}
	}
	b.Write(bytes)
	first = false

	b.WriteString("}")
	return b.Bytes(), nil
}

// Validate validates this VM disk operate remove disks
func (m *VMDiskOperateRemoveDisks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDiskIndex(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VMDiskOperateRemoveDisks) validateDiskIndex(formats strfmt.Registry) error {

	if err := validate.Required("remove_disks"+"."+"disk_index", "body", m.DiskIndex); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this VM disk operate remove disks based on context it is used
func (m *VMDiskOperateRemoveDisks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *VMDiskOperateRemoveDisks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VMDiskOperateRemoveDisks) UnmarshalBinary(b []byte) error {
	var res VMDiskOperateRemoveDisks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
