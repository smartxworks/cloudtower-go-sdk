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
)

// ContentLibraryVMTemplateUpdationParamsData content library Vm template updation params data
//
// swagger:model ContentLibraryVmTemplateUpdationParamsData
type ContentLibraryVMTemplateUpdationParamsData struct {

	// cloud init supported
	CloudInitSupported *bool `json:"cloud_init_supported,omitempty"`

	// clusters
	Clusters *ClusterWhereInput `json:"clusters,omitempty"`

	// description
	Description *string `json:"description,omitempty"`

	// name
	Name *string `json:"name,omitempty"`

	MarshalOpts *ContentLibraryVMTemplateUpdationParamsDataMarshalOpts `json:"-"`
}

type ContentLibraryVMTemplateUpdationParamsDataMarshalOpts struct {
	CloudInitSupported_Explicit_Null_When_Empty bool

	Clusters_Explicit_Null_When_Empty bool

	Description_Explicit_Null_When_Empty bool

	Name_Explicit_Null_When_Empty bool
}

func (m ContentLibraryVMTemplateUpdationParamsData) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	b.WriteString("{")

	first := true

	// handle nullable field cloud_init_supported
	if m.CloudInitSupported != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"cloud_init_supported\":")
		bytes, err := swag.WriteJSON(m.CloudInitSupported)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.CloudInitSupported_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"cloud_init_supported\":null")
		first = false
	}

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

	// handle nullable field description
	if m.Description != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"description\":")
		bytes, err := swag.WriteJSON(m.Description)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.Description_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"description\":null")
		first = false
	}

	// handle nullable field name
	if m.Name != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"name\":")
		bytes, err := swag.WriteJSON(m.Name)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.Name_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"name\":null")
		first = false
	}

	b.WriteString("}")
	return b.Bytes(), nil
}

// Validate validates this content library Vm template updation params data
func (m *ContentLibraryVMTemplateUpdationParamsData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClusters(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ContentLibraryVMTemplateUpdationParamsData) validateClusters(formats strfmt.Registry) error {
	if swag.IsZero(m.Clusters) { // not required
		return nil
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

// ContextValidate validate this content library Vm template updation params data based on the context it is used
func (m *ContentLibraryVMTemplateUpdationParamsData) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateClusters(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ContentLibraryVMTemplateUpdationParamsData) contextValidateClusters(ctx context.Context, formats strfmt.Registry) error {

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
func (m *ContentLibraryVMTemplateUpdationParamsData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ContentLibraryVMTemplateUpdationParamsData) UnmarshalBinary(b []byte) error {
	var res ContentLibraryVMTemplateUpdationParamsData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
