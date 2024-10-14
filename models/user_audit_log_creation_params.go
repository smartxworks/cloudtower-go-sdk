// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// UserAuditLogCreationParams user audit log creation params
//
// swagger:model UserAuditLogCreationParams
type UserAuditLogCreationParams struct {

	// action
	// Required: true
	// Min Length: 1
	Action *string `json:"action"`

	// cluster id
	ClusterID *string `json:"cluster_id,omitempty"`

	// finished at
	FinishedAt *string `json:"finished_at,omitempty"`

	// ip address
	IPAddress *string `json:"ip_address,omitempty"`

	// message
	// Required: true
	Message *UserAuditLogMessage `json:"message"`

	// resource id
	ResourceID *string `json:"resource_id,omitempty"`

	// resource type
	// Required: true
	// Min Length: 1
	ResourceType *string `json:"resource_type"`

	// started at
	StartedAt *string `json:"started_at,omitempty"`

	// status
	// Required: true
	Status *UserAuditLogStatus `json:"status"`

	// user id
	// Required: true
	UserID *string `json:"user_id"`
}

// Validate validates this user audit log creation params
func (m *UserAuditLogCreationParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAction(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResourceType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UserAuditLogCreationParams) validateAction(formats strfmt.Registry) error {

	if err := validate.Required("action", "body", m.Action); err != nil {
		return err
	}

	if err := validate.MinLength("action", "body", *m.Action, 1); err != nil {
		return err
	}

	return nil
}

func (m *UserAuditLogCreationParams) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("message", "body", m.Message); err != nil {
		return err
	}

	if m.Message != nil {
		if err := m.Message.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("message")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("message")
			}
			return err
		}
	}

	return nil
}

func (m *UserAuditLogCreationParams) validateResourceType(formats strfmt.Registry) error {

	if err := validate.Required("resource_type", "body", m.ResourceType); err != nil {
		return err
	}

	if err := validate.MinLength("resource_type", "body", *m.ResourceType, 1); err != nil {
		return err
	}

	return nil
}

func (m *UserAuditLogCreationParams) validateStatus(formats strfmt.Registry) error {

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

func (m *UserAuditLogCreationParams) validateUserID(formats strfmt.Registry) error {

	if err := validate.Required("user_id", "body", m.UserID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this user audit log creation params based on the context it is used
func (m *UserAuditLogCreationParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateMessage(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UserAuditLogCreationParams) contextValidateMessage(ctx context.Context, formats strfmt.Registry) error {

	if m.Message != nil {
		if err := m.Message.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("message")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("message")
			}
			return err
		}
	}

	return nil
}

func (m *UserAuditLogCreationParams) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *UserAuditLogCreationParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserAuditLogCreationParams) UnmarshalBinary(b []byte) error {
	var res UserAuditLogCreationParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}