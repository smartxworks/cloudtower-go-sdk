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

// ContentLibraryVmdkCdromModify content library vmdk cdrom modify
//
// swagger:model ContentLibraryVmdkCdromModify
type ContentLibraryVmdkCdromModify struct {

	// boot
	Boot *int32 `json:"boot,omitempty"`

	// enabled
	Enabled *bool `json:"enabled,omitempty"`

	// index
	// Required: true
	Index *int32 `json:"index"`

	// removed
	Removed *bool `json:"removed,omitempty"`
}

// Validate validates this content library vmdk cdrom modify
func (m *ContentLibraryVmdkCdromModify) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIndex(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ContentLibraryVmdkCdromModify) validateIndex(formats strfmt.Registry) error {

	if err := validate.Required("index", "body", m.Index); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this content library vmdk cdrom modify based on context it is used
func (m *ContentLibraryVmdkCdromModify) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ContentLibraryVmdkCdromModify) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ContentLibraryVmdkCdromModify) UnmarshalBinary(b []byte) error {
	var res ContentLibraryVmdkCdromModify
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
