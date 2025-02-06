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

// WithTaskVsphereEsxiAccountArray with task vsphere esxi account array
//
// swagger:model WithTask_VsphereEsxiAccount-Array_
type WithTaskVsphereEsxiAccountArray struct {

	// data
	// Required: true
	Data []*VsphereEsxiAccount `json:"data"`

	// task id
	TaskID *string `json:"task_id,omitempty"`

	MarshalOpts *WithTaskVsphereEsxiAccountArrayMarshalOpts `json:"-"`
}

type WithTaskVsphereEsxiAccountArrayMarshalOpts struct {
	TaskID_Explicit_Null_When_Empty bool
}

func (m WithTaskVsphereEsxiAccountArray) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	b.WriteString("{")

	first := true

	// handle non nullable field data without omitempty
	if !first {
		b.WriteString(",")
	}
	b.WriteString("\"data\":")
	bytes, err := swag.WriteJSON(m.Data)
	if err != nil {
		return nil, err
	}
	b.Write(bytes)
	first = false

	// handle nullable field task_id
	if m.TaskID != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"task_id\":")
		bytes, err := swag.WriteJSON(m.TaskID)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.TaskID_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"task_id\":null")
		first = false
	}

	b.WriteString("}")
	return b.Bytes(), nil
}

// Validate validates this with task vsphere esxi account array
func (m *WithTaskVsphereEsxiAccountArray) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WithTaskVsphereEsxiAccountArray) validateData(formats strfmt.Registry) error {

	if err := validate.Required("data", "body", m.Data); err != nil {
		return err
	}

	for i := 0; i < len(m.Data); i++ {
		if swag.IsZero(m.Data[i]) { // not required
			continue
		}

		if m.Data[i] != nil {
			if err := m.Data[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("data" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this with task vsphere esxi account array based on the context it is used
func (m *WithTaskVsphereEsxiAccountArray) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WithTaskVsphereEsxiAccountArray) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Data); i++ {

		if m.Data[i] != nil {
			if err := m.Data[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("data" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *WithTaskVsphereEsxiAccountArray) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WithTaskVsphereEsxiAccountArray) UnmarshalBinary(b []byte) error {
	var res WithTaskVsphereEsxiAccountArray
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
