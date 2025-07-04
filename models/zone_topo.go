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

// ZoneTopo zone topo
//
// swagger:model ZoneTopo
type ZoneTopo struct {

	// cluster
	// Required: true
	Cluster *NestedCluster `json:"cluster"`

	// cluster topo
	// Required: true
	ClusterTopo *NestedClusterTopo `json:"cluster_topo"`

	// id
	// Required: true
	ID *string `json:"id"`

	// local id
	// Required: true
	LocalID *string `json:"local_id"`

	// rack topoes
	RackTopoes []*NestedRackTopo `json:"rack_topoes,omitempty"`

	MarshalOpts *ZoneTopoMarshalOpts `json:"-"`
}

type ZoneTopoMarshalOpts struct {
	Cluster_Explicit_Null_When_Empty bool

	ClusterTopo_Explicit_Null_When_Empty bool

	ID_Explicit_Null_When_Empty bool

	LocalID_Explicit_Null_When_Empty bool

	RackTopoes_Explicit_Null_When_Empty bool
}

func (m ZoneTopo) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	b.WriteString("{")

	first := true

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

	// handle nullable field cluster_topo
	if m.ClusterTopo != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"cluster_topo\":")
		bytes, err := swag.WriteJSON(m.ClusterTopo)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.ClusterTopo_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"cluster_topo\":null")
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

	// handle non nullable field rack_topoes with omitempty
	if !swag.IsZero(m.RackTopoes) {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"rack_topoes\":")
		bytes, err := swag.WriteJSON(m.RackTopoes)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	}

	b.WriteString("}")
	return b.Bytes(), nil
}

// Validate validates this zone topo
func (m *ZoneTopo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCluster(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClusterTopo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocalID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRackTopoes(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ZoneTopo) validateCluster(formats strfmt.Registry) error {

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

func (m *ZoneTopo) validateClusterTopo(formats strfmt.Registry) error {

	if err := validate.Required("cluster_topo", "body", m.ClusterTopo); err != nil {
		return err
	}

	if m.ClusterTopo != nil {
		if err := m.ClusterTopo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cluster_topo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cluster_topo")
			}
			return err
		}
	}

	return nil
}

func (m *ZoneTopo) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *ZoneTopo) validateLocalID(formats strfmt.Registry) error {

	if err := validate.Required("local_id", "body", m.LocalID); err != nil {
		return err
	}

	return nil
}

func (m *ZoneTopo) validateRackTopoes(formats strfmt.Registry) error {
	if swag.IsZero(m.RackTopoes) { // not required
		return nil
	}

	for i := 0; i < len(m.RackTopoes); i++ {
		if swag.IsZero(m.RackTopoes[i]) { // not required
			continue
		}

		if m.RackTopoes[i] != nil {
			if err := m.RackTopoes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("rack_topoes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("rack_topoes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this zone topo based on the context it is used
func (m *ZoneTopo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCluster(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateClusterTopo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRackTopoes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ZoneTopo) contextValidateCluster(ctx context.Context, formats strfmt.Registry) error {

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

func (m *ZoneTopo) contextValidateClusterTopo(ctx context.Context, formats strfmt.Registry) error {

	if m.ClusterTopo != nil {
		if err := m.ClusterTopo.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cluster_topo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cluster_topo")
			}
			return err
		}
	}

	return nil
}

func (m *ZoneTopo) contextValidateRackTopoes(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.RackTopoes); i++ {

		if m.RackTopoes[i] != nil {
			if err := m.RackTopoes[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("rack_topoes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("rack_topoes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ZoneTopo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ZoneTopo) UnmarshalBinary(b []byte) error {
	var res ZoneTopo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
