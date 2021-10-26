// Code generated by go-swagger; DO NOT EDIT.

package gc

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

// NewGetGCLogParams creates a new GetGCLogParams object
// with the default values initialized.
func NewGetGCLogParams() *GetGCLogParams {
	var ()
	return &GetGCLogParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetGCLogParamsWithTimeout creates a new GetGCLogParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetGCLogParamsWithTimeout(timeout time.Duration) *GetGCLogParams {
	var ()
	return &GetGCLogParams{

		timeout: timeout,
	}
}

// NewGetGCLogParamsWithContext creates a new GetGCLogParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetGCLogParamsWithContext(ctx context.Context) *GetGCLogParams {
	var ()
	return &GetGCLogParams{

		Context: ctx,
	}
}

// NewGetGCLogParamsWithHTTPClient creates a new GetGCLogParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetGCLogParamsWithHTTPClient(client *http.Client) *GetGCLogParams {
	var ()
	return &GetGCLogParams{
		HTTPClient: client,
	}
}

/*GetGCLogParams contains all the parameters to send to the API endpoint
for the get g c log operation typically these are written to a http.Request
*/
type GetGCLogParams struct {

	/*XRequestID
	  An unique ID for the request

	*/
	XRequestID *string
	/*GcID
	  The ID of the gc log

	*/
	GcID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get g c log params
func (o *GetGCLogParams) WithTimeout(timeout time.Duration) *GetGCLogParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get g c log params
func (o *GetGCLogParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get g c log params
func (o *GetGCLogParams) WithContext(ctx context.Context) *GetGCLogParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get g c log params
func (o *GetGCLogParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get g c log params
func (o *GetGCLogParams) WithHTTPClient(client *http.Client) *GetGCLogParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get g c log params
func (o *GetGCLogParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the get g c log params
func (o *GetGCLogParams) WithXRequestID(xRequestID *string) *GetGCLogParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the get g c log params
func (o *GetGCLogParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithGcID adds the gcID to the get g c log params
func (o *GetGCLogParams) WithGcID(gcID int64) *GetGCLogParams {
	o.SetGcID(gcID)
	return o
}

// SetGcID adds the gcId to the get g c log params
func (o *GetGCLogParams) SetGcID(gcID int64) {
	o.GcID = gcID
}

// WriteToRequest writes these params to a swagger request
func (o *GetGCLogParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.XRequestID != nil {

		// header param X-Request-Id
		if err := r.SetHeaderParam("X-Request-Id", *o.XRequestID); err != nil {
			return err
		}

	}

	// path param gc_id
	if err := r.SetPathParam("gc_id", swag.FormatInt64(o.GcID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
