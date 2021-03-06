// Code generated by go-swagger; DO NOT EDIT.

package account_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetPublicGetAnnouncementsParams creates a new GetPublicGetAnnouncementsParams object
// with the default values initialized.
func NewGetPublicGetAnnouncementsParams() *GetPublicGetAnnouncementsParams {

	return &GetPublicGetAnnouncementsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPublicGetAnnouncementsParamsWithTimeout creates a new GetPublicGetAnnouncementsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPublicGetAnnouncementsParamsWithTimeout(timeout time.Duration) *GetPublicGetAnnouncementsParams {

	return &GetPublicGetAnnouncementsParams{

		timeout: timeout,
	}
}

// NewGetPublicGetAnnouncementsParamsWithContext creates a new GetPublicGetAnnouncementsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPublicGetAnnouncementsParamsWithContext(ctx context.Context) *GetPublicGetAnnouncementsParams {

	return &GetPublicGetAnnouncementsParams{

		Context: ctx,
	}
}

// NewGetPublicGetAnnouncementsParamsWithHTTPClient creates a new GetPublicGetAnnouncementsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPublicGetAnnouncementsParamsWithHTTPClient(client *http.Client) *GetPublicGetAnnouncementsParams {

	return &GetPublicGetAnnouncementsParams{
		HTTPClient: client,
	}
}

/*GetPublicGetAnnouncementsParams contains all the parameters to send to the API endpoint
for the get public get announcements operation typically these are written to a http.Request
*/
type GetPublicGetAnnouncementsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get public get announcements params
func (o *GetPublicGetAnnouncementsParams) WithTimeout(timeout time.Duration) *GetPublicGetAnnouncementsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get public get announcements params
func (o *GetPublicGetAnnouncementsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get public get announcements params
func (o *GetPublicGetAnnouncementsParams) WithContext(ctx context.Context) *GetPublicGetAnnouncementsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get public get announcements params
func (o *GetPublicGetAnnouncementsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get public get announcements params
func (o *GetPublicGetAnnouncementsParams) WithHTTPClient(client *http.Client) *GetPublicGetAnnouncementsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get public get announcements params
func (o *GetPublicGetAnnouncementsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetPublicGetAnnouncementsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
