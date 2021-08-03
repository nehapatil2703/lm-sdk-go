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

// NewUpdateDeviceGroupPropertyByNameParams creates a new UpdateDeviceGroupPropertyByNameParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateDeviceGroupPropertyByNameParams() *UpdateDeviceGroupPropertyByNameParams {
	return &UpdateDeviceGroupPropertyByNameParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateDeviceGroupPropertyByNameParamsWithTimeout creates a new UpdateDeviceGroupPropertyByNameParams object
// with the ability to set a timeout on a request.
func NewUpdateDeviceGroupPropertyByNameParamsWithTimeout(timeout time.Duration) *UpdateDeviceGroupPropertyByNameParams {
	return &UpdateDeviceGroupPropertyByNameParams{
		timeout: timeout,
	}
}

// NewUpdateDeviceGroupPropertyByNameParamsWithContext creates a new UpdateDeviceGroupPropertyByNameParams object
// with the ability to set a context for a request.
func NewUpdateDeviceGroupPropertyByNameParamsWithContext(ctx context.Context) *UpdateDeviceGroupPropertyByNameParams {
	return &UpdateDeviceGroupPropertyByNameParams{
		Context: ctx,
	}
}

// NewUpdateDeviceGroupPropertyByNameParamsWithHTTPClient creates a new UpdateDeviceGroupPropertyByNameParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateDeviceGroupPropertyByNameParamsWithHTTPClient(client *http.Client) *UpdateDeviceGroupPropertyByNameParams {
	return &UpdateDeviceGroupPropertyByNameParams{
		HTTPClient: client,
	}
}

/* UpdateDeviceGroupPropertyByNameParams contains all the parameters to send to the API endpoint
   for the update device group property by name operation.

   Typically these are written to a http.Request.
*/
type UpdateDeviceGroupPropertyByNameParams struct {

	// UserAgent.
	//
	// Default: "Logicmonitor/SDK: Argus Dist-v2.0.0-argus5-4-gd30bcfd-dirty"
	UserAgent *string

	// Body.
	Body *models.EntityProperty

	/* Gid.

	   group ID

	   Format: int32
	*/
	Gid int32

	// Name.
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update device group property by name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateDeviceGroupPropertyByNameParams) WithDefaults() *UpdateDeviceGroupPropertyByNameParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update device group property by name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateDeviceGroupPropertyByNameParams) SetDefaults() {
	var (
		userAgentDefault = string("Logicmonitor/SDK: Argus Dist-v2.0.0-argus5-4-gd30bcfd-dirty")
	)

	val := UpdateDeviceGroupPropertyByNameParams{
		UserAgent: &userAgentDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the update device group property by name params
func (o *UpdateDeviceGroupPropertyByNameParams) WithTimeout(timeout time.Duration) *UpdateDeviceGroupPropertyByNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update device group property by name params
func (o *UpdateDeviceGroupPropertyByNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update device group property by name params
func (o *UpdateDeviceGroupPropertyByNameParams) WithContext(ctx context.Context) *UpdateDeviceGroupPropertyByNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update device group property by name params
func (o *UpdateDeviceGroupPropertyByNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update device group property by name params
func (o *UpdateDeviceGroupPropertyByNameParams) WithHTTPClient(client *http.Client) *UpdateDeviceGroupPropertyByNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update device group property by name params
func (o *UpdateDeviceGroupPropertyByNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUserAgent adds the userAgent to the update device group property by name params
func (o *UpdateDeviceGroupPropertyByNameParams) WithUserAgent(userAgent *string) *UpdateDeviceGroupPropertyByNameParams {
	o.SetUserAgent(userAgent)
	return o
}

// SetUserAgent adds the userAgent to the update device group property by name params
func (o *UpdateDeviceGroupPropertyByNameParams) SetUserAgent(userAgent *string) {
	o.UserAgent = userAgent
}

// WithBody adds the body to the update device group property by name params
func (o *UpdateDeviceGroupPropertyByNameParams) WithBody(body *models.EntityProperty) *UpdateDeviceGroupPropertyByNameParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update device group property by name params
func (o *UpdateDeviceGroupPropertyByNameParams) SetBody(body *models.EntityProperty) {
	o.Body = body
}

// WithGid adds the gid to the update device group property by name params
func (o *UpdateDeviceGroupPropertyByNameParams) WithGid(gid int32) *UpdateDeviceGroupPropertyByNameParams {
	o.SetGid(gid)
	return o
}

// SetGid adds the gid to the update device group property by name params
func (o *UpdateDeviceGroupPropertyByNameParams) SetGid(gid int32) {
	o.Gid = gid
}

// WithName adds the name to the update device group property by name params
func (o *UpdateDeviceGroupPropertyByNameParams) WithName(name string) *UpdateDeviceGroupPropertyByNameParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the update device group property by name params
func (o *UpdateDeviceGroupPropertyByNameParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateDeviceGroupPropertyByNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param gid
	if err := r.SetPathParam("gid", swag.FormatInt32(o.Gid)); err != nil {
		return err
	}

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
