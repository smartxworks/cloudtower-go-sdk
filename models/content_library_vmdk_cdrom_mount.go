// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ContentLibraryVmdkCdromMount content library vmdk cdrom mount
//
// swagger:model ContentLibraryVmdkCdromMount
type ContentLibraryVmdkCdromMount struct {

	// boot
	Boot *int32 `json:"boot,omitempty"`

	// enabled
	Enabled *bool `json:"enabled,omitempty"`

	MarshalOpts *ContentLibraryVmdkCdromMountMarshalOpts `json:"-"`
}

type ContentLibraryVmdkCdromMountMarshalOpts struct {
	Boot_Explicit_Null_When_Empty bool

	Enabled_Explicit_Null_When_Empty bool
}

func (m ContentLibraryVmdkCdromMount) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	b.WriteString("{")

	first := true

	// handle nullable field boot
	if m.Boot != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"boot\":")
		bytes, err := swag.WriteJSON(m.Boot)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.Boot_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"boot\":null")
		first = false
	}

	// handle nullable field enabled
	if m.Enabled != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"enabled\":")
		bytes, err := swag.WriteJSON(m.Enabled)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.Enabled_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"enabled\":null")
		first = false
	}

	b.WriteString("}")
	return b.Bytes(), nil
}

// Validate validates this content library vmdk cdrom mount
func (m *ContentLibraryVmdkCdromMount) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this content library vmdk cdrom mount based on context it is used
func (m *ContentLibraryVmdkCdromMount) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ContentLibraryVmdkCdromMount) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ContentLibraryVmdkCdromMount) UnmarshalBinary(b []byte) error {
	var res ContentLibraryVmdkCdromMount
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
