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
)

// NewGetAuditLogByIDParams creates a new GetAuditLogByIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAuditLogByIDParams() *GetAuditLogByIDParams {
	return &GetAuditLogByIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAuditLogByIDParamsWithTimeout creates a new GetAuditLogByIDParams object
// with the ability to set a timeout on a request.
func NewGetAuditLogByIDParamsWithTimeout(timeout time.Duration) *GetAuditLogByIDParams {
	return &GetAuditLogByIDParams{
		timeout: timeout,
	}
}

// NewGetAuditLogByIDParamsWithContext creates a new GetAuditLogByIDParams object
// with the ability to set a context for a request.
func NewGetAuditLogByIDParamsWithContext(ctx context.Context) *GetAuditLogByIDParams {
	return &GetAuditLogByIDParams{
		Context: ctx,
	}
}

// NewGetAuditLogByIDParamsWithHTTPClient creates a new GetAuditLogByIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAuditLogByIDParamsWithHTTPClient(client *http.Client) *GetAuditLogByIDParams {
	return &GetAuditLogByIDParams{
		HTTPClient: client,
	}
}

/* GetAuditLogByIDParams contains all the parameters to send to the API endpoint
   for the get audit log by Id operation.

   Typically these are written to a http.Request.
*/
type GetAuditLogByIDParams struct {

	// UserAgent.
	//
	// Default: "Logicmonitor/SDK: Argus Dist-95bb3f4-dirty"
	UserAgent *string

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get audit log by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAuditLogByIDParams) WithDefaults() *GetAuditLogByIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get audit log by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAuditLogByIDParams) SetDefaults() {
	var (
		userAgentDefault = string("Logicmonitor/SDK: Argus Dist-95bb3f4-dirty")
	)

	val := GetAuditLogByIDParams{
		UserAgent: &userAgentDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get audit log by Id params
func (o *GetAuditLogByIDParams) WithTimeout(timeout time.Duration) *GetAuditLogByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get audit log by Id params
func (o *GetAuditLogByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get audit log by Id params
func (o *GetAuditLogByIDParams) WithContext(ctx context.Context) *GetAuditLogByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get audit log by Id params
func (o *GetAuditLogByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get audit log by Id params
func (o *GetAuditLogByIDParams) WithHTTPClient(client *http.Client) *GetAuditLogByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get audit log by Id params
func (o *GetAuditLogByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUserAgent adds the userAgent to the get audit log by Id params
func (o *GetAuditLogByIDParams) WithUserAgent(userAgent *string) *GetAuditLogByIDParams {
	o.SetUserAgent(userAgent)
	return o
}

// SetUserAgent adds the userAgent to the get audit log by Id params
func (o *GetAuditLogByIDParams) SetUserAgent(userAgent *string) {
	o.UserAgent = userAgent
}

// WithID adds the id to the get audit log by Id params
func (o *GetAuditLogByIDParams) WithID(id string) *GetAuditLogByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get audit log by Id params
func (o *GetAuditLogByIDParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetAuditLogByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.UserAgent != nil {

		// header param User-Agent
		if err := r.SetHeaderParam("User-Agent", *o.UserAgent); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
