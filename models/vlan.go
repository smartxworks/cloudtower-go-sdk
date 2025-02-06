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

// Vlan vlan
//
// swagger:model Vlan
type Vlan struct {

	// entity async status
	EntityAsyncStatus *EntityAsyncStatus `json:"entityAsyncStatus,omitempty"`

	// gateway ip
	GatewayIP *string `json:"gateway_ip,omitempty"`

	// gateway subnetmask
	GatewaySubnetmask *string `json:"gateway_subnetmask,omitempty"`

	// id
	// Required: true
	ID *string `json:"id"`

	// labels
	Labels []*NestedLabel `json:"labels,omitempty"`

	// local id
	// Required: true
	LocalID *string `json:"local_id"`

	// mode type
	ModeType *VlanModeType `json:"mode_type,omitempty"`

	// name
	// Required: true
	Name *string `json:"name"`

	// network ids
	// Required: true
	NetworkIds []string `json:"network_ids"`

	// qos burst
	QosBurst *float64 `json:"qos_burst,omitempty"`

	// qos max bandwidth
	QosMaxBandwidth *float64 `json:"qos_max_bandwidth,omitempty"`

	// qos min bandwidth
	QosMinBandwidth *float64 `json:"qos_min_bandwidth,omitempty"`

	// qos priority
	QosPriority *int32 `json:"qos_priority,omitempty"`

	// subnetmask
	Subnetmask *string `json:"subnetmask,omitempty"`

	// type
	// Required: true
	Type *NetworkType `json:"type"`

	// vds
	// Required: true
	Vds *NestedVds `json:"vds"`

	// vlan id
	// Required: true
	VlanID *int32 `json:"vlan_id"`

	// vm nics
	VMNics []*NestedVMNic `json:"vm_nics,omitempty"`

	MarshalOpts *VlanMarshalOpts `json:"-"`
}

type VlanMarshalOpts struct {
	EntityAsyncStatus_Explicit_Null_When_Empty bool

	GatewayIP_Explicit_Null_When_Empty bool

	GatewaySubnetmask_Explicit_Null_When_Empty bool

	ID_Explicit_Null_When_Empty bool

	LocalID_Explicit_Null_When_Empty bool

	ModeType_Explicit_Null_When_Empty bool

	Name_Explicit_Null_When_Empty bool

	QosBurst_Explicit_Null_When_Empty bool

	QosMaxBandwidth_Explicit_Null_When_Empty bool

	QosMinBandwidth_Explicit_Null_When_Empty bool

	QosPriority_Explicit_Null_When_Empty bool

	Subnetmask_Explicit_Null_When_Empty bool

	Type_Explicit_Null_When_Empty bool

	Vds_Explicit_Null_When_Empty bool

	VlanID_Explicit_Null_When_Empty bool
}

func (m Vlan) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	b.WriteString("{")

	first := true

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

	// handle nullable field gateway_ip
	if m.GatewayIP != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"gateway_ip\":")
		bytes, err := swag.WriteJSON(m.GatewayIP)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.GatewayIP_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"gateway_ip\":null")
		first = false
	}

	// handle nullable field gateway_subnetmask
	if m.GatewaySubnetmask != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"gateway_subnetmask\":")
		bytes, err := swag.WriteJSON(m.GatewaySubnetmask)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.GatewaySubnetmask_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"gateway_subnetmask\":null")
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

	// handle non nullable field labels with omitempty
	if swag.IsZero(m.Labels) {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"labels\":")
		bytes, err := swag.WriteJSON(m.Labels)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
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

	// handle nullable field mode_type
	if m.ModeType != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"mode_type\":")
		bytes, err := swag.WriteJSON(m.ModeType)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.ModeType_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"mode_type\":null")
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

	// handle non nullable field network_ids without omitempty
	if !first {
		b.WriteString(",")
	}
	b.WriteString("\"network_ids\":")
	bytes, err := swag.WriteJSON(m.NetworkIds)
	if err != nil {
		return nil, err
	}
	b.Write(bytes)
	first = false

	// handle nullable field qos_burst
	if m.QosBurst != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"qos_burst\":")
		bytes, err := swag.WriteJSON(m.QosBurst)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.QosBurst_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"qos_burst\":null")
		first = false
	}

	// handle nullable field qos_max_bandwidth
	if m.QosMaxBandwidth != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"qos_max_bandwidth\":")
		bytes, err := swag.WriteJSON(m.QosMaxBandwidth)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.QosMaxBandwidth_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"qos_max_bandwidth\":null")
		first = false
	}

	// handle nullable field qos_min_bandwidth
	if m.QosMinBandwidth != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"qos_min_bandwidth\":")
		bytes, err := swag.WriteJSON(m.QosMinBandwidth)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.QosMinBandwidth_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"qos_min_bandwidth\":null")
		first = false
	}

	// handle nullable field qos_priority
	if m.QosPriority != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"qos_priority\":")
		bytes, err := swag.WriteJSON(m.QosPriority)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.QosPriority_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"qos_priority\":null")
		first = false
	}

	// handle nullable field subnetmask
	if m.Subnetmask != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"subnetmask\":")
		bytes, err := swag.WriteJSON(m.Subnetmask)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.Subnetmask_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"subnetmask\":null")
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

	// handle nullable field vds
	if m.Vds != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"vds\":")
		bytes, err := swag.WriteJSON(m.Vds)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.Vds_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"vds\":null")
		first = false
	}

	// handle nullable field vlan_id
	if m.VlanID != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"vlan_id\":")
		bytes, err := swag.WriteJSON(m.VlanID)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.VlanID_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"vlan_id\":null")
		first = false
	}

	// handle non nullable field vm_nics with omitempty
	if swag.IsZero(m.VMNics) {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"vm_nics\":")
		bytes, err := swag.WriteJSON(m.VMNics)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	}

	b.WriteString("}")
	return b.Bytes(), nil
}

// Validate validates this vlan
func (m *Vlan) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEntityAsyncStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLabels(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocalID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateModeType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNetworkIds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVlanID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVMNics(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Vlan) validateEntityAsyncStatus(formats strfmt.Registry) error {
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

func (m *Vlan) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *Vlan) validateLabels(formats strfmt.Registry) error {
	if swag.IsZero(m.Labels) { // not required
		return nil
	}

	for i := 0; i < len(m.Labels); i++ {
		if swag.IsZero(m.Labels[i]) { // not required
			continue
		}

		if m.Labels[i] != nil {
			if err := m.Labels[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("labels" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("labels" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Vlan) validateLocalID(formats strfmt.Registry) error {

	if err := validate.Required("local_id", "body", m.LocalID); err != nil {
		return err
	}

	return nil
}

func (m *Vlan) validateModeType(formats strfmt.Registry) error {
	if swag.IsZero(m.ModeType) { // not required
		return nil
	}

	if m.ModeType != nil {
		if err := m.ModeType.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mode_type")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("mode_type")
			}
			return err
		}
	}

	return nil
}

func (m *Vlan) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *Vlan) validateNetworkIds(formats strfmt.Registry) error {

	if err := validate.Required("network_ids", "body", m.NetworkIds); err != nil {
		return err
	}

	return nil
}

func (m *Vlan) validateType(formats strfmt.Registry) error {

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

func (m *Vlan) validateVds(formats strfmt.Registry) error {

	if err := validate.Required("vds", "body", m.Vds); err != nil {
		return err
	}

	if m.Vds != nil {
		if err := m.Vds.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vds")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vds")
			}
			return err
		}
	}

	return nil
}

func (m *Vlan) validateVlanID(formats strfmt.Registry) error {

	if err := validate.Required("vlan_id", "body", m.VlanID); err != nil {
		return err
	}

	return nil
}

func (m *Vlan) validateVMNics(formats strfmt.Registry) error {
	if swag.IsZero(m.VMNics) { // not required
		return nil
	}

	for i := 0; i < len(m.VMNics); i++ {
		if swag.IsZero(m.VMNics[i]) { // not required
			continue
		}

		if m.VMNics[i] != nil {
			if err := m.VMNics[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("vm_nics" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("vm_nics" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this vlan based on the context it is used
func (m *Vlan) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateEntityAsyncStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLabels(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateModeType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVds(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVMNics(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Vlan) contextValidateEntityAsyncStatus(ctx context.Context, formats strfmt.Registry) error {

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

func (m *Vlan) contextValidateLabels(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Labels); i++ {

		if m.Labels[i] != nil {
			if err := m.Labels[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("labels" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("labels" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Vlan) contextValidateModeType(ctx context.Context, formats strfmt.Registry) error {

	if m.ModeType != nil {
		if err := m.ModeType.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mode_type")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("mode_type")
			}
			return err
		}
	}

	return nil
}

func (m *Vlan) contextValidateType(ctx context.Context, formats strfmt.Registry) error {

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

func (m *Vlan) contextValidateVds(ctx context.Context, formats strfmt.Registry) error {

	if m.Vds != nil {
		if err := m.Vds.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vds")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vds")
			}
			return err
		}
	}

	return nil
}

func (m *Vlan) contextValidateVMNics(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.VMNics); i++ {

		if m.VMNics[i] != nil {
			if err := m.VMNics[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("vm_nics" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("vm_nics" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Vlan) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Vlan) UnmarshalBinary(b []byte) error {
	var res Vlan
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
