// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
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
