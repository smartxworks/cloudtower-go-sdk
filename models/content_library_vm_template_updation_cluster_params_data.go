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

// ContentLibraryVMTemplateUpdationClusterParamsData content library Vm template updation cluster params data
//
// swagger:model ContentLibraryVmTemplateUpdationClusterParamsData
type ContentLibraryVMTemplateUpdationClusterParamsData struct {

	// clusters
	// Required: true
	Clusters *ClusterWhereInput `json:"clusters"`

	MarshalOpts *ContentLibraryVMTemplateUpdationClusterParamsDataMarshalOpts `json:"-"`
}

type ContentLibraryVMTemplateUpdationClusterParamsDataMarshalOpts struct {
	Clusters_Explicit_Null_When_Empty bool
}

func (m ContentLibraryVMTemplateUpdationClusterParamsData) MarshalJSON() ([]byte, error) {
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

// Validate validates this content library Vm template updation cluster params data
func (m *ContentLibraryVMTemplateUpdationClusterParamsData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClusters(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ContentLibraryVMTemplateUpdationClusterParamsData) validateClusters(formats strfmt.Registry) error {

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

// ContextValidate validate this content library Vm template updation cluster params data based on the context it is used
func (m *ContentLibraryVMTemplateUpdationClusterParamsData) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateClusters(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ContentLibraryVMTemplateUpdationClusterParamsData) contextValidateClusters(ctx context.Context, formats strfmt.Registry) error {

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
func (m *ContentLibraryVMTemplateUpdationClusterParamsData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ContentLibraryVMTemplateUpdationClusterParamsData) UnmarshalBinary(b []byte) error {
	var res ContentLibraryVMTemplateUpdationClusterParamsData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
