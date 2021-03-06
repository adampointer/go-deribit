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

// NewGetPrivateGetAccountSummaryParams creates a new GetPrivateGetAccountSummaryParams object
// with the default values initialized.
func NewGetPrivateGetAccountSummaryParams() *GetPrivateGetAccountSummaryParams {
	var ()
	return &GetPrivateGetAccountSummaryParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPrivateGetAccountSummaryParamsWithTimeout creates a new GetPrivateGetAccountSummaryParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPrivateGetAccountSummaryParamsWithTimeout(timeout time.Duration) *GetPrivateGetAccountSummaryParams {
	var ()
	return &GetPrivateGetAccountSummaryParams{

		timeout: timeout,
	}
}

// NewGetPrivateGetAccountSummaryParamsWithContext creates a new GetPrivateGetAccountSummaryParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPrivateGetAccountSummaryParamsWithContext(ctx context.Context) *GetPrivateGetAccountSummaryParams {
	var ()
	return &GetPrivateGetAccountSummaryParams{

		Context: ctx,
	}
}

// NewGetPrivateGetAccountSummaryParamsWithHTTPClient creates a new GetPrivateGetAccountSummaryParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPrivateGetAccountSummaryParamsWithHTTPClient(client *http.Client) *GetPrivateGetAccountSummaryParams {
	var ()
	return &GetPrivateGetAccountSummaryParams{
		HTTPClient: client,
	}
}

/*GetPrivateGetAccountSummaryParams contains all the parameters to send to the API endpoint
for the get private get account summary operation typically these are written to a http.Request
*/
type GetPrivateGetAccountSummaryParams struct {

	/*Currency
	  The currency symbol

	*/
	Currency string
	/*Extended
	  Include additional fields

	*/
	Extended *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get private get account summary params
func (o *GetPrivateGetAccountSummaryParams) WithTimeout(timeout time.Duration) *GetPrivateGetAccountSummaryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get private get account summary params
func (o *GetPrivateGetAccountSummaryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get private get account summary params
func (o *GetPrivateGetAccountSummaryParams) WithContext(ctx context.Context) *GetPrivateGetAccountSummaryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get private get account summary params
func (o *GetPrivateGetAccountSummaryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get private get account summary params
func (o *GetPrivateGetAccountSummaryParams) WithHTTPClient(client *http.Client) *GetPrivateGetAccountSummaryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get private get account summary params
func (o *GetPrivateGetAccountSummaryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCurrency adds the currency to the get private get account summary params
func (o *GetPrivateGetAccountSummaryParams) WithCurrency(currency string) *GetPrivateGetAccountSummaryParams {
	o.SetCurrency(currency)
	return o
}

// SetCurrency adds the currency to the get private get account summary params
func (o *GetPrivateGetAccountSummaryParams) SetCurrency(currency string) {
	o.Currency = currency
}

// WithExtended adds the extended to the get private get account summary params
func (o *GetPrivateGetAccountSummaryParams) WithExtended(extended *bool) *GetPrivateGetAccountSummaryParams {
	o.SetExtended(extended)
	return o
}

// SetExtended adds the extended to the get private get account summary params
func (o *GetPrivateGetAccountSummaryParams) SetExtended(extended *bool) {
	o.Extended = extended
}

// WriteToRequest writes these params to a swagger request
func (o *GetPrivateGetAccountSummaryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param currency
	qrCurrency := o.Currency
	qCurrency := qrCurrency
	if qCurrency != "" {
		if err := r.SetQueryParam("currency", qCurrency); err != nil {
			return err
		}
	}

	if o.Extended != nil {

		// query param extended
		var qrExtended bool
		if o.Extended != nil {
			qrExtended = *o.Extended
		}
		qExtended := swag.FormatBool(qrExtended)
		if qExtended != "" {
			if err := r.SetQueryParam("extended", qExtended); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
