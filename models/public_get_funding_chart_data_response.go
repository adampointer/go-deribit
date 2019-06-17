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

// PublicGetFundingChartDataResponse public get funding chart data response
// swagger:model public_get_funding_chart_data_response
type PublicGetFundingChartDataResponse struct {
	BaseMessage

	// result
	// Required: true
	Result *PublicGetFundingChartDataResponseAO1Result `json:"result"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *PublicGetFundingChartDataResponse) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 BaseMessage
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.BaseMessage = aO0

	// AO1
	var dataAO1 struct {
		Result *PublicGetFundingChartDataResponseAO1Result `json:"result"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.Result = dataAO1.Result

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m PublicGetFundingChartDataResponse) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.BaseMessage)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	var dataAO1 struct {
		Result *PublicGetFundingChartDataResponseAO1Result `json:"result"`
	}

	dataAO1.Result = m.Result

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this public get funding chart data response
func (m *PublicGetFundingChartDataResponse) Validate(formats strfmt.Registry) error {
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

func (m *PublicGetFundingChartDataResponse) validateResult(formats strfmt.Registry) error {

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
func (m *PublicGetFundingChartDataResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PublicGetFundingChartDataResponse) UnmarshalBinary(b []byte) error {
	var res PublicGetFundingChartDataResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PublicGetFundingChartDataResponseAO1Result public get funding chart data response a o1 result
// swagger:model PublicGetFundingChartDataResponseAO1Result
type PublicGetFundingChartDataResponseAO1Result struct {

	// Current interest
	// Required: true
	CurrentInterest *float64 `json:"current_interest"`

	// data
	// Required: true
	Data []interface{} `json:"data"`

	// Current index price
	// Required: true
	IndexPrice *float64 `json:"index_price"`

	// Current interest 8h
	// Required: true
	Interest8h *float64 `json:"interest_8h"`

	// maximal interest
	// Required: true
	Max *float64 `json:"max"`
}

// Validate validates this public get funding chart data response a o1 result
func (m *PublicGetFundingChartDataResponseAO1Result) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCurrentInterest(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIndexPrice(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInterest8h(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMax(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PublicGetFundingChartDataResponseAO1Result) validateCurrentInterest(formats strfmt.Registry) error {

	if err := validate.Required("result"+"."+"current_interest", "body", m.CurrentInterest); err != nil {
		return err
	}

	return nil
}

func (m *PublicGetFundingChartDataResponseAO1Result) validateData(formats strfmt.Registry) error {

	if err := validate.Required("result"+"."+"data", "body", m.Data); err != nil {
		return err
	}

	return nil
}

func (m *PublicGetFundingChartDataResponseAO1Result) validateIndexPrice(formats strfmt.Registry) error {

	if err := validate.Required("result"+"."+"index_price", "body", m.IndexPrice); err != nil {
		return err
	}

	return nil
}

func (m *PublicGetFundingChartDataResponseAO1Result) validateInterest8h(formats strfmt.Registry) error {

	if err := validate.Required("result"+"."+"interest_8h", "body", m.Interest8h); err != nil {
		return err
	}

	return nil
}

func (m *PublicGetFundingChartDataResponseAO1Result) validateMax(formats strfmt.Registry) error {

	if err := validate.Required("result"+"."+"max", "body", m.Max); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PublicGetFundingChartDataResponseAO1Result) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PublicGetFundingChartDataResponseAO1Result) UnmarshalBinary(b []byte) error {
	var res PublicGetFundingChartDataResponseAO1Result
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
