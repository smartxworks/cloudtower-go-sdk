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
	"github.com/go-openapi/validate"
)

// EverouteLicense everoute license
//
// swagger:model EverouteLicense
type EverouteLicense struct {

	// code
	// Required: true
	Code *string `json:"code"`

	// expire date
	// Required: true
	ExpireDate *string `json:"expire_date"`

	// id
	// Required: true
	ID *string `json:"id"`

	// max socket num
	// Required: true
	MaxSocketNum *int32 `json:"max_socket_num"`

	// serial
	// Required: true
	Serial *string `json:"serial"`

	// sign date
	// Required: true
	SignDate *string `json:"sign_date"`

	// software edition
	// Required: true
	SoftwareEdition *SoftwareEdition `json:"software_edition"`

	// type
	// Required: true
	Type *LicenseType `json:"type"`

	// uid
	// Required: true
	UID *string `json:"uid"`

	MarshalOpts *EverouteLicenseMarshalOpts `json:"-"`
}

type EverouteLicenseMarshalOpts struct {
	Code_Explicit_Null_When_Empty bool

	ExpireDate_Explicit_Null_When_Empty bool

	ID_Explicit_Null_When_Empty bool

	MaxSocketNum_Explicit_Null_When_Empty bool

	Serial_Explicit_Null_When_Empty bool

	SignDate_Explicit_Null_When_Empty bool

	SoftwareEdition_Explicit_Null_When_Empty bool

	Type_Explicit_Null_When_Empty bool

	UID_Explicit_Null_When_Empty bool
}

func (m EverouteLicense) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	b.WriteString("{")

	first := true

	// handle nullable field code
	if m.Code != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"code\":")
		bytes, err := swag.WriteJSON(m.Code)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.Code_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"code\":null")
		first = false
	}

	// handle nullable field expire_date
	if m.ExpireDate != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"expire_date\":")
		bytes, err := swag.WriteJSON(m.ExpireDate)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.ExpireDate_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"expire_date\":null")
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

	// handle nullable field max_socket_num
	if m.MaxSocketNum != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"max_socket_num\":")
		bytes, err := swag.WriteJSON(m.MaxSocketNum)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.MaxSocketNum_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"max_socket_num\":null")
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

	// handle nullable field sign_date
	if m.SignDate != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"sign_date\":")
		bytes, err := swag.WriteJSON(m.SignDate)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.SignDate_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"sign_date\":null")
		first = false
	}

	// handle nullable field software_edition
	if m.SoftwareEdition != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"software_edition\":")
		bytes, err := swag.WriteJSON(m.SoftwareEdition)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.SoftwareEdition_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"software_edition\":null")
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

	// handle nullable field uid
	if m.UID != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"uid\":")
		bytes, err := swag.WriteJSON(m.UID)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.UID_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"uid\":null")
		first = false
	}

	b.WriteString("}")
	return b.Bytes(), nil
}

// Validate validates this everoute license
func (m *EverouteLicense) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExpireDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMaxSocketNum(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSerial(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSignDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSoftwareEdition(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EverouteLicense) validateCode(formats strfmt.Registry) error {

	if err := validate.Required("code", "body", m.Code); err != nil {
		return err
	}

	return nil
}

func (m *EverouteLicense) validateExpireDate(formats strfmt.Registry) error {

	if err := validate.Required("expire_date", "body", m.ExpireDate); err != nil {
		return err
	}

	return nil
}

func (m *EverouteLicense) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *EverouteLicense) validateMaxSocketNum(formats strfmt.Registry) error {

	if err := validate.Required("max_socket_num", "body", m.MaxSocketNum); err != nil {
		return err
	}

	return nil
}

func (m *EverouteLicense) validateSerial(formats strfmt.Registry) error {

	if err := validate.Required("serial", "body", m.Serial); err != nil {
		return err
	}

	return nil
}

func (m *EverouteLicense) validateSignDate(formats strfmt.Registry) error {

	if err := validate.Required("sign_date", "body", m.SignDate); err != nil {
		return err
	}

	return nil
}

func (m *EverouteLicense) validateSoftwareEdition(formats strfmt.Registry) error {

	if err := validate.Required("software_edition", "body", m.SoftwareEdition); err != nil {
		return err
	}

	if err := validate.Required("software_edition", "body", m.SoftwareEdition); err != nil {
		return err
	}

	if m.SoftwareEdition != nil {
		if err := m.SoftwareEdition.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("software_edition")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("software_edition")
			}
			return err
		}
	}

	return nil
}

func (m *EverouteLicense) validateType(formats strfmt.Registry) error {

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

func (m *EverouteLicense) validateUID(formats strfmt.Registry) error {

	if err := validate.Required("uid", "body", m.UID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this everoute license based on the context it is used
func (m *EverouteLicense) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSoftwareEdition(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EverouteLicense) contextValidateSoftwareEdition(ctx context.Context, formats strfmt.Registry) error {

	if m.SoftwareEdition != nil {
		if err := m.SoftwareEdition.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("software_edition")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("software_edition")
			}
			return err
		}
	}

	return nil
}

func (m *EverouteLicense) contextValidateType(ctx context.Context, formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *EverouteLicense) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EverouteLicense) UnmarshalBinary(b []byte) error {
	var res EverouteLicense
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
