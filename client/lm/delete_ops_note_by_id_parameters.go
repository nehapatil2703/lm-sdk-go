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

// NewDeleteOpsNoteByIDParams creates a new DeleteOpsNoteByIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteOpsNoteByIDParams() *DeleteOpsNoteByIDParams {
	return &DeleteOpsNoteByIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteOpsNoteByIDParamsWithTimeout creates a new DeleteOpsNoteByIDParams object
// with the ability to set a timeout on a request.
func NewDeleteOpsNoteByIDParamsWithTimeout(timeout time.Duration) *DeleteOpsNoteByIDParams {
	return &DeleteOpsNoteByIDParams{
		timeout: timeout,
	}
}

// NewDeleteOpsNoteByIDParamsWithContext creates a new DeleteOpsNoteByIDParams object
// with the ability to set a context for a request.
func NewDeleteOpsNoteByIDParamsWithContext(ctx context.Context) *DeleteOpsNoteByIDParams {
	return &DeleteOpsNoteByIDParams{
		Context: ctx,
	}
}

// NewDeleteOpsNoteByIDParamsWithHTTPClient creates a new DeleteOpsNoteByIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteOpsNoteByIDParamsWithHTTPClient(client *http.Client) *DeleteOpsNoteByIDParams {
	return &DeleteOpsNoteByIDParams{
		HTTPClient: client,
	}
}

/* DeleteOpsNoteByIDParams contains all the parameters to send to the API endpoint
   for the delete ops note by Id operation.

   Typically these are written to a http.Request.
*/
type DeleteOpsNoteByIDParams struct {

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

// WithDefaults hydrates default values in the delete ops note by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteOpsNoteByIDParams) WithDefaults() *DeleteOpsNoteByIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete ops note by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteOpsNoteByIDParams) SetDefaults() {
	var (
		userAgentDefault = string("Logicmonitor/SDK: Argus Dist-95bb3f4-dirty")
	)

	val := DeleteOpsNoteByIDParams{
		UserAgent: &userAgentDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the delete ops note by Id params
func (o *DeleteOpsNoteByIDParams) WithTimeout(timeout time.Duration) *DeleteOpsNoteByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete ops note by Id params
func (o *DeleteOpsNoteByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete ops note by Id params
func (o *DeleteOpsNoteByIDParams) WithContext(ctx context.Context) *DeleteOpsNoteByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete ops note by Id params
func (o *DeleteOpsNoteByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete ops note by Id params
func (o *DeleteOpsNoteByIDParams) WithHTTPClient(client *http.Client) *DeleteOpsNoteByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete ops note by Id params
func (o *DeleteOpsNoteByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUserAgent adds the userAgent to the delete ops note by Id params
func (o *DeleteOpsNoteByIDParams) WithUserAgent(userAgent *string) *DeleteOpsNoteByIDParams {
	o.SetUserAgent(userAgent)
	return o
}

// SetUserAgent adds the userAgent to the delete ops note by Id params
func (o *DeleteOpsNoteByIDParams) SetUserAgent(userAgent *string) {
	o.UserAgent = userAgent
}

// WithID adds the id to the delete ops note by Id params
func (o *DeleteOpsNoteByIDParams) WithID(id string) *DeleteOpsNoteByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete ops note by Id params
func (o *DeleteOpsNoteByIDParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteOpsNoteByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
