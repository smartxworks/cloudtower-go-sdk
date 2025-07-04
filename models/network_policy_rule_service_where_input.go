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

// NetworkPolicyRuleServiceWhereInput network policy rule service where input
//
// swagger:model NetworkPolicyRuleServiceWhereInput
type NetworkPolicyRuleServiceWhereInput struct {

	// a n d
	AND []*NetworkPolicyRuleServiceWhereInput `json:"AND,omitempty"`

	// n o t
	NOT []*NetworkPolicyRuleServiceWhereInput `json:"NOT,omitempty"`

	// o r
	OR []*NetworkPolicyRuleServiceWhereInput `json:"OR,omitempty"`

	// description
	Description *string `json:"description,omitempty"`

	// description contains
	DescriptionContains *string `json:"description_contains,omitempty"`

	// description ends with
	DescriptionEndsWith *string `json:"description_ends_with,omitempty"`

	// description gt
	DescriptionGt *string `json:"description_gt,omitempty"`

	// description gte
	DescriptionGte *string `json:"description_gte,omitempty"`

	// description in
	DescriptionIn []string `json:"description_in,omitempty"`

	// description lt
	DescriptionLt *string `json:"description_lt,omitempty"`

	// description lte
	DescriptionLte *string `json:"description_lte,omitempty"`

	// description not
	DescriptionNot *string `json:"description_not,omitempty"`

	// description not contains
	DescriptionNotContains *string `json:"description_not_contains,omitempty"`

	// description not ends with
	DescriptionNotEndsWith *string `json:"description_not_ends_with,omitempty"`

	// description not in
	DescriptionNotIn []string `json:"description_not_in,omitempty"`

	// description not starts with
	DescriptionNotStartsWith *string `json:"description_not_starts_with,omitempty"`

	// description starts with
	DescriptionStartsWith *string `json:"description_starts_with,omitempty"`

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

	// preset
	Preset *bool `json:"preset,omitempty"`

	// preset not
	PresetNot *bool `json:"preset_not,omitempty"`

	MarshalOpts *NetworkPolicyRuleServiceWhereInputMarshalOpts `json:"-"`
}

type NetworkPolicyRuleServiceWhereInputMarshalOpts struct {
	AND_Explicit_Null_When_Empty bool

	NOT_Explicit_Null_When_Empty bool

	OR_Explicit_Null_When_Empty bool

	Description_Explicit_Null_When_Empty bool

	DescriptionContains_Explicit_Null_When_Empty bool

	DescriptionEndsWith_Explicit_Null_When_Empty bool

	DescriptionGt_Explicit_Null_When_Empty bool

	DescriptionGte_Explicit_Null_When_Empty bool

	DescriptionIn_Explicit_Null_When_Empty bool

	DescriptionLt_Explicit_Null_When_Empty bool

	DescriptionLte_Explicit_Null_When_Empty bool

	DescriptionNot_Explicit_Null_When_Empty bool

	DescriptionNotContains_Explicit_Null_When_Empty bool

	DescriptionNotEndsWith_Explicit_Null_When_Empty bool

	DescriptionNotIn_Explicit_Null_When_Empty bool

	DescriptionNotStartsWith_Explicit_Null_When_Empty bool

	DescriptionStartsWith_Explicit_Null_When_Empty bool

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

	Preset_Explicit_Null_When_Empty bool

	PresetNot_Explicit_Null_When_Empty bool
}

func (m NetworkPolicyRuleServiceWhereInput) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	b.WriteString("{")

	first := true

	// handle non nullable field AND with omitempty
	if !swag.IsZero(m.AND) {
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
	if !swag.IsZero(m.NOT) {
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
	if !swag.IsZero(m.OR) {
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

	// handle nullable field description
	if m.Description != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"description\":")
		bytes, err := swag.WriteJSON(m.Description)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.Description_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"description\":null")
		first = false
	}

	// handle nullable field description_contains
	if m.DescriptionContains != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"description_contains\":")
		bytes, err := swag.WriteJSON(m.DescriptionContains)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.DescriptionContains_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"description_contains\":null")
		first = false
	}

	// handle nullable field description_ends_with
	if m.DescriptionEndsWith != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"description_ends_with\":")
		bytes, err := swag.WriteJSON(m.DescriptionEndsWith)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.DescriptionEndsWith_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"description_ends_with\":null")
		first = false
	}

	// handle nullable field description_gt
	if m.DescriptionGt != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"description_gt\":")
		bytes, err := swag.WriteJSON(m.DescriptionGt)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.DescriptionGt_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"description_gt\":null")
		first = false
	}

	// handle nullable field description_gte
	if m.DescriptionGte != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"description_gte\":")
		bytes, err := swag.WriteJSON(m.DescriptionGte)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.DescriptionGte_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"description_gte\":null")
		first = false
	}

	// handle non nullable field description_in with omitempty
	if !swag.IsZero(m.DescriptionIn) {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"description_in\":")
		bytes, err := swag.WriteJSON(m.DescriptionIn)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	}

	// handle nullable field description_lt
	if m.DescriptionLt != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"description_lt\":")
		bytes, err := swag.WriteJSON(m.DescriptionLt)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.DescriptionLt_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"description_lt\":null")
		first = false
	}

	// handle nullable field description_lte
	if m.DescriptionLte != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"description_lte\":")
		bytes, err := swag.WriteJSON(m.DescriptionLte)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.DescriptionLte_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"description_lte\":null")
		first = false
	}

	// handle nullable field description_not
	if m.DescriptionNot != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"description_not\":")
		bytes, err := swag.WriteJSON(m.DescriptionNot)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.DescriptionNot_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"description_not\":null")
		first = false
	}

	// handle nullable field description_not_contains
	if m.DescriptionNotContains != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"description_not_contains\":")
		bytes, err := swag.WriteJSON(m.DescriptionNotContains)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.DescriptionNotContains_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"description_not_contains\":null")
		first = false
	}

	// handle nullable field description_not_ends_with
	if m.DescriptionNotEndsWith != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"description_not_ends_with\":")
		bytes, err := swag.WriteJSON(m.DescriptionNotEndsWith)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.DescriptionNotEndsWith_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"description_not_ends_with\":null")
		first = false
	}

	// handle non nullable field description_not_in with omitempty
	if !swag.IsZero(m.DescriptionNotIn) {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"description_not_in\":")
		bytes, err := swag.WriteJSON(m.DescriptionNotIn)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	}

	// handle nullable field description_not_starts_with
	if m.DescriptionNotStartsWith != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"description_not_starts_with\":")
		bytes, err := swag.WriteJSON(m.DescriptionNotStartsWith)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.DescriptionNotStartsWith_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"description_not_starts_with\":null")
		first = false
	}

	// handle nullable field description_starts_with
	if m.DescriptionStartsWith != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"description_starts_with\":")
		bytes, err := swag.WriteJSON(m.DescriptionStartsWith)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.DescriptionStartsWith_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"description_starts_with\":null")
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
	if !swag.IsZero(m.IDIn) {
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
	if !swag.IsZero(m.IDNotIn) {
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
	if !swag.IsZero(m.NameIn) {
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
	if !swag.IsZero(m.NameNotIn) {
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

	b.WriteString("}")
	return b.Bytes(), nil
}

// Validate validates this network policy rule service where input
func (m *NetworkPolicyRuleServiceWhereInput) Validate(formats strfmt.Registry) error {
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NetworkPolicyRuleServiceWhereInput) validateAND(formats strfmt.Registry) error {
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

func (m *NetworkPolicyRuleServiceWhereInput) validateNOT(formats strfmt.Registry) error {
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

func (m *NetworkPolicyRuleServiceWhereInput) validateOR(formats strfmt.Registry) error {
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

// ContextValidate validate this network policy rule service where input based on the context it is used
func (m *NetworkPolicyRuleServiceWhereInput) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NetworkPolicyRuleServiceWhereInput) contextValidateAND(ctx context.Context, formats strfmt.Registry) error {

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

func (m *NetworkPolicyRuleServiceWhereInput) contextValidateNOT(ctx context.Context, formats strfmt.Registry) error {

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

func (m *NetworkPolicyRuleServiceWhereInput) contextValidateOR(ctx context.Context, formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *NetworkPolicyRuleServiceWhereInput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NetworkPolicyRuleServiceWhereInput) UnmarshalBinary(b []byte) error {
	var res NetworkPolicyRuleServiceWhereInput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
