// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// BrickTopoCreationParams brick topo creation params
//
// swagger:model BrickTopoCreationParams
type BrickTopoCreationParams struct {

	// capacity
	Capacity *NestedCapacity `json:"capacity,omitempty"`

	// cluster id
	// Required: true
	ClusterID *string `json:"cluster_id"`

	// height
	// Required: true
	Height *int32 `json:"height"`

	// name
	// Required: true
	Name *string `json:"name"`

	// node topoes
	NodeTopoes *NodeTopoWhereInput `json:"node_topoes,omitempty"`

	// position
	// Required: true
	Position *int32 `json:"position"`

	// rack topo id
	RackTopoID *string `json:"rack_topo_id,omitempty"`

	// tag position in brick
	TagPositionInBrick []*NestedTagPosition `json:"tag_position_in_brick,omitempty"`

	MarshalOpts *BrickTopoCreationParamsMarshalOpts `json:"-"`
}

type BrickTopoCreationParamsMarshalOpts struct {
	Capacity_Explicit_Null_When_Empty bool

	ClusterID_Explicit_Null_When_Empty bool

	Height_Explicit_Null_When_Empty bool

	Name_Explicit_Null_When_Empty bool

	NodeTopoes_Explicit_Null_When_Empty bool

	Position_Explicit_Null_When_Empty bool

	RackTopoID_Explicit_Null_When_Empty bool

	TagPositionInBrick_Explicit_Null_When_Empty bool
}

func (m BrickTopoCreationParams) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	b.WriteString("{")

	first := true

	// handle nullable field capacity
	if m.Capacity != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"capacity\":")
		bytes, err := swag.WriteJSON(m.Capacity)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.Capacity_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"capacity\":null")
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

	// handle nullable field height
	if m.Height != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"height\":")
		bytes, err := swag.WriteJSON(m.Height)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.Height_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"height\":null")
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

	// handle nullable field node_topoes
	if m.NodeTopoes != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"node_topoes\":")
		bytes, err := swag.WriteJSON(m.NodeTopoes)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.NodeTopoes_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"node_topoes\":null")
		first = false
	}

	// handle nullable field position
	if m.Position != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"position\":")
		bytes, err := swag.WriteJSON(m.Position)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.Position_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"position\":null")
		first = false
	}

	// handle nullable field rack_topo_id
	if m.RackTopoID != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"rack_topo_id\":")
		bytes, err := swag.WriteJSON(m.RackTopoID)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.RackTopoID_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"rack_topo_id\":null")
		first = false
	}

	// handle non nullable field tag_position_in_brick with omitempty
	if !swag.IsZero(m.TagPositionInBrick) {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"tag_position_in_brick\":")
		bytes, err := swag.WriteJSON(m.TagPositionInBrick)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	}

	b.WriteString("}")
	return b.Bytes(), nil
}

// Validate validates this brick topo creation params
func (m *BrickTopoCreationParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCapacity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClusterID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHeight(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNodeTopoes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePosition(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTagPositionInBrick(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BrickTopoCreationParams) validateCapacity(formats strfmt.Registry) error {
	if swag.IsZero(m.Capacity) { // not required
		return nil
	}

	if m.Capacity != nil {
		if err := m.Capacity.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("capacity")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("capacity")
			}
			return err
		}
	}

	return nil
}

func (m *BrickTopoCreationParams) validateClusterID(formats strfmt.Registry) error {

	if err := validate.Required("cluster_id", "body", m.ClusterID); err != nil {
		return err
	}

	return nil
}

func (m *BrickTopoCreationParams) validateHeight(formats strfmt.Registry) error {

	if err := validate.Required("height", "body", m.Height); err != nil {
		return err
	}

	return nil
}

func (m *BrickTopoCreationParams) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *BrickTopoCreationParams) validateNodeTopoes(formats strfmt.Registry) error {
	if swag.IsZero(m.NodeTopoes) { // not required
		return nil
	}

	if m.NodeTopoes != nil {
		if err := m.NodeTopoes.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("node_topoes")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("node_topoes")
			}
			return err
		}
	}

	return nil
}

func (m *BrickTopoCreationParams) validatePosition(formats strfmt.Registry) error {

	if err := validate.Required("position", "body", m.Position); err != nil {
		return err
	}

	return nil
}

func (m *BrickTopoCreationParams) validateTagPositionInBrick(formats strfmt.Registry) error {
	if swag.IsZero(m.TagPositionInBrick) { // not required
		return nil
	}

	for i := 0; i < len(m.TagPositionInBrick); i++ {
		if swag.IsZero(m.TagPositionInBrick[i]) { // not required
			continue
		}

		if m.TagPositionInBrick[i] != nil {
			if err := m.TagPositionInBrick[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tag_position_in_brick" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("tag_position_in_brick" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this brick topo creation params based on the context it is used
func (m *BrickTopoCreationParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCapacity(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNodeTopoes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTagPositionInBrick(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BrickTopoCreationParams) contextValidateCapacity(ctx context.Context, formats strfmt.Registry) error {

	if m.Capacity != nil {
		if err := m.Capacity.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("capacity")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("capacity")
			}
			return err
		}
	}

	return nil
}

func (m *BrickTopoCreationParams) contextValidateNodeTopoes(ctx context.Context, formats strfmt.Registry) error {

	if m.NodeTopoes != nil {
		if err := m.NodeTopoes.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("node_topoes")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("node_topoes")
			}
			return err
		}
	}

	return nil
}

func (m *BrickTopoCreationParams) contextValidateTagPositionInBrick(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.TagPositionInBrick); i++ {

		if m.TagPositionInBrick[i] != nil {
			if err := m.TagPositionInBrick[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tag_position_in_brick" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("tag_position_in_brick" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *BrickTopoCreationParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BrickTopoCreationParams) UnmarshalBinary(b []byte) error {
	var res BrickTopoCreationParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
