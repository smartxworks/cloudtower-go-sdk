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

// VdsCreationParams vds creation params
//
// swagger:model VdsCreationParams
type VdsCreationParams struct {

	// bond mode
	BondMode *string `json:"bond_mode,omitempty"`

	// cluster id
	// Required: true
	ClusterID *string `json:"cluster_id"`

	// name
	// Required: true
	Name *string `json:"name"`

	// nic ids
	// Required: true
	NicIds []string `json:"nic_ids"`

	// work mode
	WorkMode *string `json:"work_mode,omitempty"`

	MarshalOpts *VdsCreationParamsMarshalOpts `json:"-"`
}

type VdsCreationParamsMarshalOpts struct {
	BondMode_Explicit_Null_When_Empty bool

	ClusterID_Explicit_Null_When_Empty bool

	Name_Explicit_Null_When_Empty bool

	NicIds_Explicit_Null_When_Empty bool

	WorkMode_Explicit_Null_When_Empty bool
}

func (m VdsCreationParams) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	b.WriteString("{")

	first := true

	// handle nullable field bond_mode
	if m.BondMode != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"bond_mode\":")
		bytes, err := swag.WriteJSON(m.BondMode)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.BondMode_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"bond_mode\":null")
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

	// handle non nullable field nic_ids without omitempty
	{
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"nic_ids\":")
		bytes, err := swag.WriteJSON(m.NicIds)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	}

	// handle nullable field work_mode
	if m.WorkMode != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"work_mode\":")
		bytes, err := swag.WriteJSON(m.WorkMode)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.WorkMode_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"work_mode\":null")
		first = false
	}

	b.WriteString("}")
	return b.Bytes(), nil
}

// Validate validates this vds creation params
func (m *VdsCreationParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClusterID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNicIds(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VdsCreationParams) validateClusterID(formats strfmt.Registry) error {

	if err := validate.Required("cluster_id", "body", m.ClusterID); err != nil {
		return err
	}

	return nil
}

func (m *VdsCreationParams) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *VdsCreationParams) validateNicIds(formats strfmt.Registry) error {

	if err := validate.Required("nic_ids", "body", m.NicIds); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this vds creation params based on context it is used
func (m *VdsCreationParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *VdsCreationParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VdsCreationParams) UnmarshalBinary(b []byte) error {
	var res VdsCreationParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
