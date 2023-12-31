// Code generated by go-swagger; DO NOT EDIT.

package proxy

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

	"github.com/Ahmed-Sermani/driver-tracking-system/proxy/proxyclient/models"
)

// NewRPCParams creates a new RPCParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRPCParams() *RPCParams {
	return &RPCParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRPCParamsWithTimeout creates a new RPCParams object
// with the ability to set a timeout on a request.
func NewRPCParamsWithTimeout(timeout time.Duration) *RPCParams {
	return &RPCParams{
		timeout: timeout,
	}
}

// NewRPCParamsWithContext creates a new RPCParams object
// with the ability to set a context for a request.
func NewRPCParamsWithContext(ctx context.Context) *RPCParams {
	return &RPCParams{
		Context: ctx,
	}
}

// NewRPCParamsWithHTTPClient creates a new RPCParams object
// with the ability to set a custom HTTPClient for a request.
func NewRPCParamsWithHTTPClient(client *http.Client) *RPCParams {
	return &RPCParams{
		HTTPClient: client,
	}
}

/*
RPCParams contains all the parameters to send to the API endpoint

	for the rpc operation.

	Typically these are written to a http.Request.
*/
type RPCParams struct {

	// Body.
	Body *models.CentrifugoProxyRPCRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the rpc params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RPCParams) WithDefaults() *RPCParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the rpc params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RPCParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the rpc params
func (o *RPCParams) WithTimeout(timeout time.Duration) *RPCParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the rpc params
func (o *RPCParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the rpc params
func (o *RPCParams) WithContext(ctx context.Context) *RPCParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the rpc params
func (o *RPCParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the rpc params
func (o *RPCParams) WithHTTPClient(client *http.Client) *RPCParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the rpc params
func (o *RPCParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the rpc params
func (o *RPCParams) WithBody(body *models.CentrifugoProxyRPCRequest) *RPCParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the rpc params
func (o *RPCParams) SetBody(body *models.CentrifugoProxyRPCRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *RPCParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
