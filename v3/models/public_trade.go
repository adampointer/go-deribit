// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PublicTrade public trade
// swagger:model public_trade
type PublicTrade struct {

	// Trade amount. For perpetual and futures - in USD units, for options it is amount of corresponding cryptocurrency contracts, e.g., BTC or ETH.
	// Required: true
	Amount *float64 `json:"amount"`

	// Trade direction of the taker
	// Required: true
	Direction Direction `json:"direction"`

	// Index Price at the moment of trade
	// Required: true
	IndexPrice *float64 `json:"index_price"`

	// instrument name
	// Required: true
	InstrumentName InstrumentName `json:"instrument_name"`

	// Option implied volatility for the price (Option only)
	Iv float64 `json:"iv,omitempty"`

	// Optional field (only for trades caused by liquidation): `"M"` when maker side of trade was under liquidation, `"T"` when taker side was under liquidation, `"MT"` when both sides of trade were under liquidation
	// Enum: [M T MT]
	Liquidation string `json:"liquidation,omitempty"`

	// The price of the trade
	// Required: true
	Price Price `json:"price"`

	// tick direction
	// Required: true
	TickDirection TickDirection `json:"tick_direction"`

	// The timestamp of the trade
	// Required: true
	Timestamp *int64 `json:"timestamp"`

	// trade id
	// Required: true
	TradeID TradeID `json:"trade_id"`

	// trade seq
	// Required: true
	TradeSeq TradeSeq `json:"trade_seq"`
}

// Validate validates this public trade
func (m *PublicTrade) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAmount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDirection(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIndexPrice(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInstrumentName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLiquidation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrice(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTickDirection(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTradeID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTradeSeq(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PublicTrade) validateAmount(formats strfmt.Registry) error {

	if err := validate.Required("amount", "body", m.Amount); err != nil {
		return err
	}

	return nil
}

func (m *PublicTrade) validateDirection(formats strfmt.Registry) error {

	if err := m.Direction.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("direction")
		}
		return err
	}

	return nil
}

func (m *PublicTrade) validateIndexPrice(formats strfmt.Registry) error {

	if err := validate.Required("index_price", "body", m.IndexPrice); err != nil {
		return err
	}

	return nil
}

func (m *PublicTrade) validateInstrumentName(formats strfmt.Registry) error {

	if err := m.InstrumentName.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("instrument_name")
		}
		return err
	}

	return nil
}

var publicTradeTypeLiquidationPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["M","T","MT"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		publicTradeTypeLiquidationPropEnum = append(publicTradeTypeLiquidationPropEnum, v)
	}
}

const (

	// PublicTradeLiquidationM captures enum value "M"
	PublicTradeLiquidationM string = "M"

	// PublicTradeLiquidationT captures enum value "T"
	PublicTradeLiquidationT string = "T"

	// PublicTradeLiquidationMT captures enum value "MT"
	PublicTradeLiquidationMT string = "MT"
)

// prop value enum
func (m *PublicTrade) validateLiquidationEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, publicTradeTypeLiquidationPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *PublicTrade) validateLiquidation(formats strfmt.Registry) error {

	if swag.IsZero(m.Liquidation) { // not required
		return nil
	}

	// value enum
	if err := m.validateLiquidationEnum("liquidation", "body", m.Liquidation); err != nil {
		return err
	}

	return nil
}

func (m *PublicTrade) validatePrice(formats strfmt.Registry) error {

	if err := m.Price.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("price")
		}
		return err
	}

	return nil
}

func (m *PublicTrade) validateTickDirection(formats strfmt.Registry) error {

	if err := m.TickDirection.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("tick_direction")
		}
		return err
	}

	return nil
}

func (m *PublicTrade) validateTimestamp(formats strfmt.Registry) error {

	if err := validate.Required("timestamp", "body", m.Timestamp); err != nil {
		return err
	}

	return nil
}

func (m *PublicTrade) validateTradeID(formats strfmt.Registry) error {

	if err := m.TradeID.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("trade_id")
		}
		return err
	}

	return nil
}

func (m *PublicTrade) validateTradeSeq(formats strfmt.Registry) error {

	if err := m.TradeSeq.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("trade_seq")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PublicTrade) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PublicTrade) UnmarshalBinary(b []byte) error {
	var res PublicTrade
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}