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

// VMGpuDetail Vm gpu detail
//
// swagger:model VmGpuDetail
type VMGpuDetail struct {

	// assigned vgpus num
	AssignedVgpusNum *int32 `json:"assigned_vgpus_num,omitempty"`

	// available vgpus num
	AvailableVgpusNum *int32 `json:"available_vgpus_num,omitempty"`

	// brand
	// Required: true
	Brand *string `json:"brand"`

	// bus location
	// Required: true
	BusLocation *string `json:"bus_location"`

	// description
	// Required: true
	Description *string `json:"description"`

	// driver info
	DriverInfo *NestedGpuDriverInfo `json:"driver_info,omitempty"`

	// entity async status
	EntityAsyncStatus *EntityAsyncStatus `json:"entityAsyncStatus,omitempty"`

	// host
	// Required: true
	Host *NestedHost `json:"host"`

	// id
	// Required: true
	ID *string `json:"id"`

	// is nvidia tools ready
	IsNvidiaToolsReady *bool `json:"is_nvidia_tools_ready,omitempty"`

	// is nvidia vfs enabled
	IsNvidiaVfsEnabled *bool `json:"is_nvidia_vfs_enabled,omitempty"`

	// is nvidia vfs supported
	IsNvidiaVfsSupported *bool `json:"is_nvidia_vfs_supported,omitempty"`

	// labels
	Labels []*NestedLabel `json:"labels,omitempty"`

	// local created at
	// Required: true
	LocalCreatedAt *string `json:"local_created_at"`

	// local id
	// Required: true
	LocalID *string `json:"local_id"`

	// mdev supported types
	MdevSupportedTypes []*NestedVgpuType `json:"mdev_supported_types,omitempty"`

	// model
	// Required: true
	Model *string `json:"model"`

	// name
	// Required: true
	Name *string `json:"name"`

	// status
	// Required: true
	Status *GpuDeviceStatus `json:"status"`

	// user usage
	UserUsage *GpuDeviceUsage `json:"user_usage,omitempty"`

	// user vgpu type id
	UserVgpuTypeID *string `json:"user_vgpu_type_id,omitempty"`

	// user vgpu type name
	UserVgpuTypeName *string `json:"user_vgpu_type_name,omitempty"`

	// vgpu instance num
	VgpuInstanceNum *int32 `json:"vgpu_instance_num,omitempty"`

	// vgpu instance on vm num
	VgpuInstanceOnVMNum *int32 `json:"vgpu_instance_on_vm_num,omitempty"`

	MarshalOpts *VMGpuDetailMarshalOpts `json:"-"`
}

type VMGpuDetailMarshalOpts struct {
	AssignedVgpusNum_Explicit_Null_When_Empty bool

	AvailableVgpusNum_Explicit_Null_When_Empty bool

	Brand_Explicit_Null_When_Empty bool

	BusLocation_Explicit_Null_When_Empty bool

	Description_Explicit_Null_When_Empty bool

	DriverInfo_Explicit_Null_When_Empty bool

	EntityAsyncStatus_Explicit_Null_When_Empty bool

	Host_Explicit_Null_When_Empty bool

	ID_Explicit_Null_When_Empty bool

	IsNvidiaToolsReady_Explicit_Null_When_Empty bool

	IsNvidiaVfsEnabled_Explicit_Null_When_Empty bool

	IsNvidiaVfsSupported_Explicit_Null_When_Empty bool

	LocalCreatedAt_Explicit_Null_When_Empty bool

	LocalID_Explicit_Null_When_Empty bool

	Model_Explicit_Null_When_Empty bool

	Name_Explicit_Null_When_Empty bool

	Status_Explicit_Null_When_Empty bool

	UserUsage_Explicit_Null_When_Empty bool

	UserVgpuTypeID_Explicit_Null_When_Empty bool

	UserVgpuTypeName_Explicit_Null_When_Empty bool

	VgpuInstanceNum_Explicit_Null_When_Empty bool

	VgpuInstanceOnVMNum_Explicit_Null_When_Empty bool
}

func (m VMGpuDetail) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	b.WriteString("{")

	first := true

	// handle nullable field assigned_vgpus_num
	if m.AssignedVgpusNum != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"assigned_vgpus_num\":")
		bytes, err := swag.WriteJSON(m.AssignedVgpusNum)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.AssignedVgpusNum_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"assigned_vgpus_num\":null")
		first = false
	}

	// handle nullable field available_vgpus_num
	if m.AvailableVgpusNum != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"available_vgpus_num\":")
		bytes, err := swag.WriteJSON(m.AvailableVgpusNum)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.AvailableVgpusNum_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"available_vgpus_num\":null")
		first = false
	}

	// handle nullable field brand
	if m.Brand != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"brand\":")
		bytes, err := swag.WriteJSON(m.Brand)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.Brand_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"brand\":null")
		first = false
	}

	// handle nullable field bus_location
	if m.BusLocation != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"bus_location\":")
		bytes, err := swag.WriteJSON(m.BusLocation)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.BusLocation_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"bus_location\":null")
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

	// handle nullable field driver_info
	if m.DriverInfo != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"driver_info\":")
		bytes, err := swag.WriteJSON(m.DriverInfo)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.DriverInfo_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"driver_info\":null")
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

	// handle nullable field host
	if m.Host != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"host\":")
		bytes, err := swag.WriteJSON(m.Host)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.Host_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"host\":null")
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

	// handle nullable field is_nvidia_tools_ready
	if m.IsNvidiaToolsReady != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"is_nvidia_tools_ready\":")
		bytes, err := swag.WriteJSON(m.IsNvidiaToolsReady)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.IsNvidiaToolsReady_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"is_nvidia_tools_ready\":null")
		first = false
	}

	// handle nullable field is_nvidia_vfs_enabled
	if m.IsNvidiaVfsEnabled != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"is_nvidia_vfs_enabled\":")
		bytes, err := swag.WriteJSON(m.IsNvidiaVfsEnabled)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.IsNvidiaVfsEnabled_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"is_nvidia_vfs_enabled\":null")
		first = false
	}

	// handle nullable field is_nvidia_vfs_supported
	if m.IsNvidiaVfsSupported != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"is_nvidia_vfs_supported\":")
		bytes, err := swag.WriteJSON(m.IsNvidiaVfsSupported)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.IsNvidiaVfsSupported_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"is_nvidia_vfs_supported\":null")
		first = false
	}

	// handle non nullable field labels with omitempty
	if swag.IsZero(m.Labels) {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"labels\":")
		bytes, err := swag.WriteJSON(m.Labels)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	}

	// handle nullable field local_created_at
	if m.LocalCreatedAt != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"local_created_at\":")
		bytes, err := swag.WriteJSON(m.LocalCreatedAt)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.LocalCreatedAt_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"local_created_at\":null")
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

	// handle non nullable field mdev_supported_types with omitempty
	if swag.IsZero(m.MdevSupportedTypes) {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"mdev_supported_types\":")
		bytes, err := swag.WriteJSON(m.MdevSupportedTypes)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	}

	// handle nullable field model
	if m.Model != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"model\":")
		bytes, err := swag.WriteJSON(m.Model)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.Model_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"model\":null")
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

	// handle nullable field status
	if m.Status != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"status\":")
		bytes, err := swag.WriteJSON(m.Status)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.Status_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"status\":null")
		first = false
	}

	// handle nullable field user_usage
	if m.UserUsage != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"user_usage\":")
		bytes, err := swag.WriteJSON(m.UserUsage)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.UserUsage_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"user_usage\":null")
		first = false
	}

	// handle nullable field user_vgpu_type_id
	if m.UserVgpuTypeID != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"user_vgpu_type_id\":")
		bytes, err := swag.WriteJSON(m.UserVgpuTypeID)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.UserVgpuTypeID_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"user_vgpu_type_id\":null")
		first = false
	}

	// handle nullable field user_vgpu_type_name
	if m.UserVgpuTypeName != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"user_vgpu_type_name\":")
		bytes, err := swag.WriteJSON(m.UserVgpuTypeName)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.UserVgpuTypeName_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"user_vgpu_type_name\":null")
		first = false
	}

	// handle nullable field vgpu_instance_num
	if m.VgpuInstanceNum != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"vgpu_instance_num\":")
		bytes, err := swag.WriteJSON(m.VgpuInstanceNum)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.VgpuInstanceNum_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"vgpu_instance_num\":null")
		first = false
	}

	// handle nullable field vgpu_instance_on_vm_num
	if m.VgpuInstanceOnVMNum != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"vgpu_instance_on_vm_num\":")
		bytes, err := swag.WriteJSON(m.VgpuInstanceOnVMNum)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.VgpuInstanceOnVMNum_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"vgpu_instance_on_vm_num\":null")
		first = false
	}

	b.WriteString("}")
	return b.Bytes(), nil
}

// Validate validates this Vm gpu detail
func (m *VMGpuDetail) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBrand(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBusLocation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDriverInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEntityAsyncStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHost(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLabels(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocalCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocalID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMdevSupportedTypes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateModel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserUsage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VMGpuDetail) validateBrand(formats strfmt.Registry) error {

	if err := validate.Required("brand", "body", m.Brand); err != nil {
		return err
	}

	return nil
}

func (m *VMGpuDetail) validateBusLocation(formats strfmt.Registry) error {

	if err := validate.Required("bus_location", "body", m.BusLocation); err != nil {
		return err
	}

	return nil
}

func (m *VMGpuDetail) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

func (m *VMGpuDetail) validateDriverInfo(formats strfmt.Registry) error {
	if swag.IsZero(m.DriverInfo) { // not required
		return nil
	}

	if m.DriverInfo != nil {
		if err := m.DriverInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("driver_info")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("driver_info")
			}
			return err
		}
	}

	return nil
}

func (m *VMGpuDetail) validateEntityAsyncStatus(formats strfmt.Registry) error {
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

func (m *VMGpuDetail) validateHost(formats strfmt.Registry) error {

	if err := validate.Required("host", "body", m.Host); err != nil {
		return err
	}

	if m.Host != nil {
		if err := m.Host.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("host")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("host")
			}
			return err
		}
	}

	return nil
}

func (m *VMGpuDetail) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *VMGpuDetail) validateLabels(formats strfmt.Registry) error {
	if swag.IsZero(m.Labels) { // not required
		return nil
	}

	for i := 0; i < len(m.Labels); i++ {
		if swag.IsZero(m.Labels[i]) { // not required
			continue
		}

		if m.Labels[i] != nil {
			if err := m.Labels[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("labels" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("labels" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *VMGpuDetail) validateLocalCreatedAt(formats strfmt.Registry) error {

	if err := validate.Required("local_created_at", "body", m.LocalCreatedAt); err != nil {
		return err
	}

	return nil
}

func (m *VMGpuDetail) validateLocalID(formats strfmt.Registry) error {

	if err := validate.Required("local_id", "body", m.LocalID); err != nil {
		return err
	}

	return nil
}

func (m *VMGpuDetail) validateMdevSupportedTypes(formats strfmt.Registry) error {
	if swag.IsZero(m.MdevSupportedTypes) { // not required
		return nil
	}

	for i := 0; i < len(m.MdevSupportedTypes); i++ {
		if swag.IsZero(m.MdevSupportedTypes[i]) { // not required
			continue
		}

		if m.MdevSupportedTypes[i] != nil {
			if err := m.MdevSupportedTypes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("mdev_supported_types" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("mdev_supported_types" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *VMGpuDetail) validateModel(formats strfmt.Registry) error {

	if err := validate.Required("model", "body", m.Model); err != nil {
		return err
	}

	return nil
}

func (m *VMGpuDetail) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *VMGpuDetail) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	if m.Status != nil {
		if err := m.Status.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("status")
			}
			return err
		}
	}

	return nil
}

func (m *VMGpuDetail) validateUserUsage(formats strfmt.Registry) error {
	if swag.IsZero(m.UserUsage) { // not required
		return nil
	}

	if m.UserUsage != nil {
		if err := m.UserUsage.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("user_usage")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("user_usage")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this Vm gpu detail based on the context it is used
func (m *VMGpuDetail) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDriverInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEntityAsyncStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateHost(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLabels(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMdevSupportedTypes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUserUsage(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VMGpuDetail) contextValidateDriverInfo(ctx context.Context, formats strfmt.Registry) error {

	if m.DriverInfo != nil {
		if err := m.DriverInfo.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("driver_info")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("driver_info")
			}
			return err
		}
	}

	return nil
}

func (m *VMGpuDetail) contextValidateEntityAsyncStatus(ctx context.Context, formats strfmt.Registry) error {

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

func (m *VMGpuDetail) contextValidateHost(ctx context.Context, formats strfmt.Registry) error {

	if m.Host != nil {
		if err := m.Host.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("host")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("host")
			}
			return err
		}
	}

	return nil
}

func (m *VMGpuDetail) contextValidateLabels(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Labels); i++ {

		if m.Labels[i] != nil {
			if err := m.Labels[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("labels" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("labels" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *VMGpuDetail) contextValidateMdevSupportedTypes(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.MdevSupportedTypes); i++ {

		if m.MdevSupportedTypes[i] != nil {
			if err := m.MdevSupportedTypes[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("mdev_supported_types" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("mdev_supported_types" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *VMGpuDetail) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

	if m.Status != nil {
		if err := m.Status.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("status")
			}
			return err
		}
	}

	return nil
}

func (m *VMGpuDetail) contextValidateUserUsage(ctx context.Context, formats strfmt.Registry) error {

	if m.UserUsage != nil {
		if err := m.UserUsage.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("user_usage")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("user_usage")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *VMGpuDetail) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VMGpuDetail) UnmarshalBinary(b []byte) error {
	var res VMGpuDetail
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
