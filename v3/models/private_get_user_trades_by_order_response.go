// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// PrivateGetUserTradesByOrderResponse private get user trades by order response
// swagger:model private_get_user_trades_by_order_response
type PrivateGetUserTradesByOrderResponse struct {
	BaseMessage

	PrivateGetUserTradesByOrderResponseAllOf1
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *PrivateGetUserTradesByOrderResponse) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 BaseMessage
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.BaseMessage = aO0

	// AO1
	var aO1 PrivateGetUserTradesByOrderResponseAllOf1
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.PrivateGetUserTradesByOrderResponseAllOf1 = aO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m PrivateGetUserTradesByOrderResponse) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.BaseMessage)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.PrivateGetUserTradesByOrderResponseAllOf1)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this private get user trades by order response
func (m *PrivateGetUserTradesByOrderResponse) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with BaseMessage
	if err := m.BaseMessage.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with PrivateGetUserTradesByOrderResponseAllOf1
	if err := m.PrivateGetUserTradesByOrderResponseAllOf1.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *PrivateGetUserTradesByOrderResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PrivateGetUserTradesByOrderResponse) UnmarshalBinary(b []byte) error {
	var res PrivateGetUserTradesByOrderResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PrivateGetUserTradesByOrderResponseAllOf1 private get user trades by order response all of1
// swagger:model PrivateGetUserTradesByOrderResponseAllOf1
type PrivateGetUserTradesByOrderResponseAllOf1 []*UserTrade

// Validate validates this private get user trades by order response all of1
func (m PrivateGetUserTradesByOrderResponseAllOf1) Validate(formats strfmt.Registry) error {
	var res []error

	for i := 0; i < len(m); i++ {
		if swag.IsZero(m[i]) { // not required
			continue
		}

		if m[i] != nil {
			if err := m[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName(strconv.Itoa(i))
				}
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}