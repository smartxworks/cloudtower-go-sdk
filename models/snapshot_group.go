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

// SnapshotGroup snapshot group
//
// swagger:model SnapshotGroup
type SnapshotGroup struct {

	// cluster
	// Required: true
	Cluster *NestedCluster `json:"cluster"`

	// deleted
	// Required: true
	Deleted *bool `json:"deleted"`

	// entity async status
	EntityAsyncStatus *EntityAsyncStatus `json:"entityAsyncStatus,omitempty"`

	// estimated recycling time
	EstimatedRecyclingTime *string `json:"estimated_recycling_time,omitempty"`

	// id
	// Required: true
	ID *string `json:"id"`

	// internal
	Internal *bool `json:"internal,omitempty"`

	// keep
	// Required: true
	Keep *bool `json:"keep"`

	// local created at
	// Required: true
	LocalCreatedAt *string `json:"local_created_at"`

	// local id
	// Required: true
	LocalID *string `json:"local_id"`

	// logical size bytes
	// Required: true
	LogicalSizeBytes *int64 `json:"logical_size_bytes"`

	// name
	// Required: true
	Name *string `json:"name"`

	// object num
	// Required: true
	ObjectNum *int32 `json:"object_num"`

	// snapshot plan task
	// Required: true
	SnapshotPlanTask *NestedSnapshotPlanTask `json:"snapshotPlanTask"`

	// vm info
	// Required: true
	VMInfo []*NestedSnapshotGroupVMInfo `json:"vm_info"`

	// vm snapshots
	VMSnapshots []*NestedVMSnapshot `json:"vm_snapshots,omitempty"`

	MarshalOpts *SnapshotGroupMarshalOpts `json:"-"`
}

type SnapshotGroupMarshalOpts struct {
	Cluster_Explicit_Null_When_Empty bool

	Deleted_Explicit_Null_When_Empty bool

	EntityAsyncStatus_Explicit_Null_When_Empty bool

	EstimatedRecyclingTime_Explicit_Null_When_Empty bool

	ID_Explicit_Null_When_Empty bool

	Internal_Explicit_Null_When_Empty bool

	Keep_Explicit_Null_When_Empty bool

	LocalCreatedAt_Explicit_Null_When_Empty bool

	LocalID_Explicit_Null_When_Empty bool

	LogicalSizeBytes_Explicit_Null_When_Empty bool

	Name_Explicit_Null_When_Empty bool

	ObjectNum_Explicit_Null_When_Empty bool

	SnapshotPlanTask_Explicit_Null_When_Empty bool

	VMInfo_Explicit_Null_When_Empty bool

	VMSnapshots_Explicit_Null_When_Empty bool
}

func (m SnapshotGroup) MarshalJSON() ([]byte, error) {
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

	// handle nullable field deleted
	if m.Deleted != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"deleted\":")
		bytes, err := swag.WriteJSON(m.Deleted)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.Deleted_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"deleted\":null")
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

	// handle nullable field estimated_recycling_time
	if m.EstimatedRecyclingTime != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"estimated_recycling_time\":")
		bytes, err := swag.WriteJSON(m.EstimatedRecyclingTime)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.EstimatedRecyclingTime_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"estimated_recycling_time\":null")
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

	// handle nullable field keep
	if m.Keep != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"keep\":")
		bytes, err := swag.WriteJSON(m.Keep)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.Keep_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"keep\":null")
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

	// handle nullable field logical_size_bytes
	if m.LogicalSizeBytes != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"logical_size_bytes\":")
		bytes, err := swag.WriteJSON(m.LogicalSizeBytes)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.LogicalSizeBytes_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"logical_size_bytes\":null")
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

	// handle nullable field object_num
	if m.ObjectNum != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"object_num\":")
		bytes, err := swag.WriteJSON(m.ObjectNum)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.ObjectNum_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"object_num\":null")
		first = false
	}

	// handle nullable field snapshotPlanTask
	if m.SnapshotPlanTask != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"snapshotPlanTask\":")
		bytes, err := swag.WriteJSON(m.SnapshotPlanTask)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.SnapshotPlanTask_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"snapshotPlanTask\":null")
		first = false
	}

	// handle non nullable field vm_info without omitempty
	{
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"vm_info\":")
		bytes, err := swag.WriteJSON(m.VMInfo)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	}

	// handle non nullable field vm_snapshots with omitempty
	if !swag.IsZero(m.VMSnapshots) {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"vm_snapshots\":")
		bytes, err := swag.WriteJSON(m.VMSnapshots)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	}

	b.WriteString("}")
	return b.Bytes(), nil
}

// Validate validates this snapshot group
func (m *SnapshotGroup) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCluster(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeleted(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEntityAsyncStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKeep(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocalCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocalID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLogicalSizeBytes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateObjectNum(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSnapshotPlanTask(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVMInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVMSnapshots(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SnapshotGroup) validateCluster(formats strfmt.Registry) error {

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

func (m *SnapshotGroup) validateDeleted(formats strfmt.Registry) error {

	if err := validate.Required("deleted", "body", m.Deleted); err != nil {
		return err
	}

	return nil
}

func (m *SnapshotGroup) validateEntityAsyncStatus(formats strfmt.Registry) error {
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

func (m *SnapshotGroup) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *SnapshotGroup) validateKeep(formats strfmt.Registry) error {

	if err := validate.Required("keep", "body", m.Keep); err != nil {
		return err
	}

	return nil
}

func (m *SnapshotGroup) validateLocalCreatedAt(formats strfmt.Registry) error {

	if err := validate.Required("local_created_at", "body", m.LocalCreatedAt); err != nil {
		return err
	}

	return nil
}

func (m *SnapshotGroup) validateLocalID(formats strfmt.Registry) error {

	if err := validate.Required("local_id", "body", m.LocalID); err != nil {
		return err
	}

	return nil
}

func (m *SnapshotGroup) validateLogicalSizeBytes(formats strfmt.Registry) error {

	if err := validate.Required("logical_size_bytes", "body", m.LogicalSizeBytes); err != nil {
		return err
	}

	return nil
}

func (m *SnapshotGroup) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *SnapshotGroup) validateObjectNum(formats strfmt.Registry) error {

	if err := validate.Required("object_num", "body", m.ObjectNum); err != nil {
		return err
	}

	return nil
}

func (m *SnapshotGroup) validateSnapshotPlanTask(formats strfmt.Registry) error {

	if err := validate.Required("snapshotPlanTask", "body", m.SnapshotPlanTask); err != nil {
		return err
	}

	if m.SnapshotPlanTask != nil {
		if err := m.SnapshotPlanTask.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("snapshotPlanTask")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("snapshotPlanTask")
			}
			return err
		}
	}

	return nil
}

func (m *SnapshotGroup) validateVMInfo(formats strfmt.Registry) error {

	if err := validate.Required("vm_info", "body", m.VMInfo); err != nil {
		return err
	}

	for i := 0; i < len(m.VMInfo); i++ {
		if swag.IsZero(m.VMInfo[i]) { // not required
			continue
		}

		if m.VMInfo[i] != nil {
			if err := m.VMInfo[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("vm_info" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("vm_info" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *SnapshotGroup) validateVMSnapshots(formats strfmt.Registry) error {
	if swag.IsZero(m.VMSnapshots) { // not required
		return nil
	}

	for i := 0; i < len(m.VMSnapshots); i++ {
		if swag.IsZero(m.VMSnapshots[i]) { // not required
			continue
		}

		if m.VMSnapshots[i] != nil {
			if err := m.VMSnapshots[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("vm_snapshots" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("vm_snapshots" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this snapshot group based on the context it is used
func (m *SnapshotGroup) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCluster(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEntityAsyncStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSnapshotPlanTask(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVMInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVMSnapshots(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SnapshotGroup) contextValidateCluster(ctx context.Context, formats strfmt.Registry) error {

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

func (m *SnapshotGroup) contextValidateEntityAsyncStatus(ctx context.Context, formats strfmt.Registry) error {

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

func (m *SnapshotGroup) contextValidateSnapshotPlanTask(ctx context.Context, formats strfmt.Registry) error {

	if m.SnapshotPlanTask != nil {
		if err := m.SnapshotPlanTask.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("snapshotPlanTask")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("snapshotPlanTask")
			}
			return err
		}
	}

	return nil
}

func (m *SnapshotGroup) contextValidateVMInfo(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.VMInfo); i++ {

		if m.VMInfo[i] != nil {
			if err := m.VMInfo[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("vm_info" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("vm_info" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *SnapshotGroup) contextValidateVMSnapshots(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.VMSnapshots); i++ {

		if m.VMSnapshots[i] != nil {
			if err := m.VMSnapshots[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("vm_snapshots" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("vm_snapshots" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *SnapshotGroup) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SnapshotGroup) UnmarshalBinary(b []byte) error {
	var res SnapshotGroup
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
