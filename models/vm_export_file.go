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

// VMExportFile Vm export file
//
// swagger:model VmExportFile
type VMExportFile struct {

	// content library vm template
	ContentLibraryVMTemplate *NestedContentLibraryVMTemplate `json:"content_library_vm_template,omitempty"`

	// created at
	// Required: true
	CreatedAt *string `json:"createdAt"`

	// damaged
	// Required: true
	Damaged *bool `json:"damaged"`

	// data port id
	// Required: true
	DataPortID *string `json:"data_port_id"`

	// entity async status
	EntityAsyncStatus *EntityAsyncStatus `json:"entityAsyncStatus,omitempty"`

	// files
	// Required: true
	Files []*VMExportFileFile `json:"files"`

	// id
	// Required: true
	ID *string `json:"id"`

	// storage cluster id
	// Required: true
	StorageClusterID *string `json:"storage_cluster_id"`

	// type
	// Required: true
	Type *VMExportFileType `json:"type"`

	// vm
	VM *NestedVM `json:"vm,omitempty"`

	// vm volume
	VMVolume *NestedVMVolume `json:"vm_volume,omitempty"`

	MarshalOpts *VMExportFileMarshalOpts `json:"-"`
}

type VMExportFileMarshalOpts struct {
	ContentLibraryVMTemplate_Explicit_Null_When_Empty bool

	CreatedAt_Explicit_Null_When_Empty bool

	Damaged_Explicit_Null_When_Empty bool

	DataPortID_Explicit_Null_When_Empty bool

	EntityAsyncStatus_Explicit_Null_When_Empty bool

	Files_Explicit_Null_When_Empty bool

	ID_Explicit_Null_When_Empty bool

	StorageClusterID_Explicit_Null_When_Empty bool

	Type_Explicit_Null_When_Empty bool

	VM_Explicit_Null_When_Empty bool

	VMVolume_Explicit_Null_When_Empty bool
}

func (m VMExportFile) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	b.WriteString("{")

	first := true

	// handle nullable field content_library_vm_template
	if m.ContentLibraryVMTemplate != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"content_library_vm_template\":")
		bytes, err := swag.WriteJSON(m.ContentLibraryVMTemplate)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.ContentLibraryVMTemplate_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"content_library_vm_template\":null")
		first = false
	}

	// handle nullable field createdAt
	if m.CreatedAt != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"createdAt\":")
		bytes, err := swag.WriteJSON(m.CreatedAt)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.CreatedAt_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"createdAt\":null")
		first = false
	}

	// handle nullable field damaged
	if m.Damaged != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"damaged\":")
		bytes, err := swag.WriteJSON(m.Damaged)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.Damaged_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"damaged\":null")
		first = false
	}

	// handle nullable field data_port_id
	if m.DataPortID != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"data_port_id\":")
		bytes, err := swag.WriteJSON(m.DataPortID)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.DataPortID_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"data_port_id\":null")
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

	// handle non nullable field files without omitempty
	{
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"files\":")
		bytes, err := swag.WriteJSON(m.Files)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
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

	// handle nullable field storage_cluster_id
	if m.StorageClusterID != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"storage_cluster_id\":")
		bytes, err := swag.WriteJSON(m.StorageClusterID)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.StorageClusterID_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"storage_cluster_id\":null")
		first = false
	}

	// handle nullable field type
	if m.Type != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"type\":")
		bytes, err := swag.WriteJSON(m.Type)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.Type_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"type\":null")
		first = false
	}

	// handle nullable field vm
	if m.VM != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"vm\":")
		bytes, err := swag.WriteJSON(m.VM)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.VM_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"vm\":null")
		first = false
	}

	// handle nullable field vm_volume
	if m.VMVolume != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"vm_volume\":")
		bytes, err := swag.WriteJSON(m.VMVolume)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.VMVolume_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"vm_volume\":null")
		first = false
	}

	b.WriteString("}")
	return b.Bytes(), nil
}

// Validate validates this Vm export file
func (m *VMExportFile) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateContentLibraryVMTemplate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDamaged(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDataPortID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEntityAsyncStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFiles(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStorageClusterID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVM(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVMVolume(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VMExportFile) validateContentLibraryVMTemplate(formats strfmt.Registry) error {
	if swag.IsZero(m.ContentLibraryVMTemplate) { // not required
		return nil
	}

	if m.ContentLibraryVMTemplate != nil {
		if err := m.ContentLibraryVMTemplate.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("content_library_vm_template")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("content_library_vm_template")
			}
			return err
		}
	}

	return nil
}

func (m *VMExportFile) validateCreatedAt(formats strfmt.Registry) error {

	if err := validate.Required("createdAt", "body", m.CreatedAt); err != nil {
		return err
	}

	return nil
}

func (m *VMExportFile) validateDamaged(formats strfmt.Registry) error {

	if err := validate.Required("damaged", "body", m.Damaged); err != nil {
		return err
	}

	return nil
}

func (m *VMExportFile) validateDataPortID(formats strfmt.Registry) error {

	if err := validate.Required("data_port_id", "body", m.DataPortID); err != nil {
		return err
	}

	return nil
}

func (m *VMExportFile) validateEntityAsyncStatus(formats strfmt.Registry) error {
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

func (m *VMExportFile) validateFiles(formats strfmt.Registry) error {

	if err := validate.Required("files", "body", m.Files); err != nil {
		return err
	}

	for i := 0; i < len(m.Files); i++ {
		if swag.IsZero(m.Files[i]) { // not required
			continue
		}

		if m.Files[i] != nil {
			if err := m.Files[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("files" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("files" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *VMExportFile) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *VMExportFile) validateStorageClusterID(formats strfmt.Registry) error {

	if err := validate.Required("storage_cluster_id", "body", m.StorageClusterID); err != nil {
		return err
	}

	return nil
}

func (m *VMExportFile) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	if m.Type != nil {
		if err := m.Type.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("type")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("type")
			}
			return err
		}
	}

	return nil
}

func (m *VMExportFile) validateVM(formats strfmt.Registry) error {
	if swag.IsZero(m.VM) { // not required
		return nil
	}

	if m.VM != nil {
		if err := m.VM.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vm")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vm")
			}
			return err
		}
	}

	return nil
}

func (m *VMExportFile) validateVMVolume(formats strfmt.Registry) error {
	if swag.IsZero(m.VMVolume) { // not required
		return nil
	}

	if m.VMVolume != nil {
		if err := m.VMVolume.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vm_volume")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vm_volume")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this Vm export file based on the context it is used
func (m *VMExportFile) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateContentLibraryVMTemplate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEntityAsyncStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFiles(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVM(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVMVolume(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VMExportFile) contextValidateContentLibraryVMTemplate(ctx context.Context, formats strfmt.Registry) error {

	if m.ContentLibraryVMTemplate != nil {
		if err := m.ContentLibraryVMTemplate.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("content_library_vm_template")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("content_library_vm_template")
			}
			return err
		}
	}

	return nil
}

func (m *VMExportFile) contextValidateEntityAsyncStatus(ctx context.Context, formats strfmt.Registry) error {

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

func (m *VMExportFile) contextValidateFiles(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Files); i++ {

		if m.Files[i] != nil {
			if err := m.Files[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("files" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("files" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *VMExportFile) contextValidateType(ctx context.Context, formats strfmt.Registry) error {

	if m.Type != nil {
		if err := m.Type.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("type")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("type")
			}
			return err
		}
	}

	return nil
}

func (m *VMExportFile) contextValidateVM(ctx context.Context, formats strfmt.Registry) error {

	if m.VM != nil {
		if err := m.VM.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vm")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vm")
			}
			return err
		}
	}

	return nil
}

func (m *VMExportFile) contextValidateVMVolume(ctx context.Context, formats strfmt.Registry) error {

	if m.VMVolume != nil {
		if err := m.VMVolume.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vm_volume")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vm_volume")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *VMExportFile) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VMExportFile) UnmarshalBinary(b []byte) error {
	var res VMExportFile
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
