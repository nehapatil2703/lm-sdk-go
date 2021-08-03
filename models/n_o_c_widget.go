// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NOCWidget n o c widget
//
// swagger:model NOCWidget
type NOCWidget struct {
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

	// Whether or not acknowledgements are displayed in the NOC widget, the default value is true
	AckChecked interface{} `json:"ackChecked,omitempty"`

	// The maximum number columns displayed in the NOC widget
	DisplayColumn int32 `json:"displayColumn,omitempty"`

	// Whether or not critical alerts are displayed in the NOC widget, the default value is true
	DisplayCriticalAlert interface{} `json:"displayCriticalAlert,omitempty"`

	// Whether or not error alerts are displayed in the NOC widget, the default value is true
	DisplayErrorAlert interface{} `json:"displayErrorAlert,omitempty"`

	// Whether or not warning alerts are displayed in the NOC widget, the default value is true
	DisplayWarnAlert interface{} `json:"displayWarnAlert,omitempty"`

	itemsField []NOCItemBase

	// Whether or not SDTs are displayed in the NOC widget, the default value is true
	SDTChecked interface{} `json:"sdtChecked,omitempty"`

	// How NOC items are sorted
	SortBy string `json:"sortBy,omitempty"`
}

// DashboardID gets the dashboard Id of this subtype
func (m *NOCWidget) DashboardID() *int32 {
	return m.dashboardIdField
}

// SetDashboardID sets the dashboard Id of this subtype
func (m *NOCWidget) SetDashboardID(val *int32) {
	m.dashboardIdField = val
}

// Description gets the description of this subtype
func (m *NOCWidget) Description() string {
	return m.descriptionField
}

// SetDescription sets the description of this subtype
func (m *NOCWidget) SetDescription(val string) {
	m.descriptionField = val
}

// ID gets the id of this subtype
func (m *NOCWidget) ID() int32 {
	return m.idField
}

// SetID sets the id of this subtype
func (m *NOCWidget) SetID(val int32) {
	m.idField = val
}

// Interval gets the interval of this subtype
func (m *NOCWidget) Interval() int32 {
	return m.intervalField
}

// SetInterval sets the interval of this subtype
func (m *NOCWidget) SetInterval(val int32) {
	m.intervalField = val
}

// LastUpdatedBy gets the last updated by of this subtype
func (m *NOCWidget) LastUpdatedBy() string {
	return m.lastUpdatedByField
}

// SetLastUpdatedBy sets the last updated by of this subtype
func (m *NOCWidget) SetLastUpdatedBy(val string) {
	m.lastUpdatedByField = val
}

// LastUpdatedOn gets the last updated on of this subtype
func (m *NOCWidget) LastUpdatedOn() int64 {
	return m.lastUpdatedOnField
}

// SetLastUpdatedOn sets the last updated on of this subtype
func (m *NOCWidget) SetLastUpdatedOn(val int64) {
	m.lastUpdatedOnField = val
}

// Name gets the name of this subtype
func (m *NOCWidget) Name() *string {
	return m.nameField
}

// SetName sets the name of this subtype
func (m *NOCWidget) SetName(val *string) {
	m.nameField = val
}

// Theme gets the theme of this subtype
func (m *NOCWidget) Theme() string {
	return m.themeField
}

// SetTheme sets the theme of this subtype
func (m *NOCWidget) SetTheme(val string) {
	m.themeField = val
}

// Timescale gets the timescale of this subtype
func (m *NOCWidget) Timescale() string {
	return m.timescaleField
}

// SetTimescale sets the timescale of this subtype
func (m *NOCWidget) SetTimescale(val string) {
	m.timescaleField = val
}

// Type gets the type of this subtype
func (m *NOCWidget) Type() string {
	return "noc"
}

// SetType sets the type of this subtype
func (m *NOCWidget) SetType(val string) {
}

// UserPermission gets the user permission of this subtype
func (m *NOCWidget) UserPermission() string {
	return m.userPermissionField
}

// SetUserPermission sets the user permission of this subtype
func (m *NOCWidget) SetUserPermission(val string) {
	m.userPermissionField = val
}

// Items gets the items of this subtype
func (m *NOCWidget) Items() []NOCItemBase {
	return m.itemsField
}

// SetItems sets the items of this subtype
func (m *NOCWidget) SetItems(val []NOCItemBase) {
	m.itemsField = val
}

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *NOCWidget) UnmarshalJSON(raw []byte) error {
	var data struct {

		// Whether or not acknowledgements are displayed in the NOC widget, the default value is true
		AckChecked interface{} `json:"ackChecked,omitempty"`

		// The maximum number columns displayed in the NOC widget
		DisplayColumn int32 `json:"displayColumn,omitempty"`

		// Whether or not critical alerts are displayed in the NOC widget, the default value is true
		DisplayCriticalAlert interface{} `json:"displayCriticalAlert,omitempty"`

		// Whether or not error alerts are displayed in the NOC widget, the default value is true
		DisplayErrorAlert interface{} `json:"displayErrorAlert,omitempty"`

		// Whether or not warning alerts are displayed in the NOC widget, the default value is true
		DisplayWarnAlert interface{} `json:"displayWarnAlert,omitempty"`

		Items json.RawMessage `json:"items"`

		// Whether or not SDTs are displayed in the NOC widget, the default value is true
		SDTChecked interface{} `json:"sdtChecked,omitempty"`

		// How NOC items are sorted
		SortBy string `json:"sortBy,omitempty"`
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

	allOfItems, err := UnmarshalNOCItemBaseSlice(bytes.NewBuffer(data.Items), runtime.JSONConsumer())
	if err != nil && err != io.EOF {
		return err
	}

	var result NOCWidget

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

	result.AckChecked = data.AckChecked
	result.DisplayColumn = data.DisplayColumn
	result.DisplayCriticalAlert = data.DisplayCriticalAlert
	result.DisplayErrorAlert = data.DisplayErrorAlert
	result.DisplayWarnAlert = data.DisplayWarnAlert
	result.itemsField = allOfItems
	result.SDTChecked = data.SDTChecked
	result.SortBy = data.SortBy

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m NOCWidget) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {

		// Whether or not acknowledgements are displayed in the NOC widget, the default value is true
		AckChecked interface{} `json:"ackChecked,omitempty"`

		// The maximum number columns displayed in the NOC widget
		DisplayColumn int32 `json:"displayColumn,omitempty"`

		// Whether or not critical alerts are displayed in the NOC widget, the default value is true
		DisplayCriticalAlert interface{} `json:"displayCriticalAlert,omitempty"`

		// Whether or not error alerts are displayed in the NOC widget, the default value is true
		DisplayErrorAlert interface{} `json:"displayErrorAlert,omitempty"`

		// Whether or not warning alerts are displayed in the NOC widget, the default value is true
		DisplayWarnAlert interface{} `json:"displayWarnAlert,omitempty"`

		// Whether or not SDTs are displayed in the NOC widget, the default value is true
		SDTChecked interface{} `json:"sdtChecked,omitempty"`

		// How NOC items are sorted
		SortBy string `json:"sortBy,omitempty"`
	}{

		AckChecked: m.AckChecked,

		DisplayColumn: m.DisplayColumn,

		DisplayCriticalAlert: m.DisplayCriticalAlert,

		DisplayErrorAlert: m.DisplayErrorAlert,

		DisplayWarnAlert: m.DisplayWarnAlert,

		SDTChecked: m.SDTChecked,

		SortBy: m.SortBy,
	})
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

		Items []NOCItemBase `json:"items"`
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

		Items: m.Items(),
	})
	if err != nil {
		return nil, err
	}

	return swag.ConcatJSON(b1, b2, b3), nil
}

// Validate validates this n o c widget
func (m *NOCWidget) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDashboardID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateItems(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NOCWidget) validateDashboardID(formats strfmt.Registry) error {

	if err := validate.Required("dashboardId", "body", m.DashboardID()); err != nil {
		return err
	}

	return nil
}

func (m *NOCWidget) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name()); err != nil {
		return err
	}

	return nil
}

func (m *NOCWidget) validateItems(formats strfmt.Registry) error {

	if err := validate.Required("items", "body", m.Items()); err != nil {
		return err
	}

	for i := 0; i < len(m.Items()); i++ {

		if err := m.itemsField[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("items" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

// ContextValidate validate this n o c widget based on the context it is used
func (m *NOCWidget) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLastUpdatedBy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLastUpdatedOn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUserPermission(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateItems(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NOCWidget) contextValidateLastUpdatedBy(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "lastUpdatedBy", "body", string(m.LastUpdatedBy())); err != nil {
		return err
	}

	return nil
}

func (m *NOCWidget) contextValidateLastUpdatedOn(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "lastUpdatedOn", "body", int64(m.LastUpdatedOn())); err != nil {
		return err
	}

	return nil
}

func (m *NOCWidget) contextValidateUserPermission(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "userPermission", "body", string(m.UserPermission())); err != nil {
		return err
	}

	return nil
}

func (m *NOCWidget) contextValidateItems(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Items()); i++ {

		if err := m.itemsField[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("items" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *NOCWidget) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NOCWidget) UnmarshalBinary(b []byte) error {
	var res NOCWidget
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
