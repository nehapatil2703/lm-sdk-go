// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"encoding/json"

	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NTLMAuthentication n t LM authentication
// swagger:model NTLMAuthentication
type NTLMAuthentication struct {
	passwordField *string

	userNameField *string

	// domain
	Domain string `json:"domain,omitempty"`
}

// Password gets the password of this subtype
func (m *NTLMAuthentication) Password() *string {
	return m.passwordField
}

// SetPassword sets the password of this subtype
func (m *NTLMAuthentication) SetPassword(val *string) {
	m.passwordField = val
}

// Type gets the type of this subtype
func (m *NTLMAuthentication) Type() string {
	return "ntlm"
}

// SetType sets the type of this subtype
func (m *NTLMAuthentication) SetType(val string) {
}

// UserName gets the user name of this subtype
func (m *NTLMAuthentication) UserName() *string {
	return m.userNameField
}

// SetUserName sets the user name of this subtype
func (m *NTLMAuthentication) SetUserName(val *string) {
	m.userNameField = val
}

// Domain gets the domain of this subtype

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *NTLMAuthentication) UnmarshalJSON(raw []byte) error {
	var data struct {

		// domain
		Domain string `json:"domain,omitempty"`
	}
	buf := bytes.NewBuffer(raw)
	dec := json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&data); err != nil {
		return err
	}

	var base struct {
		/* Just the base type fields. Used for unmashalling polymorphic types.*/

		Password *string `json:"password"`

		Type string `json:"type"`

		UserName *string `json:"userName"`
	}
	buf = bytes.NewBuffer(raw)
	dec = json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&base); err != nil {
		return err
	}

	var result NTLMAuthentication

	result.passwordField = base.Password

	if base.Type != result.Type() {
		/* Not the type we're looking for. */
		return errors.New(422, "invalid type value: %q", base.Type)
	}

	result.userNameField = base.UserName

	result.Domain = data.Domain

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m NTLMAuthentication) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {

		// domain
		Domain string `json:"domain,omitempty"`
	}{

		Domain: m.Domain,
	},
	)
	if err != nil {
		return nil, err
	}
	b2, err = json.Marshal(struct {
		Password *string `json:"password"`

		Type string `json:"type"`

		UserName *string `json:"userName"`
	}{

		Password: m.Password(),

		Type: m.Type(),

		UserName: m.UserName(),
	},
	)
	if err != nil {
		return nil, err
	}

	return swag.ConcatJSON(b1, b2, b3), nil
}

// Validate validates this n t LM authentication
func (m *NTLMAuthentication) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePassword(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NTLMAuthentication) validatePassword(formats strfmt.Registry) error {
	if err := validate.Required("password", "body", m.Password()); err != nil {
		return err
	}

	return nil
}

func (m *NTLMAuthentication) validateUserName(formats strfmt.Registry) error {
	if err := validate.Required("userName", "body", m.UserName()); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NTLMAuthentication) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NTLMAuthentication) UnmarshalBinary(b []byte) error {
	var res NTLMAuthentication
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
