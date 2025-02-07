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

// ApplicationVMSpec application Vm spec
//
// swagger:model ApplicationVmSpec
type ApplicationVMSpec struct {

	// cloud init user data
	CloudInitUserData *string `json:"cloudInitUserData,omitempty"`

	// cluster
	Cluster *string `json:"cluster,omitempty"`

	// cpu
	CPU *int32 `json:"cpu,omitempty"`

	// env
	Env []*ApplicationVMSpecEnv `json:"env,omitempty"`

	// host
	Host *string `json:"host,omitempty"`

	// image
	Image *string `json:"image,omitempty"`

	// internal
	Internal *bool `json:"internal,omitempty"`

	// memory
	Memory *int64 `json:"memory,omitempty"`

	// memory unit
	MemoryUnit *ByteUnit `json:"memory_unit,omitempty"`

	// name
	Name *string `json:"name,omitempty"`

	// network
	Network *ApplicationVMSpecNetwork `json:"network,omitempty"`

	// public keys
	PublicKeys []string `json:"publicKeys,omitempty"`

	// status
	Status *ApplicationVMSpecStatus `json:"status,omitempty"`

	// storages
	Storages []*ApplicationVMSpecStorage `json:"storages,omitempty"`

	// vm usage
	VMUsage *VMUsage `json:"vmUsage,omitempty"`

	MarshalOpts *ApplicationVMSpecMarshalOpts `json:"-"`
}

type ApplicationVMSpecMarshalOpts struct {
	CloudInitUserData_Explicit_Null_When_Empty bool

	Cluster_Explicit_Null_When_Empty bool

	CPU_Explicit_Null_When_Empty bool

	Env_Explicit_Null_When_Empty bool

	Host_Explicit_Null_When_Empty bool

	Image_Explicit_Null_When_Empty bool

	Internal_Explicit_Null_When_Empty bool

	Memory_Explicit_Null_When_Empty bool

	MemoryUnit_Explicit_Null_When_Empty bool

	Name_Explicit_Null_When_Empty bool

	Network_Explicit_Null_When_Empty bool

	PublicKeys_Explicit_Null_When_Empty bool

	Status_Explicit_Null_When_Empty bool

	Storages_Explicit_Null_When_Empty bool

	VMUsage_Explicit_Null_When_Empty bool
}

func (m ApplicationVMSpec) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	b.WriteString("{")

	first := true

	// handle nullable field cloudInitUserData
	if m.CloudInitUserData != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"cloudInitUserData\":")
		bytes, err := swag.WriteJSON(m.CloudInitUserData)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.CloudInitUserData_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"cloudInitUserData\":null")
		first = false
	}

	// handle nullable field cluster
	if m.Cluster != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"cluster\":")
		bytes, err := swag.WriteJSON(m.Cluster)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.Cluster_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"cluster\":null")
		first = false
	}

	// handle nullable field cpu
	if m.CPU != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"cpu\":")
		bytes, err := swag.WriteJSON(m.CPU)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.CPU_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"cpu\":null")
		first = false
	}

	// handle non nullable field env with omitempty
	if swag.IsZero(m.Env) {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"env\":")
		bytes, err := swag.WriteJSON(m.Env)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
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

	// handle nullable field image
	if m.Image != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"image\":")
		bytes, err := swag.WriteJSON(m.Image)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.Image_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"image\":null")
		first = false
	}

	// handle nullable field internal
	if m.Internal != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"internal\":")
		bytes, err := swag.WriteJSON(m.Internal)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.Internal_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"internal\":null")
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

	// handle nullable field network
	if m.Network != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"network\":")
		bytes, err := swag.WriteJSON(m.Network)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.Network_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"network\":null")
		first = false
	}

	// handle non nullable field publicKeys with omitempty
	if swag.IsZero(m.PublicKeys) {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"publicKeys\":")
		bytes, err := swag.WriteJSON(m.PublicKeys)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
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

	// handle non nullable field storages with omitempty
	if swag.IsZero(m.Storages) {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"storages\":")
		bytes, err := swag.WriteJSON(m.Storages)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	}

	// handle nullable field vmUsage
	if m.VMUsage != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"vmUsage\":")
		bytes, err := swag.WriteJSON(m.VMUsage)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.VMUsage_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"vmUsage\":null")
		first = false
	}

	b.WriteString("}")
	return b.Bytes(), nil
}

// Validate validates this application Vm spec
func (m *ApplicationVMSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEnv(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMemoryUnit(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNetwork(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStorages(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVMUsage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ApplicationVMSpec) validateEnv(formats strfmt.Registry) error {
	if swag.IsZero(m.Env) { // not required
		return nil
	}

	for i := 0; i < len(m.Env); i++ {
		if swag.IsZero(m.Env[i]) { // not required
			continue
		}

		if m.Env[i] != nil {
			if err := m.Env[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("env" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("env" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ApplicationVMSpec) validateMemoryUnit(formats strfmt.Registry) error {
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

func (m *ApplicationVMSpec) validateNetwork(formats strfmt.Registry) error {
	if swag.IsZero(m.Network) { // not required
		return nil
	}

	if m.Network != nil {
		if err := m.Network.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("network")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("network")
			}
			return err
		}
	}

	return nil
}

func (m *ApplicationVMSpec) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
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

func (m *ApplicationVMSpec) validateStorages(formats strfmt.Registry) error {
	if swag.IsZero(m.Storages) { // not required
		return nil
	}

	for i := 0; i < len(m.Storages); i++ {
		if swag.IsZero(m.Storages[i]) { // not required
			continue
		}

		if m.Storages[i] != nil {
			if err := m.Storages[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("storages" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("storages" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ApplicationVMSpec) validateVMUsage(formats strfmt.Registry) error {
	if swag.IsZero(m.VMUsage) { // not required
		return nil
	}

	if m.VMUsage != nil {
		if err := m.VMUsage.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vmUsage")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vmUsage")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this application Vm spec based on the context it is used
func (m *ApplicationVMSpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateEnv(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMemoryUnit(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNetwork(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStorages(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVMUsage(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ApplicationVMSpec) contextValidateEnv(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Env); i++ {

		if m.Env[i] != nil {
			if err := m.Env[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("env" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("env" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ApplicationVMSpec) contextValidateMemoryUnit(ctx context.Context, formats strfmt.Registry) error {

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

func (m *ApplicationVMSpec) contextValidateNetwork(ctx context.Context, formats strfmt.Registry) error {

	if m.Network != nil {
		if err := m.Network.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("network")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("network")
			}
			return err
		}
	}

	return nil
}

func (m *ApplicationVMSpec) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

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

func (m *ApplicationVMSpec) contextValidateStorages(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Storages); i++ {

		if m.Storages[i] != nil {
			if err := m.Storages[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("storages" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("storages" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ApplicationVMSpec) contextValidateVMUsage(ctx context.Context, formats strfmt.Registry) error {

	if m.VMUsage != nil {
		if err := m.VMUsage.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vmUsage")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vmUsage")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ApplicationVMSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ApplicationVMSpec) UnmarshalBinary(b []byte) error {
	var res ApplicationVMSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
