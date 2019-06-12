// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PublicIndexResponse public index response
// swagger:model public_index_response
type PublicIndexResponse struct {

	// result
	// Required: true
	Result *PublicIndexResponseResult `json:"result"`
}

// Validate validates this public index response
func (m *PublicIndexResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateResult(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PublicIndexResponse) validateResult(formats strfmt.Registry) error {

	if err := validate.Required("result", "body", m.Result); err != nil {
		return err
	}

	if m.Result != nil {
		if err := m.Result.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("result")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PublicIndexResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PublicIndexResponse) UnmarshalBinary(b []byte) error {
	var res PublicIndexResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PublicIndexResponseResult public index response result
// swagger:model PublicIndexResponseResult
type PublicIndexResponseResult struct {

	// The current index price for BTC-USD (only for selected currency == BTC)
	BTC float64 `json:"BTC"`

	// The current index price for ETH-USD (only for selected currency == ETH)
	ETH float64 `json:"ETH,omitempty"`

	// Estimated delivery price for the currency. For more details, see Documentation > General > Expiration Price
	// Required: true
	Edp *float64 `json:"edp"`
}

// Validate validates this public index response result
func (m *PublicIndexResponseResult) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBTC(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEdp(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PublicIndexResponseResult) validateBTC(formats strfmt.Registry) error {

	if err := validate.Required("result"+"."+"BTC", "body", m.BTC); err != nil {
		return err
	}

	return nil
}

func (m *PublicIndexResponseResult) validateEdp(formats strfmt.Registry) error {

	if err := validate.Required("result"+"."+"edp", "body", m.Edp); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PublicIndexResponseResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PublicIndexResponseResult) UnmarshalBinary(b []byte) error {
	var res PublicIndexResponseResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
