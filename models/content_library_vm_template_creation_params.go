// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ContentLibraryVMTemplateCreationParams content library Vm template creation params
//
// swagger:model ContentLibraryVmTemplateCreationParams
type ContentLibraryVMTemplateCreationParams struct {

	// cloud init supported
	CloudInitSupported *bool `json:"cloud_init_supported,omitempty"`

	// clusters
	// Required: true
	Clusters *ClusterWhereInput `json:"clusters"`

	// description
	Description *string `json:"description,omitempty"`

	// name
	// Required: true
	Name *string `json:"name"`

	// vm
	// Required: true
	VM *VMWhereUniqueInput `json:"vm"`
}

// Validate validates this content library Vm template creation params
func (m *ContentLibraryVMTemplateCreationParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClusters(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVM(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ContentLibraryVMTemplateCreationParams) validateClusters(formats strfmt.Registry) error {

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

func (m *ContentLibraryVMTemplateCreationParams) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *ContentLibraryVMTemplateCreationParams) validateVM(formats strfmt.Registry) error {

	if err := validate.Required("vm", "body", m.VM); err != nil {
		return err
	}

	if m.VM != nil {
		if err := m.VM.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vm")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vm")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this content library Vm template creation params based on the context it is used
func (m *ContentLibraryVMTemplateCreationParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateClusters(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVM(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ContentLibraryVMTemplateCreationParams) contextValidateClusters(ctx context.Context, formats strfmt.Registry) error {

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

func (m *ContentLibraryVMTemplateCreationParams) contextValidateVM(ctx context.Context, formats strfmt.Registry) error {

	if m.VM != nil {
		if err := m.VM.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vm")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vm")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ContentLibraryVMTemplateCreationParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ContentLibraryVMTemplateCreationParams) UnmarshalBinary(b []byte) error {
	var res ContentLibraryVMTemplateCreationParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
