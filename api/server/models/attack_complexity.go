// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// AttackComplexity attack complexity
//
// swagger:model AttackComplexity
type AttackComplexity string

func NewAttackComplexity(value AttackComplexity) *AttackComplexity {
	v := value
	return &v
}

const (

	// AttackComplexityLOW captures enum value "LOW"
	AttackComplexityLOW AttackComplexity = "LOW"

	// AttackComplexityHIGH captures enum value "HIGH"
	AttackComplexityHIGH AttackComplexity = "HIGH"
)

// for schema
var attackComplexityEnum []interface{}

func init() {
	var res []AttackComplexity
	if err := json.Unmarshal([]byte(`["LOW","HIGH"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		attackComplexityEnum = append(attackComplexityEnum, v)
	}
}

func (m AttackComplexity) validateAttackComplexityEnum(path, location string, value AttackComplexity) error {
	if err := validate.EnumCase(path, location, value, attackComplexityEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this attack complexity
func (m AttackComplexity) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateAttackComplexityEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this attack complexity based on context it is used
func (m AttackComplexity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
