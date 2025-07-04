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
)

// ContentLibraryVMTemplateOvfDiskOperate content library Vm template ovf disk operate
//
// swagger:model ContentLibraryVmTemplateOvfDiskOperate
type ContentLibraryVMTemplateOvfDiskOperate struct {

	// modify cd roms
	ModifyCdRoms []*ContentLibraryVmdkCdromModify `json:"modify_cd_roms,omitempty"`

	// modify vmdk disks
	ModifyVmdkDisks []*VmdkDiskModify `json:"modify_vmdk_disks,omitempty"`

	// mount new cd roms
	MountNewCdRoms []*ContentLibraryVmdkCdromMount `json:"mount_new_cd_roms,omitempty"`

	MarshalOpts *ContentLibraryVMTemplateOvfDiskOperateMarshalOpts `json:"-"`
}

type ContentLibraryVMTemplateOvfDiskOperateMarshalOpts struct {
	ModifyCdRoms_Explicit_Null_When_Empty bool

	ModifyVmdkDisks_Explicit_Null_When_Empty bool

	MountNewCdRoms_Explicit_Null_When_Empty bool
}

func (m ContentLibraryVMTemplateOvfDiskOperate) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	b.WriteString("{")

	first := true

	// handle non nullable field modify_cd_roms with omitempty
	if !swag.IsZero(m.ModifyCdRoms) {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"modify_cd_roms\":")
		bytes, err := swag.WriteJSON(m.ModifyCdRoms)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	}

	// handle non nullable field modify_vmdk_disks with omitempty
	if !swag.IsZero(m.ModifyVmdkDisks) {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"modify_vmdk_disks\":")
		bytes, err := swag.WriteJSON(m.ModifyVmdkDisks)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	}

	// handle non nullable field mount_new_cd_roms with omitempty
	if !swag.IsZero(m.MountNewCdRoms) {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"mount_new_cd_roms\":")
		bytes, err := swag.WriteJSON(m.MountNewCdRoms)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	}

	b.WriteString("}")
	return b.Bytes(), nil
}

// Validate validates this content library Vm template ovf disk operate
func (m *ContentLibraryVMTemplateOvfDiskOperate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateModifyCdRoms(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateModifyVmdkDisks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMountNewCdRoms(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ContentLibraryVMTemplateOvfDiskOperate) validateModifyCdRoms(formats strfmt.Registry) error {
	if swag.IsZero(m.ModifyCdRoms) { // not required
		return nil
	}

	for i := 0; i < len(m.ModifyCdRoms); i++ {
		if swag.IsZero(m.ModifyCdRoms[i]) { // not required
			continue
		}

		if m.ModifyCdRoms[i] != nil {
			if err := m.ModifyCdRoms[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("modify_cd_roms" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("modify_cd_roms" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ContentLibraryVMTemplateOvfDiskOperate) validateModifyVmdkDisks(formats strfmt.Registry) error {
	if swag.IsZero(m.ModifyVmdkDisks) { // not required
		return nil
	}

	for i := 0; i < len(m.ModifyVmdkDisks); i++ {
		if swag.IsZero(m.ModifyVmdkDisks[i]) { // not required
			continue
		}

		if m.ModifyVmdkDisks[i] != nil {
			if err := m.ModifyVmdkDisks[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("modify_vmdk_disks" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("modify_vmdk_disks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ContentLibraryVMTemplateOvfDiskOperate) validateMountNewCdRoms(formats strfmt.Registry) error {
	if swag.IsZero(m.MountNewCdRoms) { // not required
		return nil
	}

	for i := 0; i < len(m.MountNewCdRoms); i++ {
		if swag.IsZero(m.MountNewCdRoms[i]) { // not required
			continue
		}

		if m.MountNewCdRoms[i] != nil {
			if err := m.MountNewCdRoms[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("mount_new_cd_roms" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("mount_new_cd_roms" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this content library Vm template ovf disk operate based on the context it is used
func (m *ContentLibraryVMTemplateOvfDiskOperate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateModifyCdRoms(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateModifyVmdkDisks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMountNewCdRoms(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ContentLibraryVMTemplateOvfDiskOperate) contextValidateModifyCdRoms(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ModifyCdRoms); i++ {

		if m.ModifyCdRoms[i] != nil {
			if err := m.ModifyCdRoms[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("modify_cd_roms" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("modify_cd_roms" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ContentLibraryVMTemplateOvfDiskOperate) contextValidateModifyVmdkDisks(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ModifyVmdkDisks); i++ {

		if m.ModifyVmdkDisks[i] != nil {
			if err := m.ModifyVmdkDisks[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("modify_vmdk_disks" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("modify_vmdk_disks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ContentLibraryVMTemplateOvfDiskOperate) contextValidateMountNewCdRoms(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.MountNewCdRoms); i++ {

		if m.MountNewCdRoms[i] != nil {
			if err := m.MountNewCdRoms[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("mount_new_cd_roms" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("mount_new_cd_roms" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ContentLibraryVMTemplateOvfDiskOperate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ContentLibraryVMTemplateOvfDiskOperate) UnmarshalBinary(b []byte) error {
	var res ContentLibraryVMTemplateOvfDiskOperate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
