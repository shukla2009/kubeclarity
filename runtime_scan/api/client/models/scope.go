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

// Scope scope
//
// swagger:model Scope
type Scope string

func NewScope(value Scope) *Scope {
	v := value
	return &v
}

const (

	// ScopeUNCHANGED captures enum value "UNCHANGED"
	ScopeUNCHANGED Scope = "UNCHANGED"

	// ScopeCHANGED captures enum value "CHANGED"
	ScopeCHANGED Scope = "CHANGED"
)

// for schema
var scopeEnum []interface{}

func init() {
	var res []Scope
	if err := json.Unmarshal([]byte(`["UNCHANGED","CHANGED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		scopeEnum = append(scopeEnum, v)
	}
}

func (m Scope) validateScopeEnum(path, location string, value Scope) error {
	if err := validate.EnumCase(path, location, value, scopeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this scope
func (m Scope) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateScopeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this scope based on context it is used
func (m Scope) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
