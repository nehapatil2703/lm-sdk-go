// Code generated by go-swagger; DO NOT EDIT.

package lm

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"golang.org/x/net/context"
)

// NewGetDashboardByIDParams creates a new GetDashboardByIDParams object
// with the default values initialized.
func NewGetDashboardByIDParams() *GetDashboardByIDParams {
	var (
		formatDefault   = string("json")
		templateDefault = bool(false)
	)
	return &GetDashboardByIDParams{
		Format:   &formatDefault,
		Template: &templateDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetDashboardByIDParamsWithTimeout creates a new GetDashboardByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetDashboardByIDParamsWithTimeout(timeout time.Duration) *GetDashboardByIDParams {
	var (
		formatDefault   = string("json")
		templateDefault = bool(false)
	)
	return &GetDashboardByIDParams{
		Format:   &formatDefault,
		Template: &templateDefault,

		timeout: timeout,
	}
}

// NewGetDashboardByIDParamsWithContext creates a new GetDashboardByIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetDashboardByIDParamsWithContext(ctx context.Context) *GetDashboardByIDParams {
	var (
		formatDefault   = string("json")
		templateDefault = bool(false)
	)
	return &GetDashboardByIDParams{
		Format:   &formatDefault,
		Template: &templateDefault,

		Context: ctx,
	}
}

// NewGetDashboardByIDParamsWithHTTPClient creates a new GetDashboardByIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetDashboardByIDParamsWithHTTPClient(client *http.Client) *GetDashboardByIDParams {
	var (
		formatDefault   = string("json")
		templateDefault = bool(false)
	)
	return &GetDashboardByIDParams{
		Format:     &formatDefault,
		Template:   &templateDefault,
		HTTPClient: client,
	}
}

/*GetDashboardByIDParams contains all the parameters to send to the API endpoint
for the get dashboard by Id operation typically these are written to a http.Request
*/
type GetDashboardByIDParams struct {

	/*Fields*/
	Fields *string
	/*Format*/
	Format *string
	/*ID*/
	ID int32
	/*Template*/
	Template *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get dashboard by Id params
func (o *GetDashboardByIDParams) WithTimeout(timeout time.Duration) *GetDashboardByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get dashboard by Id params
func (o *GetDashboardByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get dashboard by Id params
func (o *GetDashboardByIDParams) WithContext(ctx context.Context) *GetDashboardByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get dashboard by Id params
func (o *GetDashboardByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get dashboard by Id params
func (o *GetDashboardByIDParams) WithHTTPClient(client *http.Client) *GetDashboardByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get dashboard by Id params
func (o *GetDashboardByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the get dashboard by Id params
func (o *GetDashboardByIDParams) WithFields(fields *string) *GetDashboardByIDParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get dashboard by Id params
func (o *GetDashboardByIDParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithFormat adds the format to the get dashboard by Id params
func (o *GetDashboardByIDParams) WithFormat(format *string) *GetDashboardByIDParams {
	o.SetFormat(format)
	return o
}

// SetFormat adds the format to the get dashboard by Id params
func (o *GetDashboardByIDParams) SetFormat(format *string) {
	o.Format = format
}

// WithID adds the id to the get dashboard by Id params
func (o *GetDashboardByIDParams) WithID(id int32) *GetDashboardByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get dashboard by Id params
func (o *GetDashboardByIDParams) SetID(id int32) {
	o.ID = id
}

// WithTemplate adds the template to the get dashboard by Id params
func (o *GetDashboardByIDParams) WithTemplate(template *bool) *GetDashboardByIDParams {
	o.SetTemplate(template)
	return o
}

// SetTemplate adds the template to the get dashboard by Id params
func (o *GetDashboardByIDParams) SetTemplate(template *bool) {
	o.Template = template
}

// WriteToRequest writes these params to a swagger request
func (o *GetDashboardByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {
	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Fields != nil {

		// query param fields
		var qrFields string
		if o.Fields != nil {
			qrFields = *o.Fields
		}
		qFields := qrFields
		if qFields != "" {
			if err := r.SetQueryParam("fields", qFields); err != nil {
				return err
			}
		}

	}

	if o.Format != nil {

		// query param format
		var qrFormat string
		if o.Format != nil {
			qrFormat = *o.Format
		}
		qFormat := qrFormat
		if qFormat != "" {
			if err := r.SetQueryParam("format", qFormat); err != nil {
				return err
			}
		}

	}

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt32(o.ID)); err != nil {
		return err
	}

	if o.Template != nil {

		// query param template
		var qrTemplate bool
		if o.Template != nil {
			qrTemplate = *o.Template
		}
		qTemplate := swag.FormatBool(qrTemplate)
		if qTemplate != "" {
			if err := r.SetQueryParam("template", qTemplate); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
