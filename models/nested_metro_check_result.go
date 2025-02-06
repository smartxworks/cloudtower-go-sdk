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

// NestedMetroCheckResult nested metro check result
//
// swagger:model NestedMetroCheckResult
type NestedMetroCheckResult struct {

	// critical
	// Required: true
	Critical *int32 `json:"critical"`

	// info
	// Required: true
	Info *int32 `json:"info"`

	// items
	// Required: true
	Items []*NestedMetroCheckItem `json:"items"`

	// notice
	// Required: true
	Notice *int32 `json:"notice"`

	// status
	// Required: true
	Status *MetroCheckStatusEnum `json:"status"`

	MarshalOpts *NestedMetroCheckResultMarshalOpts `json:"-"`
}

type NestedMetroCheckResultMarshalOpts struct {
	Critical_Explicit_Null_When_Empty bool

	Info_Explicit_Null_When_Empty bool

	Notice_Explicit_Null_When_Empty bool

	Status_Explicit_Null_When_Empty bool
}

func (m NestedMetroCheckResult) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	b.WriteString("{")

	first := true

	// handle nullable field critical
	if m.Critical != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"critical\":")
		bytes, err := swag.WriteJSON(m.Critical)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.Critical_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"critical\":null")
		first = false
	}

	// handle nullable field info
	if m.Info != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"info\":")
		bytes, err := swag.WriteJSON(m.Info)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.Info_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"info\":null")
		first = false
	}

	// handle non nullable field items without omitempty
	if !first {
		b.WriteString(",")
	}
	b.WriteString("\"items\":")
	bytes, err := swag.WriteJSON(m.Items)
	if err != nil {
		return nil, err
	}
	b.Write(bytes)
	first = false

	// handle nullable field notice
	if m.Notice != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"notice\":")
		bytes, err := swag.WriteJSON(m.Notice)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.Notice_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"notice\":null")
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

// Validate validates this nested metro check result
func (m *NestedMetroCheckResult) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCritical(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateItems(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNotice(formats); err != nil {
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

func (m *NestedMetroCheckResult) validateCritical(formats strfmt.Registry) error {

	if err := validate.Required("critical", "body", m.Critical); err != nil {
		return err
	}

	return nil
}

func (m *NestedMetroCheckResult) validateInfo(formats strfmt.Registry) error {

	if err := validate.Required("info", "body", m.Info); err != nil {
		return err
	}

	return nil
}

func (m *NestedMetroCheckResult) validateItems(formats strfmt.Registry) error {

	if err := validate.Required("items", "body", m.Items); err != nil {
		return err
	}

	for i := 0; i < len(m.Items); i++ {
		if swag.IsZero(m.Items[i]) { // not required
			continue
		}

		if m.Items[i] != nil {
			if err := m.Items[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("items" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("items" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *NestedMetroCheckResult) validateNotice(formats strfmt.Registry) error {

	if err := validate.Required("notice", "body", m.Notice); err != nil {
		return err
	}

	return nil
}

func (m *NestedMetroCheckResult) validateStatus(formats strfmt.Registry) error {

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

// ContextValidate validate this nested metro check result based on the context it is used
func (m *NestedMetroCheckResult) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateItems(ctx, formats); err != nil {
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

func (m *NestedMetroCheckResult) contextValidateItems(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Items); i++ {

		if m.Items[i] != nil {
			if err := m.Items[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("items" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("items" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *NestedMetroCheckResult) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

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
func (m *NestedMetroCheckResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NestedMetroCheckResult) UnmarshalBinary(b []byte) error {
	var res NestedMetroCheckResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
