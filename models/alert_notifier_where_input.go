// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AlertNotifierWhereInput alert notifier where input
//
// swagger:model AlertNotifierWhereInput
type AlertNotifierWhereInput struct {

	// a n d
	AND []*AlertNotifierWhereInput `json:"AND,omitempty"`

	// n o t
	NOT []*AlertNotifierWhereInput `json:"NOT,omitempty"`

	// o r
	OR []*AlertNotifierWhereInput `json:"OR,omitempty"`

	// clusters every
	ClustersEvery *ClusterWhereInput `json:"clusters_every,omitempty"`

	// clusters none
	ClustersNone *ClusterWhereInput `json:"clusters_none,omitempty"`

	// clusters some
	ClustersSome *ClusterWhereInput `json:"clusters_some,omitempty"`

	// disabled
	Disabled *bool `json:"disabled,omitempty"`

	// disabled not
	DisabledNot *bool `json:"disabled_not,omitempty"`

	// email from
	EmailFrom *string `json:"email_from,omitempty"`

	// email from contains
	EmailFromContains *string `json:"email_from_contains,omitempty"`

	// email from ends with
	EmailFromEndsWith *string `json:"email_from_ends_with,omitempty"`

	// email from gt
	EmailFromGt *string `json:"email_from_gt,omitempty"`

	// email from gte
	EmailFromGte *string `json:"email_from_gte,omitempty"`

	// email from in
	EmailFromIn []string `json:"email_from_in,omitempty"`

	// email from lt
	EmailFromLt *string `json:"email_from_lt,omitempty"`

	// email from lte
	EmailFromLte *string `json:"email_from_lte,omitempty"`

	// email from not
	EmailFromNot *string `json:"email_from_not,omitempty"`

	// email from not contains
	EmailFromNotContains *string `json:"email_from_not_contains,omitempty"`

	// email from not ends with
	EmailFromNotEndsWith *string `json:"email_from_not_ends_with,omitempty"`

	// email from not in
	EmailFromNotIn []string `json:"email_from_not_in,omitempty"`

	// email from not starts with
	EmailFromNotStartsWith *string `json:"email_from_not_starts_with,omitempty"`

	// email from starts with
	EmailFromStartsWith *string `json:"email_from_starts_with,omitempty"`

	// entity async status
	EntityAsyncStatus *EntityAsyncStatus `json:"entityAsyncStatus,omitempty"`

	// entity async status in
	EntityAsyncStatusIn []EntityAsyncStatus `json:"entityAsyncStatus_in,omitempty"`

	// entity async status not
	EntityAsyncStatusNot *EntityAsyncStatus `json:"entityAsyncStatus_not,omitempty"`

	// entity async status not in
	EntityAsyncStatusNotIn []EntityAsyncStatus `json:"entityAsyncStatus_not_in,omitempty"`

	// id
	ID *string `json:"id,omitempty"`

	// id contains
	IDContains *string `json:"id_contains,omitempty"`

	// id ends with
	IDEndsWith *string `json:"id_ends_with,omitempty"`

	// id gt
	IDGt *string `json:"id_gt,omitempty"`

	// id gte
	IDGte *string `json:"id_gte,omitempty"`

	// id in
	IDIn []string `json:"id_in,omitempty"`

	// id lt
	IDLt *string `json:"id_lt,omitempty"`

	// id lte
	IDLte *string `json:"id_lte,omitempty"`

	// id not
	IDNot *string `json:"id_not,omitempty"`

	// id not contains
	IDNotContains *string `json:"id_not_contains,omitempty"`

	// id not ends with
	IDNotEndsWith *string `json:"id_not_ends_with,omitempty"`

	// id not in
	IDNotIn []string `json:"id_not_in,omitempty"`

	// id not starts with
	IDNotStartsWith *string `json:"id_not_starts_with,omitempty"`

	// id starts with
	IDStartsWith *string `json:"id_starts_with,omitempty"`

	// language code
	LanguageCode *NotifierLanguageCode `json:"language_code,omitempty"`

	// language code in
	LanguageCodeIn []NotifierLanguageCode `json:"language_code_in,omitempty"`

	// language code not
	LanguageCodeNot *NotifierLanguageCode `json:"language_code_not,omitempty"`

	// language code not in
	LanguageCodeNotIn []NotifierLanguageCode `json:"language_code_not_in,omitempty"`

	// name
	Name *string `json:"name,omitempty"`

	// name contains
	NameContains *string `json:"name_contains,omitempty"`

	// name ends with
	NameEndsWith *string `json:"name_ends_with,omitempty"`

	// name gt
	NameGt *string `json:"name_gt,omitempty"`

	// name gte
	NameGte *string `json:"name_gte,omitempty"`

	// name in
	NameIn []string `json:"name_in,omitempty"`

	// name lt
	NameLt *string `json:"name_lt,omitempty"`

	// name lte
	NameLte *string `json:"name_lte,omitempty"`

	// name not
	NameNot *string `json:"name_not,omitempty"`

	// name not contains
	NameNotContains *string `json:"name_not_contains,omitempty"`

	// name not ends with
	NameNotEndsWith *string `json:"name_not_ends_with,omitempty"`

	// name not in
	NameNotIn []string `json:"name_not_in,omitempty"`

	// name not starts with
	NameNotStartsWith *string `json:"name_not_starts_with,omitempty"`

	// name starts with
	NameStartsWith *string `json:"name_starts_with,omitempty"`

	// security mode
	SecurityMode *NotifierSecurityMode `json:"security_mode,omitempty"`

	// security mode in
	SecurityModeIn []NotifierSecurityMode `json:"security_mode_in,omitempty"`

	// security mode not
	SecurityModeNot *NotifierSecurityMode `json:"security_mode_not,omitempty"`

	// security mode not in
	SecurityModeNotIn []NotifierSecurityMode `json:"security_mode_not_in,omitempty"`

	// smtp server config
	SMTPServerConfig *SMTPServerWhereInput `json:"smtp_server_config,omitempty"`

	// smtp server host
	SMTPServerHost *string `json:"smtp_server_host,omitempty"`

	// smtp server host contains
	SMTPServerHostContains *string `json:"smtp_server_host_contains,omitempty"`

	// smtp server host ends with
	SMTPServerHostEndsWith *string `json:"smtp_server_host_ends_with,omitempty"`

	// smtp server host gt
	SMTPServerHostGt *string `json:"smtp_server_host_gt,omitempty"`

	// smtp server host gte
	SMTPServerHostGte *string `json:"smtp_server_host_gte,omitempty"`

	// smtp server host in
	SMTPServerHostIn []string `json:"smtp_server_host_in,omitempty"`

	// smtp server host lt
	SMTPServerHostLt *string `json:"smtp_server_host_lt,omitempty"`

	// smtp server host lte
	SMTPServerHostLte *string `json:"smtp_server_host_lte,omitempty"`

	// smtp server host not
	SMTPServerHostNot *string `json:"smtp_server_host_not,omitempty"`

	// smtp server host not contains
	SMTPServerHostNotContains *string `json:"smtp_server_host_not_contains,omitempty"`

	// smtp server host not ends with
	SMTPServerHostNotEndsWith *string `json:"smtp_server_host_not_ends_with,omitempty"`

	// smtp server host not in
	SMTPServerHostNotIn []string `json:"smtp_server_host_not_in,omitempty"`

	// smtp server host not starts with
	SMTPServerHostNotStartsWith *string `json:"smtp_server_host_not_starts_with,omitempty"`

	// smtp server host starts with
	SMTPServerHostStartsWith *string `json:"smtp_server_host_starts_with,omitempty"`

	// smtp server port
	SMTPServerPort *int32 `json:"smtp_server_port,omitempty"`

	// smtp server port gt
	SMTPServerPortGt *int32 `json:"smtp_server_port_gt,omitempty"`

	// smtp server port gte
	SMTPServerPortGte *int32 `json:"smtp_server_port_gte,omitempty"`

	// smtp server port in
	SMTPServerPortIn []int32 `json:"smtp_server_port_in,omitempty"`

	// smtp server port lt
	SMTPServerPortLt *int32 `json:"smtp_server_port_lt,omitempty"`

	// smtp server port lte
	SMTPServerPortLte *int32 `json:"smtp_server_port_lte,omitempty"`

	// smtp server port not
	SMTPServerPortNot *int32 `json:"smtp_server_port_not,omitempty"`

	// smtp server port not in
	SMTPServerPortNotIn []int32 `json:"smtp_server_port_not_in,omitempty"`

	// username
	Username *string `json:"username,omitempty"`

	// username contains
	UsernameContains *string `json:"username_contains,omitempty"`

	// username ends with
	UsernameEndsWith *string `json:"username_ends_with,omitempty"`

	// username gt
	UsernameGt *string `json:"username_gt,omitempty"`

	// username gte
	UsernameGte *string `json:"username_gte,omitempty"`

	// username in
	UsernameIn []string `json:"username_in,omitempty"`

	// username lt
	UsernameLt *string `json:"username_lt,omitempty"`

	// username lte
	UsernameLte *string `json:"username_lte,omitempty"`

	// username not
	UsernameNot *string `json:"username_not,omitempty"`

	// username not contains
	UsernameNotContains *string `json:"username_not_contains,omitempty"`

	// username not ends with
	UsernameNotEndsWith *string `json:"username_not_ends_with,omitempty"`

	// username not in
	UsernameNotIn []string `json:"username_not_in,omitempty"`

	// username not starts with
	UsernameNotStartsWith *string `json:"username_not_starts_with,omitempty"`

	// username starts with
	UsernameStartsWith *string `json:"username_starts_with,omitempty"`
}

// Validate validates this alert notifier where input
func (m *AlertNotifierWhereInput) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAND(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNOT(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOR(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClustersEvery(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClustersNone(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClustersSome(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEntityAsyncStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEntityAsyncStatusIn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEntityAsyncStatusNot(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEntityAsyncStatusNotIn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLanguageCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLanguageCodeIn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLanguageCodeNot(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLanguageCodeNotIn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSecurityMode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSecurityModeIn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSecurityModeNot(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSecurityModeNotIn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSMTPServerConfig(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AlertNotifierWhereInput) validateAND(formats strfmt.Registry) error {
	if swag.IsZero(m.AND) { // not required
		return nil
	}

	for i := 0; i < len(m.AND); i++ {
		if swag.IsZero(m.AND[i]) { // not required
			continue
		}

		if m.AND[i] != nil {
			if err := m.AND[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("AND" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("AND" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AlertNotifierWhereInput) validateNOT(formats strfmt.Registry) error {
	if swag.IsZero(m.NOT) { // not required
		return nil
	}

	for i := 0; i < len(m.NOT); i++ {
		if swag.IsZero(m.NOT[i]) { // not required
			continue
		}

		if m.NOT[i] != nil {
			if err := m.NOT[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("NOT" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("NOT" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AlertNotifierWhereInput) validateOR(formats strfmt.Registry) error {
	if swag.IsZero(m.OR) { // not required
		return nil
	}

	for i := 0; i < len(m.OR); i++ {
		if swag.IsZero(m.OR[i]) { // not required
			continue
		}

		if m.OR[i] != nil {
			if err := m.OR[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("OR" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("OR" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AlertNotifierWhereInput) validateClustersEvery(formats strfmt.Registry) error {
	if swag.IsZero(m.ClustersEvery) { // not required
		return nil
	}

	if m.ClustersEvery != nil {
		if err := m.ClustersEvery.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("clusters_every")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("clusters_every")
			}
			return err
		}
	}

	return nil
}

func (m *AlertNotifierWhereInput) validateClustersNone(formats strfmt.Registry) error {
	if swag.IsZero(m.ClustersNone) { // not required
		return nil
	}

	if m.ClustersNone != nil {
		if err := m.ClustersNone.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("clusters_none")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("clusters_none")
			}
			return err
		}
	}

	return nil
}

func (m *AlertNotifierWhereInput) validateClustersSome(formats strfmt.Registry) error {
	if swag.IsZero(m.ClustersSome) { // not required
		return nil
	}

	if m.ClustersSome != nil {
		if err := m.ClustersSome.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("clusters_some")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("clusters_some")
			}
			return err
		}
	}

	return nil
}

func (m *AlertNotifierWhereInput) validateEntityAsyncStatus(formats strfmt.Registry) error {
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

func (m *AlertNotifierWhereInput) validateEntityAsyncStatusIn(formats strfmt.Registry) error {
	if swag.IsZero(m.EntityAsyncStatusIn) { // not required
		return nil
	}

	for i := 0; i < len(m.EntityAsyncStatusIn); i++ {

		if err := m.EntityAsyncStatusIn[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("entityAsyncStatus_in" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("entityAsyncStatus_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *AlertNotifierWhereInput) validateEntityAsyncStatusNot(formats strfmt.Registry) error {
	if swag.IsZero(m.EntityAsyncStatusNot) { // not required
		return nil
	}

	if m.EntityAsyncStatusNot != nil {
		if err := m.EntityAsyncStatusNot.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("entityAsyncStatus_not")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("entityAsyncStatus_not")
			}
			return err
		}
	}

	return nil
}

func (m *AlertNotifierWhereInput) validateEntityAsyncStatusNotIn(formats strfmt.Registry) error {
	if swag.IsZero(m.EntityAsyncStatusNotIn) { // not required
		return nil
	}

	for i := 0; i < len(m.EntityAsyncStatusNotIn); i++ {

		if err := m.EntityAsyncStatusNotIn[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("entityAsyncStatus_not_in" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("entityAsyncStatus_not_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *AlertNotifierWhereInput) validateLanguageCode(formats strfmt.Registry) error {
	if swag.IsZero(m.LanguageCode) { // not required
		return nil
	}

	if m.LanguageCode != nil {
		if err := m.LanguageCode.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("language_code")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("language_code")
			}
			return err
		}
	}

	return nil
}

func (m *AlertNotifierWhereInput) validateLanguageCodeIn(formats strfmt.Registry) error {
	if swag.IsZero(m.LanguageCodeIn) { // not required
		return nil
	}

	for i := 0; i < len(m.LanguageCodeIn); i++ {

		if err := m.LanguageCodeIn[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("language_code_in" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("language_code_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *AlertNotifierWhereInput) validateLanguageCodeNot(formats strfmt.Registry) error {
	if swag.IsZero(m.LanguageCodeNot) { // not required
		return nil
	}

	if m.LanguageCodeNot != nil {
		if err := m.LanguageCodeNot.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("language_code_not")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("language_code_not")
			}
			return err
		}
	}

	return nil
}

func (m *AlertNotifierWhereInput) validateLanguageCodeNotIn(formats strfmt.Registry) error {
	if swag.IsZero(m.LanguageCodeNotIn) { // not required
		return nil
	}

	for i := 0; i < len(m.LanguageCodeNotIn); i++ {

		if err := m.LanguageCodeNotIn[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("language_code_not_in" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("language_code_not_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *AlertNotifierWhereInput) validateSecurityMode(formats strfmt.Registry) error {
	if swag.IsZero(m.SecurityMode) { // not required
		return nil
	}

	if m.SecurityMode != nil {
		if err := m.SecurityMode.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("security_mode")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("security_mode")
			}
			return err
		}
	}

	return nil
}

func (m *AlertNotifierWhereInput) validateSecurityModeIn(formats strfmt.Registry) error {
	if swag.IsZero(m.SecurityModeIn) { // not required
		return nil
	}

	for i := 0; i < len(m.SecurityModeIn); i++ {

		if err := m.SecurityModeIn[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("security_mode_in" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("security_mode_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *AlertNotifierWhereInput) validateSecurityModeNot(formats strfmt.Registry) error {
	if swag.IsZero(m.SecurityModeNot) { // not required
		return nil
	}

	if m.SecurityModeNot != nil {
		if err := m.SecurityModeNot.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("security_mode_not")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("security_mode_not")
			}
			return err
		}
	}

	return nil
}

func (m *AlertNotifierWhereInput) validateSecurityModeNotIn(formats strfmt.Registry) error {
	if swag.IsZero(m.SecurityModeNotIn) { // not required
		return nil
	}

	for i := 0; i < len(m.SecurityModeNotIn); i++ {

		if err := m.SecurityModeNotIn[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("security_mode_not_in" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("security_mode_not_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *AlertNotifierWhereInput) validateSMTPServerConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.SMTPServerConfig) { // not required
		return nil
	}

	if m.SMTPServerConfig != nil {
		if err := m.SMTPServerConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("smtp_server_config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("smtp_server_config")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this alert notifier where input based on the context it is used
func (m *AlertNotifierWhereInput) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAND(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNOT(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOR(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateClustersEvery(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateClustersNone(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateClustersSome(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEntityAsyncStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEntityAsyncStatusIn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEntityAsyncStatusNot(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEntityAsyncStatusNotIn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLanguageCode(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLanguageCodeIn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLanguageCodeNot(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLanguageCodeNotIn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSecurityMode(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSecurityModeIn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSecurityModeNot(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSecurityModeNotIn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSMTPServerConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AlertNotifierWhereInput) contextValidateAND(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.AND); i++ {

		if m.AND[i] != nil {
			if err := m.AND[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("AND" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("AND" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AlertNotifierWhereInput) contextValidateNOT(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.NOT); i++ {

		if m.NOT[i] != nil {
			if err := m.NOT[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("NOT" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("NOT" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AlertNotifierWhereInput) contextValidateOR(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.OR); i++ {

		if m.OR[i] != nil {
			if err := m.OR[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("OR" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("OR" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AlertNotifierWhereInput) contextValidateClustersEvery(ctx context.Context, formats strfmt.Registry) error {

	if m.ClustersEvery != nil {
		if err := m.ClustersEvery.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("clusters_every")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("clusters_every")
			}
			return err
		}
	}

	return nil
}

func (m *AlertNotifierWhereInput) contextValidateClustersNone(ctx context.Context, formats strfmt.Registry) error {

	if m.ClustersNone != nil {
		if err := m.ClustersNone.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("clusters_none")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("clusters_none")
			}
			return err
		}
	}

	return nil
}

func (m *AlertNotifierWhereInput) contextValidateClustersSome(ctx context.Context, formats strfmt.Registry) error {

	if m.ClustersSome != nil {
		if err := m.ClustersSome.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("clusters_some")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("clusters_some")
			}
			return err
		}
	}

	return nil
}

func (m *AlertNotifierWhereInput) contextValidateEntityAsyncStatus(ctx context.Context, formats strfmt.Registry) error {

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

func (m *AlertNotifierWhereInput) contextValidateEntityAsyncStatusIn(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.EntityAsyncStatusIn); i++ {

		if err := m.EntityAsyncStatusIn[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("entityAsyncStatus_in" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("entityAsyncStatus_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *AlertNotifierWhereInput) contextValidateEntityAsyncStatusNot(ctx context.Context, formats strfmt.Registry) error {

	if m.EntityAsyncStatusNot != nil {
		if err := m.EntityAsyncStatusNot.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("entityAsyncStatus_not")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("entityAsyncStatus_not")
			}
			return err
		}
	}

	return nil
}

func (m *AlertNotifierWhereInput) contextValidateEntityAsyncStatusNotIn(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.EntityAsyncStatusNotIn); i++ {

		if err := m.EntityAsyncStatusNotIn[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("entityAsyncStatus_not_in" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("entityAsyncStatus_not_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *AlertNotifierWhereInput) contextValidateLanguageCode(ctx context.Context, formats strfmt.Registry) error {

	if m.LanguageCode != nil {
		if err := m.LanguageCode.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("language_code")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("language_code")
			}
			return err
		}
	}

	return nil
}

func (m *AlertNotifierWhereInput) contextValidateLanguageCodeIn(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.LanguageCodeIn); i++ {

		if err := m.LanguageCodeIn[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("language_code_in" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("language_code_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *AlertNotifierWhereInput) contextValidateLanguageCodeNot(ctx context.Context, formats strfmt.Registry) error {

	if m.LanguageCodeNot != nil {
		if err := m.LanguageCodeNot.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("language_code_not")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("language_code_not")
			}
			return err
		}
	}

	return nil
}

func (m *AlertNotifierWhereInput) contextValidateLanguageCodeNotIn(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.LanguageCodeNotIn); i++ {

		if err := m.LanguageCodeNotIn[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("language_code_not_in" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("language_code_not_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *AlertNotifierWhereInput) contextValidateSecurityMode(ctx context.Context, formats strfmt.Registry) error {

	if m.SecurityMode != nil {
		if err := m.SecurityMode.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("security_mode")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("security_mode")
			}
			return err
		}
	}

	return nil
}

func (m *AlertNotifierWhereInput) contextValidateSecurityModeIn(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.SecurityModeIn); i++ {

		if err := m.SecurityModeIn[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("security_mode_in" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("security_mode_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *AlertNotifierWhereInput) contextValidateSecurityModeNot(ctx context.Context, formats strfmt.Registry) error {

	if m.SecurityModeNot != nil {
		if err := m.SecurityModeNot.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("security_mode_not")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("security_mode_not")
			}
			return err
		}
	}

	return nil
}

func (m *AlertNotifierWhereInput) contextValidateSecurityModeNotIn(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.SecurityModeNotIn); i++ {

		if err := m.SecurityModeNotIn[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("security_mode_not_in" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("security_mode_not_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *AlertNotifierWhereInput) contextValidateSMTPServerConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.SMTPServerConfig != nil {
		if err := m.SMTPServerConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("smtp_server_config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("smtp_server_config")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AlertNotifierWhereInput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AlertNotifierWhereInput) UnmarshalBinary(b []byte) error {
	var res AlertNotifierWhereInput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
