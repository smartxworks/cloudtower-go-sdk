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

// MaintenanceModeVMInfo maintenance mode Vm info
//
// swagger:model MaintenanceModeVmInfo
type MaintenanceModeVMInfo struct {

	// state
	State *string `json:"state,omitempty"`

	// target host name
	TargetHostName *string `json:"target_host_name,omitempty"`

	// verify
	Verify *MaintenanceModeVerify `json:"verify,omitempty"`

	// vm ha
	VMHa *bool `json:"vm_ha,omitempty"`

	// vm name
	VMName *string `json:"vm_name,omitempty"`

	// vm state
	VMState *string `json:"vm_state,omitempty"`

	// vm uuid
	VMUUID *string `json:"vm_uuid,omitempty"`

	MarshalOpts *MaintenanceModeVMInfoMarshalOpts `json:"-"`
}

type MaintenanceModeVMInfoMarshalOpts struct {
	State_Explicit_Null_When_Empty bool

	TargetHostName_Explicit_Null_When_Empty bool

	Verify_Explicit_Null_When_Empty bool

	VMHa_Explicit_Null_When_Empty bool

	VMName_Explicit_Null_When_Empty bool

	VMState_Explicit_Null_When_Empty bool

	VMUUID_Explicit_Null_When_Empty bool
}

func (m MaintenanceModeVMInfo) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	b.WriteString("{")

	first := true

	// handle nullable field state
	if m.State != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"state\":")
		bytes, err := swag.WriteJSON(m.State)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.State_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"state\":null")
		first = false
	}

	// handle nullable field target_host_name
	if m.TargetHostName != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"target_host_name\":")
		bytes, err := swag.WriteJSON(m.TargetHostName)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.TargetHostName_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"target_host_name\":null")
		first = false
	}

	// handle nullable field verify
	if m.Verify != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"verify\":")
		bytes, err := swag.WriteJSON(m.Verify)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.Verify_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"verify\":null")
		first = false
	}

	// handle nullable field vm_ha
	if m.VMHa != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"vm_ha\":")
		bytes, err := swag.WriteJSON(m.VMHa)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.VMHa_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"vm_ha\":null")
		first = false
	}

	// handle nullable field vm_name
	if m.VMName != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"vm_name\":")
		bytes, err := swag.WriteJSON(m.VMName)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.VMName_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"vm_name\":null")
		first = false
	}

	// handle nullable field vm_state
	if m.VMState != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"vm_state\":")
		bytes, err := swag.WriteJSON(m.VMState)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.VMState_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"vm_state\":null")
		first = false
	}

	// handle nullable field vm_uuid
	if m.VMUUID != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"vm_uuid\":")
		bytes, err := swag.WriteJSON(m.VMUUID)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.VMUUID_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"vm_uuid\":null")
		first = false
	}

	b.WriteString("}")
	return b.Bytes(), nil
}

// Validate validates this maintenance mode Vm info
func (m *MaintenanceModeVMInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateVerify(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MaintenanceModeVMInfo) validateVerify(formats strfmt.Registry) error {
	if swag.IsZero(m.Verify) { // not required
		return nil
	}

	if m.Verify != nil {
		if err := m.Verify.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("verify")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("verify")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this maintenance mode Vm info based on the context it is used
func (m *MaintenanceModeVMInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateVerify(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MaintenanceModeVMInfo) contextValidateVerify(ctx context.Context, formats strfmt.Registry) error {

	if m.Verify != nil {
		if err := m.Verify.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("verify")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("verify")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MaintenanceModeVMInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MaintenanceModeVMInfo) UnmarshalBinary(b []byte) error {
	var res MaintenanceModeVMInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
