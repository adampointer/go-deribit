// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetPrivateGetOrderMarginByIdsParams creates a new GetPrivateGetOrderMarginByIdsParams object
// with the default values initialized.
func NewGetPrivateGetOrderMarginByIdsParams() *GetPrivateGetOrderMarginByIdsParams {
	var ()
	return &GetPrivateGetOrderMarginByIdsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPrivateGetOrderMarginByIdsParamsWithTimeout creates a new GetPrivateGetOrderMarginByIdsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPrivateGetOrderMarginByIdsParamsWithTimeout(timeout time.Duration) *GetPrivateGetOrderMarginByIdsParams {
	var ()
	return &GetPrivateGetOrderMarginByIdsParams{

		timeout: timeout,
	}
}

// NewGetPrivateGetOrderMarginByIdsParamsWithContext creates a new GetPrivateGetOrderMarginByIdsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPrivateGetOrderMarginByIdsParamsWithContext(ctx context.Context) *GetPrivateGetOrderMarginByIdsParams {
	var ()
	return &GetPrivateGetOrderMarginByIdsParams{

		Context: ctx,
	}
}

// NewGetPrivateGetOrderMarginByIdsParamsWithHTTPClient creates a new GetPrivateGetOrderMarginByIdsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPrivateGetOrderMarginByIdsParamsWithHTTPClient(client *http.Client) *GetPrivateGetOrderMarginByIdsParams {
	var ()
	return &GetPrivateGetOrderMarginByIdsParams{
		HTTPClient: client,
	}
}

/*GetPrivateGetOrderMarginByIdsParams contains all the parameters to send to the API endpoint
for the get private get order margin by ids operation typically these are written to a http.Request
*/
type GetPrivateGetOrderMarginByIdsParams struct {

	/*Ids
	  Ids of orders

	*/
	Ids []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get private get order margin by ids params
func (o *GetPrivateGetOrderMarginByIdsParams) WithTimeout(timeout time.Duration) *GetPrivateGetOrderMarginByIdsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get private get order margin by ids params
func (o *GetPrivateGetOrderMarginByIdsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get private get order margin by ids params
func (o *GetPrivateGetOrderMarginByIdsParams) WithContext(ctx context.Context) *GetPrivateGetOrderMarginByIdsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get private get order margin by ids params
func (o *GetPrivateGetOrderMarginByIdsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get private get order margin by ids params
func (o *GetPrivateGetOrderMarginByIdsParams) WithHTTPClient(client *http.Client) *GetPrivateGetOrderMarginByIdsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get private get order margin by ids params
func (o *GetPrivateGetOrderMarginByIdsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIds adds the ids to the get private get order margin by ids params
func (o *GetPrivateGetOrderMarginByIdsParams) WithIds(ids []string) *GetPrivateGetOrderMarginByIdsParams {
	o.SetIds(ids)
	return o
}

// SetIds adds the ids to the get private get order margin by ids params
func (o *GetPrivateGetOrderMarginByIdsParams) SetIds(ids []string) {
	o.Ids = ids
}

// WriteToRequest writes these params to a swagger request
func (o *GetPrivateGetOrderMarginByIdsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	valuesIds := o.Ids

	joinedIds := swag.JoinByFormat(valuesIds, "multi")
	// query array param ids
	if err := r.SetQueryParam("ids", joinedIds...); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
