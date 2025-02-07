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

// GetClusterMetricInput get cluster metric input
//
// swagger:model GetClusterMetricInput
type GetClusterMetricInput struct {

	// clusters
	// Required: true
	Clusters *ClusterWhereInput `json:"clusters"`

	// metrics
	// Required: true
	Metrics []string `json:"metrics"`

	// range
	// Required: true
	Range *string `json:"range"`

	MarshalOpts *GetClusterMetricInputMarshalOpts `json:"-"`
}

type GetClusterMetricInputMarshalOpts struct {
	Clusters_Explicit_Null_When_Empty bool

	Metrics_Explicit_Null_When_Empty bool

	Range_Explicit_Null_When_Empty bool
}

func (m GetClusterMetricInput) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	b.WriteString("{")

	first := true

	// handle nullable field clusters
	if m.Clusters != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"clusters\":")
		bytes, err := swag.WriteJSON(m.Clusters)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.Clusters_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"clusters\":null")
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

// Validate validates this get cluster metric input
func (m *GetClusterMetricInput) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClusters(formats); err != nil {
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

func (m *GetClusterMetricInput) validateClusters(formats strfmt.Registry) error {

	if err := validate.Required("clusters", "body", m.Clusters); err != nil {
		return err
	}

	if m.Clusters != nil {
		if err := m.Clusters.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("clusters")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("clusters")
			}
			return err
		}
	}

	return nil
}

func (m *GetClusterMetricInput) validateMetrics(formats strfmt.Registry) error {

	if err := validate.Required("metrics", "body", m.Metrics); err != nil {
		return err
	}

	return nil
}

func (m *GetClusterMetricInput) validateRange(formats strfmt.Registry) error {

	if err := validate.Required("range", "body", m.Range); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this get cluster metric input based on the context it is used
func (m *GetClusterMetricInput) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateClusters(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetClusterMetricInput) contextValidateClusters(ctx context.Context, formats strfmt.Registry) error {

	if m.Clusters != nil {
		if err := m.Clusters.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("clusters")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("clusters")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetClusterMetricInput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetClusterMetricInput) UnmarshalBinary(b []byte) error {
	var res GetClusterMetricInput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
