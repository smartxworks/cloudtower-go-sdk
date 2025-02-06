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

// VMAddDiskParams Vm add disk params
//
// swagger:model VmAddDiskParams
type VMAddDiskParams struct {

	// data
	// Required: true
	Data *VMAddDiskParamsData `json:"data"`

	// where
	// Required: true
	Where *VMWhereInput `json:"where"`

	MarshalOpts *VMAddDiskParamsMarshalOpts `json:"-"`
}

type VMAddDiskParamsMarshalOpts struct {
	Data_Explicit_Null_When_Empty bool

	Where_Explicit_Null_When_Empty bool
}

func (m VMAddDiskParams) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	b.WriteString("{")

	first := true

	// handle nullable field data
	if m.Data != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"data\":")
		bytes, err := swag.WriteJSON(m.Data)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.Data_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"data\":null")
		first = false
	}

	// handle nullable field where
	if m.Where != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"where\":")
		bytes, err := swag.WriteJSON(m.Where)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.Where_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"where\":null")
		first = false
	}

	b.WriteString("}")
	return b.Bytes(), nil
}

// Validate validates this Vm add disk params
func (m *VMAddDiskParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWhere(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VMAddDiskParams) validateData(formats strfmt.Registry) error {

	if err := validate.Required("data", "body", m.Data); err != nil {
		return err
	}

	if m.Data != nil {
		if err := m.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("data")
			}
			return err
		}
	}

	return nil
}

func (m *VMAddDiskParams) validateWhere(formats strfmt.Registry) error {

	if err := validate.Required("where", "body", m.Where); err != nil {
		return err
	}

	if m.Where != nil {
		if err := m.Where.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("where")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("where")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this Vm add disk params based on the context it is used
func (m *VMAddDiskParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateWhere(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VMAddDiskParams) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if m.Data != nil {
		if err := m.Data.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("data")
			}
			return err
		}
	}

	return nil
}

func (m *VMAddDiskParams) contextValidateWhere(ctx context.Context, formats strfmt.Registry) error {

	if m.Where != nil {
		if err := m.Where.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("where")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("where")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *VMAddDiskParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VMAddDiskParams) UnmarshalBinary(b []byte) error {
	var res VMAddDiskParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// VMAddDiskParamsData VM add disk params data
//
// swagger:model VMAddDiskParamsData
type VMAddDiskParamsData struct {

	// io policy
	IoPolicy *VMDiskIoPolicy `json:"io_policy,omitempty"`

	// max bandwidth
	MaxBandwidth *int64 `json:"max_bandwidth,omitempty"`

	// max bandwidth policy
	MaxBandwidthPolicy *VMDiskIoRestrictType `json:"max_bandwidth_policy,omitempty"`

	// max bandwidth unit
	MaxBandwidthUnit *BPSUnit `json:"max_bandwidth_unit,omitempty"`

	// max iops
	MaxIops *int64 `json:"max_iops,omitempty"`

	// max iops policy
	MaxIopsPolicy *VMDiskIoRestrictType `json:"max_iops_policy,omitempty"`

	// vm disks
	// Required: true
	VMDisks *VMAddDiskParamsDataVMDisks `json:"vm_disks"`

	MarshalOpts *VMAddDiskParamsDataMarshalOpts `json:"-"`
}

type VMAddDiskParamsDataMarshalOpts struct {
	IoPolicy_Explicit_Null_When_Empty bool

	MaxBandwidth_Explicit_Null_When_Empty bool

	MaxBandwidthPolicy_Explicit_Null_When_Empty bool

	MaxBandwidthUnit_Explicit_Null_When_Empty bool

	MaxIops_Explicit_Null_When_Empty bool

	MaxIopsPolicy_Explicit_Null_When_Empty bool

	VMDisks_Explicit_Null_When_Empty bool
}

func (m VMAddDiskParamsData) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	b.WriteString("{")

	first := true

	// handle nullable field io_policy
	if m.IoPolicy != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"io_policy\":")
		bytes, err := swag.WriteJSON(m.IoPolicy)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.IoPolicy_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"io_policy\":null")
		first = false
	}

	// handle nullable field max_bandwidth
	if m.MaxBandwidth != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"max_bandwidth\":")
		bytes, err := swag.WriteJSON(m.MaxBandwidth)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.MaxBandwidth_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"max_bandwidth\":null")
		first = false
	}

	// handle nullable field max_bandwidth_policy
	if m.MaxBandwidthPolicy != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"max_bandwidth_policy\":")
		bytes, err := swag.WriteJSON(m.MaxBandwidthPolicy)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.MaxBandwidthPolicy_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"max_bandwidth_policy\":null")
		first = false
	}

	// handle nullable field max_bandwidth_unit
	if m.MaxBandwidthUnit != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"max_bandwidth_unit\":")
		bytes, err := swag.WriteJSON(m.MaxBandwidthUnit)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.MaxBandwidthUnit_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"max_bandwidth_unit\":null")
		first = false
	}

	// handle nullable field max_iops
	if m.MaxIops != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"max_iops\":")
		bytes, err := swag.WriteJSON(m.MaxIops)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.MaxIops_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"max_iops\":null")
		first = false
	}

	// handle nullable field max_iops_policy
	if m.MaxIopsPolicy != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"max_iops_policy\":")
		bytes, err := swag.WriteJSON(m.MaxIopsPolicy)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.MaxIopsPolicy_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"max_iops_policy\":null")
		first = false
	}

	// handle nullable field vm_disks
	if m.VMDisks != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"vm_disks\":")
		bytes, err := swag.WriteJSON(m.VMDisks)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.VMDisks_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"vm_disks\":null")
		first = false
	}

	b.WriteString("}")
	return b.Bytes(), nil
}

// Validate validates this VM add disk params data
func (m *VMAddDiskParamsData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIoPolicy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMaxBandwidthPolicy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMaxBandwidthUnit(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMaxIopsPolicy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVMDisks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VMAddDiskParamsData) validateIoPolicy(formats strfmt.Registry) error {
	if swag.IsZero(m.IoPolicy) { // not required
		return nil
	}

	if m.IoPolicy != nil {
		if err := m.IoPolicy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data" + "." + "io_policy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("data" + "." + "io_policy")
			}
			return err
		}
	}

	return nil
}

func (m *VMAddDiskParamsData) validateMaxBandwidthPolicy(formats strfmt.Registry) error {
	if swag.IsZero(m.MaxBandwidthPolicy) { // not required
		return nil
	}

	if m.MaxBandwidthPolicy != nil {
		if err := m.MaxBandwidthPolicy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data" + "." + "max_bandwidth_policy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("data" + "." + "max_bandwidth_policy")
			}
			return err
		}
	}

	return nil
}

func (m *VMAddDiskParamsData) validateMaxBandwidthUnit(formats strfmt.Registry) error {
	if swag.IsZero(m.MaxBandwidthUnit) { // not required
		return nil
	}

	if m.MaxBandwidthUnit != nil {
		if err := m.MaxBandwidthUnit.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data" + "." + "max_bandwidth_unit")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("data" + "." + "max_bandwidth_unit")
			}
			return err
		}
	}

	return nil
}

func (m *VMAddDiskParamsData) validateMaxIopsPolicy(formats strfmt.Registry) error {
	if swag.IsZero(m.MaxIopsPolicy) { // not required
		return nil
	}

	if m.MaxIopsPolicy != nil {
		if err := m.MaxIopsPolicy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data" + "." + "max_iops_policy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("data" + "." + "max_iops_policy")
			}
			return err
		}
	}

	return nil
}

func (m *VMAddDiskParamsData) validateVMDisks(formats strfmt.Registry) error {

	if err := validate.Required("data"+"."+"vm_disks", "body", m.VMDisks); err != nil {
		return err
	}

	if m.VMDisks != nil {
		if err := m.VMDisks.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data" + "." + "vm_disks")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("data" + "." + "vm_disks")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this VM add disk params data based on the context it is used
func (m *VMAddDiskParamsData) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateIoPolicy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMaxBandwidthPolicy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMaxBandwidthUnit(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMaxIopsPolicy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVMDisks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VMAddDiskParamsData) contextValidateIoPolicy(ctx context.Context, formats strfmt.Registry) error {

	if m.IoPolicy != nil {
		if err := m.IoPolicy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data" + "." + "io_policy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("data" + "." + "io_policy")
			}
			return err
		}
	}

	return nil
}

func (m *VMAddDiskParamsData) contextValidateMaxBandwidthPolicy(ctx context.Context, formats strfmt.Registry) error {

	if m.MaxBandwidthPolicy != nil {
		if err := m.MaxBandwidthPolicy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data" + "." + "max_bandwidth_policy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("data" + "." + "max_bandwidth_policy")
			}
			return err
		}
	}

	return nil
}

func (m *VMAddDiskParamsData) contextValidateMaxBandwidthUnit(ctx context.Context, formats strfmt.Registry) error {

	if m.MaxBandwidthUnit != nil {
		if err := m.MaxBandwidthUnit.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data" + "." + "max_bandwidth_unit")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("data" + "." + "max_bandwidth_unit")
			}
			return err
		}
	}

	return nil
}

func (m *VMAddDiskParamsData) contextValidateMaxIopsPolicy(ctx context.Context, formats strfmt.Registry) error {

	if m.MaxIopsPolicy != nil {
		if err := m.MaxIopsPolicy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data" + "." + "max_iops_policy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("data" + "." + "max_iops_policy")
			}
			return err
		}
	}

	return nil
}

func (m *VMAddDiskParamsData) contextValidateVMDisks(ctx context.Context, formats strfmt.Registry) error {

	if m.VMDisks != nil {
		if err := m.VMDisks.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data" + "." + "vm_disks")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("data" + "." + "vm_disks")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *VMAddDiskParamsData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VMAddDiskParamsData) UnmarshalBinary(b []byte) error {
	var res VMAddDiskParamsData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// VMAddDiskParamsDataVMDisks VM add disk params data VM disks
//
// swagger:model VMAddDiskParamsDataVMDisks
type VMAddDiskParamsDataVMDisks struct {

	// mount disks
	MountDisks []*MountDisksParams `json:"mount_disks,omitempty"`

	// mount new create disks
	MountNewCreateDisks []*MountNewCreateDisksParams `json:"mount_new_create_disks,omitempty"`

	MarshalOpts *VMAddDiskParamsDataVMDisksMarshalOpts `json:"-"`
}

type VMAddDiskParamsDataVMDisksMarshalOpts struct {
	MountDisks_Explicit_Null_When_Empty bool

	MountNewCreateDisks_Explicit_Null_When_Empty bool
}

func (m VMAddDiskParamsDataVMDisks) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	b.WriteString("{")

	first := true

	// handle non nullable field mount_disks with omitempty
	if swag.IsZero(m.MountDisks) {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"mount_disks\":")
		bytes, err := swag.WriteJSON(m.MountDisks)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	}

	// handle non nullable field mount_new_create_disks with omitempty
	if swag.IsZero(m.MountNewCreateDisks) {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"mount_new_create_disks\":")
		bytes, err := swag.WriteJSON(m.MountNewCreateDisks)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	}

	b.WriteString("}")
	return b.Bytes(), nil
}

// Validate validates this VM add disk params data VM disks
func (m *VMAddDiskParamsDataVMDisks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMountDisks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMountNewCreateDisks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VMAddDiskParamsDataVMDisks) validateMountDisks(formats strfmt.Registry) error {
	if swag.IsZero(m.MountDisks) { // not required
		return nil
	}

	for i := 0; i < len(m.MountDisks); i++ {
		if swag.IsZero(m.MountDisks[i]) { // not required
			continue
		}

		if m.MountDisks[i] != nil {
			if err := m.MountDisks[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("data" + "." + "vm_disks" + "." + "mount_disks" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("data" + "." + "vm_disks" + "." + "mount_disks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *VMAddDiskParamsDataVMDisks) validateMountNewCreateDisks(formats strfmt.Registry) error {
	if swag.IsZero(m.MountNewCreateDisks) { // not required
		return nil
	}

	for i := 0; i < len(m.MountNewCreateDisks); i++ {
		if swag.IsZero(m.MountNewCreateDisks[i]) { // not required
			continue
		}

		if m.MountNewCreateDisks[i] != nil {
			if err := m.MountNewCreateDisks[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("data" + "." + "vm_disks" + "." + "mount_new_create_disks" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("data" + "." + "vm_disks" + "." + "mount_new_create_disks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this VM add disk params data VM disks based on the context it is used
func (m *VMAddDiskParamsDataVMDisks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateMountDisks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMountNewCreateDisks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VMAddDiskParamsDataVMDisks) contextValidateMountDisks(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.MountDisks); i++ {

		if m.MountDisks[i] != nil {
			if err := m.MountDisks[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("data" + "." + "vm_disks" + "." + "mount_disks" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("data" + "." + "vm_disks" + "." + "mount_disks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *VMAddDiskParamsDataVMDisks) contextValidateMountNewCreateDisks(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.MountNewCreateDisks); i++ {

		if m.MountNewCreateDisks[i] != nil {
			if err := m.MountNewCreateDisks[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("data" + "." + "vm_disks" + "." + "mount_new_create_disks" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("data" + "." + "vm_disks" + "." + "mount_new_create_disks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *VMAddDiskParamsDataVMDisks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VMAddDiskParamsDataVMDisks) UnmarshalBinary(b []byte) error {
	var res VMAddDiskParamsDataVMDisks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
