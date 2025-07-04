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

// LogCollectionServiceGroupParams log collection service group params
//
// swagger:model LogCollectionServiceGroupParams
type LogCollectionServiceGroupParams struct {

	// group name
	// Required: true
	GroupName *string `json:"group_name"`

	// services
	Services []string `json:"services,omitempty"`

	MarshalOpts *LogCollectionServiceGroupParamsMarshalOpts `json:"-"`
}

type LogCollectionServiceGroupParamsMarshalOpts struct {
	GroupName_Explicit_Null_When_Empty bool

	Services_Explicit_Null_When_Empty bool
}

func (m LogCollectionServiceGroupParams) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	b.WriteString("{")

	first := true

	// handle nullable field group_name
	if m.GroupName != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"group_name\":")
		bytes, err := swag.WriteJSON(m.GroupName)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.GroupName_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"group_name\":null")
		first = false
	}

	// handle non nullable field services with omitempty
	if !swag.IsZero(m.Services) {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"services\":")
		bytes, err := swag.WriteJSON(m.Services)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	}

	b.WriteString("}")
	return b.Bytes(), nil
}

// Validate validates this log collection service group params
func (m *LogCollectionServiceGroupParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGroupName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LogCollectionServiceGroupParams) validateGroupName(formats strfmt.Registry) error {

	if err := validate.Required("group_name", "body", m.GroupName); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this log collection service group params based on context it is used
func (m *LogCollectionServiceGroupParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LogCollectionServiceGroupParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LogCollectionServiceGroupParams) UnmarshalBinary(b []byte) error {
	var res LogCollectionServiceGroupParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
