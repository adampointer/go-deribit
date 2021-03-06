// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CurrencyDetail currency detail
// swagger:model currency_detail
type CurrencyDetail struct {

	// The type of the currency.
	// Required: true
	// Enum: [BITCOIN ETHER]
	CoinType *string `json:"coin_type"`

	// The abbreviation of the currency. This abbreviation is used elsewhere in the API to identify the currency.
	// Required: true
	Currency *string `json:"currency"`

	// The full name for the currency.
	// Required: true
	CurrencyLong *string `json:"currency_long"`

	// False if deposit address creation is disabled
	DisabledDepositAddressCreation bool `json:"disabled_deposit_address_creation,omitempty"`

	// fee precision
	FeePrecision int64 `json:"fee_precision,omitempty"`

	// Minimum number of block chain confirmations before deposit is accepted.
	// Required: true
	MinConfirmations *int64 `json:"min_confirmations"`

	// The minimum transaction fee paid for withdrawals
	MinWithdrawalFee float64 `json:"min_withdrawal_fee,omitempty"`

	// The total transaction fee paid for withdrawals
	// Required: true
	WithdrawalFee *float64 `json:"withdrawal_fee"`

	// withdrawal priorities
	WithdrawalPriorities []*KeyNumberPair `json:"withdrawal_priorities"`
}

// Validate validates this currency detail
func (m *CurrencyDetail) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCoinType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCurrency(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCurrencyLong(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMinConfirmations(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWithdrawalFee(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWithdrawalPriorities(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var currencyDetailTypeCoinTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["BITCOIN","ETHER"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		currencyDetailTypeCoinTypePropEnum = append(currencyDetailTypeCoinTypePropEnum, v)
	}
}

const (

	// CurrencyDetailCoinTypeBITCOIN captures enum value "BITCOIN"
	CurrencyDetailCoinTypeBITCOIN string = "BITCOIN"

	// CurrencyDetailCoinTypeETHER captures enum value "ETHER"
	CurrencyDetailCoinTypeETHER string = "ETHER"
)

// prop value enum
func (m *CurrencyDetail) validateCoinTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, currencyDetailTypeCoinTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *CurrencyDetail) validateCoinType(formats strfmt.Registry) error {

	if err := validate.Required("coin_type", "body", m.CoinType); err != nil {
		return err
	}

	// value enum
	if err := m.validateCoinTypeEnum("coin_type", "body", *m.CoinType); err != nil {
		return err
	}

	return nil
}

func (m *CurrencyDetail) validateCurrency(formats strfmt.Registry) error {

	if err := validate.Required("currency", "body", m.Currency); err != nil {
		return err
	}

	return nil
}

func (m *CurrencyDetail) validateCurrencyLong(formats strfmt.Registry) error {

	if err := validate.Required("currency_long", "body", m.CurrencyLong); err != nil {
		return err
	}

	return nil
}

func (m *CurrencyDetail) validateMinConfirmations(formats strfmt.Registry) error {

	if err := validate.Required("min_confirmations", "body", m.MinConfirmations); err != nil {
		return err
	}

	return nil
}

func (m *CurrencyDetail) validateWithdrawalFee(formats strfmt.Registry) error {

	if err := validate.Required("withdrawal_fee", "body", m.WithdrawalFee); err != nil {
		return err
	}

	return nil
}

func (m *CurrencyDetail) validateWithdrawalPriorities(formats strfmt.Registry) error {

	if swag.IsZero(m.WithdrawalPriorities) { // not required
		return nil
	}

	for i := 0; i < len(m.WithdrawalPriorities); i++ {
		if swag.IsZero(m.WithdrawalPriorities[i]) { // not required
			continue
		}

		if m.WithdrawalPriorities[i] != nil {
			if err := m.WithdrawalPriorities[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("withdrawal_priorities" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *CurrencyDetail) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CurrencyDetail) UnmarshalBinary(b []byte) error {
	var res CurrencyDetail
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
