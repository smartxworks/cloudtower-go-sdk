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

// ContentLibraryVMTemplateImportParams content library Vm template import params
//
// swagger:model ContentLibraryVmTemplateImportParams
type ContentLibraryVMTemplateImportParams struct {

	// cluster id
	// Required: true
	ClusterID *string `json:"cluster_id"`

	// cpu cores
	CPUCores *int32 `json:"cpu_cores,omitempty"`

	// cpu sockets
	CPUSockets *int32 `json:"cpu_sockets,omitempty"`

	// description
	Description *string `json:"description,omitempty"`

	// disk operate
	DiskOperate *ContentLibraryVMTemplateOvfDiskOperate `json:"disk_operate,omitempty"`

	// ha
	Ha *bool `json:"ha,omitempty"`

	// memory
	Memory *int64 `json:"memory,omitempty"`

	// memory unit
	MemoryUnit *ByteUnit `json:"memory_unit,omitempty"`

	// name
	// Required: true
	Name *string `json:"name"`

	// parsed ovf
	// Required: true
	ParsedOvf *ParsedOVF `json:"parsed_ovf"`

	// upload tasks
	// Required: true
	UploadTasks []string `json:"upload_tasks"`

	// vcpu
	Vcpu *int32 `json:"vcpu,omitempty"`

	// vm nics
	VMNics []*ContentLibraryImportVMNic `json:"vm_nics,omitempty"`

	MarshalOpts *ContentLibraryVMTemplateImportParamsMarshalOpts `json:"-"`
}

type ContentLibraryVMTemplateImportParamsMarshalOpts struct {
	ClusterID_Explicit_Null_When_Empty bool

	CPUCores_Explicit_Null_When_Empty bool

	CPUSockets_Explicit_Null_When_Empty bool

	Description_Explicit_Null_When_Empty bool

	DiskOperate_Explicit_Null_When_Empty bool

	Ha_Explicit_Null_When_Empty bool

	Memory_Explicit_Null_When_Empty bool

	MemoryUnit_Explicit_Null_When_Empty bool

	Name_Explicit_Null_When_Empty bool

	ParsedOvf_Explicit_Null_When_Empty bool

	UploadTasks_Explicit_Null_When_Empty bool

	Vcpu_Explicit_Null_When_Empty bool

	VMNics_Explicit_Null_When_Empty bool
}

func (m ContentLibraryVMTemplateImportParams) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	b.WriteString("{")

	first := true

	// handle nullable field cluster_id
	if m.ClusterID != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"cluster_id\":")
		bytes, err := swag.WriteJSON(m.ClusterID)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.ClusterID_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"cluster_id\":null")
		first = false
	}

	// handle nullable field cpu_cores
	if m.CPUCores != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"cpu_cores\":")
		bytes, err := swag.WriteJSON(m.CPUCores)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.CPUCores_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"cpu_cores\":null")
		first = false
	}

	// handle nullable field cpu_sockets
	if m.CPUSockets != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"cpu_sockets\":")
		bytes, err := swag.WriteJSON(m.CPUSockets)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.CPUSockets_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"cpu_sockets\":null")
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

	// handle nullable field disk_operate
	if m.DiskOperate != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"disk_operate\":")
		bytes, err := swag.WriteJSON(m.DiskOperate)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.DiskOperate_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"disk_operate\":null")
		first = false
	}

	// handle nullable field ha
	if m.Ha != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"ha\":")
		bytes, err := swag.WriteJSON(m.Ha)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.Ha_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"ha\":null")
		first = false
	}

	// handle nullable field memory
	if m.Memory != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"memory\":")
		bytes, err := swag.WriteJSON(m.Memory)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.Memory_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"memory\":null")
		first = false
	}

	// handle nullable field memory_unit
	if m.MemoryUnit != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"memory_unit\":")
		bytes, err := swag.WriteJSON(m.MemoryUnit)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.MemoryUnit_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"memory_unit\":null")
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

	// handle nullable field parsed_ovf
	if m.ParsedOvf != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"parsed_ovf\":")
		bytes, err := swag.WriteJSON(m.ParsedOvf)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.ParsedOvf_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"parsed_ovf\":null")
		first = false
	}

	// handle non nullable field upload_tasks without omitempty
	{
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"upload_tasks\":")
		bytes, err := swag.WriteJSON(m.UploadTasks)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	}

	// handle nullable field vcpu
	if m.Vcpu != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"vcpu\":")
		bytes, err := swag.WriteJSON(m.Vcpu)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.Vcpu_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"vcpu\":null")
		first = false
	}

	// handle non nullable field vm_nics with omitempty
	if swag.IsZero(m.VMNics) {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"vm_nics\":")
		bytes, err := swag.WriteJSON(m.VMNics)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	}

	b.WriteString("}")
	return b.Bytes(), nil
}

// Validate validates this content library Vm template import params
func (m *ContentLibraryVMTemplateImportParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClusterID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDiskOperate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMemoryUnit(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateParsedOvf(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUploadTasks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVMNics(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ContentLibraryVMTemplateImportParams) validateClusterID(formats strfmt.Registry) error {

	if err := validate.Required("cluster_id", "body", m.ClusterID); err != nil {
		return err
	}

	return nil
}

func (m *ContentLibraryVMTemplateImportParams) validateDiskOperate(formats strfmt.Registry) error {
	if swag.IsZero(m.DiskOperate) { // not required
		return nil
	}

	if m.DiskOperate != nil {
		if err := m.DiskOperate.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("disk_operate")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("disk_operate")
			}
			return err
		}
	}

	return nil
}

func (m *ContentLibraryVMTemplateImportParams) validateMemoryUnit(formats strfmt.Registry) error {
	if swag.IsZero(m.MemoryUnit) { // not required
		return nil
	}

	if m.MemoryUnit != nil {
		if err := m.MemoryUnit.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("memory_unit")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("memory_unit")
			}
			return err
		}
	}

	return nil
}

func (m *ContentLibraryVMTemplateImportParams) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *ContentLibraryVMTemplateImportParams) validateParsedOvf(formats strfmt.Registry) error {

	if err := validate.Required("parsed_ovf", "body", m.ParsedOvf); err != nil {
		return err
	}

	if m.ParsedOvf != nil {
		if err := m.ParsedOvf.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("parsed_ovf")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("parsed_ovf")
			}
			return err
		}
	}

	return nil
}

func (m *ContentLibraryVMTemplateImportParams) validateUploadTasks(formats strfmt.Registry) error {

	if err := validate.Required("upload_tasks", "body", m.UploadTasks); err != nil {
		return err
	}

	return nil
}

func (m *ContentLibraryVMTemplateImportParams) validateVMNics(formats strfmt.Registry) error {
	if swag.IsZero(m.VMNics) { // not required
		return nil
	}

	for i := 0; i < len(m.VMNics); i++ {
		if swag.IsZero(m.VMNics[i]) { // not required
			continue
		}

		if m.VMNics[i] != nil {
			if err := m.VMNics[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("vm_nics" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("vm_nics" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this content library Vm template import params based on the context it is used
func (m *ContentLibraryVMTemplateImportParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDiskOperate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMemoryUnit(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateParsedOvf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVMNics(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ContentLibraryVMTemplateImportParams) contextValidateDiskOperate(ctx context.Context, formats strfmt.Registry) error {

	if m.DiskOperate != nil {
		if err := m.DiskOperate.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("disk_operate")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("disk_operate")
			}
			return err
		}
	}

	return nil
}

func (m *ContentLibraryVMTemplateImportParams) contextValidateMemoryUnit(ctx context.Context, formats strfmt.Registry) error {

	if m.MemoryUnit != nil {
		if err := m.MemoryUnit.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("memory_unit")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("memory_unit")
			}
			return err
		}
	}

	return nil
}

func (m *ContentLibraryVMTemplateImportParams) contextValidateParsedOvf(ctx context.Context, formats strfmt.Registry) error {

	if m.ParsedOvf != nil {
		if err := m.ParsedOvf.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("parsed_ovf")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("parsed_ovf")
			}
			return err
		}
	}

	return nil
}

func (m *ContentLibraryVMTemplateImportParams) contextValidateVMNics(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.VMNics); i++ {

		if m.VMNics[i] != nil {
			if err := m.VMNics[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("vm_nics" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("vm_nics" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ContentLibraryVMTemplateImportParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ContentLibraryVMTemplateImportParams) UnmarshalBinary(b []byte) error {
	var res ContentLibraryVMTemplateImportParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
