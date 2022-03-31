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

// PrivilegesRequired privileges required
//
// swagger:model PrivilegesRequired
type PrivilegesRequired string

func NewPrivilegesRequired(value PrivilegesRequired) *PrivilegesRequired {
	v := value
	return &v
}

const (

	// PrivilegesRequiredNONE captures enum value "NONE"
	PrivilegesRequiredNONE PrivilegesRequired = "NONE"

	// PrivilegesRequiredLOW captures enum value "LOW"
	PrivilegesRequiredLOW PrivilegesRequired = "LOW"

	// PrivilegesRequiredHIGH captures enum value "HIGH"
	PrivilegesRequiredHIGH PrivilegesRequired = "HIGH"
)

// for schema
var privilegesRequiredEnum []interface{}

func init() {
	var res []PrivilegesRequired
	if err := json.Unmarshal([]byte(`["NONE","LOW","HIGH"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		privilegesRequiredEnum = append(privilegesRequiredEnum, v)
	}
}

func (m PrivilegesRequired) validatePrivilegesRequiredEnum(path, location string, value PrivilegesRequired) error {
	if err := validate.EnumCase(path, location, value, privilegesRequiredEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this privileges required
func (m PrivilegesRequired) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validatePrivilegesRequiredEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this privileges required based on context it is used
func (m PrivilegesRequired) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
