//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2022 SeMI Technologies B.V. All rights reserved.
//
//  CONTACT: hello@semi.technology
//

// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// WhereFilter Filter search results using a where filter
//
// swagger:model WhereFilter
type WhereFilter struct {

	// combine multiple where filters, requires 'And' or 'Or' operator
	Operands []*WhereFilter `json:"operands"`

	// operator to use
	// Enum: [And Or Equal Like Not NotEqual GreaterThan GreaterThanEqual LessThan LessThanEqual WithinGeoRange]
	Operator string `json:"operator,omitempty"`

	// path to the property currently being filtered
	Path []string `json:"path"`

	// value as boolean
	ValueBoolean *bool `json:"valueBoolean,omitempty"`

	// value as date (as string)
	ValueDate *string `json:"valueDate,omitempty"`

	// value as geo coordinates and distance
	ValueGeoRange *WhereFilterGeoRange `json:"valueGeoRange,omitempty"`

	// value as integer
	ValueInt *int64 `json:"valueInt,omitempty"`

	// value as number/float
	ValueNumber *float64 `json:"valueNumber,omitempty"`

	// value as string
	ValueString *string `json:"valueString,omitempty"`

	// value as text (on text props)
	ValueText *string `json:"valueText,omitempty"`
}

// Validate validates this where filter
func (m *WhereFilter) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOperands(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOperator(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValueGeoRange(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WhereFilter) validateOperands(formats strfmt.Registry) error {

	if swag.IsZero(m.Operands) { // not required
		return nil
	}

	for i := 0; i < len(m.Operands); i++ {
		if swag.IsZero(m.Operands[i]) { // not required
			continue
		}

		if m.Operands[i] != nil {
			if err := m.Operands[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("operands" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var whereFilterTypeOperatorPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["And","Or","Equal","Like","Not","NotEqual","GreaterThan","GreaterThanEqual","LessThan","LessThanEqual","WithinGeoRange"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		whereFilterTypeOperatorPropEnum = append(whereFilterTypeOperatorPropEnum, v)
	}
}

const (

	// WhereFilterOperatorAnd captures enum value "And"
	WhereFilterOperatorAnd string = "And"

	// WhereFilterOperatorOr captures enum value "Or"
	WhereFilterOperatorOr string = "Or"

	// WhereFilterOperatorEqual captures enum value "Equal"
	WhereFilterOperatorEqual string = "Equal"

	// WhereFilterOperatorLike captures enum value "Like"
	WhereFilterOperatorLike string = "Like"

	// WhereFilterOperatorNot captures enum value "Not"
	WhereFilterOperatorNot string = "Not"

	// WhereFilterOperatorNotEqual captures enum value "NotEqual"
	WhereFilterOperatorNotEqual string = "NotEqual"

	// WhereFilterOperatorGreaterThan captures enum value "GreaterThan"
	WhereFilterOperatorGreaterThan string = "GreaterThan"

	// WhereFilterOperatorGreaterThanEqual captures enum value "GreaterThanEqual"
	WhereFilterOperatorGreaterThanEqual string = "GreaterThanEqual"

	// WhereFilterOperatorLessThan captures enum value "LessThan"
	WhereFilterOperatorLessThan string = "LessThan"

	// WhereFilterOperatorLessThanEqual captures enum value "LessThanEqual"
	WhereFilterOperatorLessThanEqual string = "LessThanEqual"

	// WhereFilterOperatorWithinGeoRange captures enum value "WithinGeoRange"
	WhereFilterOperatorWithinGeoRange string = "WithinGeoRange"
)

// prop value enum
func (m *WhereFilter) validateOperatorEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, whereFilterTypeOperatorPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *WhereFilter) validateOperator(formats strfmt.Registry) error {

	if swag.IsZero(m.Operator) { // not required
		return nil
	}

	// value enum
	if err := m.validateOperatorEnum("operator", "body", m.Operator); err != nil {
		return err
	}

	return nil
}

func (m *WhereFilter) validateValueGeoRange(formats strfmt.Registry) error {

	if swag.IsZero(m.ValueGeoRange) { // not required
		return nil
	}

	if m.ValueGeoRange != nil {
		if err := m.ValueGeoRange.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("valueGeoRange")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *WhereFilter) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WhereFilter) UnmarshalBinary(b []byte) error {
	var res WhereFilter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
