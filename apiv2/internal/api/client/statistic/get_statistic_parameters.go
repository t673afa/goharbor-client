// Code generated by go-swagger; DO NOT EDIT.

package statistic

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

// NewGetStatisticParams creates a new GetStatisticParams object
// with the default values initialized.
func NewGetStatisticParams() *GetStatisticParams {
	var ()
	return &GetStatisticParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetStatisticParamsWithTimeout creates a new GetStatisticParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetStatisticParamsWithTimeout(timeout time.Duration) *GetStatisticParams {
	var ()
	return &GetStatisticParams{

		timeout: timeout,
	}
}

// NewGetStatisticParamsWithContext creates a new GetStatisticParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetStatisticParamsWithContext(ctx context.Context) *GetStatisticParams {
	var ()
	return &GetStatisticParams{

		Context: ctx,
	}
}

// NewGetStatisticParamsWithHTTPClient creates a new GetStatisticParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetStatisticParamsWithHTTPClient(client *http.Client) *GetStatisticParams {
	var ()
	return &GetStatisticParams{
		HTTPClient: client,
	}
}

/*GetStatisticParams contains all the parameters to send to the API endpoint
for the get statistic operation typically these are written to a http.Request
*/
type GetStatisticParams struct {

	/*XRequestID
	  An unique ID for the request

	*/
	XRequestID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get statistic params
func (o *GetStatisticParams) WithTimeout(timeout time.Duration) *GetStatisticParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get statistic params
func (o *GetStatisticParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get statistic params
func (o *GetStatisticParams) WithContext(ctx context.Context) *GetStatisticParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get statistic params
func (o *GetStatisticParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get statistic params
func (o *GetStatisticParams) WithHTTPClient(client *http.Client) *GetStatisticParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get statistic params
func (o *GetStatisticParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the get statistic params
func (o *GetStatisticParams) WithXRequestID(xRequestID *string) *GetStatisticParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the get statistic params
func (o *GetStatisticParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WriteToRequest writes these params to a swagger request
func (o *GetStatisticParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
