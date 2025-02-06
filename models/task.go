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

// Task task
//
// swagger:model Task
type Task struct {

	// args
	// Required: true
	Args interface{} `json:"args"`

	// cluster
	Cluster *NestedCluster `json:"cluster,omitempty"`

	// description
	// Required: true
	Description *string `json:"description"`

	// error code
	ErrorCode *string `json:"error_code,omitempty"`

	// error message
	ErrorMessage *string `json:"error_message,omitempty"`

	// finished at
	FinishedAt *string `json:"finished_at,omitempty"`

	// id
	// Required: true
	ID *string `json:"id"`

	// internal
	// Required: true
	Internal *bool `json:"internal"`

	// key
	Key *string `json:"key,omitempty"`

	// local created at
	// Required: true
	LocalCreatedAt *string `json:"local_created_at"`

	// progress
	// Required: true
	Progress *float64 `json:"progress"`

	// resource id
	ResourceID *string `json:"resource_id,omitempty"`

	// resource mutation
	ResourceMutation *string `json:"resource_mutation,omitempty"`

	// resource rollback error
	ResourceRollbackError *string `json:"resource_rollback_error,omitempty"`

	// resource rollback retry count
	ResourceRollbackRetryCount *int32 `json:"resource_rollback_retry_count,omitempty"`

	// resource rollbacked
	ResourceRollbacked *bool `json:"resource_rollbacked,omitempty"`

	// resource type
	ResourceType *string `json:"resource_type,omitempty"`

	// snapshot
	// Required: true
	Snapshot *string `json:"snapshot"`

	// started at
	StartedAt *string `json:"started_at,omitempty"`

	// status
	// Required: true
	Status *TaskStatus `json:"status"`

	// steps
	// Required: true
	Steps []*NestedStep `json:"steps"`

	// type
	Type *TaskType `json:"type,omitempty"`

	// user
	User *NestedUser `json:"user,omitempty"`

	MarshalOpts *TaskMarshalOpts `json:"-"`
}

type TaskMarshalOpts struct {
	Cluster_Explicit_Null_When_Empty bool

	Description_Explicit_Null_When_Empty bool

	ErrorCode_Explicit_Null_When_Empty bool

	ErrorMessage_Explicit_Null_When_Empty bool

	FinishedAt_Explicit_Null_When_Empty bool

	ID_Explicit_Null_When_Empty bool

	Internal_Explicit_Null_When_Empty bool

	Key_Explicit_Null_When_Empty bool

	LocalCreatedAt_Explicit_Null_When_Empty bool

	Progress_Explicit_Null_When_Empty bool

	ResourceID_Explicit_Null_When_Empty bool

	ResourceMutation_Explicit_Null_When_Empty bool

	ResourceRollbackError_Explicit_Null_When_Empty bool

	ResourceRollbackRetryCount_Explicit_Null_When_Empty bool

	ResourceRollbacked_Explicit_Null_When_Empty bool

	ResourceType_Explicit_Null_When_Empty bool

	Snapshot_Explicit_Null_When_Empty bool

	StartedAt_Explicit_Null_When_Empty bool

	Status_Explicit_Null_When_Empty bool

	Type_Explicit_Null_When_Empty bool

	User_Explicit_Null_When_Empty bool
}

func (m Task) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	b.WriteString("{")

	first := true

	// handle non nullable field args with omitempty
	if swag.IsZero(m.Args) {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"args\":")
		bytes, err := swag.WriteJSON(m.Args)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
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

	// handle nullable field error_code
	if m.ErrorCode != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"error_code\":")
		bytes, err := swag.WriteJSON(m.ErrorCode)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.ErrorCode_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"error_code\":null")
		first = false
	}

	// handle nullable field error_message
	if m.ErrorMessage != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"error_message\":")
		bytes, err := swag.WriteJSON(m.ErrorMessage)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.ErrorMessage_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"error_message\":null")
		first = false
	}

	// handle nullable field finished_at
	if m.FinishedAt != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"finished_at\":")
		bytes, err := swag.WriteJSON(m.FinishedAt)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.FinishedAt_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"finished_at\":null")
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

	// handle nullable field key
	if m.Key != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"key\":")
		bytes, err := swag.WriteJSON(m.Key)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.Key_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"key\":null")
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

	// handle nullable field progress
	if m.Progress != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"progress\":")
		bytes, err := swag.WriteJSON(m.Progress)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.Progress_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"progress\":null")
		first = false
	}

	// handle nullable field resource_id
	if m.ResourceID != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"resource_id\":")
		bytes, err := swag.WriteJSON(m.ResourceID)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.ResourceID_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"resource_id\":null")
		first = false
	}

	// handle nullable field resource_mutation
	if m.ResourceMutation != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"resource_mutation\":")
		bytes, err := swag.WriteJSON(m.ResourceMutation)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.ResourceMutation_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"resource_mutation\":null")
		first = false
	}

	// handle nullable field resource_rollback_error
	if m.ResourceRollbackError != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"resource_rollback_error\":")
		bytes, err := swag.WriteJSON(m.ResourceRollbackError)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.ResourceRollbackError_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"resource_rollback_error\":null")
		first = false
	}

	// handle nullable field resource_rollback_retry_count
	if m.ResourceRollbackRetryCount != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"resource_rollback_retry_count\":")
		bytes, err := swag.WriteJSON(m.ResourceRollbackRetryCount)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.ResourceRollbackRetryCount_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"resource_rollback_retry_count\":null")
		first = false
	}

	// handle nullable field resource_rollbacked
	if m.ResourceRollbacked != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"resource_rollbacked\":")
		bytes, err := swag.WriteJSON(m.ResourceRollbacked)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.ResourceRollbacked_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"resource_rollbacked\":null")
		first = false
	}

	// handle nullable field resource_type
	if m.ResourceType != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"resource_type\":")
		bytes, err := swag.WriteJSON(m.ResourceType)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.ResourceType_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"resource_type\":null")
		first = false
	}

	// handle nullable field snapshot
	if m.Snapshot != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"snapshot\":")
		bytes, err := swag.WriteJSON(m.Snapshot)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.Snapshot_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"snapshot\":null")
		first = false
	}

	// handle nullable field started_at
	if m.StartedAt != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"started_at\":")
		bytes, err := swag.WriteJSON(m.StartedAt)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.StartedAt_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"started_at\":null")
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

	// handle non nullable field steps without omitempty
	if !first {
		b.WriteString(",")
	}
	b.WriteString("\"steps\":")
	bytes, err := swag.WriteJSON(m.Steps)
	if err != nil {
		return nil, err
	}
	b.Write(bytes)
	first = false

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

	// handle nullable field user
	if m.User != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"user\":")
		bytes, err := swag.WriteJSON(m.User)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.User_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"user\":null")
		first = false
	}

	b.WriteString("}")
	return b.Bytes(), nil
}

// Validate validates this task
func (m *Task) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateArgs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCluster(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInternal(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocalCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProgress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSnapshot(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSteps(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUser(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Task) validateArgs(formats strfmt.Registry) error {

	if m.Args == nil {
		return errors.Required("args", "body", nil)
	}

	return nil
}

func (m *Task) validateCluster(formats strfmt.Registry) error {
	if swag.IsZero(m.Cluster) { // not required
		return nil
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

func (m *Task) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

func (m *Task) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *Task) validateInternal(formats strfmt.Registry) error {

	if err := validate.Required("internal", "body", m.Internal); err != nil {
		return err
	}

	return nil
}

func (m *Task) validateLocalCreatedAt(formats strfmt.Registry) error {

	if err := validate.Required("local_created_at", "body", m.LocalCreatedAt); err != nil {
		return err
	}

	return nil
}

func (m *Task) validateProgress(formats strfmt.Registry) error {

	if err := validate.Required("progress", "body", m.Progress); err != nil {
		return err
	}

	return nil
}

func (m *Task) validateSnapshot(formats strfmt.Registry) error {

	if err := validate.Required("snapshot", "body", m.Snapshot); err != nil {
		return err
	}

	return nil
}

func (m *Task) validateStatus(formats strfmt.Registry) error {

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

func (m *Task) validateSteps(formats strfmt.Registry) error {

	if err := validate.Required("steps", "body", m.Steps); err != nil {
		return err
	}

	for i := 0; i < len(m.Steps); i++ {
		if swag.IsZero(m.Steps[i]) { // not required
			continue
		}

		if m.Steps[i] != nil {
			if err := m.Steps[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("steps" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("steps" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Task) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
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

func (m *Task) validateUser(formats strfmt.Registry) error {
	if swag.IsZero(m.User) { // not required
		return nil
	}

	if m.User != nil {
		if err := m.User.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("user")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("user")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this task based on the context it is used
func (m *Task) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCluster(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSteps(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUser(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Task) contextValidateCluster(ctx context.Context, formats strfmt.Registry) error {

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

func (m *Task) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

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

func (m *Task) contextValidateSteps(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Steps); i++ {

		if m.Steps[i] != nil {
			if err := m.Steps[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("steps" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("steps" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Task) contextValidateType(ctx context.Context, formats strfmt.Registry) error {

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

func (m *Task) contextValidateUser(ctx context.Context, formats strfmt.Registry) error {

	if m.User != nil {
		if err := m.User.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("user")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("user")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Task) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Task) UnmarshalBinary(b []byte) error {
	var res Task
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
