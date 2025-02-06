// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// BackupRestoreExecutionNetworkMapping backup restore execution network mapping
//
// swagger:model BackupRestoreExecutionNetworkMapping
type BackupRestoreExecutionNetworkMapping struct {

	// typename
	// Enum: [BackupRestoreExecutionNetworkMapping]
	Typename *string `json:"__typename,omitempty"`

	// dst vlan id
	// Required: true
	DstVlanID *string `json:"dst_vlan_id"`

	// src vlan id
	// Required: true
	SrcVlanID *string `json:"src_vlan_id"`

	MarshalOpts *BackupRestoreExecutionNetworkMappingMarshalOpts `json:"-"`
}

type BackupRestoreExecutionNetworkMappingMarshalOpts struct {
	Typename_Explicit_Null_When_Empty bool

	DstVlanID_Explicit_Null_When_Empty bool

	SrcVlanID_Explicit_Null_When_Empty bool
}

func (m BackupRestoreExecutionNetworkMapping) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	b.WriteString("{")

	first := true

	// handle nullable field __typename
	if m.Typename != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"__typename\":")
		bytes, err := swag.WriteJSON(m.Typename)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.Typename_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"__typename\":null")
		first = false
	}

	// handle nullable field dst_vlan_id
	if m.DstVlanID != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"dst_vlan_id\":")
		bytes, err := swag.WriteJSON(m.DstVlanID)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.DstVlanID_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"dst_vlan_id\":null")
		first = false
	}

	// handle nullable field src_vlan_id
	if m.SrcVlanID != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"src_vlan_id\":")
		bytes, err := swag.WriteJSON(m.SrcVlanID)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.SrcVlanID_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"src_vlan_id\":null")
		first = false
	}

	b.WriteString("}")
	return b.Bytes(), nil
}

// Validate validates this backup restore execution network mapping
func (m *BackupRestoreExecutionNetworkMapping) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTypename(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDstVlanID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSrcVlanID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var backupRestoreExecutionNetworkMappingTypeTypenamePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["BackupRestoreExecutionNetworkMapping"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		backupRestoreExecutionNetworkMappingTypeTypenamePropEnum = append(backupRestoreExecutionNetworkMappingTypeTypenamePropEnum, v)
	}
}

const (

	// BackupRestoreExecutionNetworkMappingTypenameBackupRestoreExecutionNetworkMapping captures enum value "BackupRestoreExecutionNetworkMapping"
	BackupRestoreExecutionNetworkMappingTypenameBackupRestoreExecutionNetworkMapping string = "BackupRestoreExecutionNetworkMapping"
)

// prop value enum
func (m *BackupRestoreExecutionNetworkMapping) validateTypenameEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, backupRestoreExecutionNetworkMappingTypeTypenamePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *BackupRestoreExecutionNetworkMapping) validateTypename(formats strfmt.Registry) error {
	if swag.IsZero(m.Typename) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypenameEnum("__typename", "body", *m.Typename); err != nil {
		return err
	}

	return nil
}

func (m *BackupRestoreExecutionNetworkMapping) validateDstVlanID(formats strfmt.Registry) error {

	if err := validate.Required("dst_vlan_id", "body", m.DstVlanID); err != nil {
		return err
	}

	return nil
}

func (m *BackupRestoreExecutionNetworkMapping) validateSrcVlanID(formats strfmt.Registry) error {

	if err := validate.Required("src_vlan_id", "body", m.SrcVlanID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this backup restore execution network mapping based on context it is used
func (m *BackupRestoreExecutionNetworkMapping) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *BackupRestoreExecutionNetworkMapping) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BackupRestoreExecutionNetworkMapping) UnmarshalBinary(b []byte) error {
	var res BackupRestoreExecutionNetworkMapping
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
