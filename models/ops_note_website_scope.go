// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// OpsNoteWebsiteScope ops note website scope
//
// swagger:model OpsNoteWebsiteScope
type OpsNoteWebsiteScope struct {

	// full path
	FullPath string `json:"fullPath,omitempty"`

	// group Id
	// Example: 74
	GroupID int32 `json:"groupId,omitempty"`

	// website Id
	// Example: 87
	WebsiteID int32 `json:"websiteId,omitempty"`

	// website name
	WebsiteName string `json:"websiteName,omitempty"`
}

// Type gets the type of this subtype
func (m *OpsNoteWebsiteScope) Type() string {
	return "website"
}

// SetType sets the type of this subtype
func (m *OpsNoteWebsiteScope) SetType(val string) {
}

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *OpsNoteWebsiteScope) UnmarshalJSON(raw []byte) error {
	var data struct {

		// full path
		FullPath string `json:"fullPath,omitempty"`

		// group Id
		// Example: 74
		GroupID int32 `json:"groupId,omitempty"`

		// website Id
		// Example: 87
		WebsiteID int32 `json:"websiteId,omitempty"`

		// website name
		WebsiteName string `json:"websiteName,omitempty"`
	}
	buf := bytes.NewBuffer(raw)
	dec := json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&data); err != nil {
		return err
	}

	var base struct {
		/* Just the base type fields. Used for unmashalling polymorphic types.*/

		Type string `json:"type"`
	}
	buf = bytes.NewBuffer(raw)
	dec = json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&base); err != nil {
		return err
	}

	var result OpsNoteWebsiteScope

	if base.Type != result.Type() {
		/* Not the type we're looking for. */
		return errors.New(422, "invalid type value: %q", base.Type)
	}

	result.FullPath = data.FullPath
	result.GroupID = data.GroupID
	result.WebsiteID = data.WebsiteID
	result.WebsiteName = data.WebsiteName

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m OpsNoteWebsiteScope) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {

		// full path
		FullPath string `json:"fullPath,omitempty"`

		// group Id
		// Example: 74
		GroupID int32 `json:"groupId,omitempty"`

		// website Id
		// Example: 87
		WebsiteID int32 `json:"websiteId,omitempty"`

		// website name
		WebsiteName string `json:"websiteName,omitempty"`
	}{

		FullPath: m.FullPath,

		GroupID: m.GroupID,

		WebsiteID: m.WebsiteID,

		WebsiteName: m.WebsiteName,
	})
	if err != nil {
		return nil, err
	}
	b2, err = json.Marshal(struct {
		Type string `json:"type"`
	}{

		Type: m.Type(),
	})
	if err != nil {
		return nil, err
	}

	return swag.ConcatJSON(b1, b2, b3), nil
}

// Validate validates this ops note website scope
func (m *OpsNoteWebsiteScope) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validate this ops note website scope based on the context it is used
func (m *OpsNoteWebsiteScope) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *OpsNoteWebsiteScope) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpsNoteWebsiteScope) UnmarshalBinary(b []byte) error {
	var res OpsNoteWebsiteScope
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
