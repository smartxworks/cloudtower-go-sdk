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

// DiscoveredHost discovered host
//
// swagger:model DiscoveredHost
type DiscoveredHost struct {

	// all flash
	// Required: true
	AllFlash *bool `json:"all_flash"`

	// deployed
	Deployed *bool `json:"deployed,omitempty"`

	// dimms
	Dimms []*NestedDiscoveredHostDimms `json:"dimms,omitempty"`

	// disks
	// Required: true
	Disks []*NestedDiscoveredHostDisk `json:"disks"`

	// host ip
	// Required: true
	HostIP *string `json:"host_ip"`

	// host uuid
	// Required: true
	HostUUID *string `json:"host_uuid"`

	// hostname
	// Required: true
	Hostname *string `json:"hostname"`

	// ifaces
	// Required: true
	Ifaces []*NestedDiscoveredHostIface `json:"ifaces"`

	// ipmi ip
	IpmiIP *string `json:"ipmi_ip,omitempty"`

	// is os in raid1
	IsOsInRaid1 *bool `json:"is_os_in_raid1,omitempty"`

	// product
	Product *string `json:"product,omitempty"`

	// serial
	// Required: true
	Serial *string `json:"serial"`

	// sockets
	// Required: true
	Sockets *int32 `json:"sockets"`

	// version
	// Required: true
	Version *string `json:"version"`

	// zbs spec
	ZbsSpec *string `json:"zbs_spec,omitempty"`

	MarshalOpts *DiscoveredHostMarshalOpts `json:"-"`
}

type DiscoveredHostMarshalOpts struct {
	AllFlash_Explicit_Null_When_Empty bool

	Deployed_Explicit_Null_When_Empty bool

	Dimms_Explicit_Null_When_Empty bool

	Disks_Explicit_Null_When_Empty bool

	HostIP_Explicit_Null_When_Empty bool

	HostUUID_Explicit_Null_When_Empty bool

	Hostname_Explicit_Null_When_Empty bool

	Ifaces_Explicit_Null_When_Empty bool

	IpmiIP_Explicit_Null_When_Empty bool

	IsOsInRaid1_Explicit_Null_When_Empty bool

	Product_Explicit_Null_When_Empty bool

	Serial_Explicit_Null_When_Empty bool

	Sockets_Explicit_Null_When_Empty bool

	Version_Explicit_Null_When_Empty bool

	ZbsSpec_Explicit_Null_When_Empty bool
}

func (m DiscoveredHost) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	b.WriteString("{")

	first := true

	// handle nullable field all_flash
	if m.AllFlash != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"all_flash\":")
		bytes, err := swag.WriteJSON(m.AllFlash)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.AllFlash_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"all_flash\":null")
		first = false
	}

	// handle nullable field deployed
	if m.Deployed != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"deployed\":")
		bytes, err := swag.WriteJSON(m.Deployed)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.Deployed_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"deployed\":null")
		first = false
	}

	// handle non nullable field dimms with omitempty
	if !swag.IsZero(m.Dimms) {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"dimms\":")
		bytes, err := swag.WriteJSON(m.Dimms)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	}

	// handle non nullable field disks without omitempty
	{
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"disks\":")
		bytes, err := swag.WriteJSON(m.Disks)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	}

	// handle nullable field host_ip
	if m.HostIP != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"host_ip\":")
		bytes, err := swag.WriteJSON(m.HostIP)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.HostIP_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"host_ip\":null")
		first = false
	}

	// handle nullable field host_uuid
	if m.HostUUID != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"host_uuid\":")
		bytes, err := swag.WriteJSON(m.HostUUID)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.HostUUID_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"host_uuid\":null")
		first = false
	}

	// handle nullable field hostname
	if m.Hostname != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"hostname\":")
		bytes, err := swag.WriteJSON(m.Hostname)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.Hostname_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"hostname\":null")
		first = false
	}

	// handle non nullable field ifaces without omitempty
	{
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"ifaces\":")
		bytes, err := swag.WriteJSON(m.Ifaces)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	}

	// handle nullable field ipmi_ip
	if m.IpmiIP != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"ipmi_ip\":")
		bytes, err := swag.WriteJSON(m.IpmiIP)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.IpmiIP_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"ipmi_ip\":null")
		first = false
	}

	// handle nullable field is_os_in_raid1
	if m.IsOsInRaid1 != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"is_os_in_raid1\":")
		bytes, err := swag.WriteJSON(m.IsOsInRaid1)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.IsOsInRaid1_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"is_os_in_raid1\":null")
		first = false
	}

	// handle nullable field product
	if m.Product != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"product\":")
		bytes, err := swag.WriteJSON(m.Product)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.Product_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"product\":null")
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

	// handle nullable field sockets
	if m.Sockets != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"sockets\":")
		bytes, err := swag.WriteJSON(m.Sockets)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.Sockets_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"sockets\":null")
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

	// handle nullable field zbs_spec
	if m.ZbsSpec != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"zbs_spec\":")
		bytes, err := swag.WriteJSON(m.ZbsSpec)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.ZbsSpec_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"zbs_spec\":null")
		first = false
	}

	b.WriteString("}")
	return b.Bytes(), nil
}

// Validate validates this discovered host
func (m *DiscoveredHost) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAllFlash(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDimms(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDisks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHostIP(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHostUUID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHostname(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIfaces(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSerial(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSockets(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVersion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DiscoveredHost) validateAllFlash(formats strfmt.Registry) error {

	if err := validate.Required("all_flash", "body", m.AllFlash); err != nil {
		return err
	}

	return nil
}

func (m *DiscoveredHost) validateDimms(formats strfmt.Registry) error {
	if swag.IsZero(m.Dimms) { // not required
		return nil
	}

	for i := 0; i < len(m.Dimms); i++ {
		if swag.IsZero(m.Dimms[i]) { // not required
			continue
		}

		if m.Dimms[i] != nil {
			if err := m.Dimms[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("dimms" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("dimms" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DiscoveredHost) validateDisks(formats strfmt.Registry) error {

	if err := validate.Required("disks", "body", m.Disks); err != nil {
		return err
	}

	for i := 0; i < len(m.Disks); i++ {
		if swag.IsZero(m.Disks[i]) { // not required
			continue
		}

		if m.Disks[i] != nil {
			if err := m.Disks[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("disks" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("disks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DiscoveredHost) validateHostIP(formats strfmt.Registry) error {

	if err := validate.Required("host_ip", "body", m.HostIP); err != nil {
		return err
	}

	return nil
}

func (m *DiscoveredHost) validateHostUUID(formats strfmt.Registry) error {

	if err := validate.Required("host_uuid", "body", m.HostUUID); err != nil {
		return err
	}

	return nil
}

func (m *DiscoveredHost) validateHostname(formats strfmt.Registry) error {

	if err := validate.Required("hostname", "body", m.Hostname); err != nil {
		return err
	}

	return nil
}

func (m *DiscoveredHost) validateIfaces(formats strfmt.Registry) error {

	if err := validate.Required("ifaces", "body", m.Ifaces); err != nil {
		return err
	}

	for i := 0; i < len(m.Ifaces); i++ {
		if swag.IsZero(m.Ifaces[i]) { // not required
			continue
		}

		if m.Ifaces[i] != nil {
			if err := m.Ifaces[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ifaces" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ifaces" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DiscoveredHost) validateSerial(formats strfmt.Registry) error {

	if err := validate.Required("serial", "body", m.Serial); err != nil {
		return err
	}

	return nil
}

func (m *DiscoveredHost) validateSockets(formats strfmt.Registry) error {

	if err := validate.Required("sockets", "body", m.Sockets); err != nil {
		return err
	}

	return nil
}

func (m *DiscoveredHost) validateVersion(formats strfmt.Registry) error {

	if err := validate.Required("version", "body", m.Version); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this discovered host based on the context it is used
func (m *DiscoveredHost) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDimms(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDisks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIfaces(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DiscoveredHost) contextValidateDimms(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Dimms); i++ {

		if m.Dimms[i] != nil {
			if err := m.Dimms[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("dimms" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("dimms" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DiscoveredHost) contextValidateDisks(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Disks); i++ {

		if m.Disks[i] != nil {
			if err := m.Disks[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("disks" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("disks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DiscoveredHost) contextValidateIfaces(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Ifaces); i++ {

		if m.Ifaces[i] != nil {
			if err := m.Ifaces[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ifaces" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ifaces" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *DiscoveredHost) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DiscoveredHost) UnmarshalBinary(b []byte) error {
	var res DiscoveredHost
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
