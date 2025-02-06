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

// VMCreateVMFromContentLibraryTemplateBatchParams Vm create Vm from content library template batch params
//
// swagger:model VmCreateVmFromContentLibraryTemplateBatchParams
type VMCreateVMFromContentLibraryTemplateBatchParams struct {

	// template id
	// Required: true
	TemplateID *string `json:"template_id"`

	// vms
	// Required: true
	Vms []*VMCreateVMFromContentLibraryTemplateBatchVMParams `json:"vms"`

	MarshalOpts *VMCreateVMFromContentLibraryTemplateBatchParamsMarshalOpts `json:"-"`
}

type VMCreateVMFromContentLibraryTemplateBatchParamsMarshalOpts struct {
	TemplateID_Explicit_Null_When_Empty bool

	Vms_Explicit_Null_When_Empty bool
}

func (m VMCreateVMFromContentLibraryTemplateBatchParams) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	b.WriteString("{")

	first := true

	// handle nullable field template_id
	if m.TemplateID != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"template_id\":")
		bytes, err := swag.WriteJSON(m.TemplateID)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.TemplateID_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"template_id\":null")
		first = false
	}

	// handle non nullable field vms without omitempty
	if !first {
		b.WriteString(",")
	}
	b.WriteString("\"vms\":")
	{
		bytes, err := swag.WriteJSON(m.Vms)
		if err != nil {
			return nil, err
		}
	}
	b.Write(bytes)
	first = false

	b.WriteString("}")
	return b.Bytes(), nil
}

// Validate validates this Vm create Vm from content library template batch params
func (m *VMCreateVMFromContentLibraryTemplateBatchParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTemplateID(formats); err != nil {
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

func (m *VMCreateVMFromContentLibraryTemplateBatchParams) validateTemplateID(formats strfmt.Registry) error {

	if err := validate.Required("template_id", "body", m.TemplateID); err != nil {
		return err
	}

	return nil
}

func (m *VMCreateVMFromContentLibraryTemplateBatchParams) validateVms(formats strfmt.Registry) error {

	if err := validate.Required("vms", "body", m.Vms); err != nil {
		return err
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

// ContextValidate validate this Vm create Vm from content library template batch params based on the context it is used
func (m *VMCreateVMFromContentLibraryTemplateBatchParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateVms(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VMCreateVMFromContentLibraryTemplateBatchParams) contextValidateVms(ctx context.Context, formats strfmt.Registry) error {

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
func (m *VMCreateVMFromContentLibraryTemplateBatchParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VMCreateVMFromContentLibraryTemplateBatchParams) UnmarshalBinary(b []byte) error {
	var res VMCreateVMFromContentLibraryTemplateBatchParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
