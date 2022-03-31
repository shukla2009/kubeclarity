// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// OrganizationalContact organizational contact
//
// swagger:model OrganizationalContact
type OrganizationalContact struct {

	// email
	Email string `json:"email,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// phone
	Phone string `json:"phone,omitempty"`
}

// Validate validates this organizational contact
func (m *OrganizationalContact) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this organizational contact based on context it is used
func (m *OrganizationalContact) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *OrganizationalContact) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OrganizationalContact) UnmarshalBinary(b []byte) error {
	var res OrganizationalContact
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
