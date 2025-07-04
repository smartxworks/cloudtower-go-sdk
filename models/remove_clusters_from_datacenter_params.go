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

// RemoveClustersFromDatacenterParams remove clusters from datacenter params
//
// swagger:model RemoveClustersFromDatacenterParams
type RemoveClustersFromDatacenterParams struct {

	// data
	// Required: true
	Data *RemoveClustersFromDatacenterParamsData `json:"data"`

	// where
	// Required: true
	Where *DatacenterWhereUniqueInput `json:"where"`

	MarshalOpts *RemoveClustersFromDatacenterParamsMarshalOpts `json:"-"`
}

type RemoveClustersFromDatacenterParamsMarshalOpts struct {
	Data_Explicit_Null_When_Empty bool

	Where_Explicit_Null_When_Empty bool
}

func (m RemoveClustersFromDatacenterParams) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	b.WriteString("{")

	first := true

	// handle nullable field data
	if m.Data != nil {
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
	} else if m.MarshalOpts != nil && m.MarshalOpts.Data_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"data\":null")
		first = false
	}

	// handle nullable field where
	if m.Where != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"where\":")
		bytes, err := swag.WriteJSON(m.Where)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.Where_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"where\":null")
		first = false
	}

	b.WriteString("}")
	return b.Bytes(), nil
}

// Validate validates this remove clusters from datacenter params
func (m *RemoveClustersFromDatacenterParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWhere(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RemoveClustersFromDatacenterParams) validateData(formats strfmt.Registry) error {

	if err := validate.Required("data", "body", m.Data); err != nil {
		return err
	}

	if m.Data != nil {
		if err := m.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("data")
			}
			return err
		}
	}

	return nil
}

func (m *RemoveClustersFromDatacenterParams) validateWhere(formats strfmt.Registry) error {

	if err := validate.Required("where", "body", m.Where); err != nil {
		return err
	}

	if m.Where != nil {
		if err := m.Where.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("where")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("where")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this remove clusters from datacenter params based on the context it is used
func (m *RemoveClustersFromDatacenterParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateWhere(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RemoveClustersFromDatacenterParams) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if m.Data != nil {
		if err := m.Data.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("data")
			}
			return err
		}
	}

	return nil
}

func (m *RemoveClustersFromDatacenterParams) contextValidateWhere(ctx context.Context, formats strfmt.Registry) error {

	if m.Where != nil {
		if err := m.Where.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("where")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("where")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RemoveClustersFromDatacenterParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RemoveClustersFromDatacenterParams) UnmarshalBinary(b []byte) error {
	var res RemoveClustersFromDatacenterParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// RemoveClustersFromDatacenterParamsData remove clusters from datacenter params data
//
// swagger:model RemoveClustersFromDatacenterParamsData
type RemoveClustersFromDatacenterParamsData struct {

	// clusters
	Clusters *ClusterWhereInput `json:"clusters,omitempty"`

	MarshalOpts *RemoveClustersFromDatacenterParamsDataMarshalOpts `json:"-"`
}

type RemoveClustersFromDatacenterParamsDataMarshalOpts struct {
	Clusters_Explicit_Null_When_Empty bool
}

func (m RemoveClustersFromDatacenterParamsData) MarshalJSON() ([]byte, error) {
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

	b.WriteString("}")
	return b.Bytes(), nil
}

// Validate validates this remove clusters from datacenter params data
func (m *RemoveClustersFromDatacenterParamsData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClusters(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RemoveClustersFromDatacenterParamsData) validateClusters(formats strfmt.Registry) error {
	if swag.IsZero(m.Clusters) { // not required
		return nil
	}

	if m.Clusters != nil {
		if err := m.Clusters.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data" + "." + "clusters")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("data" + "." + "clusters")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this remove clusters from datacenter params data based on the context it is used
func (m *RemoveClustersFromDatacenterParamsData) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateClusters(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RemoveClustersFromDatacenterParamsData) contextValidateClusters(ctx context.Context, formats strfmt.Registry) error {

	if m.Clusters != nil {
		if err := m.Clusters.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data" + "." + "clusters")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("data" + "." + "clusters")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RemoveClustersFromDatacenterParamsData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RemoveClustersFromDatacenterParamsData) UnmarshalBinary(b []byte) error {
	var res RemoveClustersFromDatacenterParamsData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
