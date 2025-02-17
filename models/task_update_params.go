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

// TaskUpdateParams task update params
//
// swagger:model TaskUpdateParams
type TaskUpdateParams struct {

	// data
	// Required: true
	Data *TaskUpdateParamsData `json:"data"`

	// where
	// Required: true
	Where *TaskWhereInput `json:"where"`
}

// Validate validates this task update params
func (m *TaskUpdateParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWhere(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TaskUpdateParams) validateData(formats strfmt.Registry) error {

	if err := validate.Required("data", "body", m.Data); err != nil {
		return err
	}

	if m.Data != nil {
		if err := m.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("data")
			}
			return err
		}
	}

	return nil
}

func (m *TaskUpdateParams) validateWhere(formats strfmt.Registry) error {

	if err := validate.Required("where", "body", m.Where); err != nil {
		return err
	}

	if m.Where != nil {
		if err := m.Where.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("where")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("where")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this task update params based on the context it is used
func (m *TaskUpdateParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateWhere(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TaskUpdateParams) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if m.Data != nil {
		if err := m.Data.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("data")
			}
			return err
		}
	}

	return nil
}

func (m *TaskUpdateParams) contextValidateWhere(ctx context.Context, formats strfmt.Registry) error {

	if m.Where != nil {
		if err := m.Where.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("where")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("where")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TaskUpdateParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TaskUpdateParams) UnmarshalBinary(b []byte) error {
	var res TaskUpdateParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// TaskUpdateParamsData task update params data
//
// swagger:model TaskUpdateParamsData
type TaskUpdateParamsData struct {

	// args
	Args interface{} `json:"args,omitempty"`

	// cluster id
	ClusterID *string `json:"cluster_id,omitempty"`

	// description
	Description *TaskDescription `json:"description,omitempty"`

	// error code
	ErrorCode *string `json:"error_code,omitempty"`

	// error message
	ErrorMessage *string `json:"error_message,omitempty"`

	// finished at
	FinishedAt *string `json:"finished_at,omitempty"`

	// key
	Key *string `json:"key,omitempty"`

	// progress
	Progress *float32 `json:"progress,omitempty"`

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
	Snapshot *string `json:"snapshot,omitempty"`

	// started at
	StartedAt *string `json:"started_at,omitempty"`

	// status
	Status *TaskStatus `json:"status,omitempty"`

	// steps
	Steps []*TaskStepCreationParams `json:"steps,omitempty"`

	// type
	Type *TaskType `json:"type,omitempty"`

	// user id
	UserID *string `json:"user_id,omitempty"`
}

// Validate validates this task update params data
func (m *TaskUpdateParamsData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDescription(formats); err != nil {
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TaskUpdateParamsData) validateDescription(formats strfmt.Registry) error {
	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if m.Description != nil {
		if err := m.Description.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data" + "." + "description")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("data" + "." + "description")
			}
			return err
		}
	}

	return nil
}

func (m *TaskUpdateParamsData) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if m.Status != nil {
		if err := m.Status.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data" + "." + "status")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("data" + "." + "status")
			}
			return err
		}
	}

	return nil
}

func (m *TaskUpdateParamsData) validateSteps(formats strfmt.Registry) error {
	if swag.IsZero(m.Steps) { // not required
		return nil
	}

	for i := 0; i < len(m.Steps); i++ {
		if swag.IsZero(m.Steps[i]) { // not required
			continue
		}

		if m.Steps[i] != nil {
			if err := m.Steps[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("data" + "." + "steps" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("data" + "." + "steps" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *TaskUpdateParamsData) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if m.Type != nil {
		if err := m.Type.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data" + "." + "type")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("data" + "." + "type")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this task update params data based on the context it is used
func (m *TaskUpdateParamsData) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDescription(ctx, formats); err != nil {
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TaskUpdateParamsData) contextValidateDescription(ctx context.Context, formats strfmt.Registry) error {

	if m.Description != nil {
		if err := m.Description.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data" + "." + "description")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("data" + "." + "description")
			}
			return err
		}
	}

	return nil
}

func (m *TaskUpdateParamsData) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

	if m.Status != nil {
		if err := m.Status.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data" + "." + "status")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("data" + "." + "status")
			}
			return err
		}
	}

	return nil
}

func (m *TaskUpdateParamsData) contextValidateSteps(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Steps); i++ {

		if m.Steps[i] != nil {
			if err := m.Steps[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("data" + "." + "steps" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("data" + "." + "steps" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *TaskUpdateParamsData) contextValidateType(ctx context.Context, formats strfmt.Registry) error {

	if m.Type != nil {
		if err := m.Type.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data" + "." + "type")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("data" + "." + "type")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TaskUpdateParamsData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TaskUpdateParamsData) UnmarshalBinary(b []byte) error {
	var res TaskUpdateParamsData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
