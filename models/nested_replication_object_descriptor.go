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

// NestedReplicationObjectDescriptor nested replication object descriptor
//
// swagger:model NestedReplicationObjectDescriptor
type NestedReplicationObjectDescriptor struct {

	// cluster local id
	// Required: true
	ClusterLocalID *string `json:"cluster_local_id"`

	// cluster name
	ClusterName *string `json:"cluster_name,omitempty"`

	// object local id
	// Required: true
	ObjectLocalID *string `json:"object_local_id"`

	// object name
	ObjectName *string `json:"object_name,omitempty"`

	// parent object local id
	ParentObjectLocalID *string `json:"parent_object_local_id,omitempty"`

	// parent object name
	ParentObjectName *string `json:"parent_object_name,omitempty"`

	// tower deploy id
	TowerDeployID *string `json:"tower_deploy_id,omitempty"`

	// zbs volume id
	ZbsVolumeID *string `json:"zbs_volume_id,omitempty"`

	MarshalOpts *NestedReplicationObjectDescriptorMarshalOpts `json:"-"`
}

type NestedReplicationObjectDescriptorMarshalOpts struct {
	ClusterLocalID_Explicit_Null_When_Empty bool

	ClusterName_Explicit_Null_When_Empty bool

	ObjectLocalID_Explicit_Null_When_Empty bool

	ObjectName_Explicit_Null_When_Empty bool

	ParentObjectLocalID_Explicit_Null_When_Empty bool

	ParentObjectName_Explicit_Null_When_Empty bool

	TowerDeployID_Explicit_Null_When_Empty bool

	ZbsVolumeID_Explicit_Null_When_Empty bool
}

func (m NestedReplicationObjectDescriptor) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	b.WriteString("{")

	first := true

	// handle nullable field cluster_local_id
	if m.ClusterLocalID != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"cluster_local_id\":")
		bytes, err := swag.WriteJSON(m.ClusterLocalID)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.ClusterLocalID_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"cluster_local_id\":null")
		first = false
	}

	// handle nullable field cluster_name
	if m.ClusterName != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"cluster_name\":")
		bytes, err := swag.WriteJSON(m.ClusterName)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.ClusterName_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"cluster_name\":null")
		first = false
	}

	// handle nullable field object_local_id
	if m.ObjectLocalID != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"object_local_id\":")
		bytes, err := swag.WriteJSON(m.ObjectLocalID)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.ObjectLocalID_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"object_local_id\":null")
		first = false
	}

	// handle nullable field object_name
	if m.ObjectName != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"object_name\":")
		bytes, err := swag.WriteJSON(m.ObjectName)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.ObjectName_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"object_name\":null")
		first = false
	}

	// handle nullable field parent_object_local_id
	if m.ParentObjectLocalID != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"parent_object_local_id\":")
		bytes, err := swag.WriteJSON(m.ParentObjectLocalID)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.ParentObjectLocalID_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"parent_object_local_id\":null")
		first = false
	}

	// handle nullable field parent_object_name
	if m.ParentObjectName != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"parent_object_name\":")
		bytes, err := swag.WriteJSON(m.ParentObjectName)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.ParentObjectName_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"parent_object_name\":null")
		first = false
	}

	// handle nullable field tower_deploy_id
	if m.TowerDeployID != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"tower_deploy_id\":")
		bytes, err := swag.WriteJSON(m.TowerDeployID)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.TowerDeployID_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"tower_deploy_id\":null")
		first = false
	}

	// handle nullable field zbs_volume_id
	if m.ZbsVolumeID != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"zbs_volume_id\":")
		bytes, err := swag.WriteJSON(m.ZbsVolumeID)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.ZbsVolumeID_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"zbs_volume_id\":null")
		first = false
	}

	b.WriteString("}")
	return b.Bytes(), nil
}

// Validate validates this nested replication object descriptor
func (m *NestedReplicationObjectDescriptor) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClusterLocalID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateObjectLocalID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NestedReplicationObjectDescriptor) validateClusterLocalID(formats strfmt.Registry) error {

	if err := validate.Required("cluster_local_id", "body", m.ClusterLocalID); err != nil {
		return err
	}

	return nil
}

func (m *NestedReplicationObjectDescriptor) validateObjectLocalID(formats strfmt.Registry) error {

	if err := validate.Required("object_local_id", "body", m.ObjectLocalID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this nested replication object descriptor based on context it is used
func (m *NestedReplicationObjectDescriptor) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NestedReplicationObjectDescriptor) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NestedReplicationObjectDescriptor) UnmarshalBinary(b []byte) error {
	var res NestedReplicationObjectDescriptor
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
