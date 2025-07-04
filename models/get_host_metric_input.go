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

// GetHostMetricInput get host metric input
//
// swagger:model GetHostMetricInput
type GetHostMetricInput struct {

	// hosts
	// Required: true
	Hosts *HostWhereInput `json:"hosts"`

	// metrics
	// Required: true
	Metrics []string `json:"metrics"`

	// range
	// Required: true
	Range *string `json:"range"`

	MarshalOpts *GetHostMetricInputMarshalOpts `json:"-"`
}

type GetHostMetricInputMarshalOpts struct {
	Hosts_Explicit_Null_When_Empty bool

	Metrics_Explicit_Null_When_Empty bool

	Range_Explicit_Null_When_Empty bool
}

func (m GetHostMetricInput) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	b.WriteString("{")

	first := true

	// handle nullable field hosts
	if m.Hosts != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"hosts\":")
		bytes, err := swag.WriteJSON(m.Hosts)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.Hosts_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"hosts\":null")
		first = false
	}

	// handle non nullable field metrics without omitempty
	{
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"metrics\":")
		bytes, err := swag.WriteJSON(m.Metrics)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	}

	// handle nullable field range
	if m.Range != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"range\":")
		bytes, err := swag.WriteJSON(m.Range)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.Range_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"range\":null")
		first = false
	}

	b.WriteString("}")
	return b.Bytes(), nil
}

// Validate validates this get host metric input
func (m *GetHostMetricInput) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHosts(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMetrics(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRange(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetHostMetricInput) validateHosts(formats strfmt.Registry) error {

	if err := validate.Required("hosts", "body", m.Hosts); err != nil {
		return err
	}

	if m.Hosts != nil {
		if err := m.Hosts.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("hosts")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("hosts")
			}
			return err
		}
	}

	return nil
}

func (m *GetHostMetricInput) validateMetrics(formats strfmt.Registry) error {

	if err := validate.Required("metrics", "body", m.Metrics); err != nil {
		return err
	}

	return nil
}

func (m *GetHostMetricInput) validateRange(formats strfmt.Registry) error {

	if err := validate.Required("range", "body", m.Range); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this get host metric input based on the context it is used
func (m *GetHostMetricInput) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateHosts(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetHostMetricInput) contextValidateHosts(ctx context.Context, formats strfmt.Registry) error {

	if m.Hosts != nil {
		if err := m.Hosts.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("hosts")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("hosts")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetHostMetricInput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetHostMetricInput) UnmarshalBinary(b []byte) error {
	var res GetHostMetricInput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
