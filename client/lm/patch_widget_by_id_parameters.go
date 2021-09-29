// Code generated by go-swagger; DO NOT EDIT.

package lm

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/logicmonitor/lm-sdk-go/models"
)

// NewPatchWidgetByIDParams creates a new PatchWidgetByIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPatchWidgetByIDParams() *PatchWidgetByIDParams {
	return &PatchWidgetByIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPatchWidgetByIDParamsWithTimeout creates a new PatchWidgetByIDParams object
// with the ability to set a timeout on a request.
func NewPatchWidgetByIDParamsWithTimeout(timeout time.Duration) *PatchWidgetByIDParams {
	return &PatchWidgetByIDParams{
		timeout: timeout,
	}
}

// NewPatchWidgetByIDParamsWithContext creates a new PatchWidgetByIDParams object
// with the ability to set a context for a request.
func NewPatchWidgetByIDParamsWithContext(ctx context.Context) *PatchWidgetByIDParams {
	return &PatchWidgetByIDParams{
		Context: ctx,
	}
}

// NewPatchWidgetByIDParamsWithHTTPClient creates a new PatchWidgetByIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewPatchWidgetByIDParamsWithHTTPClient(client *http.Client) *PatchWidgetByIDParams {
	return &PatchWidgetByIDParams{
		HTTPClient: client,
	}
}

/* PatchWidgetByIDParams contains all the parameters to send to the API endpoint
   for the patch widget by Id operation.

   Typically these are written to a http.Request.
*/
type PatchWidgetByIDParams struct {

	// PatchFields.
	PatchFields *string

	// UserAgent.
	//
	// Default: "Logicmonitor/SDK: Argus Dist-95bb3f4-dirty"
	UserAgent *string

	// Body.
	Body models.Widget

	// ID.
	//
	// Format: int32
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the patch widget by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchWidgetByIDParams) WithDefaults() *PatchWidgetByIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the patch widget by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchWidgetByIDParams) SetDefaults() {
	var (
		userAgentDefault = string("Logicmonitor/SDK: Argus Dist-95bb3f4-dirty")
	)

	val := PatchWidgetByIDParams{
		UserAgent: &userAgentDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the patch widget by Id params
func (o *PatchWidgetByIDParams) WithTimeout(timeout time.Duration) *PatchWidgetByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch widget by Id params
func (o *PatchWidgetByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch widget by Id params
func (o *PatchWidgetByIDParams) WithContext(ctx context.Context) *PatchWidgetByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch widget by Id params
func (o *PatchWidgetByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch widget by Id params
func (o *PatchWidgetByIDParams) WithHTTPClient(client *http.Client) *PatchWidgetByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch widget by Id params
func (o *PatchWidgetByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPatchFields adds the patchFields to the patch widget by Id params
func (o *PatchWidgetByIDParams) WithPatchFields(patchFields *string) *PatchWidgetByIDParams {
	o.SetPatchFields(patchFields)
	return o
}

// SetPatchFields adds the patchFields to the patch widget by Id params
func (o *PatchWidgetByIDParams) SetPatchFields(patchFields *string) {
	o.PatchFields = patchFields
}

// WithUserAgent adds the userAgent to the patch widget by Id params
func (o *PatchWidgetByIDParams) WithUserAgent(userAgent *string) *PatchWidgetByIDParams {
	o.SetUserAgent(userAgent)
	return o
}

// SetUserAgent adds the userAgent to the patch widget by Id params
func (o *PatchWidgetByIDParams) SetUserAgent(userAgent *string) {
	o.UserAgent = userAgent
}

// WithBody adds the body to the patch widget by Id params
func (o *PatchWidgetByIDParams) WithBody(body models.Widget) *PatchWidgetByIDParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch widget by Id params
func (o *PatchWidgetByIDParams) SetBody(body models.Widget) {
	o.Body = body
}

// WithID adds the id to the patch widget by Id params
func (o *PatchWidgetByIDParams) WithID(id int32) *PatchWidgetByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the patch widget by Id params
func (o *PatchWidgetByIDParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PatchWidgetByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.PatchFields != nil {

		// query param PatchFields
		var qrPatchFields string

		if o.PatchFields != nil {
			qrPatchFields = *o.PatchFields
		}
		qPatchFields := qrPatchFields
		if qPatchFields != "" {

			if err := r.SetQueryParam("PatchFields", qPatchFields); err != nil {
				return err
			}
		}
	}

	if o.UserAgent != nil {

		// header param User-Agent
		if err := r.SetHeaderParam("User-Agent", *o.UserAgent); err != nil {
			return err
		}
	}
	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt32(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
