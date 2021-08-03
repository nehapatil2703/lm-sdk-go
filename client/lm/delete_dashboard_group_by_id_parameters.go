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

// NewDeleteDashboardGroupByIDParams creates a new DeleteDashboardGroupByIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteDashboardGroupByIDParams() *DeleteDashboardGroupByIDParams {
	return &DeleteDashboardGroupByIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteDashboardGroupByIDParamsWithTimeout creates a new DeleteDashboardGroupByIDParams object
// with the ability to set a timeout on a request.
func NewDeleteDashboardGroupByIDParamsWithTimeout(timeout time.Duration) *DeleteDashboardGroupByIDParams {
	return &DeleteDashboardGroupByIDParams{
		timeout: timeout,
	}
}

// NewDeleteDashboardGroupByIDParamsWithContext creates a new DeleteDashboardGroupByIDParams object
// with the ability to set a context for a request.
func NewDeleteDashboardGroupByIDParamsWithContext(ctx context.Context) *DeleteDashboardGroupByIDParams {
	return &DeleteDashboardGroupByIDParams{
		Context: ctx,
	}
}

// NewDeleteDashboardGroupByIDParamsWithHTTPClient creates a new DeleteDashboardGroupByIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteDashboardGroupByIDParamsWithHTTPClient(client *http.Client) *DeleteDashboardGroupByIDParams {
	return &DeleteDashboardGroupByIDParams{
		HTTPClient: client,
	}
}

/* DeleteDashboardGroupByIDParams contains all the parameters to send to the API endpoint
   for the delete dashboard group by Id operation.

   Typically these are written to a http.Request.
*/
type DeleteDashboardGroupByIDParams struct {

	// UserAgent.
	//
	// Default: "Logicmonitor/SDK: Argus Dist-v2.0.0-argus5-4-gd30bcfd-dirty"
	UserAgent *string

	// AllowNonEmptyGroup.
	AllowNonEmptyGroup *bool

	// ID.
	//
	// Format: int32
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete dashboard group by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteDashboardGroupByIDParams) WithDefaults() *DeleteDashboardGroupByIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete dashboard group by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteDashboardGroupByIDParams) SetDefaults() {
	var (
		userAgentDefault = string("Logicmonitor/SDK: Argus Dist-v2.0.0-argus5-4-gd30bcfd-dirty")

		allowNonEmptyGroupDefault = bool(false)
	)

	val := DeleteDashboardGroupByIDParams{
		UserAgent:          &userAgentDefault,
		AllowNonEmptyGroup: &allowNonEmptyGroupDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the delete dashboard group by Id params
func (o *DeleteDashboardGroupByIDParams) WithTimeout(timeout time.Duration) *DeleteDashboardGroupByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete dashboard group by Id params
func (o *DeleteDashboardGroupByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete dashboard group by Id params
func (o *DeleteDashboardGroupByIDParams) WithContext(ctx context.Context) *DeleteDashboardGroupByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete dashboard group by Id params
func (o *DeleteDashboardGroupByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete dashboard group by Id params
func (o *DeleteDashboardGroupByIDParams) WithHTTPClient(client *http.Client) *DeleteDashboardGroupByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete dashboard group by Id params
func (o *DeleteDashboardGroupByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUserAgent adds the userAgent to the delete dashboard group by Id params
func (o *DeleteDashboardGroupByIDParams) WithUserAgent(userAgent *string) *DeleteDashboardGroupByIDParams {
	o.SetUserAgent(userAgent)
	return o
}

// SetUserAgent adds the userAgent to the delete dashboard group by Id params
func (o *DeleteDashboardGroupByIDParams) SetUserAgent(userAgent *string) {
	o.UserAgent = userAgent
}

// WithAllowNonEmptyGroup adds the allowNonEmptyGroup to the delete dashboard group by Id params
func (o *DeleteDashboardGroupByIDParams) WithAllowNonEmptyGroup(allowNonEmptyGroup *bool) *DeleteDashboardGroupByIDParams {
	o.SetAllowNonEmptyGroup(allowNonEmptyGroup)
	return o
}

// SetAllowNonEmptyGroup adds the allowNonEmptyGroup to the delete dashboard group by Id params
func (o *DeleteDashboardGroupByIDParams) SetAllowNonEmptyGroup(allowNonEmptyGroup *bool) {
	o.AllowNonEmptyGroup = allowNonEmptyGroup
}

// WithID adds the id to the delete dashboard group by Id params
func (o *DeleteDashboardGroupByIDParams) WithID(id int32) *DeleteDashboardGroupByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete dashboard group by Id params
func (o *DeleteDashboardGroupByIDParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteDashboardGroupByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.AllowNonEmptyGroup != nil {

		// query param allowNonEmptyGroup
		var qrAllowNonEmptyGroup bool

		if o.AllowNonEmptyGroup != nil {
			qrAllowNonEmptyGroup = *o.AllowNonEmptyGroup
		}
		qAllowNonEmptyGroup := swag.FormatBool(qrAllowNonEmptyGroup)
		if qAllowNonEmptyGroup != "" {

			if err := r.SetQueryParam("allowNonEmptyGroup", qAllowNonEmptyGroup); err != nil {
				return err
			}
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
