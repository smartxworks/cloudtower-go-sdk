// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// WithTaskDisassociateSystemServiceFromObsServiceResult with task disassociate system service from obs service result
//
// swagger:model WithTask_DisassociateSystemServiceFromObsServiceResult_
type WithTaskDisassociateSystemServiceFromObsServiceResult struct {

	// data
	// Required: true
	Data *DisassociateSystemServiceFromObsServiceResult `json:"data"`

	// task id
	TaskID *string `json:"task_id,omitempty"`
}

// Validate validates this with task disassociate system service from obs service result
func (m *WithTaskDisassociateSystemServiceFromObsServiceResult) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WithTaskDisassociateSystemServiceFromObsServiceResult) validateData(formats strfmt.Registry) error {

	if err := validate.Required("data", "body", m.Data); err != nil {
		return err
	}

	if m.Data != nil {
		if err := m.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("data")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this with task disassociate system service from obs service result based on the context it is used
func (m *WithTaskDisassociateSystemServiceFromObsServiceResult) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WithTaskDisassociateSystemServiceFromObsServiceResult) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if m.Data != nil {
		if err := m.Data.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *WithTaskDisassociateSystemServiceFromObsServiceResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WithTaskDisassociateSystemServiceFromObsServiceResult) UnmarshalBinary(b []byte) error {
	var res WithTaskDisassociateSystemServiceFromObsServiceResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
