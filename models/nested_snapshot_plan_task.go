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

// NestedSnapshotPlanTask nested snapshot plan task
//
// swagger:model NestedSnapshotPlanTask
type NestedSnapshotPlanTask struct {

	// id
	// Required: true
	ID *string `json:"id"`

	MarshalOpts *NestedSnapshotPlanTaskMarshalOpts `json:"-"`
}

type NestedSnapshotPlanTaskMarshalOpts struct {
	ID_Explicit_Null_When_Empty bool
}

func (m NestedSnapshotPlanTask) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	b.WriteString("{")

	first := true

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

	b.WriteString("}")
	return b.Bytes(), nil
}

// Validate validates this nested snapshot plan task
func (m *NestedSnapshotPlanTask) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NestedSnapshotPlanTask) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this nested snapshot plan task based on context it is used
func (m *NestedSnapshotPlanTask) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NestedSnapshotPlanTask) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NestedSnapshotPlanTask) UnmarshalBinary(b []byte) error {
	var res NestedSnapshotPlanTask
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
