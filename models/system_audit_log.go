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

// SystemAuditLog system audit log
//
// swagger:model SystemAuditLog
type SystemAuditLog struct {

	// action
	// Required: true
	Action *string `json:"action"`

	// cluster
	Cluster *NestedCluster `json:"cluster,omitempty"`

	// finished at
	FinishedAt *string `json:"finished_at,omitempty"`

	// id
	// Required: true
	ID *string `json:"id"`

	// local created at
	LocalCreatedAt *string `json:"local_created_at,omitempty"`

	// local id
	// Required: true
	LocalID *string `json:"local_id"`

	// message
	// Required: true
	Message *string `json:"message"`

	// resource id
	ResourceID *string `json:"resource_id,omitempty"`

	// status
	Status *UserAuditLogStatus `json:"status,omitempty"`

	MarshalOpts *SystemAuditLogMarshalOpts `json:"-"`
}

type SystemAuditLogMarshalOpts struct {
	Action_Explicit_Null_When_Empty bool

	Cluster_Explicit_Null_When_Empty bool

	FinishedAt_Explicit_Null_When_Empty bool

	ID_Explicit_Null_When_Empty bool

	LocalCreatedAt_Explicit_Null_When_Empty bool

	LocalID_Explicit_Null_When_Empty bool

	Message_Explicit_Null_When_Empty bool

	ResourceID_Explicit_Null_When_Empty bool

	Status_Explicit_Null_When_Empty bool
}

func (m SystemAuditLog) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	b.WriteString("{")

	first := true

	// handle nullable field action
	if m.Action != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"action\":")
		bytes, err := swag.WriteJSON(m.Action)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.Action_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"action\":null")
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

	// handle nullable field message
	if m.Message != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"message\":")
		bytes, err := swag.WriteJSON(m.Message)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.Message_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"message\":null")
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

	b.WriteString("}")
	return b.Bytes(), nil
}

// Validate validates this system audit log
func (m *SystemAuditLog) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAction(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCluster(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocalID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SystemAuditLog) validateAction(formats strfmt.Registry) error {

	if err := validate.Required("action", "body", m.Action); err != nil {
		return err
	}

	return nil
}

func (m *SystemAuditLog) validateCluster(formats strfmt.Registry) error {
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

func (m *SystemAuditLog) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *SystemAuditLog) validateLocalID(formats strfmt.Registry) error {

	if err := validate.Required("local_id", "body", m.LocalID); err != nil {
		return err
	}

	return nil
}

func (m *SystemAuditLog) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("message", "body", m.Message); err != nil {
		return err
	}

	return nil
}

func (m *SystemAuditLog) validateStatus(formats strfmt.Registry) error {
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

// ContextValidate validate this system audit log based on the context it is used
func (m *SystemAuditLog) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCluster(ctx, formats); err != nil {
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

func (m *SystemAuditLog) contextValidateCluster(ctx context.Context, formats strfmt.Registry) error {

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

func (m *SystemAuditLog) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

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
func (m *SystemAuditLog) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SystemAuditLog) UnmarshalBinary(b []byte) error {
	var res SystemAuditLog
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
