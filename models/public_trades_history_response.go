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

// PublicTradesHistoryResponse public trades history response
// swagger:model public_trades_history_response
type PublicTradesHistoryResponse struct {
	BaseMessage

	// result
	// Required: true
	Result *PublicTradesHistoryResponseAO1Result `json:"result"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *PublicTradesHistoryResponse) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 BaseMessage
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.BaseMessage = aO0

	// AO1
	var dataAO1 struct {
		Result *PublicTradesHistoryResponseAO1Result `json:"result"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.Result = dataAO1.Result

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m PublicTradesHistoryResponse) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.BaseMessage)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	var dataAO1 struct {
		Result *PublicTradesHistoryResponseAO1Result `json:"result"`
	}

	dataAO1.Result = m.Result

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this public trades history response
func (m *PublicTradesHistoryResponse) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with BaseMessage
	if err := m.BaseMessage.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResult(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PublicTradesHistoryResponse) validateResult(formats strfmt.Registry) error {

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
func (m *PublicTradesHistoryResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PublicTradesHistoryResponse) UnmarshalBinary(b []byte) error {
	var res PublicTradesHistoryResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PublicTradesHistoryResponseAO1Result public trades history response a o1 result
// swagger:model PublicTradesHistoryResponseAO1Result
type PublicTradesHistoryResponseAO1Result struct {

	// has more
	// Required: true
	HasMore *bool `json:"has_more"`

	// trades
	// Required: true
	Trades []*PublicTrade `json:"trades"`
}

// Validate validates this public trades history response a o1 result
func (m *PublicTradesHistoryResponseAO1Result) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHasMore(formats); err != nil {
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

func (m *PublicTradesHistoryResponseAO1Result) validateHasMore(formats strfmt.Registry) error {

	if err := validate.Required("result"+"."+"has_more", "body", m.HasMore); err != nil {
		return err
	}

	return nil
}

func (m *PublicTradesHistoryResponseAO1Result) validateTrades(formats strfmt.Registry) error {

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
func (m *PublicTradesHistoryResponseAO1Result) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PublicTradesHistoryResponseAO1Result) UnmarshalBinary(b []byte) error {
	var res PublicTradesHistoryResponseAO1Result
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
