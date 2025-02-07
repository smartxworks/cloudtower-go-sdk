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

// ClusterTopo cluster topo
//
// swagger:model ClusterTopo
type ClusterTopo struct {

	// brick topoes
	BrickTopoes []*NestedBrickTopo `json:"brick_topoes,omitempty"`

	// cluster
	// Required: true
	Cluster *NestedCluster `json:"cluster"`

	// id
	// Required: true
	ID *string `json:"id"`

	// local id
	// Required: true
	LocalID *string `json:"local_id"`

	// name
	// Required: true
	Name *string `json:"name"`

	// node topoes
	NodeTopoes []*NestedNodeTopo `json:"node_topoes,omitempty"`

	// zone topoes
	ZoneTopoes []*NestedZoneTopo `json:"zone_topoes,omitempty"`

	MarshalOpts *ClusterTopoMarshalOpts `json:"-"`
}

type ClusterTopoMarshalOpts struct {
	BrickTopoes_Explicit_Null_When_Empty bool

	Cluster_Explicit_Null_When_Empty bool

	ID_Explicit_Null_When_Empty bool

	LocalID_Explicit_Null_When_Empty bool

	Name_Explicit_Null_When_Empty bool

	NodeTopoes_Explicit_Null_When_Empty bool

	ZoneTopoes_Explicit_Null_When_Empty bool
}

func (m ClusterTopo) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	b.WriteString("{")

	first := true

	// handle non nullable field brick_topoes with omitempty
	if !swag.IsZero(m.BrickTopoes) {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"brick_topoes\":")
		bytes, err := swag.WriteJSON(m.BrickTopoes)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	}

	// handle nullable field cluster
	if m.Cluster != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"cluster\":")
		bytes, err := swag.WriteJSON(m.Cluster)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.Cluster_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"cluster\":null")
		first = false
	}

	// handle nullable field id
	if m.ID != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"id\":")
		bytes, err := swag.WriteJSON(m.ID)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.ID_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"id\":null")
		first = false
	}

	// handle nullable field local_id
	if m.LocalID != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"local_id\":")
		bytes, err := swag.WriteJSON(m.LocalID)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.LocalID_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"local_id\":null")
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

	// handle non nullable field node_topoes with omitempty
	if !swag.IsZero(m.NodeTopoes) {
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
	}

	// handle non nullable field zone_topoes with omitempty
	if !swag.IsZero(m.ZoneTopoes) {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"zone_topoes\":")
		bytes, err := swag.WriteJSON(m.ZoneTopoes)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	}

	b.WriteString("}")
	return b.Bytes(), nil
}

// Validate validates this cluster topo
func (m *ClusterTopo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBrickTopoes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCluster(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocalID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNodeTopoes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateZoneTopoes(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterTopo) validateBrickTopoes(formats strfmt.Registry) error {
	if swag.IsZero(m.BrickTopoes) { // not required
		return nil
	}

	for i := 0; i < len(m.BrickTopoes); i++ {
		if swag.IsZero(m.BrickTopoes[i]) { // not required
			continue
		}

		if m.BrickTopoes[i] != nil {
			if err := m.BrickTopoes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("brick_topoes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("brick_topoes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ClusterTopo) validateCluster(formats strfmt.Registry) error {

	if err := validate.Required("cluster", "body", m.Cluster); err != nil {
		return err
	}

	if m.Cluster != nil {
		if err := m.Cluster.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cluster")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cluster")
			}
			return err
		}
	}

	return nil
}

func (m *ClusterTopo) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *ClusterTopo) validateLocalID(formats strfmt.Registry) error {

	if err := validate.Required("local_id", "body", m.LocalID); err != nil {
		return err
	}

	return nil
}

func (m *ClusterTopo) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *ClusterTopo) validateNodeTopoes(formats strfmt.Registry) error {
	if swag.IsZero(m.NodeTopoes) { // not required
		return nil
	}

	for i := 0; i < len(m.NodeTopoes); i++ {
		if swag.IsZero(m.NodeTopoes[i]) { // not required
			continue
		}

		if m.NodeTopoes[i] != nil {
			if err := m.NodeTopoes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("node_topoes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("node_topoes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ClusterTopo) validateZoneTopoes(formats strfmt.Registry) error {
	if swag.IsZero(m.ZoneTopoes) { // not required
		return nil
	}

	for i := 0; i < len(m.ZoneTopoes); i++ {
		if swag.IsZero(m.ZoneTopoes[i]) { // not required
			continue
		}

		if m.ZoneTopoes[i] != nil {
			if err := m.ZoneTopoes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("zone_topoes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("zone_topoes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this cluster topo based on the context it is used
func (m *ClusterTopo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBrickTopoes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCluster(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNodeTopoes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateZoneTopoes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterTopo) contextValidateBrickTopoes(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.BrickTopoes); i++ {

		if m.BrickTopoes[i] != nil {
			if err := m.BrickTopoes[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("brick_topoes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("brick_topoes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ClusterTopo) contextValidateCluster(ctx context.Context, formats strfmt.Registry) error {

	if m.Cluster != nil {
		if err := m.Cluster.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cluster")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cluster")
			}
			return err
		}
	}

	return nil
}

func (m *ClusterTopo) contextValidateNodeTopoes(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.NodeTopoes); i++ {

		if m.NodeTopoes[i] != nil {
			if err := m.NodeTopoes[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("node_topoes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("node_topoes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ClusterTopo) contextValidateZoneTopoes(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ZoneTopoes); i++ {

		if m.ZoneTopoes[i] != nil {
			if err := m.ZoneTopoes[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("zone_topoes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("zone_topoes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ClusterTopo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterTopo) UnmarshalBinary(b []byte) error {
	var res ClusterTopo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
