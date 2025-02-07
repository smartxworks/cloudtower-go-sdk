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
)

// UserRoleNextWhereInput user role next where input
//
// swagger:model UserRoleNextWhereInput
type UserRoleNextWhereInput struct {

	// a n d
	AND []*UserRoleNextWhereInput `json:"AND,omitempty"`

	// n o t
	NOT []*UserRoleNextWhereInput `json:"NOT,omitempty"`

	// o r
	OR []*UserRoleNextWhereInput `json:"OR,omitempty"`

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

	// platform
	Platform *UserRolePlatform `json:"platform,omitempty"`

	// platform in
	PlatformIn []UserRolePlatform `json:"platform_in,omitempty"`

	// platform not
	PlatformNot *UserRolePlatform `json:"platform_not,omitempty"`

	// platform not in
	PlatformNotIn []UserRolePlatform `json:"platform_not_in,omitempty"`

	// preset
	Preset *UserRolePreset `json:"preset,omitempty"`

	// preset in
	PresetIn []UserRolePreset `json:"preset_in,omitempty"`

	// preset not
	PresetNot *UserRolePreset `json:"preset_not,omitempty"`

	// preset not in
	PresetNotIn []UserRolePreset `json:"preset_not_in,omitempty"`

	// users every
	UsersEvery *UserWhereInput `json:"users_every,omitempty"`

	// users none
	UsersNone *UserWhereInput `json:"users_none,omitempty"`

	// users some
	UsersSome *UserWhereInput `json:"users_some,omitempty"`

	MarshalOpts *UserRoleNextWhereInputMarshalOpts `json:"-"`
}

type UserRoleNextWhereInputMarshalOpts struct {
	AND_Explicit_Null_When_Empty bool

	NOT_Explicit_Null_When_Empty bool

	OR_Explicit_Null_When_Empty bool

	ID_Explicit_Null_When_Empty bool

	IDContains_Explicit_Null_When_Empty bool

	IDEndsWith_Explicit_Null_When_Empty bool

	IDGt_Explicit_Null_When_Empty bool

	IDGte_Explicit_Null_When_Empty bool

	IDIn_Explicit_Null_When_Empty bool

	IDLt_Explicit_Null_When_Empty bool

	IDLte_Explicit_Null_When_Empty bool

	IDNot_Explicit_Null_When_Empty bool

	IDNotContains_Explicit_Null_When_Empty bool

	IDNotEndsWith_Explicit_Null_When_Empty bool

	IDNotIn_Explicit_Null_When_Empty bool

	IDNotStartsWith_Explicit_Null_When_Empty bool

	IDStartsWith_Explicit_Null_When_Empty bool

	Name_Explicit_Null_When_Empty bool

	NameContains_Explicit_Null_When_Empty bool

	NameEndsWith_Explicit_Null_When_Empty bool

	NameGt_Explicit_Null_When_Empty bool

	NameGte_Explicit_Null_When_Empty bool

	NameIn_Explicit_Null_When_Empty bool

	NameLt_Explicit_Null_When_Empty bool

	NameLte_Explicit_Null_When_Empty bool

	NameNot_Explicit_Null_When_Empty bool

	NameNotContains_Explicit_Null_When_Empty bool

	NameNotEndsWith_Explicit_Null_When_Empty bool

	NameNotIn_Explicit_Null_When_Empty bool

	NameNotStartsWith_Explicit_Null_When_Empty bool

	NameStartsWith_Explicit_Null_When_Empty bool

	Platform_Explicit_Null_When_Empty bool

	PlatformIn_Explicit_Null_When_Empty bool

	PlatformNot_Explicit_Null_When_Empty bool

	PlatformNotIn_Explicit_Null_When_Empty bool

	Preset_Explicit_Null_When_Empty bool

	PresetIn_Explicit_Null_When_Empty bool

	PresetNot_Explicit_Null_When_Empty bool

	PresetNotIn_Explicit_Null_When_Empty bool

	UsersEvery_Explicit_Null_When_Empty bool

	UsersNone_Explicit_Null_When_Empty bool

	UsersSome_Explicit_Null_When_Empty bool
}

func (m UserRoleNextWhereInput) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	b.WriteString("{")

	first := true

	// handle non nullable field AND with omitempty
	if swag.IsZero(m.AND) {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"AND\":")
		bytes, err := swag.WriteJSON(m.AND)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	}

	// handle non nullable field NOT with omitempty
	if swag.IsZero(m.NOT) {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"NOT\":")
		bytes, err := swag.WriteJSON(m.NOT)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	}

	// handle non nullable field OR with omitempty
	if swag.IsZero(m.OR) {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"OR\":")
		bytes, err := swag.WriteJSON(m.OR)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
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

	// handle nullable field id_contains
	if m.IDContains != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"id_contains\":")
		bytes, err := swag.WriteJSON(m.IDContains)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.IDContains_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"id_contains\":null")
		first = false
	}

	// handle nullable field id_ends_with
	if m.IDEndsWith != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"id_ends_with\":")
		bytes, err := swag.WriteJSON(m.IDEndsWith)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.IDEndsWith_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"id_ends_with\":null")
		first = false
	}

	// handle nullable field id_gt
	if m.IDGt != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"id_gt\":")
		bytes, err := swag.WriteJSON(m.IDGt)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.IDGt_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"id_gt\":null")
		first = false
	}

	// handle nullable field id_gte
	if m.IDGte != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"id_gte\":")
		bytes, err := swag.WriteJSON(m.IDGte)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.IDGte_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"id_gte\":null")
		first = false
	}

	// handle non nullable field id_in with omitempty
	if swag.IsZero(m.IDIn) {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"id_in\":")
		bytes, err := swag.WriteJSON(m.IDIn)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	}

	// handle nullable field id_lt
	if m.IDLt != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"id_lt\":")
		bytes, err := swag.WriteJSON(m.IDLt)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.IDLt_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"id_lt\":null")
		first = false
	}

	// handle nullable field id_lte
	if m.IDLte != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"id_lte\":")
		bytes, err := swag.WriteJSON(m.IDLte)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.IDLte_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"id_lte\":null")
		first = false
	}

	// handle nullable field id_not
	if m.IDNot != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"id_not\":")
		bytes, err := swag.WriteJSON(m.IDNot)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.IDNot_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"id_not\":null")
		first = false
	}

	// handle nullable field id_not_contains
	if m.IDNotContains != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"id_not_contains\":")
		bytes, err := swag.WriteJSON(m.IDNotContains)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.IDNotContains_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"id_not_contains\":null")
		first = false
	}

	// handle nullable field id_not_ends_with
	if m.IDNotEndsWith != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"id_not_ends_with\":")
		bytes, err := swag.WriteJSON(m.IDNotEndsWith)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.IDNotEndsWith_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"id_not_ends_with\":null")
		first = false
	}

	// handle non nullable field id_not_in with omitempty
	if swag.IsZero(m.IDNotIn) {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"id_not_in\":")
		bytes, err := swag.WriteJSON(m.IDNotIn)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	}

	// handle nullable field id_not_starts_with
	if m.IDNotStartsWith != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"id_not_starts_with\":")
		bytes, err := swag.WriteJSON(m.IDNotStartsWith)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.IDNotStartsWith_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"id_not_starts_with\":null")
		first = false
	}

	// handle nullable field id_starts_with
	if m.IDStartsWith != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"id_starts_with\":")
		bytes, err := swag.WriteJSON(m.IDStartsWith)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.IDStartsWith_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"id_starts_with\":null")
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

	// handle nullable field name_contains
	if m.NameContains != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"name_contains\":")
		bytes, err := swag.WriteJSON(m.NameContains)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.NameContains_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"name_contains\":null")
		first = false
	}

	// handle nullable field name_ends_with
	if m.NameEndsWith != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"name_ends_with\":")
		bytes, err := swag.WriteJSON(m.NameEndsWith)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.NameEndsWith_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"name_ends_with\":null")
		first = false
	}

	// handle nullable field name_gt
	if m.NameGt != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"name_gt\":")
		bytes, err := swag.WriteJSON(m.NameGt)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.NameGt_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"name_gt\":null")
		first = false
	}

	// handle nullable field name_gte
	if m.NameGte != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"name_gte\":")
		bytes, err := swag.WriteJSON(m.NameGte)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.NameGte_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"name_gte\":null")
		first = false
	}

	// handle non nullable field name_in with omitempty
	if swag.IsZero(m.NameIn) {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"name_in\":")
		bytes, err := swag.WriteJSON(m.NameIn)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	}

	// handle nullable field name_lt
	if m.NameLt != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"name_lt\":")
		bytes, err := swag.WriteJSON(m.NameLt)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.NameLt_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"name_lt\":null")
		first = false
	}

	// handle nullable field name_lte
	if m.NameLte != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"name_lte\":")
		bytes, err := swag.WriteJSON(m.NameLte)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.NameLte_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"name_lte\":null")
		first = false
	}

	// handle nullable field name_not
	if m.NameNot != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"name_not\":")
		bytes, err := swag.WriteJSON(m.NameNot)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.NameNot_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"name_not\":null")
		first = false
	}

	// handle nullable field name_not_contains
	if m.NameNotContains != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"name_not_contains\":")
		bytes, err := swag.WriteJSON(m.NameNotContains)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.NameNotContains_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"name_not_contains\":null")
		first = false
	}

	// handle nullable field name_not_ends_with
	if m.NameNotEndsWith != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"name_not_ends_with\":")
		bytes, err := swag.WriteJSON(m.NameNotEndsWith)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.NameNotEndsWith_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"name_not_ends_with\":null")
		first = false
	}

	// handle non nullable field name_not_in with omitempty
	if swag.IsZero(m.NameNotIn) {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"name_not_in\":")
		bytes, err := swag.WriteJSON(m.NameNotIn)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	}

	// handle nullable field name_not_starts_with
	if m.NameNotStartsWith != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"name_not_starts_with\":")
		bytes, err := swag.WriteJSON(m.NameNotStartsWith)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.NameNotStartsWith_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"name_not_starts_with\":null")
		first = false
	}

	// handle nullable field name_starts_with
	if m.NameStartsWith != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"name_starts_with\":")
		bytes, err := swag.WriteJSON(m.NameStartsWith)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.NameStartsWith_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"name_starts_with\":null")
		first = false
	}

	// handle nullable field platform
	if m.Platform != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"platform\":")
		bytes, err := swag.WriteJSON(m.Platform)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.Platform_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"platform\":null")
		first = false
	}

	// handle non nullable field platform_in with omitempty
	if swag.IsZero(m.PlatformIn) {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"platform_in\":")
		bytes, err := swag.WriteJSON(m.PlatformIn)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	}

	// handle nullable field platform_not
	if m.PlatformNot != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"platform_not\":")
		bytes, err := swag.WriteJSON(m.PlatformNot)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.PlatformNot_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"platform_not\":null")
		first = false
	}

	// handle non nullable field platform_not_in with omitempty
	if swag.IsZero(m.PlatformNotIn) {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"platform_not_in\":")
		bytes, err := swag.WriteJSON(m.PlatformNotIn)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	}

	// handle nullable field preset
	if m.Preset != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"preset\":")
		bytes, err := swag.WriteJSON(m.Preset)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.Preset_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"preset\":null")
		first = false
	}

	// handle non nullable field preset_in with omitempty
	if swag.IsZero(m.PresetIn) {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"preset_in\":")
		bytes, err := swag.WriteJSON(m.PresetIn)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	}

	// handle nullable field preset_not
	if m.PresetNot != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"preset_not\":")
		bytes, err := swag.WriteJSON(m.PresetNot)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.PresetNot_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"preset_not\":null")
		first = false
	}

	// handle non nullable field preset_not_in with omitempty
	if swag.IsZero(m.PresetNotIn) {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"preset_not_in\":")
		bytes, err := swag.WriteJSON(m.PresetNotIn)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	}

	// handle nullable field users_every
	if m.UsersEvery != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"users_every\":")
		bytes, err := swag.WriteJSON(m.UsersEvery)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.UsersEvery_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"users_every\":null")
		first = false
	}

	// handle nullable field users_none
	if m.UsersNone != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"users_none\":")
		bytes, err := swag.WriteJSON(m.UsersNone)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.UsersNone_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"users_none\":null")
		first = false
	}

	// handle nullable field users_some
	if m.UsersSome != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"users_some\":")
		bytes, err := swag.WriteJSON(m.UsersSome)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.UsersSome_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"users_some\":null")
		first = false
	}

	b.WriteString("}")
	return b.Bytes(), nil
}

// Validate validates this user role next where input
func (m *UserRoleNextWhereInput) Validate(formats strfmt.Registry) error {
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

	if err := m.validatePlatform(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePlatformIn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePlatformNot(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePlatformNotIn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePreset(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePresetIn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePresetNot(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePresetNotIn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUsersEvery(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUsersNone(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUsersSome(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UserRoleNextWhereInput) validateAND(formats strfmt.Registry) error {
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

func (m *UserRoleNextWhereInput) validateNOT(formats strfmt.Registry) error {
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

func (m *UserRoleNextWhereInput) validateOR(formats strfmt.Registry) error {
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

func (m *UserRoleNextWhereInput) validatePlatform(formats strfmt.Registry) error {
	if swag.IsZero(m.Platform) { // not required
		return nil
	}

	if m.Platform != nil {
		if err := m.Platform.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("platform")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("platform")
			}
			return err
		}
	}

	return nil
}

func (m *UserRoleNextWhereInput) validatePlatformIn(formats strfmt.Registry) error {
	if swag.IsZero(m.PlatformIn) { // not required
		return nil
	}

	for i := 0; i < len(m.PlatformIn); i++ {

		if err := m.PlatformIn[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("platform_in" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("platform_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *UserRoleNextWhereInput) validatePlatformNot(formats strfmt.Registry) error {
	if swag.IsZero(m.PlatformNot) { // not required
		return nil
	}

	if m.PlatformNot != nil {
		if err := m.PlatformNot.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("platform_not")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("platform_not")
			}
			return err
		}
	}

	return nil
}

func (m *UserRoleNextWhereInput) validatePlatformNotIn(formats strfmt.Registry) error {
	if swag.IsZero(m.PlatformNotIn) { // not required
		return nil
	}

	for i := 0; i < len(m.PlatformNotIn); i++ {

		if err := m.PlatformNotIn[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("platform_not_in" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("platform_not_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *UserRoleNextWhereInput) validatePreset(formats strfmt.Registry) error {
	if swag.IsZero(m.Preset) { // not required
		return nil
	}

	if m.Preset != nil {
		if err := m.Preset.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("preset")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("preset")
			}
			return err
		}
	}

	return nil
}

func (m *UserRoleNextWhereInput) validatePresetIn(formats strfmt.Registry) error {
	if swag.IsZero(m.PresetIn) { // not required
		return nil
	}

	for i := 0; i < len(m.PresetIn); i++ {

		if err := m.PresetIn[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("preset_in" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("preset_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *UserRoleNextWhereInput) validatePresetNot(formats strfmt.Registry) error {
	if swag.IsZero(m.PresetNot) { // not required
		return nil
	}

	if m.PresetNot != nil {
		if err := m.PresetNot.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("preset_not")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("preset_not")
			}
			return err
		}
	}

	return nil
}

func (m *UserRoleNextWhereInput) validatePresetNotIn(formats strfmt.Registry) error {
	if swag.IsZero(m.PresetNotIn) { // not required
		return nil
	}

	for i := 0; i < len(m.PresetNotIn); i++ {

		if err := m.PresetNotIn[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("preset_not_in" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("preset_not_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *UserRoleNextWhereInput) validateUsersEvery(formats strfmt.Registry) error {
	if swag.IsZero(m.UsersEvery) { // not required
		return nil
	}

	if m.UsersEvery != nil {
		if err := m.UsersEvery.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("users_every")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("users_every")
			}
			return err
		}
	}

	return nil
}

func (m *UserRoleNextWhereInput) validateUsersNone(formats strfmt.Registry) error {
	if swag.IsZero(m.UsersNone) { // not required
		return nil
	}

	if m.UsersNone != nil {
		if err := m.UsersNone.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("users_none")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("users_none")
			}
			return err
		}
	}

	return nil
}

func (m *UserRoleNextWhereInput) validateUsersSome(formats strfmt.Registry) error {
	if swag.IsZero(m.UsersSome) { // not required
		return nil
	}

	if m.UsersSome != nil {
		if err := m.UsersSome.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("users_some")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("users_some")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this user role next where input based on the context it is used
func (m *UserRoleNextWhereInput) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

	if err := m.contextValidatePlatform(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePlatformIn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePlatformNot(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePlatformNotIn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePreset(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePresetIn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePresetNot(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePresetNotIn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUsersEvery(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUsersNone(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUsersSome(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UserRoleNextWhereInput) contextValidateAND(ctx context.Context, formats strfmt.Registry) error {

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

func (m *UserRoleNextWhereInput) contextValidateNOT(ctx context.Context, formats strfmt.Registry) error {

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

func (m *UserRoleNextWhereInput) contextValidateOR(ctx context.Context, formats strfmt.Registry) error {

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

func (m *UserRoleNextWhereInput) contextValidatePlatform(ctx context.Context, formats strfmt.Registry) error {

	if m.Platform != nil {
		if err := m.Platform.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("platform")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("platform")
			}
			return err
		}
	}

	return nil
}

func (m *UserRoleNextWhereInput) contextValidatePlatformIn(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.PlatformIn); i++ {

		if err := m.PlatformIn[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("platform_in" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("platform_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *UserRoleNextWhereInput) contextValidatePlatformNot(ctx context.Context, formats strfmt.Registry) error {

	if m.PlatformNot != nil {
		if err := m.PlatformNot.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("platform_not")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("platform_not")
			}
			return err
		}
	}

	return nil
}

func (m *UserRoleNextWhereInput) contextValidatePlatformNotIn(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.PlatformNotIn); i++ {

		if err := m.PlatformNotIn[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("platform_not_in" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("platform_not_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *UserRoleNextWhereInput) contextValidatePreset(ctx context.Context, formats strfmt.Registry) error {

	if m.Preset != nil {
		if err := m.Preset.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("preset")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("preset")
			}
			return err
		}
	}

	return nil
}

func (m *UserRoleNextWhereInput) contextValidatePresetIn(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.PresetIn); i++ {

		if err := m.PresetIn[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("preset_in" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("preset_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *UserRoleNextWhereInput) contextValidatePresetNot(ctx context.Context, formats strfmt.Registry) error {

	if m.PresetNot != nil {
		if err := m.PresetNot.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("preset_not")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("preset_not")
			}
			return err
		}
	}

	return nil
}

func (m *UserRoleNextWhereInput) contextValidatePresetNotIn(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.PresetNotIn); i++ {

		if err := m.PresetNotIn[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("preset_not_in" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("preset_not_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *UserRoleNextWhereInput) contextValidateUsersEvery(ctx context.Context, formats strfmt.Registry) error {

	if m.UsersEvery != nil {
		if err := m.UsersEvery.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("users_every")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("users_every")
			}
			return err
		}
	}

	return nil
}

func (m *UserRoleNextWhereInput) contextValidateUsersNone(ctx context.Context, formats strfmt.Registry) error {

	if m.UsersNone != nil {
		if err := m.UsersNone.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("users_none")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("users_none")
			}
			return err
		}
	}

	return nil
}

func (m *UserRoleNextWhereInput) contextValidateUsersSome(ctx context.Context, formats strfmt.Registry) error {

	if m.UsersSome != nil {
		if err := m.UsersSome.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("users_some")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("users_some")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UserRoleNextWhereInput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserRoleNextWhereInput) UnmarshalBinary(b []byte) error {
	var res UserRoleNextWhereInput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
