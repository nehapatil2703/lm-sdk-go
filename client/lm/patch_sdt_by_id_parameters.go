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
	models "github.com/logicmonitor/lm-sdk-go/models"
	"golang.org/x/net/context"
)

// NewPatchSDTByIDParams creates a new PatchSDTByIDParams object
// with the default values initialized.
func NewPatchSDTByIDParams() *PatchSDTByIDParams {
	var ()
	return &PatchSDTByIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchSDTByIDParamsWithTimeout creates a new PatchSDTByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchSDTByIDParamsWithTimeout(timeout time.Duration) *PatchSDTByIDParams {
	var ()
	return &PatchSDTByIDParams{

		timeout: timeout,
	}
}

// NewPatchSDTByIDParamsWithContext creates a new PatchSDTByIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchSDTByIDParamsWithContext(ctx context.Context) *PatchSDTByIDParams {
	var ()
	return &PatchSDTByIDParams{

		Context: ctx,
	}
}

// NewPatchSDTByIDParamsWithHTTPClient creates a new PatchSDTByIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchSDTByIDParamsWithHTTPClient(client *http.Client) *PatchSDTByIDParams {
	var ()
	return &PatchSDTByIDParams{
		HTTPClient: client,
	}
}

/*PatchSDTByIDParams contains all the parameters to send to the API endpoint
for the patch SDT by Id operation typically these are written to a http.Request
*/
type PatchSDTByIDParams struct {

	/*Body*/
	Body models.SDT
	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch SDT by Id params
func (o *PatchSDTByIDParams) WithTimeout(timeout time.Duration) *PatchSDTByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch SDT by Id params
func (o *PatchSDTByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch SDT by Id params
func (o *PatchSDTByIDParams) WithContext(ctx context.Context) *PatchSDTByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch SDT by Id params
func (o *PatchSDTByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch SDT by Id params
func (o *PatchSDTByIDParams) WithHTTPClient(client *http.Client) *PatchSDTByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch SDT by Id params
func (o *PatchSDTByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch SDT by Id params
func (o *PatchSDTByIDParams) WithBody(body models.SDT) *PatchSDTByIDParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch SDT by Id params
func (o *PatchSDTByIDParams) SetBody(body models.SDT) {
	o.Body = body
}

// WithID adds the id to the patch SDT by Id params
func (o *PatchSDTByIDParams) WithID(id string) *PatchSDTByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the patch SDT by Id params
func (o *PatchSDTByIDParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PatchSDTByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {
	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
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
