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

// VirtualPrivateCloudEdgeGatewayGroup virtual private cloud edge gateway group
//
// swagger:model VirtualPrivateCloudEdgeGatewayGroup
type VirtualPrivateCloudEdgeGatewayGroup struct {

	// active edge gateway ids
	// Required: true
	ActiveEdgeGatewayIds []string `json:"active_edge_gateway_ids"`

	// description
	Description *string `json:"description,omitempty"`

	// edge gateways
	EdgeGateways []*NestedVirtualPrivateCloudEdgeGateway `json:"edge_gateways,omitempty"`

	// entity async status
	EntityAsyncStatus *EntityAsyncStatus `json:"entityAsyncStatus,omitempty"`

	// id
	// Required: true
	ID *string `json:"id"`

	// name
	// Required: true
	Name *string `json:"name"`

	// primary edge gateway id
	PrimaryEdgeGatewayID *string `json:"primary_edge_gateway_id,omitempty"`

	MarshalOpts *VirtualPrivateCloudEdgeGatewayGroupMarshalOpts `json:"-"`
}

type VirtualPrivateCloudEdgeGatewayGroupMarshalOpts struct {
	ActiveEdgeGatewayIds_Explicit_Null_When_Empty bool

	Description_Explicit_Null_When_Empty bool

	EdgeGateways_Explicit_Null_When_Empty bool

	EntityAsyncStatus_Explicit_Null_When_Empty bool

	ID_Explicit_Null_When_Empty bool

	Name_Explicit_Null_When_Empty bool

	PrimaryEdgeGatewayID_Explicit_Null_When_Empty bool
}

func (m VirtualPrivateCloudEdgeGatewayGroup) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	b.WriteString("{")

	first := true

	// handle non nullable field active_edge_gateway_ids without omitempty
	{
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"active_edge_gateway_ids\":")
		bytes, err := swag.WriteJSON(m.ActiveEdgeGatewayIds)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	}

	// handle nullable field description
	if m.Description != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"description\":")
		bytes, err := swag.WriteJSON(m.Description)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.Description_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"description\":null")
		first = false
	}

	// handle non nullable field edge_gateways with omitempty
	if !swag.IsZero(m.EdgeGateways) {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"edge_gateways\":")
		bytes, err := swag.WriteJSON(m.EdgeGateways)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	}

	// handle nullable field entityAsyncStatus
	if m.EntityAsyncStatus != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"entityAsyncStatus\":")
		bytes, err := swag.WriteJSON(m.EntityAsyncStatus)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.EntityAsyncStatus_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"entityAsyncStatus\":null")
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

	// handle nullable field primary_edge_gateway_id
	if m.PrimaryEdgeGatewayID != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"primary_edge_gateway_id\":")
		bytes, err := swag.WriteJSON(m.PrimaryEdgeGatewayID)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.PrimaryEdgeGatewayID_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"primary_edge_gateway_id\":null")
		first = false
	}

	b.WriteString("}")
	return b.Bytes(), nil
}

// Validate validates this virtual private cloud edge gateway group
func (m *VirtualPrivateCloudEdgeGatewayGroup) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActiveEdgeGatewayIds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEdgeGateways(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEntityAsyncStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VirtualPrivateCloudEdgeGatewayGroup) validateActiveEdgeGatewayIds(formats strfmt.Registry) error {

	if err := validate.Required("active_edge_gateway_ids", "body", m.ActiveEdgeGatewayIds); err != nil {
		return err
	}

	return nil
}

func (m *VirtualPrivateCloudEdgeGatewayGroup) validateEdgeGateways(formats strfmt.Registry) error {
	if swag.IsZero(m.EdgeGateways) { // not required
		return nil
	}

	for i := 0; i < len(m.EdgeGateways); i++ {
		if swag.IsZero(m.EdgeGateways[i]) { // not required
			continue
		}

		if m.EdgeGateways[i] != nil {
			if err := m.EdgeGateways[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("edge_gateways" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("edge_gateways" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *VirtualPrivateCloudEdgeGatewayGroup) validateEntityAsyncStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.EntityAsyncStatus) { // not required
		return nil
	}

	if m.EntityAsyncStatus != nil {
		if err := m.EntityAsyncStatus.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("entityAsyncStatus")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("entityAsyncStatus")
			}
			return err
		}
	}

	return nil
}

func (m *VirtualPrivateCloudEdgeGatewayGroup) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *VirtualPrivateCloudEdgeGatewayGroup) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this virtual private cloud edge gateway group based on the context it is used
func (m *VirtualPrivateCloudEdgeGatewayGroup) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateEdgeGateways(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEntityAsyncStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VirtualPrivateCloudEdgeGatewayGroup) contextValidateEdgeGateways(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.EdgeGateways); i++ {

		if m.EdgeGateways[i] != nil {
			if err := m.EdgeGateways[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("edge_gateways" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("edge_gateways" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *VirtualPrivateCloudEdgeGatewayGroup) contextValidateEntityAsyncStatus(ctx context.Context, formats strfmt.Registry) error {

	if m.EntityAsyncStatus != nil {
		if err := m.EntityAsyncStatus.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("entityAsyncStatus")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("entityAsyncStatus")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *VirtualPrivateCloudEdgeGatewayGroup) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VirtualPrivateCloudEdgeGatewayGroup) UnmarshalBinary(b []byte) error {
	var res VirtualPrivateCloudEdgeGatewayGroup
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
