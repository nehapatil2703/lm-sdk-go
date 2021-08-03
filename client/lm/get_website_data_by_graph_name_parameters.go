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

// NewGetWebsiteDataByGraphNameParams creates a new GetWebsiteDataByGraphNameParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetWebsiteDataByGraphNameParams() *GetWebsiteDataByGraphNameParams {
	return &GetWebsiteDataByGraphNameParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetWebsiteDataByGraphNameParamsWithTimeout creates a new GetWebsiteDataByGraphNameParams object
// with the ability to set a timeout on a request.
func NewGetWebsiteDataByGraphNameParamsWithTimeout(timeout time.Duration) *GetWebsiteDataByGraphNameParams {
	return &GetWebsiteDataByGraphNameParams{
		timeout: timeout,
	}
}

// NewGetWebsiteDataByGraphNameParamsWithContext creates a new GetWebsiteDataByGraphNameParams object
// with the ability to set a context for a request.
func NewGetWebsiteDataByGraphNameParamsWithContext(ctx context.Context) *GetWebsiteDataByGraphNameParams {
	return &GetWebsiteDataByGraphNameParams{
		Context: ctx,
	}
}

// NewGetWebsiteDataByGraphNameParamsWithHTTPClient creates a new GetWebsiteDataByGraphNameParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetWebsiteDataByGraphNameParamsWithHTTPClient(client *http.Client) *GetWebsiteDataByGraphNameParams {
	return &GetWebsiteDataByGraphNameParams{
		HTTPClient: client,
	}
}

/* GetWebsiteDataByGraphNameParams contains all the parameters to send to the API endpoint
   for the get website data by graph name operation.

   Typically these are written to a http.Request.
*/
type GetWebsiteDataByGraphNameParams struct {

	// UserAgent.
	//
	// Default: "Logicmonitor/SDK: Argus Dist-v2.0.0-argus5-4-gd30bcfd-dirty"
	UserAgent *string

	// End.
	//
	// Format: int64
	End *int64

	// Format.
	Format *string

	// GraphName.
	GraphName string

	// ID.
	//
	// Format: int32
	ID int32

	// Start.
	//
	// Format: int64
	Start *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get website data by graph name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetWebsiteDataByGraphNameParams) WithDefaults() *GetWebsiteDataByGraphNameParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get website data by graph name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetWebsiteDataByGraphNameParams) SetDefaults() {
	var (
		userAgentDefault = string("Logicmonitor/SDK: Argus Dist-v2.0.0-argus5-4-gd30bcfd-dirty")

		endDefault = int64(0)

		startDefault = int64(0)
	)

	val := GetWebsiteDataByGraphNameParams{
		UserAgent: &userAgentDefault,
		End:       &endDefault,
		Start:     &startDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get website data by graph name params
func (o *GetWebsiteDataByGraphNameParams) WithTimeout(timeout time.Duration) *GetWebsiteDataByGraphNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get website data by graph name params
func (o *GetWebsiteDataByGraphNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get website data by graph name params
func (o *GetWebsiteDataByGraphNameParams) WithContext(ctx context.Context) *GetWebsiteDataByGraphNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get website data by graph name params
func (o *GetWebsiteDataByGraphNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get website data by graph name params
func (o *GetWebsiteDataByGraphNameParams) WithHTTPClient(client *http.Client) *GetWebsiteDataByGraphNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get website data by graph name params
func (o *GetWebsiteDataByGraphNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUserAgent adds the userAgent to the get website data by graph name params
func (o *GetWebsiteDataByGraphNameParams) WithUserAgent(userAgent *string) *GetWebsiteDataByGraphNameParams {
	o.SetUserAgent(userAgent)
	return o
}

// SetUserAgent adds the userAgent to the get website data by graph name params
func (o *GetWebsiteDataByGraphNameParams) SetUserAgent(userAgent *string) {
	o.UserAgent = userAgent
}

// WithEnd adds the end to the get website data by graph name params
func (o *GetWebsiteDataByGraphNameParams) WithEnd(end *int64) *GetWebsiteDataByGraphNameParams {
	o.SetEnd(end)
	return o
}

// SetEnd adds the end to the get website data by graph name params
func (o *GetWebsiteDataByGraphNameParams) SetEnd(end *int64) {
	o.End = end
}

// WithFormat adds the format to the get website data by graph name params
func (o *GetWebsiteDataByGraphNameParams) WithFormat(format *string) *GetWebsiteDataByGraphNameParams {
	o.SetFormat(format)
	return o
}

// SetFormat adds the format to the get website data by graph name params
func (o *GetWebsiteDataByGraphNameParams) SetFormat(format *string) {
	o.Format = format
}

// WithGraphName adds the graphName to the get website data by graph name params
func (o *GetWebsiteDataByGraphNameParams) WithGraphName(graphName string) *GetWebsiteDataByGraphNameParams {
	o.SetGraphName(graphName)
	return o
}

// SetGraphName adds the graphName to the get website data by graph name params
func (o *GetWebsiteDataByGraphNameParams) SetGraphName(graphName string) {
	o.GraphName = graphName
}

// WithID adds the id to the get website data by graph name params
func (o *GetWebsiteDataByGraphNameParams) WithID(id int32) *GetWebsiteDataByGraphNameParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get website data by graph name params
func (o *GetWebsiteDataByGraphNameParams) SetID(id int32) {
	o.ID = id
}

// WithStart adds the start to the get website data by graph name params
func (o *GetWebsiteDataByGraphNameParams) WithStart(start *int64) *GetWebsiteDataByGraphNameParams {
	o.SetStart(start)
	return o
}

// SetStart adds the start to the get website data by graph name params
func (o *GetWebsiteDataByGraphNameParams) SetStart(start *int64) {
	o.Start = start
}

// WriteToRequest writes these params to a swagger request
func (o *GetWebsiteDataByGraphNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.End != nil {

		// query param end
		var qrEnd int64

		if o.End != nil {
			qrEnd = *o.End
		}
		qEnd := swag.FormatInt64(qrEnd)
		if qEnd != "" {

			if err := r.SetQueryParam("end", qEnd); err != nil {
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

	// path param graphName
	if err := r.SetPathParam("graphName", o.GraphName); err != nil {
		return err
	}

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt32(o.ID)); err != nil {
		return err
	}

	if o.Start != nil {

		// query param start
		var qrStart int64

		if o.Start != nil {
			qrStart = *o.Start
		}
		qStart := swag.FormatInt64(qrStart)
		if qStart != "" {

			if err := r.SetQueryParam("start", qStart); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
