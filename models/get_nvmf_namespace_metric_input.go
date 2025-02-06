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

// GetNvmfNamespaceMetricInput get nvmf namespace metric input
//
// swagger:model GetNvmfNamespaceMetricInput
type GetNvmfNamespaceMetricInput struct {

	// hosts
	Hosts *HostWhereInput `json:"hosts,omitempty"`

	// metrics
	// Required: true
	Metrics []string `json:"metrics"`

	// nvmf namespaces
	// Required: true
	NvmfNamespaces *NvmfNamespaceWhereInput `json:"nvmfNamespaces"`

	// range
	// Required: true
	Range *string `json:"range"`

	MarshalOpts *GetNvmfNamespaceMetricInputMarshalOpts `json:"-"`
}

type GetNvmfNamespaceMetricInputMarshalOpts struct {
	Hosts_Explicit_Null_When_Empty bool

	NvmfNamespaces_Explicit_Null_When_Empty bool

	Range_Explicit_Null_When_Empty bool
}

func (m GetNvmfNamespaceMetricInput) MarshalJSON() ([]byte, error) {
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

	// handle nullable field nvmfNamespaces
	if m.NvmfNamespaces != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"nvmfNamespaces\":")
		bytes, err := swag.WriteJSON(m.NvmfNamespaces)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.NvmfNamespaces_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"nvmfNamespaces\":null")
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

// Validate validates this get nvmf namespace metric input
func (m *GetNvmfNamespaceMetricInput) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHosts(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMetrics(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNvmfNamespaces(formats); err != nil {
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

func (m *GetNvmfNamespaceMetricInput) validateHosts(formats strfmt.Registry) error {
	if swag.IsZero(m.Hosts) { // not required
		return nil
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

func (m *GetNvmfNamespaceMetricInput) validateMetrics(formats strfmt.Registry) error {

	if err := validate.Required("metrics", "body", m.Metrics); err != nil {
		return err
	}

	return nil
}

func (m *GetNvmfNamespaceMetricInput) validateNvmfNamespaces(formats strfmt.Registry) error {

	if err := validate.Required("nvmfNamespaces", "body", m.NvmfNamespaces); err != nil {
		return err
	}

	if m.NvmfNamespaces != nil {
		if err := m.NvmfNamespaces.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("nvmfNamespaces")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("nvmfNamespaces")
			}
			return err
		}
	}

	return nil
}

func (m *GetNvmfNamespaceMetricInput) validateRange(formats strfmt.Registry) error {

	if err := validate.Required("range", "body", m.Range); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this get nvmf namespace metric input based on the context it is used
func (m *GetNvmfNamespaceMetricInput) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateHosts(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNvmfNamespaces(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetNvmfNamespaceMetricInput) contextValidateHosts(ctx context.Context, formats strfmt.Registry) error {

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

func (m *GetNvmfNamespaceMetricInput) contextValidateNvmfNamespaces(ctx context.Context, formats strfmt.Registry) error {

	if m.NvmfNamespaces != nil {
		if err := m.NvmfNamespaces.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("nvmfNamespaces")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("nvmfNamespaces")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetNvmfNamespaceMetricInput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetNvmfNamespaceMetricInput) UnmarshalBinary(b []byte) error {
	var res GetNvmfNamespaceMetricInput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
