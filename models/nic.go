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

// Nic nic
//
// swagger:model Nic
type Nic struct {

	// driver
	Driver *string `json:"driver,omitempty"`

	// driver state
	DriverState *NicDriverState `json:"driver_state,omitempty"`

	// entity async status
	EntityAsyncStatus *EntityAsyncStatus `json:"entityAsyncStatus,omitempty"`

	// gateway ip
	GatewayIP *string `json:"gateway_ip,omitempty"`

	// host
	// Required: true
	Host *NestedHost `json:"host"`

	// ibdev
	Ibdev *string `json:"ibdev,omitempty"`

	// id
	// Required: true
	ID *string `json:"id"`

	// iommu status
	IommuStatus *IommuStatus `json:"iommu_status,omitempty"`

	// ip address
	IPAddress *string `json:"ip_address,omitempty"`

	// is sriov
	IsSriov *bool `json:"is_sriov,omitempty"`

	// labels
	Labels []*NestedLabel `json:"labels,omitempty"`

	// local id
	// Required: true
	LocalID *string `json:"local_id"`

	// mac address
	// Required: true
	MacAddress *string `json:"mac_address"`

	// max vf num
	MaxVfNum *int32 `json:"max_vf_num,omitempty"`

	// model
	Model *string `json:"model,omitempty"`

	// mtu
	// Required: true
	Mtu *int32 `json:"mtu"`

	// name
	// Required: true
	Name *string `json:"name"`

	// nic uuid
	NicUUID *string `json:"nic_uuid,omitempty"`

	// physical
	// Required: true
	Physical *bool `json:"physical"`

	// rdma enabled
	RdmaEnabled *bool `json:"rdma_enabled,omitempty"`

	// running
	// Required: true
	Running *bool `json:"running"`

	// speed
	Speed *int64 `json:"speed,omitempty"`

	// subnet mask
	SubnetMask *string `json:"subnet_mask,omitempty"`

	// total vf num
	TotalVfNum *int32 `json:"total_vf_num,omitempty"`

	// type
	Type *NetworkType `json:"type,omitempty"`

	// up
	// Required: true
	Up *bool `json:"up"`

	// used vf num
	UsedVfNum *int32 `json:"used_vf_num,omitempty"`

	// user usage
	UserUsage *NicUserUsage `json:"user_usage,omitempty"`

	// vds
	Vds *NestedVds `json:"vds,omitempty"`

	// vms
	Vms []*NestedVM `json:"vms,omitempty"`

	MarshalOpts *NicMarshalOpts `json:"-"`
}

type NicMarshalOpts struct {
	Driver_Explicit_Null_When_Empty bool

	DriverState_Explicit_Null_When_Empty bool

	EntityAsyncStatus_Explicit_Null_When_Empty bool

	GatewayIP_Explicit_Null_When_Empty bool

	Host_Explicit_Null_When_Empty bool

	Ibdev_Explicit_Null_When_Empty bool

	ID_Explicit_Null_When_Empty bool

	IommuStatus_Explicit_Null_When_Empty bool

	IPAddress_Explicit_Null_When_Empty bool

	IsSriov_Explicit_Null_When_Empty bool

	Labels_Explicit_Null_When_Empty bool

	LocalID_Explicit_Null_When_Empty bool

	MacAddress_Explicit_Null_When_Empty bool

	MaxVfNum_Explicit_Null_When_Empty bool

	Model_Explicit_Null_When_Empty bool

	Mtu_Explicit_Null_When_Empty bool

	Name_Explicit_Null_When_Empty bool

	NicUUID_Explicit_Null_When_Empty bool

	Physical_Explicit_Null_When_Empty bool

	RdmaEnabled_Explicit_Null_When_Empty bool

	Running_Explicit_Null_When_Empty bool

	Speed_Explicit_Null_When_Empty bool

	SubnetMask_Explicit_Null_When_Empty bool

	TotalVfNum_Explicit_Null_When_Empty bool

	Type_Explicit_Null_When_Empty bool

	Up_Explicit_Null_When_Empty bool

	UsedVfNum_Explicit_Null_When_Empty bool

	UserUsage_Explicit_Null_When_Empty bool

	Vds_Explicit_Null_When_Empty bool

	Vms_Explicit_Null_When_Empty bool
}

func (m Nic) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	b.WriteString("{")

	first := true

	// handle nullable field driver
	if m.Driver != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"driver\":")
		bytes, err := swag.WriteJSON(m.Driver)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.Driver_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"driver\":null")
		first = false
	}

	// handle nullable field driver_state
	if m.DriverState != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"driver_state\":")
		bytes, err := swag.WriteJSON(m.DriverState)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.DriverState_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"driver_state\":null")
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

	// handle nullable field host
	if m.Host != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"host\":")
		bytes, err := swag.WriteJSON(m.Host)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.Host_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"host\":null")
		first = false
	}

	// handle nullable field ibdev
	if m.Ibdev != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"ibdev\":")
		bytes, err := swag.WriteJSON(m.Ibdev)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.Ibdev_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"ibdev\":null")
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

	// handle nullable field iommu_status
	if m.IommuStatus != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"iommu_status\":")
		bytes, err := swag.WriteJSON(m.IommuStatus)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.IommuStatus_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"iommu_status\":null")
		first = false
	}

	// handle nullable field ip_address
	if m.IPAddress != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"ip_address\":")
		bytes, err := swag.WriteJSON(m.IPAddress)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.IPAddress_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"ip_address\":null")
		first = false
	}

	// handle nullable field is_sriov
	if m.IsSriov != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"is_sriov\":")
		bytes, err := swag.WriteJSON(m.IsSriov)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.IsSriov_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"is_sriov\":null")
		first = false
	}

	// handle non nullable field labels with omitempty
	if !swag.IsZero(m.Labels) {
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

	// handle nullable field mac_address
	if m.MacAddress != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"mac_address\":")
		bytes, err := swag.WriteJSON(m.MacAddress)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.MacAddress_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"mac_address\":null")
		first = false
	}

	// handle nullable field max_vf_num
	if m.MaxVfNum != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"max_vf_num\":")
		bytes, err := swag.WriteJSON(m.MaxVfNum)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.MaxVfNum_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"max_vf_num\":null")
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

	// handle nullable field mtu
	if m.Mtu != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"mtu\":")
		bytes, err := swag.WriteJSON(m.Mtu)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.Mtu_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"mtu\":null")
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

	// handle nullable field nic_uuid
	if m.NicUUID != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"nic_uuid\":")
		bytes, err := swag.WriteJSON(m.NicUUID)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.NicUUID_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"nic_uuid\":null")
		first = false
	}

	// handle nullable field physical
	if m.Physical != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"physical\":")
		bytes, err := swag.WriteJSON(m.Physical)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.Physical_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"physical\":null")
		first = false
	}

	// handle nullable field rdma_enabled
	if m.RdmaEnabled != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"rdma_enabled\":")
		bytes, err := swag.WriteJSON(m.RdmaEnabled)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.RdmaEnabled_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"rdma_enabled\":null")
		first = false
	}

	// handle nullable field running
	if m.Running != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"running\":")
		bytes, err := swag.WriteJSON(m.Running)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.Running_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"running\":null")
		first = false
	}

	// handle nullable field speed
	if m.Speed != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"speed\":")
		bytes, err := swag.WriteJSON(m.Speed)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.Speed_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"speed\":null")
		first = false
	}

	// handle nullable field subnet_mask
	if m.SubnetMask != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"subnet_mask\":")
		bytes, err := swag.WriteJSON(m.SubnetMask)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.SubnetMask_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"subnet_mask\":null")
		first = false
	}

	// handle nullable field total_vf_num
	if m.TotalVfNum != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"total_vf_num\":")
		bytes, err := swag.WriteJSON(m.TotalVfNum)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.TotalVfNum_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"total_vf_num\":null")
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

	// handle nullable field up
	if m.Up != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"up\":")
		bytes, err := swag.WriteJSON(m.Up)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.Up_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"up\":null")
		first = false
	}

	// handle nullable field used_vf_num
	if m.UsedVfNum != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"used_vf_num\":")
		bytes, err := swag.WriteJSON(m.UsedVfNum)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.UsedVfNum_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"used_vf_num\":null")
		first = false
	}

	// handle nullable field user_usage
	if m.UserUsage != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"user_usage\":")
		bytes, err := swag.WriteJSON(m.UserUsage)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.UserUsage_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"user_usage\":null")
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

	// handle non nullable field vms with omitempty
	if !swag.IsZero(m.Vms) {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"vms\":")
		bytes, err := swag.WriteJSON(m.Vms)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	}

	b.WriteString("}")
	return b.Bytes(), nil
}

// Validate validates this nic
func (m *Nic) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDriverState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEntityAsyncStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHost(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIommuStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLabels(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocalID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMacAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMtu(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePhysical(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRunning(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserUsage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVms(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Nic) validateDriverState(formats strfmt.Registry) error {
	if swag.IsZero(m.DriverState) { // not required
		return nil
	}

	if m.DriverState != nil {
		if err := m.DriverState.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("driver_state")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("driver_state")
			}
			return err
		}
	}

	return nil
}

func (m *Nic) validateEntityAsyncStatus(formats strfmt.Registry) error {
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

func (m *Nic) validateHost(formats strfmt.Registry) error {

	if err := validate.Required("host", "body", m.Host); err != nil {
		return err
	}

	if m.Host != nil {
		if err := m.Host.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("host")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("host")
			}
			return err
		}
	}

	return nil
}

func (m *Nic) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *Nic) validateIommuStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.IommuStatus) { // not required
		return nil
	}

	if m.IommuStatus != nil {
		if err := m.IommuStatus.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("iommu_status")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("iommu_status")
			}
			return err
		}
	}

	return nil
}

func (m *Nic) validateLabels(formats strfmt.Registry) error {
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

func (m *Nic) validateLocalID(formats strfmt.Registry) error {

	if err := validate.Required("local_id", "body", m.LocalID); err != nil {
		return err
	}

	return nil
}

func (m *Nic) validateMacAddress(formats strfmt.Registry) error {

	if err := validate.Required("mac_address", "body", m.MacAddress); err != nil {
		return err
	}

	return nil
}

func (m *Nic) validateMtu(formats strfmt.Registry) error {

	if err := validate.Required("mtu", "body", m.Mtu); err != nil {
		return err
	}

	return nil
}

func (m *Nic) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *Nic) validatePhysical(formats strfmt.Registry) error {

	if err := validate.Required("physical", "body", m.Physical); err != nil {
		return err
	}

	return nil
}

func (m *Nic) validateRunning(formats strfmt.Registry) error {

	if err := validate.Required("running", "body", m.Running); err != nil {
		return err
	}

	return nil
}

func (m *Nic) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
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

func (m *Nic) validateUp(formats strfmt.Registry) error {

	if err := validate.Required("up", "body", m.Up); err != nil {
		return err
	}

	return nil
}

func (m *Nic) validateUserUsage(formats strfmt.Registry) error {
	if swag.IsZero(m.UserUsage) { // not required
		return nil
	}

	if m.UserUsage != nil {
		if err := m.UserUsage.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("user_usage")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("user_usage")
			}
			return err
		}
	}

	return nil
}

func (m *Nic) validateVds(formats strfmt.Registry) error {
	if swag.IsZero(m.Vds) { // not required
		return nil
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

func (m *Nic) validateVms(formats strfmt.Registry) error {
	if swag.IsZero(m.Vms) { // not required
		return nil
	}

	for i := 0; i < len(m.Vms); i++ {
		if swag.IsZero(m.Vms[i]) { // not required
			continue
		}

		if m.Vms[i] != nil {
			if err := m.Vms[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("vms" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("vms" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this nic based on the context it is used
func (m *Nic) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDriverState(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEntityAsyncStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateHost(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIommuStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLabels(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUserUsage(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVds(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVms(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Nic) contextValidateDriverState(ctx context.Context, formats strfmt.Registry) error {

	if m.DriverState != nil {
		if err := m.DriverState.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("driver_state")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("driver_state")
			}
			return err
		}
	}

	return nil
}

func (m *Nic) contextValidateEntityAsyncStatus(ctx context.Context, formats strfmt.Registry) error {

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

func (m *Nic) contextValidateHost(ctx context.Context, formats strfmt.Registry) error {

	if m.Host != nil {
		if err := m.Host.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("host")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("host")
			}
			return err
		}
	}

	return nil
}

func (m *Nic) contextValidateIommuStatus(ctx context.Context, formats strfmt.Registry) error {

	if m.IommuStatus != nil {
		if err := m.IommuStatus.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("iommu_status")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("iommu_status")
			}
			return err
		}
	}

	return nil
}

func (m *Nic) contextValidateLabels(ctx context.Context, formats strfmt.Registry) error {

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

func (m *Nic) contextValidateType(ctx context.Context, formats strfmt.Registry) error {

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

func (m *Nic) contextValidateUserUsage(ctx context.Context, formats strfmt.Registry) error {

	if m.UserUsage != nil {
		if err := m.UserUsage.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("user_usage")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("user_usage")
			}
			return err
		}
	}

	return nil
}

func (m *Nic) contextValidateVds(ctx context.Context, formats strfmt.Registry) error {

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

func (m *Nic) contextValidateVms(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Vms); i++ {

		if m.Vms[i] != nil {
			if err := m.Vms[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("vms" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("vms" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Nic) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Nic) UnmarshalBinary(b []byte) error {
	var res Nic
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
