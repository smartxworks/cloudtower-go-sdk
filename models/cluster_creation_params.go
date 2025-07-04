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

// ClusterCreationParams cluster creation params
//
// swagger:model ClusterCreationParams
type ClusterCreationParams struct {

	// datacenter id
	DatacenterID *string `json:"datacenter_id,omitempty"`

	// ip
	// Required: true
	IP *string `json:"ip"`

	// password
	// Required: true
	Password *string `json:"password"`

	// primary zone datacenter id
	PrimaryZoneDatacenterID *string `json:"primary_zone_datacenter_id,omitempty"`

	// secondary zone datacenter id
	SecondaryZoneDatacenterID *string `json:"secondary_zone_datacenter_id,omitempty"`

	// username
	// Required: true
	Username *string `json:"username"`

	MarshalOpts *ClusterCreationParamsMarshalOpts `json:"-"`
}

type ClusterCreationParamsMarshalOpts struct {
	DatacenterID_Explicit_Null_When_Empty bool

	IP_Explicit_Null_When_Empty bool

	Password_Explicit_Null_When_Empty bool

	PrimaryZoneDatacenterID_Explicit_Null_When_Empty bool

	SecondaryZoneDatacenterID_Explicit_Null_When_Empty bool

	Username_Explicit_Null_When_Empty bool
}

func (m ClusterCreationParams) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	b.WriteString("{")

	first := true

	// handle nullable field datacenter_id
	if m.DatacenterID != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"datacenter_id\":")
		bytes, err := swag.WriteJSON(m.DatacenterID)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.DatacenterID_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"datacenter_id\":null")
		first = false
	}

	// handle nullable field ip
	if m.IP != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"ip\":")
		bytes, err := swag.WriteJSON(m.IP)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.IP_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"ip\":null")
		first = false
	}

	// handle nullable field password
	if m.Password != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"password\":")
		bytes, err := swag.WriteJSON(m.Password)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.Password_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"password\":null")
		first = false
	}

	// handle nullable field primary_zone_datacenter_id
	if m.PrimaryZoneDatacenterID != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"primary_zone_datacenter_id\":")
		bytes, err := swag.WriteJSON(m.PrimaryZoneDatacenterID)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.PrimaryZoneDatacenterID_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"primary_zone_datacenter_id\":null")
		first = false
	}

	// handle nullable field secondary_zone_datacenter_id
	if m.SecondaryZoneDatacenterID != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"secondary_zone_datacenter_id\":")
		bytes, err := swag.WriteJSON(m.SecondaryZoneDatacenterID)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.SecondaryZoneDatacenterID_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"secondary_zone_datacenter_id\":null")
		first = false
	}

	// handle nullable field username
	if m.Username != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"username\":")
		bytes, err := swag.WriteJSON(m.Username)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.Username_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"username\":null")
		first = false
	}

	b.WriteString("}")
	return b.Bytes(), nil
}

// Validate validates this cluster creation params
func (m *ClusterCreationParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIP(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePassword(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUsername(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterCreationParams) validateIP(formats strfmt.Registry) error {

	if err := validate.Required("ip", "body", m.IP); err != nil {
		return err
	}

	return nil
}

func (m *ClusterCreationParams) validatePassword(formats strfmt.Registry) error {

	if err := validate.Required("password", "body", m.Password); err != nil {
		return err
	}

	return nil
}

func (m *ClusterCreationParams) validateUsername(formats strfmt.Registry) error {

	if err := validate.Required("username", "body", m.Username); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this cluster creation params based on context it is used
func (m *ClusterCreationParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ClusterCreationParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterCreationParams) UnmarshalBinary(b []byte) error {
	var res ClusterCreationParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
