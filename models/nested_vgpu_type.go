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

// NestedVgpuType nested vgpu type
//
// swagger:model NestedVgpuType
type NestedVgpuType struct {

	// framebuffer
	Framebuffer *float64 `json:"framebuffer,omitempty"`

	// max instance
	MaxInstance *int32 `json:"max_instance,omitempty"`

	// max resolution
	MaxResolution *string `json:"max_resolution,omitempty"`

	// name
	Name *string `json:"name,omitempty"`

	// vgpu type id
	VgpuTypeID *string `json:"vgpu_type_id,omitempty"`

	MarshalOpts *NestedVgpuTypeMarshalOpts `json:"-"`
}

type NestedVgpuTypeMarshalOpts struct {
	Framebuffer_Explicit_Null_When_Empty bool

	MaxInstance_Explicit_Null_When_Empty bool

	MaxResolution_Explicit_Null_When_Empty bool

	Name_Explicit_Null_When_Empty bool

	VgpuTypeID_Explicit_Null_When_Empty bool
}

func (m NestedVgpuType) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	b.WriteString("{")

	first := true

	// handle nullable field framebuffer
	if m.Framebuffer != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"framebuffer\":")
		bytes, err := swag.WriteJSON(m.Framebuffer)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.Framebuffer_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"framebuffer\":null")
		first = false
	}

	// handle nullable field max_instance
	if m.MaxInstance != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"max_instance\":")
		bytes, err := swag.WriteJSON(m.MaxInstance)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.MaxInstance_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"max_instance\":null")
		first = false
	}

	// handle nullable field max_resolution
	if m.MaxResolution != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"max_resolution\":")
		bytes, err := swag.WriteJSON(m.MaxResolution)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.MaxResolution_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"max_resolution\":null")
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

	// handle nullable field vgpu_type_id
	if m.VgpuTypeID != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"vgpu_type_id\":")
		bytes, err := swag.WriteJSON(m.VgpuTypeID)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.VgpuTypeID_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"vgpu_type_id\":null")
		first = false
	}

	b.WriteString("}")
	return b.Bytes(), nil
}

// Validate validates this nested vgpu type
func (m *NestedVgpuType) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this nested vgpu type based on context it is used
func (m *NestedVgpuType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NestedVgpuType) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NestedVgpuType) UnmarshalBinary(b []byte) error {
	var res NestedVgpuType
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
