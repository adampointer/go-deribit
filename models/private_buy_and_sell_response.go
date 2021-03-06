// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PrivateBuyAndSellResponse private buy and sell response
// swagger:model private_buy_and_sell_response
type PrivateBuyAndSellResponse struct {

	// result
	// Required: true
	Result *PrivateBuyAndSellResponseResult `json:"result"`
}

// Validate validates this private buy and sell response
func (m *PrivateBuyAndSellResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateResult(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PrivateBuyAndSellResponse) validateResult(formats strfmt.Registry) error {

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
func (m *PrivateBuyAndSellResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PrivateBuyAndSellResponse) UnmarshalBinary(b []byte) error {
	var res PrivateBuyAndSellResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PrivateBuyAndSellResponseResult private buy and sell response result
// swagger:model PrivateBuyAndSellResponseResult
type PrivateBuyAndSellResponseResult struct {

	// order
	// Required: true
	Order *Order `json:"order"`

	// trades
	// Required: true
	Trades []*UserTrade `json:"trades"`
}

// Validate validates this private buy and sell response result
func (m *PrivateBuyAndSellResponseResult) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOrder(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTrades(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PrivateBuyAndSellResponseResult) validateOrder(formats strfmt.Registry) error {

	if err := validate.Required("result"+"."+"order", "body", m.Order); err != nil {
		return err
	}

	if m.Order != nil {
		if err := m.Order.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("result" + "." + "order")
			}
			return err
		}
	}

	return nil
}

func (m *PrivateBuyAndSellResponseResult) validateTrades(formats strfmt.Registry) error {

	if err := validate.Required("result"+"."+"trades", "body", m.Trades); err != nil {
		return err
	}

	for i := 0; i < len(m.Trades); i++ {
		if swag.IsZero(m.Trades[i]) { // not required
			continue
		}

		if m.Trades[i] != nil {
			if err := m.Trades[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("result" + "." + "trades" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PrivateBuyAndSellResponseResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PrivateBuyAndSellResponseResult) UnmarshalBinary(b []byte) error {
	var res PrivateBuyAndSellResponseResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
