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

// NestedGpuDriverInfo nested gpu driver info
//
// swagger:model NestedGpuDriverInfo
type NestedGpuDriverInfo struct {

	// filename
	Filename *string `json:"filename,omitempty"`

	// name
	Name *string `json:"name,omitempty"`

	// rhelversion
	Rhelversion *string `json:"rhelversion,omitempty"`

	// srcversion
	Srcversion *string `json:"srcversion,omitempty"`

	// supported
	Supported *string `json:"supported,omitempty"`

	// vermagic
	Vermagic *string `json:"vermagic,omitempty"`

	// version
	Version *string `json:"version,omitempty"`

	MarshalOpts *NestedGpuDriverInfoMarshalOpts `json:"-"`
}

type NestedGpuDriverInfoMarshalOpts struct {
	Filename_Explicit_Null_When_Empty bool

	Name_Explicit_Null_When_Empty bool

	Rhelversion_Explicit_Null_When_Empty bool

	Srcversion_Explicit_Null_When_Empty bool

	Supported_Explicit_Null_When_Empty bool

	Vermagic_Explicit_Null_When_Empty bool

	Version_Explicit_Null_When_Empty bool
}

func (m NestedGpuDriverInfo) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	b.WriteString("{")

	first := true

	// handle nullable field filename
	if m.Filename != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"filename\":")
		bytes, err := swag.WriteJSON(m.Filename)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.Filename_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"filename\":null")
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

	// handle nullable field rhelversion
	if m.Rhelversion != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"rhelversion\":")
		bytes, err := swag.WriteJSON(m.Rhelversion)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.Rhelversion_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"rhelversion\":null")
		first = false
	}

	// handle nullable field srcversion
	if m.Srcversion != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"srcversion\":")
		bytes, err := swag.WriteJSON(m.Srcversion)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.Srcversion_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"srcversion\":null")
		first = false
	}

	// handle nullable field supported
	if m.Supported != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"supported\":")
		bytes, err := swag.WriteJSON(m.Supported)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.Supported_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"supported\":null")
		first = false
	}

	// handle nullable field vermagic
	if m.Vermagic != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"vermagic\":")
		bytes, err := swag.WriteJSON(m.Vermagic)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.Vermagic_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"vermagic\":null")
		first = false
	}

	// handle nullable field version
	if m.Version != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"version\":")
		bytes, err := swag.WriteJSON(m.Version)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.Version_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"version\":null")
		first = false
	}

	b.WriteString("}")
	return b.Bytes(), nil
}

// Validate validates this nested gpu driver info
func (m *NestedGpuDriverInfo) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this nested gpu driver info based on context it is used
func (m *NestedGpuDriverInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NestedGpuDriverInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NestedGpuDriverInfo) UnmarshalBinary(b []byte) error {
	var res NestedGpuDriverInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
