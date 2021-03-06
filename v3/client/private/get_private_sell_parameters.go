// Code generated by go-swagger; DO NOT EDIT.

package private

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

// NewGetPrivateSellParams creates a new GetPrivateSellParams object
// with the default values initialized.
func NewGetPrivateSellParams() *GetPrivateSellParams {
	var (
		maxShowDefault     = float64(1)
		postOnlyDefault    = bool(true)
		reduceOnlyDefault  = bool(false)
		timeInForceDefault = string("good_til_cancelled")
	)
	return &GetPrivateSellParams{
		MaxShow:     &maxShowDefault,
		PostOnly:    &postOnlyDefault,
		ReduceOnly:  &reduceOnlyDefault,
		TimeInForce: &timeInForceDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPrivateSellParamsWithTimeout creates a new GetPrivateSellParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPrivateSellParamsWithTimeout(timeout time.Duration) *GetPrivateSellParams {
	var (
		maxShowDefault     = float64(1)
		postOnlyDefault    = bool(true)
		reduceOnlyDefault  = bool(false)
		timeInForceDefault = string("good_til_cancelled")
	)
	return &GetPrivateSellParams{
		MaxShow:     &maxShowDefault,
		PostOnly:    &postOnlyDefault,
		ReduceOnly:  &reduceOnlyDefault,
		TimeInForce: &timeInForceDefault,

		timeout: timeout,
	}
}

// NewGetPrivateSellParamsWithContext creates a new GetPrivateSellParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPrivateSellParamsWithContext(ctx context.Context) *GetPrivateSellParams {
	var (
		maxShowDefault     = float64(1)
		postOnlyDefault    = bool(true)
		reduceOnlyDefault  = bool(false)
		timeInForceDefault = string("good_til_cancelled")
	)
	return &GetPrivateSellParams{
		MaxShow:     &maxShowDefault,
		PostOnly:    &postOnlyDefault,
		ReduceOnly:  &reduceOnlyDefault,
		TimeInForce: &timeInForceDefault,

		Context: ctx,
	}
}

// NewGetPrivateSellParamsWithHTTPClient creates a new GetPrivateSellParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPrivateSellParamsWithHTTPClient(client *http.Client) *GetPrivateSellParams {
	var (
		maxShowDefault     = float64(1)
		postOnlyDefault    = bool(true)
		reduceOnlyDefault  = bool(false)
		timeInForceDefault = string("good_til_cancelled")
	)
	return &GetPrivateSellParams{
		MaxShow:     &maxShowDefault,
		PostOnly:    &postOnlyDefault,
		ReduceOnly:  &reduceOnlyDefault,
		TimeInForce: &timeInForceDefault,
		HTTPClient:  client,
	}
}

/*GetPrivateSellParams contains all the parameters to send to the API endpoint
for the get private sell operation typically these are written to a http.Request
*/
type GetPrivateSellParams struct {

	/*Advanced
	  Advanced option order type. (Only for options)

	*/
	Advanced *string
	/*Amount
	  It represents the requested order size. For perpetual and futures the amount is in USD units, for options it is amount of corresponding cryptocurrency contracts, e.g., BTC or ETH

	*/
	Amount float64
	/*InstrumentName
	  Instrument name

	*/
	InstrumentName string
	/*Label
	  user defined label for the order (maximum 32 characters)

	*/
	Label *string
	/*MaxShow
	  Maximum amount within an order to be shown to other customers, `0` for invisible order

	*/
	MaxShow *float64
	/*PostOnly
	  <p>If true, the order is considered post-only. If the new price would cause the order to be filled immediately (as taker), the price will be changed to be just below the bid.</p> <p>Only valid in combination with time_in_force=`"good_til_cancelled"`</p>

	*/
	PostOnly *bool
	/*Price
	  <p>The order price in base currency (Only for limit and stop_limit orders)</p> <p>When adding order with advanced=usd, the field price should be the option price value in USD.</p> <p>When adding order with advanced=implv, the field price should be a value of implied volatility in percentages. For example,  price=100, means implied volatility of 100%</p>

	*/
	Price *float64
	/*ReduceOnly
	  If `true`, the order is considered reduce-only which is intended to only reduce a current position

	*/
	ReduceOnly *bool
	/*StopPrice
	  Stop price, required for stop limit orders (Only for stop orders)

	*/
	StopPrice *float64
	/*TimeInForce
	  <p>Specifies how long the order remains in effect. Default `"good_til_cancelled"`</p> <ul> <li>`"good_til_cancelled"` - unfilled order remains in order book until cancelled</li> <li>`"fill_or_kill"` - execute a transaction immediately and completely or not at all</li> <li>`"immediate_or_cancel"` - execute a transaction immediately, and any portion of the order that cannot be immediately filled is cancelled</li> </ul>

	*/
	TimeInForce *string
	/*Trigger
	  Defines trigger type, required for `"stop_limit"` order type

	*/
	Trigger *string
	/*Type
	  The order type, default: `"limit"`

	*/
	Type *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get private sell params
func (o *GetPrivateSellParams) WithTimeout(timeout time.Duration) *GetPrivateSellParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get private sell params
func (o *GetPrivateSellParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get private sell params
func (o *GetPrivateSellParams) WithContext(ctx context.Context) *GetPrivateSellParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get private sell params
func (o *GetPrivateSellParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get private sell params
func (o *GetPrivateSellParams) WithHTTPClient(client *http.Client) *GetPrivateSellParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get private sell params
func (o *GetPrivateSellParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAdvanced adds the advanced to the get private sell params
func (o *GetPrivateSellParams) WithAdvanced(advanced *string) *GetPrivateSellParams {
	o.SetAdvanced(advanced)
	return o
}

// SetAdvanced adds the advanced to the get private sell params
func (o *GetPrivateSellParams) SetAdvanced(advanced *string) {
	o.Advanced = advanced
}

// WithAmount adds the amount to the get private sell params
func (o *GetPrivateSellParams) WithAmount(amount float64) *GetPrivateSellParams {
	o.SetAmount(amount)
	return o
}

// SetAmount adds the amount to the get private sell params
func (o *GetPrivateSellParams) SetAmount(amount float64) {
	o.Amount = amount
}

// WithInstrumentName adds the instrumentName to the get private sell params
func (o *GetPrivateSellParams) WithInstrumentName(instrumentName string) *GetPrivateSellParams {
	o.SetInstrumentName(instrumentName)
	return o
}

// SetInstrumentName adds the instrumentName to the get private sell params
func (o *GetPrivateSellParams) SetInstrumentName(instrumentName string) {
	o.InstrumentName = instrumentName
}

// WithLabel adds the label to the get private sell params
func (o *GetPrivateSellParams) WithLabel(label *string) *GetPrivateSellParams {
	o.SetLabel(label)
	return o
}

// SetLabel adds the label to the get private sell params
func (o *GetPrivateSellParams) SetLabel(label *string) {
	o.Label = label
}

// WithMaxShow adds the maxShow to the get private sell params
func (o *GetPrivateSellParams) WithMaxShow(maxShow *float64) *GetPrivateSellParams {
	o.SetMaxShow(maxShow)
	return o
}

// SetMaxShow adds the maxShow to the get private sell params
func (o *GetPrivateSellParams) SetMaxShow(maxShow *float64) {
	o.MaxShow = maxShow
}

// WithPostOnly adds the postOnly to the get private sell params
func (o *GetPrivateSellParams) WithPostOnly(postOnly *bool) *GetPrivateSellParams {
	o.SetPostOnly(postOnly)
	return o
}

// SetPostOnly adds the postOnly to the get private sell params
func (o *GetPrivateSellParams) SetPostOnly(postOnly *bool) {
	o.PostOnly = postOnly
}

// WithPrice adds the price to the get private sell params
func (o *GetPrivateSellParams) WithPrice(price *float64) *GetPrivateSellParams {
	o.SetPrice(price)
	return o
}

// SetPrice adds the price to the get private sell params
func (o *GetPrivateSellParams) SetPrice(price *float64) {
	o.Price = price
}

// WithReduceOnly adds the reduceOnly to the get private sell params
func (o *GetPrivateSellParams) WithReduceOnly(reduceOnly *bool) *GetPrivateSellParams {
	o.SetReduceOnly(reduceOnly)
	return o
}

// SetReduceOnly adds the reduceOnly to the get private sell params
func (o *GetPrivateSellParams) SetReduceOnly(reduceOnly *bool) {
	o.ReduceOnly = reduceOnly
}

// WithStopPrice adds the stopPrice to the get private sell params
func (o *GetPrivateSellParams) WithStopPrice(stopPrice *float64) *GetPrivateSellParams {
	o.SetStopPrice(stopPrice)
	return o
}

// SetStopPrice adds the stopPrice to the get private sell params
func (o *GetPrivateSellParams) SetStopPrice(stopPrice *float64) {
	o.StopPrice = stopPrice
}

// WithTimeInForce adds the timeInForce to the get private sell params
func (o *GetPrivateSellParams) WithTimeInForce(timeInForce *string) *GetPrivateSellParams {
	o.SetTimeInForce(timeInForce)
	return o
}

// SetTimeInForce adds the timeInForce to the get private sell params
func (o *GetPrivateSellParams) SetTimeInForce(timeInForce *string) {
	o.TimeInForce = timeInForce
}

// WithTrigger adds the trigger to the get private sell params
func (o *GetPrivateSellParams) WithTrigger(trigger *string) *GetPrivateSellParams {
	o.SetTrigger(trigger)
	return o
}

// SetTrigger adds the trigger to the get private sell params
func (o *GetPrivateSellParams) SetTrigger(trigger *string) {
	o.Trigger = trigger
}

// WithType adds the typeVar to the get private sell params
func (o *GetPrivateSellParams) WithType(typeVar *string) *GetPrivateSellParams {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the get private sell params
func (o *GetPrivateSellParams) SetType(typeVar *string) {
	o.Type = typeVar
}

// WriteToRequest writes these params to a swagger request
func (o *GetPrivateSellParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Advanced != nil {

		// query param advanced
		var qrAdvanced string
		if o.Advanced != nil {
			qrAdvanced = *o.Advanced
		}
		qAdvanced := qrAdvanced
		if qAdvanced != "" {
			if err := r.SetQueryParam("advanced", qAdvanced); err != nil {
				return err
			}
		}

	}

	// query param amount
	qrAmount := o.Amount
	qAmount := swag.FormatFloat64(qrAmount)
	if qAmount != "" {
		if err := r.SetQueryParam("amount", qAmount); err != nil {
			return err
		}
	}

	// query param instrument_name
	qrInstrumentName := o.InstrumentName
	qInstrumentName := qrInstrumentName
	if qInstrumentName != "" {
		if err := r.SetQueryParam("instrument_name", qInstrumentName); err != nil {
			return err
		}
	}

	if o.Label != nil {

		// query param label
		var qrLabel string
		if o.Label != nil {
			qrLabel = *o.Label
		}
		qLabel := qrLabel
		if qLabel != "" {
			if err := r.SetQueryParam("label", qLabel); err != nil {
				return err
			}
		}

	}

	if o.MaxShow != nil {

		// query param max_show
		var qrMaxShow float64
		if o.MaxShow != nil {
			qrMaxShow = *o.MaxShow
		}
		qMaxShow := swag.FormatFloat64(qrMaxShow)
		if qMaxShow != "" {
			if err := r.SetQueryParam("max_show", qMaxShow); err != nil {
				return err
			}
		}

	}

	if o.PostOnly != nil {

		// query param post_only
		var qrPostOnly bool
		if o.PostOnly != nil {
			qrPostOnly = *o.PostOnly
		}
		qPostOnly := swag.FormatBool(qrPostOnly)
		if qPostOnly != "" {
			if err := r.SetQueryParam("post_only", qPostOnly); err != nil {
				return err
			}
		}

	}

	if o.Price != nil {

		// query param price
		var qrPrice float64
		if o.Price != nil {
			qrPrice = *o.Price
		}
		qPrice := swag.FormatFloat64(qrPrice)
		if qPrice != "" {
			if err := r.SetQueryParam("price", qPrice); err != nil {
				return err
			}
		}

	}

	if o.ReduceOnly != nil {

		// query param reduce_only
		var qrReduceOnly bool
		if o.ReduceOnly != nil {
			qrReduceOnly = *o.ReduceOnly
		}
		qReduceOnly := swag.FormatBool(qrReduceOnly)
		if qReduceOnly != "" {
			if err := r.SetQueryParam("reduce_only", qReduceOnly); err != nil {
				return err
			}
		}

	}

	if o.StopPrice != nil {

		// query param stop_price
		var qrStopPrice float64
		if o.StopPrice != nil {
			qrStopPrice = *o.StopPrice
		}
		qStopPrice := swag.FormatFloat64(qrStopPrice)
		if qStopPrice != "" {
			if err := r.SetQueryParam("stop_price", qStopPrice); err != nil {
				return err
			}
		}

	}

	if o.TimeInForce != nil {

		// query param time_in_force
		var qrTimeInForce string
		if o.TimeInForce != nil {
			qrTimeInForce = *o.TimeInForce
		}
		qTimeInForce := qrTimeInForce
		if qTimeInForce != "" {
			if err := r.SetQueryParam("time_in_force", qTimeInForce); err != nil {
				return err
			}
		}

	}

	if o.Trigger != nil {

		// query param trigger
		var qrTrigger string
		if o.Trigger != nil {
			qrTrigger = *o.Trigger
		}
		qTrigger := qrTrigger
		if qTrigger != "" {
			if err := r.SetQueryParam("trigger", qTrigger); err != nil {
				return err
			}
		}

	}

	if o.Type != nil {

		// query param type
		var qrType string
		if o.Type != nil {
			qrType = *o.Type
		}
		qType := qrType
		if qType != "" {
			if err := r.SetQueryParam("type", qType); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
