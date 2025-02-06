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

// HostCreationParamsData host creation params data
//
// swagger:model HostCreationParamsData
type HostCreationParamsData struct {

	// disks
	// Required: true
	Disks []*HostBatchCreateDiskInput `json:"disks"`

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
	Ifaces []*HostBatchCreateIfaceInput `json:"ifaces"`

	// ipmi
	Ipmi *HostBatchCreateIpmiInput `json:"ipmi,omitempty"`

	// platform ip
	PlatformIP *string `json:"platform_ip,omitempty"`

	// platform password
	PlatformPassword *string `json:"platform_password,omitempty"`

	// platform username
	PlatformUsername *string `json:"platform_username,omitempty"`

	// vdses
	Vdses []*HostVdsConfig `json:"vdses,omitempty"`

	// zbs spec
	ZbsSpec *ZbsSpec `json:"zbs_spec,omitempty"`

	MarshalOpts *HostCreationParamsDataMarshalOpts `json:"-"`
}

type HostCreationParamsDataMarshalOpts struct {
	HostIP_Explicit_Null_When_Empty bool

	HostUUID_Explicit_Null_When_Empty bool

	Hostname_Explicit_Null_When_Empty bool

	Ipmi_Explicit_Null_When_Empty bool

	PlatformIP_Explicit_Null_When_Empty bool

	PlatformPassword_Explicit_Null_When_Empty bool

	PlatformUsername_Explicit_Null_When_Empty bool

	ZbsSpec_Explicit_Null_When_Empty bool
}

func (m HostCreationParamsData) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	b.WriteString("{")

	first := true

	// handle non nullable field disks without omitempty
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

	// handle nullable field ipmi
	if m.Ipmi != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"ipmi\":")
		bytes, err := swag.WriteJSON(m.Ipmi)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.Ipmi_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"ipmi\":null")
		first = false
	}

	// handle nullable field platform_ip
	if m.PlatformIP != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"platform_ip\":")
		bytes, err := swag.WriteJSON(m.PlatformIP)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.PlatformIP_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"platform_ip\":null")
		first = false
	}

	// handle nullable field platform_password
	if m.PlatformPassword != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"platform_password\":")
		bytes, err := swag.WriteJSON(m.PlatformPassword)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.PlatformPassword_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"platform_password\":null")
		first = false
	}

	// handle nullable field platform_username
	if m.PlatformUsername != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"platform_username\":")
		bytes, err := swag.WriteJSON(m.PlatformUsername)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.PlatformUsername_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"platform_username\":null")
		first = false
	}

	// handle non nullable field vdses with omitempty
	if swag.IsZero(m.Vdses) {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"vdses\":")
		bytes, err := swag.WriteJSON(m.Vdses)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
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

// Validate validates this host creation params data
func (m *HostCreationParamsData) Validate(formats strfmt.Registry) error {
	var res []error

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

	if err := m.validateIpmi(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVdses(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateZbsSpec(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HostCreationParamsData) validateDisks(formats strfmt.Registry) error {

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

func (m *HostCreationParamsData) validateHostIP(formats strfmt.Registry) error {

	if err := validate.Required("host_ip", "body", m.HostIP); err != nil {
		return err
	}

	return nil
}

func (m *HostCreationParamsData) validateHostUUID(formats strfmt.Registry) error {

	if err := validate.Required("host_uuid", "body", m.HostUUID); err != nil {
		return err
	}

	return nil
}

func (m *HostCreationParamsData) validateHostname(formats strfmt.Registry) error {

	if err := validate.Required("hostname", "body", m.Hostname); err != nil {
		return err
	}

	return nil
}

func (m *HostCreationParamsData) validateIfaces(formats strfmt.Registry) error {

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

func (m *HostCreationParamsData) validateIpmi(formats strfmt.Registry) error {
	if swag.IsZero(m.Ipmi) { // not required
		return nil
	}

	if m.Ipmi != nil {
		if err := m.Ipmi.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ipmi")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ipmi")
			}
			return err
		}
	}

	return nil
}

func (m *HostCreationParamsData) validateVdses(formats strfmt.Registry) error {
	if swag.IsZero(m.Vdses) { // not required
		return nil
	}

	for i := 0; i < len(m.Vdses); i++ {
		if swag.IsZero(m.Vdses[i]) { // not required
			continue
		}

		if m.Vdses[i] != nil {
			if err := m.Vdses[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("vdses" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("vdses" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *HostCreationParamsData) validateZbsSpec(formats strfmt.Registry) error {
	if swag.IsZero(m.ZbsSpec) { // not required
		return nil
	}

	if m.ZbsSpec != nil {
		if err := m.ZbsSpec.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("zbs_spec")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("zbs_spec")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this host creation params data based on the context it is used
func (m *HostCreationParamsData) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDisks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIfaces(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIpmi(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVdses(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateZbsSpec(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HostCreationParamsData) contextValidateDisks(ctx context.Context, formats strfmt.Registry) error {

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

func (m *HostCreationParamsData) contextValidateIfaces(ctx context.Context, formats strfmt.Registry) error {

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

func (m *HostCreationParamsData) contextValidateIpmi(ctx context.Context, formats strfmt.Registry) error {

	if m.Ipmi != nil {
		if err := m.Ipmi.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ipmi")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ipmi")
			}
			return err
		}
	}

	return nil
}

func (m *HostCreationParamsData) contextValidateVdses(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Vdses); i++ {

		if m.Vdses[i] != nil {
			if err := m.Vdses[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("vdses" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("vdses" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *HostCreationParamsData) contextValidateZbsSpec(ctx context.Context, formats strfmt.Registry) error {

	if m.ZbsSpec != nil {
		if err := m.ZbsSpec.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("zbs_spec")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("zbs_spec")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HostCreationParamsData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HostCreationParamsData) UnmarshalBinary(b []byte) error {
	var res HostCreationParamsData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
