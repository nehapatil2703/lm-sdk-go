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

// NewGetDeviceGroupDatasourceListParams creates a new GetDeviceGroupDatasourceListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetDeviceGroupDatasourceListParams() *GetDeviceGroupDatasourceListParams {
	return &GetDeviceGroupDatasourceListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetDeviceGroupDatasourceListParamsWithTimeout creates a new GetDeviceGroupDatasourceListParams object
// with the ability to set a timeout on a request.
func NewGetDeviceGroupDatasourceListParamsWithTimeout(timeout time.Duration) *GetDeviceGroupDatasourceListParams {
	return &GetDeviceGroupDatasourceListParams{
		timeout: timeout,
	}
}

// NewGetDeviceGroupDatasourceListParamsWithContext creates a new GetDeviceGroupDatasourceListParams object
// with the ability to set a context for a request.
func NewGetDeviceGroupDatasourceListParamsWithContext(ctx context.Context) *GetDeviceGroupDatasourceListParams {
	return &GetDeviceGroupDatasourceListParams{
		Context: ctx,
	}
}

// NewGetDeviceGroupDatasourceListParamsWithHTTPClient creates a new GetDeviceGroupDatasourceListParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetDeviceGroupDatasourceListParamsWithHTTPClient(client *http.Client) *GetDeviceGroupDatasourceListParams {
	return &GetDeviceGroupDatasourceListParams{
		HTTPClient: client,
	}
}

/* GetDeviceGroupDatasourceListParams contains all the parameters to send to the API endpoint
   for the get device group datasource list operation.

   Typically these are written to a http.Request.
*/
type GetDeviceGroupDatasourceListParams struct {

	// UserAgent.
	//
	// Default: "Logicmonitor/SDK: Argus Dist-v2.0.0-argus5-4-gd30bcfd-dirty"
	UserAgent *string

	// DeviceGroupID.
	//
	// Format: int32
	DeviceGroupID int32

	// Fields.
	Fields *string

	// Filter.
	Filter *string

	// IncludeDisabledDataSourceWithoutInstance.
	IncludeDisabledDataSourceWithoutInstance *bool

	// Offset.
	//
	// Format: int32
	Offset *int32

	// Size.
	//
	// Format: int32
	// Default: 50
	Size *int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get device group datasource list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDeviceGroupDatasourceListParams) WithDefaults() *GetDeviceGroupDatasourceListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get device group datasource list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDeviceGroupDatasourceListParams) SetDefaults() {
	var (
		userAgentDefault = string("Logicmonitor/SDK: Argus Dist-v2.0.0-argus5-4-gd30bcfd-dirty")

		includeDisabledDataSourceWithoutInstanceDefault = bool(false)

		offsetDefault = int32(0)

		sizeDefault = int32(50)
	)

	val := GetDeviceGroupDatasourceListParams{
		UserAgent:                                &userAgentDefault,
		IncludeDisabledDataSourceWithoutInstance: &includeDisabledDataSourceWithoutInstanceDefault,
		Offset:                                   &offsetDefault,
		Size:                                     &sizeDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get device group datasource list params
func (o *GetDeviceGroupDatasourceListParams) WithTimeout(timeout time.Duration) *GetDeviceGroupDatasourceListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get device group datasource list params
func (o *GetDeviceGroupDatasourceListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get device group datasource list params
func (o *GetDeviceGroupDatasourceListParams) WithContext(ctx context.Context) *GetDeviceGroupDatasourceListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get device group datasource list params
func (o *GetDeviceGroupDatasourceListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get device group datasource list params
func (o *GetDeviceGroupDatasourceListParams) WithHTTPClient(client *http.Client) *GetDeviceGroupDatasourceListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get device group datasource list params
func (o *GetDeviceGroupDatasourceListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUserAgent adds the userAgent to the get device group datasource list params
func (o *GetDeviceGroupDatasourceListParams) WithUserAgent(userAgent *string) *GetDeviceGroupDatasourceListParams {
	o.SetUserAgent(userAgent)
	return o
}

// SetUserAgent adds the userAgent to the get device group datasource list params
func (o *GetDeviceGroupDatasourceListParams) SetUserAgent(userAgent *string) {
	o.UserAgent = userAgent
}

// WithDeviceGroupID adds the deviceGroupID to the get device group datasource list params
func (o *GetDeviceGroupDatasourceListParams) WithDeviceGroupID(deviceGroupID int32) *GetDeviceGroupDatasourceListParams {
	o.SetDeviceGroupID(deviceGroupID)
	return o
}

// SetDeviceGroupID adds the deviceGroupId to the get device group datasource list params
func (o *GetDeviceGroupDatasourceListParams) SetDeviceGroupID(deviceGroupID int32) {
	o.DeviceGroupID = deviceGroupID
}

// WithFields adds the fields to the get device group datasource list params
func (o *GetDeviceGroupDatasourceListParams) WithFields(fields *string) *GetDeviceGroupDatasourceListParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get device group datasource list params
func (o *GetDeviceGroupDatasourceListParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithFilter adds the filter to the get device group datasource list params
func (o *GetDeviceGroupDatasourceListParams) WithFilter(filter *string) *GetDeviceGroupDatasourceListParams {
	o.SetFilter(filter)
	return o
}

// SetFilter adds the filter to the get device group datasource list params
func (o *GetDeviceGroupDatasourceListParams) SetFilter(filter *string) {
	o.Filter = filter
}

// WithIncludeDisabledDataSourceWithoutInstance adds the includeDisabledDataSourceWithoutInstance to the get device group datasource list params
func (o *GetDeviceGroupDatasourceListParams) WithIncludeDisabledDataSourceWithoutInstance(includeDisabledDataSourceWithoutInstance *bool) *GetDeviceGroupDatasourceListParams {
	o.SetIncludeDisabledDataSourceWithoutInstance(includeDisabledDataSourceWithoutInstance)
	return o
}

// SetIncludeDisabledDataSourceWithoutInstance adds the includeDisabledDataSourceWithoutInstance to the get device group datasource list params
func (o *GetDeviceGroupDatasourceListParams) SetIncludeDisabledDataSourceWithoutInstance(includeDisabledDataSourceWithoutInstance *bool) {
	o.IncludeDisabledDataSourceWithoutInstance = includeDisabledDataSourceWithoutInstance
}

// WithOffset adds the offset to the get device group datasource list params
func (o *GetDeviceGroupDatasourceListParams) WithOffset(offset *int32) *GetDeviceGroupDatasourceListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the get device group datasource list params
func (o *GetDeviceGroupDatasourceListParams) SetOffset(offset *int32) {
	o.Offset = offset
}

// WithSize adds the size to the get device group datasource list params
func (o *GetDeviceGroupDatasourceListParams) WithSize(size *int32) *GetDeviceGroupDatasourceListParams {
	o.SetSize(size)
	return o
}

// SetSize adds the size to the get device group datasource list params
func (o *GetDeviceGroupDatasourceListParams) SetSize(size *int32) {
	o.Size = size
}

// WriteToRequest writes these params to a swagger request
func (o *GetDeviceGroupDatasourceListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param deviceGroupId
	if err := r.SetPathParam("deviceGroupId", swag.FormatInt32(o.DeviceGroupID)); err != nil {
		return err
	}

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

	if o.Filter != nil {

		// query param filter
		var qrFilter string

		if o.Filter != nil {
			qrFilter = *o.Filter
		}
		qFilter := qrFilter
		if qFilter != "" {

			if err := r.SetQueryParam("filter", qFilter); err != nil {
				return err
			}
		}
	}

	if o.IncludeDisabledDataSourceWithoutInstance != nil {

		// query param includeDisabledDataSourceWithoutInstance
		var qrIncludeDisabledDataSourceWithoutInstance bool

		if o.IncludeDisabledDataSourceWithoutInstance != nil {
			qrIncludeDisabledDataSourceWithoutInstance = *o.IncludeDisabledDataSourceWithoutInstance
		}
		qIncludeDisabledDataSourceWithoutInstance := swag.FormatBool(qrIncludeDisabledDataSourceWithoutInstance)
		if qIncludeDisabledDataSourceWithoutInstance != "" {

			if err := r.SetQueryParam("includeDisabledDataSourceWithoutInstance", qIncludeDisabledDataSourceWithoutInstance); err != nil {
				return err
			}
		}
	}

	if o.Offset != nil {

		// query param offset
		var qrOffset int32

		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt32(qrOffset)
		if qOffset != "" {

			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}
	}

	if o.Size != nil {

		// query param size
		var qrSize int32

		if o.Size != nil {
			qrSize = *o.Size
		}
		qSize := swag.FormatInt32(qrSize)
		if qSize != "" {

			if err := r.SetQueryParam("size", qSize); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
