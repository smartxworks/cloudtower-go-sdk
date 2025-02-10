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

// GetScvmServiceMetricInput get scvm service metric input
//
// swagger:model GetScvmServiceMetricInput
type GetScvmServiceMetricInput struct {

	// hosts
	// Required: true
	Hosts *HostWhereInput `json:"hosts"`

	// metrics
	// Required: true
	Metrics []string `json:"metrics"`

	// range
	// Required: true
	Range *string `json:"range"`

	// services
	// Required: true
	Services []string `json:"services"`

	MarshalOpts *GetScvmServiceMetricInputMarshalOpts `json:"-"`
}

type GetScvmServiceMetricInputMarshalOpts struct {
	Hosts_Explicit_Null_When_Empty bool

	Metrics_Explicit_Null_When_Empty bool

	Range_Explicit_Null_When_Empty bool

	Services_Explicit_Null_When_Empty bool
}

func (m GetScvmServiceMetricInput) MarshalJSON() ([]byte, error) {
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

	// handle non nullable field services without omitempty
	{
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

// Validate validates this get scvm service metric input
func (m *GetScvmServiceMetricInput) Validate(formats strfmt.Registry) error {
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

	if err := m.validateServices(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetScvmServiceMetricInput) validateHosts(formats strfmt.Registry) error {

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

func (m *GetScvmServiceMetricInput) validateMetrics(formats strfmt.Registry) error {

	if err := validate.Required("metrics", "body", m.Metrics); err != nil {
		return err
	}

	return nil
}

func (m *GetScvmServiceMetricInput) validateRange(formats strfmt.Registry) error {

	if err := validate.Required("range", "body", m.Range); err != nil {
		return err
	}

	return nil
}

func (m *GetScvmServiceMetricInput) validateServices(formats strfmt.Registry) error {

	if err := validate.Required("services", "body", m.Services); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this get scvm service metric input based on the context it is used
func (m *GetScvmServiceMetricInput) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateHosts(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetScvmServiceMetricInput) contextValidateHosts(ctx context.Context, formats strfmt.Registry) error {

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
func (m *GetScvmServiceMetricInput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetScvmServiceMetricInput) UnmarshalBinary(b []byte) error {
	var res GetScvmServiceMetricInput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
