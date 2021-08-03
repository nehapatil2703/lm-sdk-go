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

// NewUpdateDeviceGroupDatasourceByIDParams creates a new UpdateDeviceGroupDatasourceByIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateDeviceGroupDatasourceByIDParams() *UpdateDeviceGroupDatasourceByIDParams {
	return &UpdateDeviceGroupDatasourceByIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateDeviceGroupDatasourceByIDParamsWithTimeout creates a new UpdateDeviceGroupDatasourceByIDParams object
// with the ability to set a timeout on a request.
func NewUpdateDeviceGroupDatasourceByIDParamsWithTimeout(timeout time.Duration) *UpdateDeviceGroupDatasourceByIDParams {
	return &UpdateDeviceGroupDatasourceByIDParams{
		timeout: timeout,
	}
}

// NewUpdateDeviceGroupDatasourceByIDParamsWithContext creates a new UpdateDeviceGroupDatasourceByIDParams object
// with the ability to set a context for a request.
func NewUpdateDeviceGroupDatasourceByIDParamsWithContext(ctx context.Context) *UpdateDeviceGroupDatasourceByIDParams {
	return &UpdateDeviceGroupDatasourceByIDParams{
		Context: ctx,
	}
}

// NewUpdateDeviceGroupDatasourceByIDParamsWithHTTPClient creates a new UpdateDeviceGroupDatasourceByIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateDeviceGroupDatasourceByIDParamsWithHTTPClient(client *http.Client) *UpdateDeviceGroupDatasourceByIDParams {
	return &UpdateDeviceGroupDatasourceByIDParams{
		HTTPClient: client,
	}
}

/* UpdateDeviceGroupDatasourceByIDParams contains all the parameters to send to the API endpoint
   for the update device group datasource by Id operation.

   Typically these are written to a http.Request.
*/
type UpdateDeviceGroupDatasourceByIDParams struct {

	// UserAgent.
	//
	// Default: "Logicmonitor/SDK: Argus Dist-v2.0.0-argus5-4-gd30bcfd-dirty"
	UserAgent *string

	// Body.
	Body *models.DeviceGroupDataSource

	// DeviceGroupID.
	//
	// Format: int32
	DeviceGroupID int32

	// ID.
	//
	// Format: int32
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update device group datasource by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateDeviceGroupDatasourceByIDParams) WithDefaults() *UpdateDeviceGroupDatasourceByIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update device group datasource by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateDeviceGroupDatasourceByIDParams) SetDefaults() {
	var (
		userAgentDefault = string("Logicmonitor/SDK: Argus Dist-v2.0.0-argus5-4-gd30bcfd-dirty")
	)

	val := UpdateDeviceGroupDatasourceByIDParams{
		UserAgent: &userAgentDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the update device group datasource by Id params
func (o *UpdateDeviceGroupDatasourceByIDParams) WithTimeout(timeout time.Duration) *UpdateDeviceGroupDatasourceByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update device group datasource by Id params
func (o *UpdateDeviceGroupDatasourceByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update device group datasource by Id params
func (o *UpdateDeviceGroupDatasourceByIDParams) WithContext(ctx context.Context) *UpdateDeviceGroupDatasourceByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update device group datasource by Id params
func (o *UpdateDeviceGroupDatasourceByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update device group datasource by Id params
func (o *UpdateDeviceGroupDatasourceByIDParams) WithHTTPClient(client *http.Client) *UpdateDeviceGroupDatasourceByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update device group datasource by Id params
func (o *UpdateDeviceGroupDatasourceByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUserAgent adds the userAgent to the update device group datasource by Id params
func (o *UpdateDeviceGroupDatasourceByIDParams) WithUserAgent(userAgent *string) *UpdateDeviceGroupDatasourceByIDParams {
	o.SetUserAgent(userAgent)
	return o
}

// SetUserAgent adds the userAgent to the update device group datasource by Id params
func (o *UpdateDeviceGroupDatasourceByIDParams) SetUserAgent(userAgent *string) {
	o.UserAgent = userAgent
}

// WithBody adds the body to the update device group datasource by Id params
func (o *UpdateDeviceGroupDatasourceByIDParams) WithBody(body *models.DeviceGroupDataSource) *UpdateDeviceGroupDatasourceByIDParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update device group datasource by Id params
func (o *UpdateDeviceGroupDatasourceByIDParams) SetBody(body *models.DeviceGroupDataSource) {
	o.Body = body
}

// WithDeviceGroupID adds the deviceGroupID to the update device group datasource by Id params
func (o *UpdateDeviceGroupDatasourceByIDParams) WithDeviceGroupID(deviceGroupID int32) *UpdateDeviceGroupDatasourceByIDParams {
	o.SetDeviceGroupID(deviceGroupID)
	return o
}

// SetDeviceGroupID adds the deviceGroupId to the update device group datasource by Id params
func (o *UpdateDeviceGroupDatasourceByIDParams) SetDeviceGroupID(deviceGroupID int32) {
	o.DeviceGroupID = deviceGroupID
}

// WithID adds the id to the update device group datasource by Id params
func (o *UpdateDeviceGroupDatasourceByIDParams) WithID(id int32) *UpdateDeviceGroupDatasourceByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update device group datasource by Id params
func (o *UpdateDeviceGroupDatasourceByIDParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateDeviceGroupDatasourceByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param deviceGroupId
	if err := r.SetPathParam("deviceGroupId", swag.FormatInt32(o.DeviceGroupID)); err != nil {
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
