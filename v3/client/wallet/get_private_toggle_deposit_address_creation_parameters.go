// Code generated by go-swagger; DO NOT EDIT.

package wallet

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

// NewGetPrivateToggleDepositAddressCreationParams creates a new GetPrivateToggleDepositAddressCreationParams object
// with the default values initialized.
func NewGetPrivateToggleDepositAddressCreationParams() *GetPrivateToggleDepositAddressCreationParams {
	var ()
	return &GetPrivateToggleDepositAddressCreationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPrivateToggleDepositAddressCreationParamsWithTimeout creates a new GetPrivateToggleDepositAddressCreationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPrivateToggleDepositAddressCreationParamsWithTimeout(timeout time.Duration) *GetPrivateToggleDepositAddressCreationParams {
	var ()
	return &GetPrivateToggleDepositAddressCreationParams{

		timeout: timeout,
	}
}

// NewGetPrivateToggleDepositAddressCreationParamsWithContext creates a new GetPrivateToggleDepositAddressCreationParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPrivateToggleDepositAddressCreationParamsWithContext(ctx context.Context) *GetPrivateToggleDepositAddressCreationParams {
	var ()
	return &GetPrivateToggleDepositAddressCreationParams{

		Context: ctx,
	}
}

// NewGetPrivateToggleDepositAddressCreationParamsWithHTTPClient creates a new GetPrivateToggleDepositAddressCreationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPrivateToggleDepositAddressCreationParamsWithHTTPClient(client *http.Client) *GetPrivateToggleDepositAddressCreationParams {
	var ()
	return &GetPrivateToggleDepositAddressCreationParams{
		HTTPClient: client,
	}
}

/*GetPrivateToggleDepositAddressCreationParams contains all the parameters to send to the API endpoint
for the get private toggle deposit address creation operation typically these are written to a http.Request
*/
type GetPrivateToggleDepositAddressCreationParams struct {

	/*Currency
	  The currency symbol

	*/
	Currency string
	/*State*/
	State bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get private toggle deposit address creation params
func (o *GetPrivateToggleDepositAddressCreationParams) WithTimeout(timeout time.Duration) *GetPrivateToggleDepositAddressCreationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get private toggle deposit address creation params
func (o *GetPrivateToggleDepositAddressCreationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get private toggle deposit address creation params
func (o *GetPrivateToggleDepositAddressCreationParams) WithContext(ctx context.Context) *GetPrivateToggleDepositAddressCreationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get private toggle deposit address creation params
func (o *GetPrivateToggleDepositAddressCreationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get private toggle deposit address creation params
func (o *GetPrivateToggleDepositAddressCreationParams) WithHTTPClient(client *http.Client) *GetPrivateToggleDepositAddressCreationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get private toggle deposit address creation params
func (o *GetPrivateToggleDepositAddressCreationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCurrency adds the currency to the get private toggle deposit address creation params
func (o *GetPrivateToggleDepositAddressCreationParams) WithCurrency(currency string) *GetPrivateToggleDepositAddressCreationParams {
	o.SetCurrency(currency)
	return o
}

// SetCurrency adds the currency to the get private toggle deposit address creation params
func (o *GetPrivateToggleDepositAddressCreationParams) SetCurrency(currency string) {
	o.Currency = currency
}

// WithState adds the state to the get private toggle deposit address creation params
func (o *GetPrivateToggleDepositAddressCreationParams) WithState(state bool) *GetPrivateToggleDepositAddressCreationParams {
	o.SetState(state)
	return o
}

// SetState adds the state to the get private toggle deposit address creation params
func (o *GetPrivateToggleDepositAddressCreationParams) SetState(state bool) {
	o.State = state
}

// WriteToRequest writes these params to a swagger request
func (o *GetPrivateToggleDepositAddressCreationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// query param state
	qrState := o.State
	qState := swag.FormatBool(qrState)
	if qState != "" {
		if err := r.SetQueryParam("state", qState); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
