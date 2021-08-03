// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// GenerateReportResult generate report result
//
// swagger:model GenerateReportResult
type GenerateReportResult struct {

	// The id of the report
	// Required: true
	ReportID *int32 `json:"reportId"`

	// The url of the generated report
	// Read Only: true
	Resulturl string `json:"resulturl,omitempty"`

	// The task id of the generating process
	// Required: true
	TaskID *int64 `json:"taskId"`
}

// Validate validates this generate report result
func (m *GenerateReportResult) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateReportID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTaskID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GenerateReportResult) validateReportID(formats strfmt.Registry) error {

	if err := validate.Required("reportId", "body", m.ReportID); err != nil {
		return err
	}

	return nil
}

func (m *GenerateReportResult) validateTaskID(formats strfmt.Registry) error {

	if err := validate.Required("taskId", "body", m.TaskID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this generate report result based on the context it is used
func (m *GenerateReportResult) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateResulturl(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GenerateReportResult) contextValidateResulturl(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "resulturl", "body", string(m.Resulturl)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GenerateReportResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GenerateReportResult) UnmarshalBinary(b []byte) error {
	var res GenerateReportResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
