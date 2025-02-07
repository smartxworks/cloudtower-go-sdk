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

// VMTemplateCreationParams Vm template creation params
//
// swagger:model VmTemplateCreationParams
type VMTemplateCreationParams struct {

	// cloud init supported
	// Required: true
	CloudInitSupported *bool `json:"cloud_init_supported"`

	// cluster id
	ClusterID *string `json:"cluster_id,omitempty"`

	// description
	Description *string `json:"description,omitempty"`

	// name
	// Required: true
	Name *string `json:"name"`

	// vm id
	// Required: true
	VMID *string `json:"vm_id"`

	MarshalOpts *VMTemplateCreationParamsMarshalOpts `json:"-"`
}

type VMTemplateCreationParamsMarshalOpts struct {
	CloudInitSupported_Explicit_Null_When_Empty bool

	ClusterID_Explicit_Null_When_Empty bool

	Description_Explicit_Null_When_Empty bool

	Name_Explicit_Null_When_Empty bool

	VMID_Explicit_Null_When_Empty bool
}

func (m VMTemplateCreationParams) MarshalJSON() ([]byte, error) {
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

	// handle nullable field cluster_id
	if m.ClusterID != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"cluster_id\":")
		bytes, err := swag.WriteJSON(m.ClusterID)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.ClusterID_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"cluster_id\":null")
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

	// handle nullable field vm_id
	if m.VMID != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"vm_id\":")
		bytes, err := swag.WriteJSON(m.VMID)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.VMID_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"vm_id\":null")
		first = false
	}

	b.WriteString("}")
	return b.Bytes(), nil
}

// Validate validates this Vm template creation params
func (m *VMTemplateCreationParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCloudInitSupported(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVMID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VMTemplateCreationParams) validateCloudInitSupported(formats strfmt.Registry) error {

	if err := validate.Required("cloud_init_supported", "body", m.CloudInitSupported); err != nil {
		return err
	}

	return nil
}

func (m *VMTemplateCreationParams) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *VMTemplateCreationParams) validateVMID(formats strfmt.Registry) error {

	if err := validate.Required("vm_id", "body", m.VMID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this Vm template creation params based on context it is used
func (m *VMTemplateCreationParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *VMTemplateCreationParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VMTemplateCreationParams) UnmarshalBinary(b []byte) error {
	var res VMTemplateCreationParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
