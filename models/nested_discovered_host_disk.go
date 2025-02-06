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

// NestedDiscoveredHostDisk nested discovered host disk
//
// swagger:model NestedDiscoveredHostDisk
type NestedDiscoveredHostDisk struct {

	// dimm ids
	DimmIds []string `json:"dimm_ids,omitempty"`

	// drive
	// Required: true
	Drive *string `json:"drive"`

	// function
	Function *DiskFunction `json:"function,omitempty"`

	// model
	// Required: true
	Model *string `json:"model"`

	// numa node
	NumaNode *int32 `json:"numa_node,omitempty"`

	// persistent memory type
	PersistentMemoryType *string `json:"persistent_memory_type,omitempty"`

	// serial
	// Required: true
	Serial *string `json:"serial"`

	// size
	// Required: true
	Size *int64 `json:"size"`

	// type
	// Required: true
	Type *DiskType `json:"type"`

	MarshalOpts *NestedDiscoveredHostDiskMarshalOpts `json:"-"`
}

type NestedDiscoveredHostDiskMarshalOpts struct {
	Drive_Explicit_Null_When_Empty bool

	Function_Explicit_Null_When_Empty bool

	Model_Explicit_Null_When_Empty bool

	NumaNode_Explicit_Null_When_Empty bool

	PersistentMemoryType_Explicit_Null_When_Empty bool

	Serial_Explicit_Null_When_Empty bool

	Size_Explicit_Null_When_Empty bool

	Type_Explicit_Null_When_Empty bool
}

func (m NestedDiscoveredHostDisk) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	b.WriteString("{")

	first := true

	// handle non nullable field dimm_ids with omitempty
	if swag.IsZero(m.DimmIds) {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"dimm_ids\":")
		bytes, err := swag.WriteJSON(m.DimmIds)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	}

	// handle nullable field drive
	if m.Drive != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"drive\":")
		bytes, err := swag.WriteJSON(m.Drive)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.Drive_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"drive\":null")
		first = false
	}

	// handle nullable field function
	if m.Function != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"function\":")
		bytes, err := swag.WriteJSON(m.Function)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.Function_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"function\":null")
		first = false
	}

	// handle nullable field model
	if m.Model != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"model\":")
		bytes, err := swag.WriteJSON(m.Model)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.Model_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"model\":null")
		first = false
	}

	// handle nullable field numa_node
	if m.NumaNode != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"numa_node\":")
		bytes, err := swag.WriteJSON(m.NumaNode)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.NumaNode_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"numa_node\":null")
		first = false
	}

	// handle nullable field persistent_memory_type
	if m.PersistentMemoryType != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"persistent_memory_type\":")
		bytes, err := swag.WriteJSON(m.PersistentMemoryType)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.PersistentMemoryType_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"persistent_memory_type\":null")
		first = false
	}

	// handle nullable field serial
	if m.Serial != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"serial\":")
		bytes, err := swag.WriteJSON(m.Serial)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.Serial_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"serial\":null")
		first = false
	}

	// handle nullable field size
	if m.Size != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"size\":")
		bytes, err := swag.WriteJSON(m.Size)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.Size_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"size\":null")
		first = false
	}

	// handle nullable field type
	if m.Type != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"type\":")
		bytes, err := swag.WriteJSON(m.Type)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.Type_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"type\":null")
		first = false
	}

	b.WriteString("}")
	return b.Bytes(), nil
}

// Validate validates this nested discovered host disk
func (m *NestedDiscoveredHostDisk) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDrive(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFunction(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateModel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSerial(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSize(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NestedDiscoveredHostDisk) validateDrive(formats strfmt.Registry) error {

	if err := validate.Required("drive", "body", m.Drive); err != nil {
		return err
	}

	return nil
}

func (m *NestedDiscoveredHostDisk) validateFunction(formats strfmt.Registry) error {
	if swag.IsZero(m.Function) { // not required
		return nil
	}

	if m.Function != nil {
		if err := m.Function.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("function")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("function")
			}
			return err
		}
	}

	return nil
}

func (m *NestedDiscoveredHostDisk) validateModel(formats strfmt.Registry) error {

	if err := validate.Required("model", "body", m.Model); err != nil {
		return err
	}

	return nil
}

func (m *NestedDiscoveredHostDisk) validateSerial(formats strfmt.Registry) error {

	if err := validate.Required("serial", "body", m.Serial); err != nil {
		return err
	}

	return nil
}

func (m *NestedDiscoveredHostDisk) validateSize(formats strfmt.Registry) error {

	if err := validate.Required("size", "body", m.Size); err != nil {
		return err
	}

	return nil
}

func (m *NestedDiscoveredHostDisk) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	if m.Type != nil {
		if err := m.Type.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("type")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("type")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this nested discovered host disk based on the context it is used
func (m *NestedDiscoveredHostDisk) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateFunction(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NestedDiscoveredHostDisk) contextValidateFunction(ctx context.Context, formats strfmt.Registry) error {

	if m.Function != nil {
		if err := m.Function.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("function")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("function")
			}
			return err
		}
	}

	return nil
}

func (m *NestedDiscoveredHostDisk) contextValidateType(ctx context.Context, formats strfmt.Registry) error {

	if m.Type != nil {
		if err := m.Type.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("type")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("type")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NestedDiscoveredHostDisk) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NestedDiscoveredHostDisk) UnmarshalBinary(b []byte) error {
	var res NestedDiscoveredHostDisk
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
