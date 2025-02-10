// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ElfDataStore elf data store
//
// swagger:model ElfDataStore
type ElfDataStore struct {

	// cluster
	// Required: true
	Cluster *NestedCluster `json:"cluster"`

	// description
	// Required: true
	Description *string `json:"description"`

	// external use
	// Required: true
	ExternalUse *bool `json:"external_use"`

	// id
	// Required: true
	ID *string `json:"id"`

	// internal
	// Required: true
	Internal *bool `json:"internal"`

	// ip whitelist
	// Required: true
	IPWhitelist *string `json:"ip_whitelist"`

	// iscsi target
	IscsiTarget *NestedIscsiTarget `json:"iscsi_target,omitempty"`

	// local id
	// Required: true
	LocalID *string `json:"local_id"`

	// name
	// Required: true
	Name *string `json:"name"`

	// nfs export
	NfsExport *NestedNfsExport `json:"nfs_export,omitempty"`

	// nvmf subsystem
	NvmfSubsystem *NestedNvmfSubsystem `json:"nvmf_subsystem,omitempty"`

	// replica num
	// Required: true
	ReplicaNum *int32 `json:"replica_num"`

	// thin provision
	// Required: true
	ThinProvision *bool `json:"thin_provision"`

	// type
	// Required: true
	Type *ElfDataStoreType `json:"type"`

	MarshalOpts *ElfDataStoreMarshalOpts `json:"-"`
}

type ElfDataStoreMarshalOpts struct {
	Cluster_Explicit_Null_When_Empty bool

	Description_Explicit_Null_When_Empty bool

	ExternalUse_Explicit_Null_When_Empty bool

	ID_Explicit_Null_When_Empty bool

	Internal_Explicit_Null_When_Empty bool

	IPWhitelist_Explicit_Null_When_Empty bool

	IscsiTarget_Explicit_Null_When_Empty bool

	LocalID_Explicit_Null_When_Empty bool

	Name_Explicit_Null_When_Empty bool

	NfsExport_Explicit_Null_When_Empty bool

	NvmfSubsystem_Explicit_Null_When_Empty bool

	ReplicaNum_Explicit_Null_When_Empty bool

	ThinProvision_Explicit_Null_When_Empty bool

	Type_Explicit_Null_When_Empty bool
}

func (m ElfDataStore) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	b.WriteString("{")

	first := true

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

	// handle nullable field external_use
	if m.ExternalUse != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"external_use\":")
		bytes, err := swag.WriteJSON(m.ExternalUse)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.ExternalUse_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"external_use\":null")
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

	// handle nullable field ip_whitelist
	if m.IPWhitelist != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"ip_whitelist\":")
		bytes, err := swag.WriteJSON(m.IPWhitelist)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.IPWhitelist_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"ip_whitelist\":null")
		first = false
	}

	// handle nullable field iscsi_target
	if m.IscsiTarget != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"iscsi_target\":")
		bytes, err := swag.WriteJSON(m.IscsiTarget)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.IscsiTarget_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"iscsi_target\":null")
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

	// handle nullable field nfs_export
	if m.NfsExport != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"nfs_export\":")
		bytes, err := swag.WriteJSON(m.NfsExport)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.NfsExport_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"nfs_export\":null")
		first = false
	}

	// handle nullable field nvmf_subsystem
	if m.NvmfSubsystem != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"nvmf_subsystem\":")
		bytes, err := swag.WriteJSON(m.NvmfSubsystem)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.NvmfSubsystem_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"nvmf_subsystem\":null")
		first = false
	}

	// handle nullable field replica_num
	if m.ReplicaNum != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"replica_num\":")
		bytes, err := swag.WriteJSON(m.ReplicaNum)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.ReplicaNum_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"replica_num\":null")
		first = false
	}

	// handle nullable field thin_provision
	if m.ThinProvision != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"thin_provision\":")
		bytes, err := swag.WriteJSON(m.ThinProvision)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.ThinProvision_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"thin_provision\":null")
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

	b.WriteString("}")
	return b.Bytes(), nil
}

// Validate validates this elf data store
func (m *ElfDataStore) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCluster(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExternalUse(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInternal(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIPWhitelist(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIscsiTarget(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocalID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNfsExport(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNvmfSubsystem(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReplicaNum(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateThinProvision(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ElfDataStore) validateCluster(formats strfmt.Registry) error {

	if err := validate.Required("cluster", "body", m.Cluster); err != nil {
		return err
	}

	if m.Cluster != nil {
		if err := m.Cluster.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cluster")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cluster")
			}
			return err
		}
	}

	return nil
}

func (m *ElfDataStore) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

func (m *ElfDataStore) validateExternalUse(formats strfmt.Registry) error {

	if err := validate.Required("external_use", "body", m.ExternalUse); err != nil {
		return err
	}

	return nil
}

func (m *ElfDataStore) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *ElfDataStore) validateInternal(formats strfmt.Registry) error {

	if err := validate.Required("internal", "body", m.Internal); err != nil {
		return err
	}

	return nil
}

func (m *ElfDataStore) validateIPWhitelist(formats strfmt.Registry) error {

	if err := validate.Required("ip_whitelist", "body", m.IPWhitelist); err != nil {
		return err
	}

	return nil
}

func (m *ElfDataStore) validateIscsiTarget(formats strfmt.Registry) error {
	if swag.IsZero(m.IscsiTarget) { // not required
		return nil
	}

	if m.IscsiTarget != nil {
		if err := m.IscsiTarget.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("iscsi_target")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("iscsi_target")
			}
			return err
		}
	}

	return nil
}

func (m *ElfDataStore) validateLocalID(formats strfmt.Registry) error {

	if err := validate.Required("local_id", "body", m.LocalID); err != nil {
		return err
	}

	return nil
}

func (m *ElfDataStore) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *ElfDataStore) validateNfsExport(formats strfmt.Registry) error {
	if swag.IsZero(m.NfsExport) { // not required
		return nil
	}

	if m.NfsExport != nil {
		if err := m.NfsExport.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("nfs_export")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("nfs_export")
			}
			return err
		}
	}

	return nil
}

func (m *ElfDataStore) validateNvmfSubsystem(formats strfmt.Registry) error {
	if swag.IsZero(m.NvmfSubsystem) { // not required
		return nil
	}

	if m.NvmfSubsystem != nil {
		if err := m.NvmfSubsystem.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("nvmf_subsystem")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("nvmf_subsystem")
			}
			return err
		}
	}

	return nil
}

func (m *ElfDataStore) validateReplicaNum(formats strfmt.Registry) error {

	if err := validate.Required("replica_num", "body", m.ReplicaNum); err != nil {
		return err
	}

	return nil
}

func (m *ElfDataStore) validateThinProvision(formats strfmt.Registry) error {

	if err := validate.Required("thin_provision", "body", m.ThinProvision); err != nil {
		return err
	}

	return nil
}

func (m *ElfDataStore) validateType(formats strfmt.Registry) error {

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

// ContextValidate validate this elf data store based on the context it is used
func (m *ElfDataStore) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCluster(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIscsiTarget(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNfsExport(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNvmfSubsystem(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ElfDataStore) contextValidateCluster(ctx context.Context, formats strfmt.Registry) error {

	if m.Cluster != nil {
		if err := m.Cluster.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cluster")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cluster")
			}
			return err
		}
	}

	return nil
}

func (m *ElfDataStore) contextValidateIscsiTarget(ctx context.Context, formats strfmt.Registry) error {

	if m.IscsiTarget != nil {
		if err := m.IscsiTarget.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("iscsi_target")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("iscsi_target")
			}
			return err
		}
	}

	return nil
}

func (m *ElfDataStore) contextValidateNfsExport(ctx context.Context, formats strfmt.Registry) error {

	if m.NfsExport != nil {
		if err := m.NfsExport.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("nfs_export")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("nfs_export")
			}
			return err
		}
	}

	return nil
}

func (m *ElfDataStore) contextValidateNvmfSubsystem(ctx context.Context, formats strfmt.Registry) error {

	if m.NvmfSubsystem != nil {
		if err := m.NvmfSubsystem.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("nvmf_subsystem")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("nvmf_subsystem")
			}
			return err
		}
	}

	return nil
}

func (m *ElfDataStore) contextValidateType(ctx context.Context, formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *ElfDataStore) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ElfDataStore) UnmarshalBinary(b []byte) error {
	var res ElfDataStore
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
