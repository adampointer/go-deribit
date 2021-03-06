// Code generated by go-swagger; DO NOT EDIT.

package trading

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

// NewGetPrivateGetUserTradesByInstrumentParams creates a new GetPrivateGetUserTradesByInstrumentParams object
// with the default values initialized.
func NewGetPrivateGetUserTradesByInstrumentParams() *GetPrivateGetUserTradesByInstrumentParams {
	var ()
	return &GetPrivateGetUserTradesByInstrumentParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPrivateGetUserTradesByInstrumentParamsWithTimeout creates a new GetPrivateGetUserTradesByInstrumentParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPrivateGetUserTradesByInstrumentParamsWithTimeout(timeout time.Duration) *GetPrivateGetUserTradesByInstrumentParams {
	var ()
	return &GetPrivateGetUserTradesByInstrumentParams{

		timeout: timeout,
	}
}

// NewGetPrivateGetUserTradesByInstrumentParamsWithContext creates a new GetPrivateGetUserTradesByInstrumentParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPrivateGetUserTradesByInstrumentParamsWithContext(ctx context.Context) *GetPrivateGetUserTradesByInstrumentParams {
	var ()
	return &GetPrivateGetUserTradesByInstrumentParams{

		Context: ctx,
	}
}

// NewGetPrivateGetUserTradesByInstrumentParamsWithHTTPClient creates a new GetPrivateGetUserTradesByInstrumentParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPrivateGetUserTradesByInstrumentParamsWithHTTPClient(client *http.Client) *GetPrivateGetUserTradesByInstrumentParams {
	var ()
	return &GetPrivateGetUserTradesByInstrumentParams{
		HTTPClient: client,
	}
}

/*GetPrivateGetUserTradesByInstrumentParams contains all the parameters to send to the API endpoint
for the get private get user trades by instrument operation typically these are written to a http.Request
*/
type GetPrivateGetUserTradesByInstrumentParams struct {

	/*Count
	  Number of requested items, default - `10`

	*/
	Count *int64
	/*EndSeq
	  The sequence number of the last trade to be returned

	*/
	EndSeq *int64
	/*IncludeOld
	  Include trades older than 7 days, default - `false`

	*/
	IncludeOld *bool
	/*InstrumentName
	  Instrument name

	*/
	InstrumentName string
	/*Sorting
	  Direction of results sorting (`default` value means no sorting, results will be returned in order in which they left the database)

	*/
	Sorting *string
	/*StartSeq
	  The sequence number of the first trade to be returned

	*/
	StartSeq *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get private get user trades by instrument params
func (o *GetPrivateGetUserTradesByInstrumentParams) WithTimeout(timeout time.Duration) *GetPrivateGetUserTradesByInstrumentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get private get user trades by instrument params
func (o *GetPrivateGetUserTradesByInstrumentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get private get user trades by instrument params
func (o *GetPrivateGetUserTradesByInstrumentParams) WithContext(ctx context.Context) *GetPrivateGetUserTradesByInstrumentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get private get user trades by instrument params
func (o *GetPrivateGetUserTradesByInstrumentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get private get user trades by instrument params
func (o *GetPrivateGetUserTradesByInstrumentParams) WithHTTPClient(client *http.Client) *GetPrivateGetUserTradesByInstrumentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get private get user trades by instrument params
func (o *GetPrivateGetUserTradesByInstrumentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCount adds the count to the get private get user trades by instrument params
func (o *GetPrivateGetUserTradesByInstrumentParams) WithCount(count *int64) *GetPrivateGetUserTradesByInstrumentParams {
	o.SetCount(count)
	return o
}

// SetCount adds the count to the get private get user trades by instrument params
func (o *GetPrivateGetUserTradesByInstrumentParams) SetCount(count *int64) {
	o.Count = count
}

// WithEndSeq adds the endSeq to the get private get user trades by instrument params
func (o *GetPrivateGetUserTradesByInstrumentParams) WithEndSeq(endSeq *int64) *GetPrivateGetUserTradesByInstrumentParams {
	o.SetEndSeq(endSeq)
	return o
}

// SetEndSeq adds the endSeq to the get private get user trades by instrument params
func (o *GetPrivateGetUserTradesByInstrumentParams) SetEndSeq(endSeq *int64) {
	o.EndSeq = endSeq
}

// WithIncludeOld adds the includeOld to the get private get user trades by instrument params
func (o *GetPrivateGetUserTradesByInstrumentParams) WithIncludeOld(includeOld *bool) *GetPrivateGetUserTradesByInstrumentParams {
	o.SetIncludeOld(includeOld)
	return o
}

// SetIncludeOld adds the includeOld to the get private get user trades by instrument params
func (o *GetPrivateGetUserTradesByInstrumentParams) SetIncludeOld(includeOld *bool) {
	o.IncludeOld = includeOld
}

// WithInstrumentName adds the instrumentName to the get private get user trades by instrument params
func (o *GetPrivateGetUserTradesByInstrumentParams) WithInstrumentName(instrumentName string) *GetPrivateGetUserTradesByInstrumentParams {
	o.SetInstrumentName(instrumentName)
	return o
}

// SetInstrumentName adds the instrumentName to the get private get user trades by instrument params
func (o *GetPrivateGetUserTradesByInstrumentParams) SetInstrumentName(instrumentName string) {
	o.InstrumentName = instrumentName
}

// WithSorting adds the sorting to the get private get user trades by instrument params
func (o *GetPrivateGetUserTradesByInstrumentParams) WithSorting(sorting *string) *GetPrivateGetUserTradesByInstrumentParams {
	o.SetSorting(sorting)
	return o
}

// SetSorting adds the sorting to the get private get user trades by instrument params
func (o *GetPrivateGetUserTradesByInstrumentParams) SetSorting(sorting *string) {
	o.Sorting = sorting
}

// WithStartSeq adds the startSeq to the get private get user trades by instrument params
func (o *GetPrivateGetUserTradesByInstrumentParams) WithStartSeq(startSeq *int64) *GetPrivateGetUserTradesByInstrumentParams {
	o.SetStartSeq(startSeq)
	return o
}

// SetStartSeq adds the startSeq to the get private get user trades by instrument params
func (o *GetPrivateGetUserTradesByInstrumentParams) SetStartSeq(startSeq *int64) {
	o.StartSeq = startSeq
}

// WriteToRequest writes these params to a swagger request
func (o *GetPrivateGetUserTradesByInstrumentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Count != nil {

		// query param count
		var qrCount int64
		if o.Count != nil {
			qrCount = *o.Count
		}
		qCount := swag.FormatInt64(qrCount)
		if qCount != "" {
			if err := r.SetQueryParam("count", qCount); err != nil {
				return err
			}
		}

	}

	if o.EndSeq != nil {

		// query param end_seq
		var qrEndSeq int64
		if o.EndSeq != nil {
			qrEndSeq = *o.EndSeq
		}
		qEndSeq := swag.FormatInt64(qrEndSeq)
		if qEndSeq != "" {
			if err := r.SetQueryParam("end_seq", qEndSeq); err != nil {
				return err
			}
		}

	}

	if o.IncludeOld != nil {

		// query param include_old
		var qrIncludeOld bool
		if o.IncludeOld != nil {
			qrIncludeOld = *o.IncludeOld
		}
		qIncludeOld := swag.FormatBool(qrIncludeOld)
		if qIncludeOld != "" {
			if err := r.SetQueryParam("include_old", qIncludeOld); err != nil {
				return err
			}
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

	if o.Sorting != nil {

		// query param sorting
		var qrSorting string
		if o.Sorting != nil {
			qrSorting = *o.Sorting
		}
		qSorting := qrSorting
		if qSorting != "" {
			if err := r.SetQueryParam("sorting", qSorting); err != nil {
				return err
			}
		}

	}

	if o.StartSeq != nil {

		// query param start_seq
		var qrStartSeq int64
		if o.StartSeq != nil {
			qrStartSeq = *o.StartSeq
		}
		qStartSeq := swag.FormatInt64(qrStartSeq)
		if qStartSeq != "" {
			if err := r.SetQueryParam("start_seq", qStartSeq); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
