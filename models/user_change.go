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

// UserChange user change
// swagger:model user_change
type UserChange struct {

	// instrument name
	InstrumentName InstrumentName `json:"instrument_name,omitempty"`

	// orders
	Orders []*Order `json:"orders"`

	// position
	Position []*Position `json:"position"`

	// trades
	Trades []*UserTrade `json:"trades"`
}

// Validate validates this user change
func (m *UserChange) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateInstrumentName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrders(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePosition(formats); err != nil {
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

func (m *UserChange) validateInstrumentName(formats strfmt.Registry) error {

	if swag.IsZero(m.InstrumentName) { // not required
		return nil
	}

	if err := m.InstrumentName.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("instrument_name")
		}
		return err
	}

	return nil
}

func (m *UserChange) validateOrders(formats strfmt.Registry) error {

	if swag.IsZero(m.Orders) { // not required
		return nil
	}

	for i := 0; i < len(m.Orders); i++ {
		if swag.IsZero(m.Orders[i]) { // not required
			continue
		}

		if m.Orders[i] != nil {
			if err := m.Orders[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("orders" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *UserChange) validatePosition(formats strfmt.Registry) error {

	if swag.IsZero(m.Position) { // not required
		return nil
	}

	for i := 0; i < len(m.Position); i++ {
		if swag.IsZero(m.Position[i]) { // not required
			continue
		}

		if m.Position[i] != nil {
			if err := m.Position[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("position" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *UserChange) validateTrades(formats strfmt.Registry) error {

	if swag.IsZero(m.Trades) { // not required
		return nil
	}

	for i := 0; i < len(m.Trades); i++ {
		if swag.IsZero(m.Trades[i]) { // not required
			continue
		}

		if m.Trades[i] != nil {
			if err := m.Trades[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("trades" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *UserChange) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserChange) UnmarshalBinary(b []byte) error {
	var res UserChange
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
