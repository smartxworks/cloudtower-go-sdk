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

// GetVirtualPrivateCloudRouteTablesConnectionRequestBody get virtual private cloud route tables connection request body
//
// swagger:model GetVirtualPrivateCloudRouteTablesConnectionRequestBody
type GetVirtualPrivateCloudRouteTablesConnectionRequestBody struct {

	// after
	After *string `json:"after,omitempty"`

	// before
	Before *string `json:"before,omitempty"`

	// first
	First *int32 `json:"first,omitempty"`

	// last
	Last *int32 `json:"last,omitempty"`

	// order by
	OrderBy *VirtualPrivateCloudRouteTableOrderByInput `json:"orderBy,omitempty"`

	// skip
	Skip *int32 `json:"skip,omitempty"`

	// where
	Where *VirtualPrivateCloudRouteTableWhereInput `json:"where,omitempty"`

	MarshalOpts *GetVirtualPrivateCloudRouteTablesConnectionRequestBodyMarshalOpts `json:"-"`
}

type GetVirtualPrivateCloudRouteTablesConnectionRequestBodyMarshalOpts struct {
	After_Explicit_Null_When_Empty bool

	Before_Explicit_Null_When_Empty bool

	First_Explicit_Null_When_Empty bool

	Last_Explicit_Null_When_Empty bool

	OrderBy_Explicit_Null_When_Empty bool

	Skip_Explicit_Null_When_Empty bool

	Where_Explicit_Null_When_Empty bool
}

func (m GetVirtualPrivateCloudRouteTablesConnectionRequestBody) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	b.WriteString("{")

	first := true

	// handle nullable field after
	if m.After != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"after\":")
		bytes, err := swag.WriteJSON(m.After)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.After_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"after\":null")
		first = false
	}

	// handle nullable field before
	if m.Before != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"before\":")
		bytes, err := swag.WriteJSON(m.Before)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.Before_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"before\":null")
		first = false
	}

	// handle nullable field first
	if m.First != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"first\":")
		bytes, err := swag.WriteJSON(m.First)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.First_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"first\":null")
		first = false
	}

	// handle nullable field last
	if m.Last != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"last\":")
		bytes, err := swag.WriteJSON(m.Last)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.Last_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"last\":null")
		first = false
	}

	// handle nullable field orderBy
	if m.OrderBy != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"orderBy\":")
		bytes, err := swag.WriteJSON(m.OrderBy)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.OrderBy_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"orderBy\":null")
		first = false
	}

	// handle nullable field skip
	if m.Skip != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"skip\":")
		bytes, err := swag.WriteJSON(m.Skip)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.Skip_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"skip\":null")
		first = false
	}

	// handle nullable field where
	if m.Where != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"where\":")
		bytes, err := swag.WriteJSON(m.Where)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.Where_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"where\":null")
		first = false
	}

	b.WriteString("}")
	return b.Bytes(), nil
}

// Validate validates this get virtual private cloud route tables connection request body
func (m *GetVirtualPrivateCloudRouteTablesConnectionRequestBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOrderBy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWhere(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetVirtualPrivateCloudRouteTablesConnectionRequestBody) validateOrderBy(formats strfmt.Registry) error {
	if swag.IsZero(m.OrderBy) { // not required
		return nil
	}

	if m.OrderBy != nil {
		if err := m.OrderBy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("orderBy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("orderBy")
			}
			return err
		}
	}

	return nil
}

func (m *GetVirtualPrivateCloudRouteTablesConnectionRequestBody) validateWhere(formats strfmt.Registry) error {
	if swag.IsZero(m.Where) { // not required
		return nil
	}

	if m.Where != nil {
		if err := m.Where.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("where")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("where")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get virtual private cloud route tables connection request body based on the context it is used
func (m *GetVirtualPrivateCloudRouteTablesConnectionRequestBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateOrderBy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateWhere(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetVirtualPrivateCloudRouteTablesConnectionRequestBody) contextValidateOrderBy(ctx context.Context, formats strfmt.Registry) error {

	if m.OrderBy != nil {
		if err := m.OrderBy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("orderBy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("orderBy")
			}
			return err
		}
	}

	return nil
}

func (m *GetVirtualPrivateCloudRouteTablesConnectionRequestBody) contextValidateWhere(ctx context.Context, formats strfmt.Registry) error {

	if m.Where != nil {
		if err := m.Where.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("where")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("where")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetVirtualPrivateCloudRouteTablesConnectionRequestBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetVirtualPrivateCloudRouteTablesConnectionRequestBody) UnmarshalBinary(b []byte) error {
	var res GetVirtualPrivateCloudRouteTablesConnectionRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
