// Code generated by go-swagger; DO NOT EDIT.

package internal

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new internal API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for internal API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetPrivateDisableTfaWithRecoveryCode disables t f a with one time recovery code
*/
func (a *Client) GetPrivateDisableTfaWithRecoveryCode(params *GetPrivateDisableTfaWithRecoveryCodeParams) (*GetPrivateDisableTfaWithRecoveryCodeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPrivateDisableTfaWithRecoveryCodeParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetPrivateDisableTfaWithRecoveryCode",
		Method:             "GET",
		PathPattern:        "/private/disable_tfa_with_recovery_code",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetPrivateDisableTfaWithRecoveryCodeReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetPrivateDisableTfaWithRecoveryCodeOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetPrivateDisableTfaWithRecoveryCode: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetPublicGetFooter gets information to be displayed in the footer of the website
*/
func (a *Client) GetPublicGetFooter(params *GetPublicGetFooterParams) (*GetPublicGetFooterOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPublicGetFooterParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetPublicGetFooter",
		Method:             "GET",
		PathPattern:        "/public/get_footer",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetPublicGetFooterReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetPublicGetFooterOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetPublicGetFooter: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetPublicGetOptionMarkPrices retrives market prices and its implied volatility of options instruments
*/
func (a *Client) GetPublicGetOptionMarkPrices(params *GetPublicGetOptionMarkPricesParams) (*GetPublicGetOptionMarkPricesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPublicGetOptionMarkPricesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetPublicGetOptionMarkPrices",
		Method:             "GET",
		PathPattern:        "/public/get_option_mark_prices",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetPublicGetOptionMarkPricesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetPublicGetOptionMarkPricesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetPublicGetOptionMarkPrices: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
