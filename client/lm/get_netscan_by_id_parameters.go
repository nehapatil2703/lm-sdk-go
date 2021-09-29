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
)

// NewGetNetscanByIDParams creates a new GetNetscanByIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetNetscanByIDParams() *GetNetscanByIDParams {
	return &GetNetscanByIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetNetscanByIDParamsWithTimeout creates a new GetNetscanByIDParams object
// with the ability to set a timeout on a request.
func NewGetNetscanByIDParamsWithTimeout(timeout time.Duration) *GetNetscanByIDParams {
	return &GetNetscanByIDParams{
		timeout: timeout,
	}
}

// NewGetNetscanByIDParamsWithContext creates a new GetNetscanByIDParams object
// with the ability to set a context for a request.
func NewGetNetscanByIDParamsWithContext(ctx context.Context) *GetNetscanByIDParams {
	return &GetNetscanByIDParams{
		Context: ctx,
	}
}

// NewGetNetscanByIDParamsWithHTTPClient creates a new GetNetscanByIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetNetscanByIDParamsWithHTTPClient(client *http.Client) *GetNetscanByIDParams {
	return &GetNetscanByIDParams{
		HTTPClient: client,
	}
}

/* GetNetscanByIDParams contains all the parameters to send to the API endpoint
   for the get netscan by Id operation.

   Typically these are written to a http.Request.
*/
type GetNetscanByIDParams struct {

	// UserAgent.
	//
	// Default: "Logicmonitor/SDK: Argus Dist-95bb3f4-dirty"
	UserAgent *string

	// ID.
	//
	// Format: int32
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get netscan by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNetscanByIDParams) WithDefaults() *GetNetscanByIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get netscan by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNetscanByIDParams) SetDefaults() {
	var (
		userAgentDefault = string("Logicmonitor/SDK: Argus Dist-95bb3f4-dirty")
	)

	val := GetNetscanByIDParams{
		UserAgent: &userAgentDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get netscan by Id params
func (o *GetNetscanByIDParams) WithTimeout(timeout time.Duration) *GetNetscanByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get netscan by Id params
func (o *GetNetscanByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get netscan by Id params
func (o *GetNetscanByIDParams) WithContext(ctx context.Context) *GetNetscanByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get netscan by Id params
func (o *GetNetscanByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get netscan by Id params
func (o *GetNetscanByIDParams) WithHTTPClient(client *http.Client) *GetNetscanByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get netscan by Id params
func (o *GetNetscanByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUserAgent adds the userAgent to the get netscan by Id params
func (o *GetNetscanByIDParams) WithUserAgent(userAgent *string) *GetNetscanByIDParams {
	o.SetUserAgent(userAgent)
	return o
}

// SetUserAgent adds the userAgent to the get netscan by Id params
func (o *GetNetscanByIDParams) SetUserAgent(userAgent *string) {
	o.UserAgent = userAgent
}

// WithID adds the id to the get netscan by Id params
func (o *GetNetscanByIDParams) WithID(id int32) *GetNetscanByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get netscan by Id params
func (o *GetNetscanByIDParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetNetscanByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
	if err := r.SetPathParam("id", swag.FormatInt32(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
