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

// NestedSnapshotGroupVMInfo nested snapshot group Vm info
//
// swagger:model NestedSnapshotGroupVmInfo
type NestedSnapshotGroupVMInfo struct {

	// disks
	// Required: true
	Disks []*NestedSnapshotGroupVMDiskInfo `json:"disks"`

	// vm id
	// Required: true
	VMID *string `json:"vm_id"`

	// vm name
	// Required: true
	VMName *string `json:"vm_name"`

	// vm snapshot status
	// Required: true
	VMSnapshotStatus *ProtectSnapshotStatus `json:"vm_snapshot_status"`

	MarshalOpts *NestedSnapshotGroupVMInfoMarshalOpts `json:"-"`
}

type NestedSnapshotGroupVMInfoMarshalOpts struct {
	Disks_Explicit_Null_When_Empty bool

	VMID_Explicit_Null_When_Empty bool

	VMName_Explicit_Null_When_Empty bool

	VMSnapshotStatus_Explicit_Null_When_Empty bool
}

func (m NestedSnapshotGroupVMInfo) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	b.WriteString("{")

	first := true

	// handle non nullable field disks without omitempty
	if !first {
		b.WriteString(",")
	}
	b.WriteString("\"disks\":")
	{
		bytes, err := swag.WriteJSON(m.Disks)
		if err != nil {
			return nil, err
		}
	}
	b.Write(bytes)
	first = false

	// handle nullable field vm_id
	if m.VMID != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"vm_id\":")
		bytes, err := swag.WriteJSON(m.VMID)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.VMID_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"vm_id\":null")
		first = false
	}

	// handle nullable field vm_name
	if m.VMName != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"vm_name\":")
		bytes, err := swag.WriteJSON(m.VMName)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.VMName_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"vm_name\":null")
		first = false
	}

	// handle nullable field vm_snapshot_status
	if m.VMSnapshotStatus != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"vm_snapshot_status\":")
		bytes, err := swag.WriteJSON(m.VMSnapshotStatus)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.VMSnapshotStatus_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"vm_snapshot_status\":null")
		first = false
	}

	b.WriteString("}")
	return b.Bytes(), nil
}

// Validate validates this nested snapshot group Vm info
func (m *NestedSnapshotGroupVMInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDisks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVMID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVMName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVMSnapshotStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NestedSnapshotGroupVMInfo) validateDisks(formats strfmt.Registry) error {

	if err := validate.Required("disks", "body", m.Disks); err != nil {
		return err
	}

	for i := 0; i < len(m.Disks); i++ {
		if swag.IsZero(m.Disks[i]) { // not required
			continue
		}

		if m.Disks[i] != nil {
			if err := m.Disks[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("disks" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("disks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *NestedSnapshotGroupVMInfo) validateVMID(formats strfmt.Registry) error {

	if err := validate.Required("vm_id", "body", m.VMID); err != nil {
		return err
	}

	return nil
}

func (m *NestedSnapshotGroupVMInfo) validateVMName(formats strfmt.Registry) error {

	if err := validate.Required("vm_name", "body", m.VMName); err != nil {
		return err
	}

	return nil
}

func (m *NestedSnapshotGroupVMInfo) validateVMSnapshotStatus(formats strfmt.Registry) error {

	if err := validate.Required("vm_snapshot_status", "body", m.VMSnapshotStatus); err != nil {
		return err
	}

	if err := validate.Required("vm_snapshot_status", "body", m.VMSnapshotStatus); err != nil {
		return err
	}

	if m.VMSnapshotStatus != nil {
		if err := m.VMSnapshotStatus.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vm_snapshot_status")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vm_snapshot_status")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this nested snapshot group Vm info based on the context it is used
func (m *NestedSnapshotGroupVMInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDisks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVMSnapshotStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NestedSnapshotGroupVMInfo) contextValidateDisks(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Disks); i++ {

		if m.Disks[i] != nil {
			if err := m.Disks[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("disks" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("disks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *NestedSnapshotGroupVMInfo) contextValidateVMSnapshotStatus(ctx context.Context, formats strfmt.Registry) error {

	if m.VMSnapshotStatus != nil {
		if err := m.VMSnapshotStatus.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vm_snapshot_status")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vm_snapshot_status")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NestedSnapshotGroupVMInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NestedSnapshotGroupVMInfo) UnmarshalBinary(b []byte) error {
	var res NestedSnapshotGroupVMInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
