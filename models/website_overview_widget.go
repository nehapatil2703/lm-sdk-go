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

// WebsiteOverviewWidget website overview widget
// swagger:model WebsiteOverviewWidget
type WebsiteOverviewWidget struct {
	dashboardIdField *int32

	descriptionField string

	idField int32

	intervalField int32

	lastUpdatedByField string

	lastUpdatedOnField int64

	nameField *string

	themeField string

	timescaleField string

	userPermissionField string

	// geo info
	// Read Only: true
	GeoInfo string `json:"geoInfo,omitempty"`

	// graph
	// Required: true
	Graph *string `json:"graph"`

	// website Id
	WebsiteID int32 `json:"websiteId,omitempty"`

	// website name
	// Read Only: true
	WebsiteName string `json:"websiteName,omitempty"`
}

// DashboardID gets the dashboard Id of this subtype
func (m *WebsiteOverviewWidget) DashboardID() *int32 {
	return m.dashboardIdField
}

// SetDashboardID sets the dashboard Id of this subtype
func (m *WebsiteOverviewWidget) SetDashboardID(val *int32) {
	m.dashboardIdField = val
}

// Description gets the description of this subtype
func (m *WebsiteOverviewWidget) Description() string {
	return m.descriptionField
}

// SetDescription sets the description of this subtype
func (m *WebsiteOverviewWidget) SetDescription(val string) {
	m.descriptionField = val
}

// ID gets the id of this subtype
func (m *WebsiteOverviewWidget) ID() int32 {
	return m.idField
}

// SetID sets the id of this subtype
func (m *WebsiteOverviewWidget) SetID(val int32) {
	m.idField = val
}

// Interval gets the interval of this subtype
func (m *WebsiteOverviewWidget) Interval() int32 {
	return m.intervalField
}

// SetInterval sets the interval of this subtype
func (m *WebsiteOverviewWidget) SetInterval(val int32) {
	m.intervalField = val
}

// LastUpdatedBy gets the last updated by of this subtype
func (m *WebsiteOverviewWidget) LastUpdatedBy() string {
	return m.lastUpdatedByField
}

// SetLastUpdatedBy sets the last updated by of this subtype
func (m *WebsiteOverviewWidget) SetLastUpdatedBy(val string) {
	m.lastUpdatedByField = val
}

// LastUpdatedOn gets the last updated on of this subtype
func (m *WebsiteOverviewWidget) LastUpdatedOn() int64 {
	return m.lastUpdatedOnField
}

// SetLastUpdatedOn sets the last updated on of this subtype
func (m *WebsiteOverviewWidget) SetLastUpdatedOn(val int64) {
	m.lastUpdatedOnField = val
}

// Name gets the name of this subtype
func (m *WebsiteOverviewWidget) Name() *string {
	return m.nameField
}

// SetName sets the name of this subtype
func (m *WebsiteOverviewWidget) SetName(val *string) {
	m.nameField = val
}

// Theme gets the theme of this subtype
func (m *WebsiteOverviewWidget) Theme() string {
	return m.themeField
}

// SetTheme sets the theme of this subtype
func (m *WebsiteOverviewWidget) SetTheme(val string) {
	m.themeField = val
}

// Timescale gets the timescale of this subtype
func (m *WebsiteOverviewWidget) Timescale() string {
	return m.timescaleField
}

// SetTimescale sets the timescale of this subtype
func (m *WebsiteOverviewWidget) SetTimescale(val string) {
	m.timescaleField = val
}

// Type gets the type of this subtype
func (m *WebsiteOverviewWidget) Type() string {
	return "websiteOverview"
}

// SetType sets the type of this subtype
func (m *WebsiteOverviewWidget) SetType(val string) {
}

// UserPermission gets the user permission of this subtype
func (m *WebsiteOverviewWidget) UserPermission() string {
	return m.userPermissionField
}

// SetUserPermission sets the user permission of this subtype
func (m *WebsiteOverviewWidget) SetUserPermission(val string) {
	m.userPermissionField = val
}

// GeoInfo gets the geo info of this subtype

// Graph gets the graph of this subtype

// WebsiteID gets the website Id of this subtype

// WebsiteName gets the website name of this subtype

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *WebsiteOverviewWidget) UnmarshalJSON(raw []byte) error {
	var data struct {

		// geo info
		// Read Only: true
		GeoInfo string `json:"geoInfo,omitempty"`

		// graph
		// Required: true
		Graph *string `json:"graph"`

		// website Id
		WebsiteID int32 `json:"websiteId,omitempty"`

		// website name
		// Read Only: true
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

		DashboardID *int32 `json:"dashboardId"`

		Description string `json:"description,omitempty"`

		ID int32 `json:"id,omitempty"`

		Interval int32 `json:"interval,omitempty"`

		LastUpdatedBy string `json:"lastUpdatedBy,omitempty"`

		LastUpdatedOn int64 `json:"lastUpdatedOn,omitempty"`

		Name *string `json:"name"`

		Theme string `json:"theme,omitempty"`

		Timescale string `json:"timescale,omitempty"`

		Type string `json:"type"`

		UserPermission string `json:"userPermission,omitempty"`
	}
	buf = bytes.NewBuffer(raw)
	dec = json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&base); err != nil {
		return err
	}

	var result WebsiteOverviewWidget

	result.dashboardIdField = base.DashboardID

	result.descriptionField = base.Description

	result.idField = base.ID

	result.intervalField = base.Interval

	result.lastUpdatedByField = base.LastUpdatedBy

	result.lastUpdatedOnField = base.LastUpdatedOn

	result.nameField = base.Name

	result.themeField = base.Theme

	result.timescaleField = base.Timescale

	if base.Type != result.Type() {
		/* Not the type we're looking for. */
		return errors.New(422, "invalid type value: %q", base.Type)
	}

	result.userPermissionField = base.UserPermission

	result.GeoInfo = data.GeoInfo

	result.Graph = data.Graph

	result.WebsiteID = data.WebsiteID

	result.WebsiteName = data.WebsiteName

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m WebsiteOverviewWidget) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {

		// geo info
		// Read Only: true
		GeoInfo string `json:"geoInfo,omitempty"`

		// graph
		// Required: true
		Graph *string `json:"graph"`

		// website Id
		WebsiteID int32 `json:"websiteId,omitempty"`

		// website name
		// Read Only: true
		WebsiteName string `json:"websiteName,omitempty"`
	}{

		GeoInfo: m.GeoInfo,

		Graph: m.Graph,

		WebsiteID: m.WebsiteID,

		WebsiteName: m.WebsiteName,
	},
	)
	if err != nil {
		return nil, err
	}
	b2, err = json.Marshal(struct {
		DashboardID *int32 `json:"dashboardId"`

		Description string `json:"description,omitempty"`

		ID int32 `json:"id,omitempty"`

		Interval int32 `json:"interval,omitempty"`

		LastUpdatedBy string `json:"lastUpdatedBy,omitempty"`

		LastUpdatedOn int64 `json:"lastUpdatedOn,omitempty"`

		Name *string `json:"name"`

		Theme string `json:"theme,omitempty"`

		Timescale string `json:"timescale,omitempty"`

		Type string `json:"type"`

		UserPermission string `json:"userPermission,omitempty"`
	}{

		DashboardID: m.DashboardID(),

		Description: m.Description(),

		ID: m.ID(),

		Interval: m.Interval(),

		LastUpdatedBy: m.LastUpdatedBy(),

		LastUpdatedOn: m.LastUpdatedOn(),

		Name: m.Name(),

		Theme: m.Theme(),

		Timescale: m.Timescale(),

		Type: m.Type(),

		UserPermission: m.UserPermission(),
	},
	)
	if err != nil {
		return nil, err
	}

	return swag.ConcatJSON(b1, b2, b3), nil
}

// Validate validates this website overview widget
func (m *WebsiteOverviewWidget) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDashboardID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGraph(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WebsiteOverviewWidget) validateDashboardID(formats strfmt.Registry) error {
	if err := validate.Required("dashboardId", "body", m.DashboardID()); err != nil {
		return err
	}

	return nil
}

func (m *WebsiteOverviewWidget) validateName(formats strfmt.Registry) error {
	if err := validate.Required("name", "body", m.Name()); err != nil {
		return err
	}

	return nil
}

func (m *WebsiteOverviewWidget) validateGraph(formats strfmt.Registry) error {
	if err := validate.Required("graph", "body", m.Graph); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *WebsiteOverviewWidget) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WebsiteOverviewWidget) UnmarshalBinary(b []byte) error {
	var res WebsiteOverviewWidget
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
