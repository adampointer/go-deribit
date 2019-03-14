package deribit

import (
	"fmt"
	"github.com/adampointer/go-deribit/models/private"
	"github.com/adampointer/go-deribit/models/public"
	"github.com/mitchellh/mapstructure"
)

// PrivateBuy makes a request to private/buy
func (e *Exchange) PrivateBuy(params *private.BuyRequest) (*private.BuyResponse, error) {
	res, err := e.rpcRequest("private/buy", params)
	if err != nil {
		return nil, err
	}
	var ret private.BuyResponse
	if err := mapstructure.Decode(res.Result, &ret); err != nil {
		return nil, fmt.Errorf("error decoding result: %s", err)
	}
	return &ret, nil
}

// PrivateCancel makes a request to private/cancel
func (e *Exchange) PrivateCancel(params *private.CancelRequest) (*private.CancelResponse, error) {
	res, err := e.rpcRequest("private/cancel", params)
	if err != nil {
		return nil, err
	}
	var ret private.CancelResponse
	if err := mapstructure.Decode(res.Result, &ret); err != nil {
		return nil, fmt.Errorf("error decoding result: %s", err)
	}
	return &ret, nil
}

// PrivateCancelAll makes a request to private/cancel_all
func (e *Exchange) PrivateCancelAll(params *private.CancelAllRequest) (*private.CancelAllResponse, error) {
	res, err := e.rpcRequest("private/cancel_all", params)
	if err != nil {
		return nil, err
	}
	var ret private.CancelAllResponse
	if err := mapstructure.Decode(res.Result, &ret); err != nil {
		return nil, fmt.Errorf("error decoding result: %s", err)
	}
	return &ret, nil
}

// PrivateCancelAllByCurrency makes a request to private/cancel_all_by_currency
func (e *Exchange) PrivateCancelAllByCurrency(params *private.CancelAllByCurrencyRequest) (*private.CancelAllByCurrencyResponse, error) {
	res, err := e.rpcRequest("private/cancel_all_by_currency", params)
	if err != nil {
		return nil, err
	}
	var ret private.CancelAllByCurrencyResponse
	if err := mapstructure.Decode(res.Result, &ret); err != nil {
		return nil, fmt.Errorf("error decoding result: %s", err)
	}
	return &ret, nil
}

// PrivateCancelAllByInstrument makes a request to private/cancel_all_by_instrument
func (e *Exchange) PrivateCancelAllByInstrument(params *private.CancelAllByInstrumentRequest) (*private.CancelAllByInstrumentResponse, error) {
	res, err := e.rpcRequest("private/cancel_all_by_instrument", params)
	if err != nil {
		return nil, err
	}
	var ret private.CancelAllByInstrumentResponse
	if err := mapstructure.Decode(res.Result, &ret); err != nil {
		return nil, fmt.Errorf("error decoding result: %s", err)
	}
	return &ret, nil
}

// PrivateCancelTransferById makes a request to private/cancel_transfer_by_id
func (e *Exchange) PrivateCancelTransferById(params *private.CancelTransferByIdRequest) (*private.CancelTransferByIdResponse, error) {
	res, err := e.rpcRequest("private/cancel_transfer_by_id", params)
	if err != nil {
		return nil, err
	}
	var ret private.CancelTransferByIdResponse
	if err := mapstructure.Decode(res.Result, &ret); err != nil {
		return nil, fmt.Errorf("error decoding result: %s", err)
	}
	return &ret, nil
}

// PrivateCancelWithdrawal makes a request to private/cancel_withdrawal
func (e *Exchange) PrivateCancelWithdrawal(params *private.CancelWithdrawalRequest) (*private.CancelWithdrawalResponse, error) {
	res, err := e.rpcRequest("private/cancel_withdrawal", params)
	if err != nil {
		return nil, err
	}
	var ret private.CancelWithdrawalResponse
	if err := mapstructure.Decode(res.Result, &ret); err != nil {
		return nil, fmt.Errorf("error decoding result: %s", err)
	}
	return &ret, nil
}

// PrivateChangeSubaccountName makes a request to private/change_subaccount_name
func (e *Exchange) PrivateChangeSubaccountName(params *private.ChangeSubaccountNameRequest) (*private.ChangeSubaccountNameResponse, error) {
	res, err := e.rpcRequest("private/change_subaccount_name", params)
	if err != nil {
		return nil, err
	}
	var ret private.ChangeSubaccountNameResponse
	if err := mapstructure.Decode(res.Result, &ret); err != nil {
		return nil, fmt.Errorf("error decoding result: %s", err)
	}
	return &ret, nil
}

// PrivateClosePosition makes a request to private/close_position
func (e *Exchange) PrivateClosePosition(params *private.ClosePositionRequest) (*private.ClosePositionResponse, error) {
	res, err := e.rpcRequest("private/close_position", params)
	if err != nil {
		return nil, err
	}
	var ret private.ClosePositionResponse
	if err := mapstructure.Decode(res.Result, &ret); err != nil {
		return nil, fmt.Errorf("error decoding result: %s", err)
	}
	return &ret, nil
}

// PrivateCreateDepositAddress makes a request to private/create_deposit_address
func (e *Exchange) PrivateCreateDepositAddress(params *private.CreateDepositAddressRequest) (*private.CreateDepositAddressResponse, error) {
	res, err := e.rpcRequest("private/create_deposit_address", params)
	if err != nil {
		return nil, err
	}
	var ret private.CreateDepositAddressResponse
	if err := mapstructure.Decode(res.Result, &ret); err != nil {
		return nil, fmt.Errorf("error decoding result: %s", err)
	}
	return &ret, nil
}

// PrivateCreateSubaccount makes a request to private/create_subaccount
func (e *Exchange) PrivateCreateSubaccount(params *private.CreateSubaccountRequest) (*private.CreateSubaccountResponse, error) {
	res, err := e.rpcRequest("private/create_subaccount", params)
	if err != nil {
		return nil, err
	}
	var ret private.CreateSubaccountResponse
	if err := mapstructure.Decode(res.Result, &ret); err != nil {
		return nil, fmt.Errorf("error decoding result: %s", err)
	}
	return &ret, nil
}

// PrivateDisableCancelOnDisconnect makes a request to private/disable_cancel_on_disconnect
func (e *Exchange) PrivateDisableCancelOnDisconnect(params *private.DisableCancelOnDisconnectRequest) (*private.DisableCancelOnDisconnectResponse, error) {
	res, err := e.rpcRequest("private/disable_cancel_on_disconnect", params)
	if err != nil {
		return nil, err
	}
	var ret private.DisableCancelOnDisconnectResponse
	if err := mapstructure.Decode(res.Result, &ret); err != nil {
		return nil, fmt.Errorf("error decoding result: %s", err)
	}
	return &ret, nil
}

// PrivateDisableTfaForSubaccount makes a request to private/disable_tfa_for_subaccount
func (e *Exchange) PrivateDisableTfaForSubaccount(params *private.DisableTfaForSubaccountRequest) (*private.DisableTfaForSubaccountResponse, error) {
	res, err := e.rpcRequest("private/disable_tfa_for_subaccount", params)
	if err != nil {
		return nil, err
	}
	var ret private.DisableTfaForSubaccountResponse
	if err := mapstructure.Decode(res.Result, &ret); err != nil {
		return nil, fmt.Errorf("error decoding result: %s", err)
	}
	return &ret, nil
}

// PrivateEdit makes a request to private/edit
func (e *Exchange) PrivateEdit(params *private.EditRequest) (*private.EditResponse, error) {
	res, err := e.rpcRequest("private/edit", params)
	if err != nil {
		return nil, err
	}
	var ret private.EditResponse
	if err := mapstructure.Decode(res.Result, &ret); err != nil {
		return nil, fmt.Errorf("error decoding result: %s", err)
	}
	return &ret, nil
}

// PrivateEnableCancelOnDisconnect makes a request to private/enable_cancel_on_disconnect
func (e *Exchange) PrivateEnableCancelOnDisconnect(params *private.EnableCancelOnDisconnectRequest) (*private.EnableCancelOnDisconnectResponse, error) {
	res, err := e.rpcRequest("private/enable_cancel_on_disconnect", params)
	if err != nil {
		return nil, err
	}
	var ret private.EnableCancelOnDisconnectResponse
	if err := mapstructure.Decode(res.Result, &ret); err != nil {
		return nil, fmt.Errorf("error decoding result: %s", err)
	}
	return &ret, nil
}

// PrivateGetAccountSummary makes a request to private/get_account_summary
func (e *Exchange) PrivateGetAccountSummary(params *private.GetAccountSummaryRequest) (*private.GetAccountSummaryResponse, error) {
	res, err := e.rpcRequest("private/get_account_summary", params)
	if err != nil {
		return nil, err
	}
	var ret private.GetAccountSummaryResponse
	if err := mapstructure.Decode(res.Result, &ret); err != nil {
		return nil, fmt.Errorf("error decoding result: %s", err)
	}
	return &ret, nil
}

// PrivateGetCurrentDepositAddress makes a request to private/get_current_deposit_address
func (e *Exchange) PrivateGetCurrentDepositAddress(params *private.GetCurrentDepositAddressRequest) (*private.GetCurrentDepositAddressResponse, error) {
	res, err := e.rpcRequest("private/get_current_deposit_address", params)
	if err != nil {
		return nil, err
	}
	var ret private.GetCurrentDepositAddressResponse
	if err := mapstructure.Decode(res.Result, &ret); err != nil {
		return nil, fmt.Errorf("error decoding result: %s", err)
	}
	return &ret, nil
}

// PrivateGetDeposits makes a request to private/get_deposits
func (e *Exchange) PrivateGetDeposits(params *private.GetDepositsRequest) (*private.GetDepositsResponse, error) {
	res, err := e.rpcRequest("private/get_deposits", params)
	if err != nil {
		return nil, err
	}
	var ret private.GetDepositsResponse
	if err := mapstructure.Decode(res.Result, &ret); err != nil {
		return nil, fmt.Errorf("error decoding result: %s", err)
	}
	return &ret, nil
}

// PrivateGetEmailLanguage makes a request to private/get_email_language
func (e *Exchange) PrivateGetEmailLanguage(params *private.GetEmailLanguageRequest) (*private.GetEmailLanguageResponse, error) {
	res, err := e.rpcRequest("private/get_email_language", params)
	if err != nil {
		return nil, err
	}
	var ret private.GetEmailLanguageResponse
	if err := mapstructure.Decode(res.Result, &ret); err != nil {
		return nil, fmt.Errorf("error decoding result: %s", err)
	}
	return &ret, nil
}

// PrivateGetMargins makes a request to private/get_margins
func (e *Exchange) PrivateGetMargins(params *private.GetMarginsRequest) (*private.GetMarginsResponse, error) {
	res, err := e.rpcRequest("private/get_margins", params)
	if err != nil {
		return nil, err
	}
	var ret private.GetMarginsResponse
	if err := mapstructure.Decode(res.Result, &ret); err != nil {
		return nil, fmt.Errorf("error decoding result: %s", err)
	}
	return &ret, nil
}

// PrivateGetNewAnnouncements makes a request to private/get_new_announcements
func (e *Exchange) PrivateGetNewAnnouncements(params *private.GetNewAnnouncementsRequest) (*private.GetNewAnnouncementsResponse, error) {
	res, err := e.rpcRequest("private/get_new_announcements", params)
	if err != nil {
		return nil, err
	}
	var ret private.GetNewAnnouncementsResponse
	if err := mapstructure.Decode(res.Result, &ret); err != nil {
		return nil, fmt.Errorf("error decoding result: %s", err)
	}
	return &ret, nil
}

// PrivateGetOpenOrdersByCurrency makes a request to private/get_open_orders_by_currency
func (e *Exchange) PrivateGetOpenOrdersByCurrency(params *private.GetOpenOrdersByCurrencyRequest) (*private.GetOpenOrdersByCurrencyResponse, error) {
	res, err := e.rpcRequest("private/get_open_orders_by_currency", params)
	if err != nil {
		return nil, err
	}
	var ret private.GetOpenOrdersByCurrencyResponse
	if err := mapstructure.Decode(res.Result, &ret); err != nil {
		return nil, fmt.Errorf("error decoding result: %s", err)
	}
	return &ret, nil
}

// PrivateGetOpenOrdersByInstrument makes a request to private/get_open_orders_by_instrument
func (e *Exchange) PrivateGetOpenOrdersByInstrument(params *private.GetOpenOrdersByInstrumentRequest) (*private.GetOpenOrdersByInstrumentResponse, error) {
	res, err := e.rpcRequest("private/get_open_orders_by_instrument", params)
	if err != nil {
		return nil, err
	}
	var ret private.GetOpenOrdersByInstrumentResponse
	if err := mapstructure.Decode(res.Result, &ret); err != nil {
		return nil, fmt.Errorf("error decoding result: %s", err)
	}
	return &ret, nil
}

// PrivateGetOrderHistoryByCurrency makes a request to private/get_order_history_by_currency
func (e *Exchange) PrivateGetOrderHistoryByCurrency(params *private.GetOrderHistoryByCurrencyRequest) (*private.GetOrderHistoryByCurrencyResponse, error) {
	res, err := e.rpcRequest("private/get_order_history_by_currency", params)
	if err != nil {
		return nil, err
	}
	var ret private.GetOrderHistoryByCurrencyResponse
	if err := mapstructure.Decode(res.Result, &ret); err != nil {
		return nil, fmt.Errorf("error decoding result: %s", err)
	}
	return &ret, nil
}

// PrivateGetOrderHistoryByInstrument makes a request to private/get_order_history_by_instrument
func (e *Exchange) PrivateGetOrderHistoryByInstrument(params *private.GetOrderHistoryByInstrumentRequest) (*private.GetOrderHistoryByInstrumentResponse, error) {
	res, err := e.rpcRequest("private/get_order_history_by_instrument", params)
	if err != nil {
		return nil, err
	}
	var ret private.GetOrderHistoryByInstrumentResponse
	if err := mapstructure.Decode(res.Result, &ret); err != nil {
		return nil, fmt.Errorf("error decoding result: %s", err)
	}
	return &ret, nil
}

// PrivateGetOrderMarginByIds makes a request to private/get_order_margin_by_ids
func (e *Exchange) PrivateGetOrderMarginByIds(params *private.GetOrderMarginByIdsRequest) (*private.GetOrderMarginByIdsResponse, error) {
	res, err := e.rpcRequest("private/get_order_margin_by_ids", params)
	if err != nil {
		return nil, err
	}
	var ret private.GetOrderMarginByIdsResponse
	if err := mapstructure.Decode(res.Result, &ret); err != nil {
		return nil, fmt.Errorf("error decoding result: %s", err)
	}
	return &ret, nil
}

// PrivateGetOrderState makes a request to private/get_order_state
func (e *Exchange) PrivateGetOrderState(params *private.GetOrderStateRequest) (*private.GetOrderStateResponse, error) {
	res, err := e.rpcRequest("private/get_order_state", params)
	if err != nil {
		return nil, err
	}
	var ret private.GetOrderStateResponse
	if err := mapstructure.Decode(res.Result, &ret); err != nil {
		return nil, fmt.Errorf("error decoding result: %s", err)
	}
	return &ret, nil
}

// PrivateGetPosition makes a request to private/get_position
func (e *Exchange) PrivateGetPosition(params *private.GetPositionRequest) (*private.GetPositionResponse, error) {
	res, err := e.rpcRequest("private/get_position", params)
	if err != nil {
		return nil, err
	}
	var ret private.GetPositionResponse
	if err := mapstructure.Decode(res.Result, &ret); err != nil {
		return nil, fmt.Errorf("error decoding result: %s", err)
	}
	return &ret, nil
}

// PrivateGetPositions makes a request to private/get_positions
func (e *Exchange) PrivateGetPositions(params *private.GetPositionsRequest) (*private.GetPositionsResponse, error) {
	res, err := e.rpcRequest("private/get_positions", params)
	if err != nil {
		return nil, err
	}
	var ret private.GetPositionsResponse
	if err := mapstructure.Decode(res.Result, &ret); err != nil {
		return nil, fmt.Errorf("error decoding result: %s", err)
	}
	return &ret, nil
}

// PrivateGetSettlementHistoryByCurrency makes a request to private/get_settlement_history_by_currency
func (e *Exchange) PrivateGetSettlementHistoryByCurrency(params *private.GetSettlementHistoryByCurrencyRequest) (*private.GetSettlementHistoryByCurrencyResponse, error) {
	res, err := e.rpcRequest("private/get_settlement_history_by_currency", params)
	if err != nil {
		return nil, err
	}
	var ret private.GetSettlementHistoryByCurrencyResponse
	if err := mapstructure.Decode(res.Result, &ret); err != nil {
		return nil, fmt.Errorf("error decoding result: %s", err)
	}
	return &ret, nil
}

// PrivateGetSettlementHistoryByInstrument makes a request to private/get_settlement_history_by_instrument
func (e *Exchange) PrivateGetSettlementHistoryByInstrument(params *private.GetSettlementHistoryByInstrumentRequest) (*private.GetSettlementHistoryByInstrumentResponse, error) {
	res, err := e.rpcRequest("private/get_settlement_history_by_instrument", params)
	if err != nil {
		return nil, err
	}
	var ret private.GetSettlementHistoryByInstrumentResponse
	if err := mapstructure.Decode(res.Result, &ret); err != nil {
		return nil, fmt.Errorf("error decoding result: %s", err)
	}
	return &ret, nil
}

// PrivateGetSubaccounts makes a request to private/get_subaccounts
func (e *Exchange) PrivateGetSubaccounts(params *private.GetSubaccountsRequest) (*private.GetSubaccountsResponse, error) {
	res, err := e.rpcRequest("private/get_subaccounts", params)
	if err != nil {
		return nil, err
	}
	var ret private.GetSubaccountsResponse
	if err := mapstructure.Decode(res.Result, &ret); err != nil {
		return nil, fmt.Errorf("error decoding result: %s", err)
	}
	return &ret, nil
}

// PrivateGetTransfers makes a request to private/get_transfers
func (e *Exchange) PrivateGetTransfers(params *private.GetTransfersRequest) (*private.GetTransfersResponse, error) {
	res, err := e.rpcRequest("private/get_transfers", params)
	if err != nil {
		return nil, err
	}
	var ret private.GetTransfersResponse
	if err := mapstructure.Decode(res.Result, &ret); err != nil {
		return nil, fmt.Errorf("error decoding result: %s", err)
	}
	return &ret, nil
}

// PrivateGetUserTradesByCurrency makes a request to private/get_user_trades_by_currency
func (e *Exchange) PrivateGetUserTradesByCurrency(params *private.GetUserTradesByCurrencyRequest) (*private.GetUserTradesByCurrencyResponse, error) {
	res, err := e.rpcRequest("private/get_user_trades_by_currency", params)
	if err != nil {
		return nil, err
	}
	var ret private.GetUserTradesByCurrencyResponse
	if err := mapstructure.Decode(res.Result, &ret); err != nil {
		return nil, fmt.Errorf("error decoding result: %s", err)
	}
	return &ret, nil
}

// PrivateGetUserTradesByCurrencyAndTime makes a request to private/get_user_trades_by_currency_and_time
func (e *Exchange) PrivateGetUserTradesByCurrencyAndTime(params *private.GetUserTradesByCurrencyAndTimeRequest) (*private.GetUserTradesByCurrencyAndTimeResponse, error) {
	res, err := e.rpcRequest("private/get_user_trades_by_currency_and_time", params)
	if err != nil {
		return nil, err
	}
	var ret private.GetUserTradesByCurrencyAndTimeResponse
	if err := mapstructure.Decode(res.Result, &ret); err != nil {
		return nil, fmt.Errorf("error decoding result: %s", err)
	}
	return &ret, nil
}

// PrivateGetUserTradesByInstrument makes a request to private/get_user_trades_by_instrument
func (e *Exchange) PrivateGetUserTradesByInstrument(params *private.GetUserTradesByInstrumentRequest) (*private.GetUserTradesByInstrumentResponse, error) {
	res, err := e.rpcRequest("private/get_user_trades_by_instrument", params)
	if err != nil {
		return nil, err
	}
	var ret private.GetUserTradesByInstrumentResponse
	if err := mapstructure.Decode(res.Result, &ret); err != nil {
		return nil, fmt.Errorf("error decoding result: %s", err)
	}
	return &ret, nil
}

// PrivateGetUserTradesByInstrumentAndTime makes a request to private/get_user_trades_by_instrument_and_time
func (e *Exchange) PrivateGetUserTradesByInstrumentAndTime(params *private.GetUserTradesByInstrumentAndTimeRequest) (*private.GetUserTradesByInstrumentAndTimeResponse, error) {
	res, err := e.rpcRequest("private/get_user_trades_by_instrument_and_time", params)
	if err != nil {
		return nil, err
	}
	var ret private.GetUserTradesByInstrumentAndTimeResponse
	if err := mapstructure.Decode(res.Result, &ret); err != nil {
		return nil, fmt.Errorf("error decoding result: %s", err)
	}
	return &ret, nil
}

// PrivateGetUserTradesByOrder makes a request to private/get_user_trades_by_order
func (e *Exchange) PrivateGetUserTradesByOrder(params *private.GetUserTradesByOrderRequest) (*private.GetUserTradesByOrderResponse, error) {
	res, err := e.rpcRequest("private/get_user_trades_by_order", params)
	if err != nil {
		return nil, err
	}
	var ret private.GetUserTradesByOrderResponse
	if err := mapstructure.Decode(res.Result, &ret); err != nil {
		return nil, fmt.Errorf("error decoding result: %s", err)
	}
	return &ret, nil
}

// PrivateGetWithdrawals makes a request to private/get_withdrawals
func (e *Exchange) PrivateGetWithdrawals(params *private.GetWithdrawalsRequest) (*private.GetWithdrawalsResponse, error) {
	res, err := e.rpcRequest("private/get_withdrawals", params)
	if err != nil {
		return nil, err
	}
	var ret private.GetWithdrawalsResponse
	if err := mapstructure.Decode(res.Result, &ret); err != nil {
		return nil, fmt.Errorf("error decoding result: %s", err)
	}
	return &ret, nil
}

// PrivateSell makes a request to private/sell
func (e *Exchange) PrivateSell(params *private.SellRequest) (*private.SellResponse, error) {
	res, err := e.rpcRequest("private/sell", params)
	if err != nil {
		return nil, err
	}
	var ret private.SellResponse
	if err := mapstructure.Decode(res.Result, &ret); err != nil {
		return nil, fmt.Errorf("error decoding result: %s", err)
	}
	return &ret, nil
}

// PrivateSetAnnouncementAsRead makes a request to private/set_announcement_as_read
func (e *Exchange) PrivateSetAnnouncementAsRead(params *private.SetAnnouncementAsReadRequest) (*private.SetAnnouncementAsReadResponse, error) {
	res, err := e.rpcRequest("private/set_announcement_as_read", params)
	if err != nil {
		return nil, err
	}
	var ret private.SetAnnouncementAsReadResponse
	if err := mapstructure.Decode(res.Result, &ret); err != nil {
		return nil, fmt.Errorf("error decoding result: %s", err)
	}
	return &ret, nil
}

// PrivateSetEmailForSubaccount makes a request to private/set_email_for_subaccount
func (e *Exchange) PrivateSetEmailForSubaccount(params *private.SetEmailForSubaccountRequest) (*private.SetEmailForSubaccountResponse, error) {
	res, err := e.rpcRequest("private/set_email_for_subaccount", params)
	if err != nil {
		return nil, err
	}
	var ret private.SetEmailForSubaccountResponse
	if err := mapstructure.Decode(res.Result, &ret); err != nil {
		return nil, fmt.Errorf("error decoding result: %s", err)
	}
	return &ret, nil
}

// PrivateSetEmailLanguage makes a request to private/set_email_language
func (e *Exchange) PrivateSetEmailLanguage(params *private.SetEmailLanguageRequest) (*private.SetEmailLanguageResponse, error) {
	res, err := e.rpcRequest("private/set_email_language", params)
	if err != nil {
		return nil, err
	}
	var ret private.SetEmailLanguageResponse
	if err := mapstructure.Decode(res.Result, &ret); err != nil {
		return nil, fmt.Errorf("error decoding result: %s", err)
	}
	return &ret, nil
}

// PrivateSetPasswordForSubaccount makes a request to private/set_password_for_subaccount
func (e *Exchange) PrivateSetPasswordForSubaccount(params *private.SetPasswordForSubaccountRequest) (*private.SetPasswordForSubaccountResponse, error) {
	res, err := e.rpcRequest("private/set_password_for_subaccount", params)
	if err != nil {
		return nil, err
	}
	var ret private.SetPasswordForSubaccountResponse
	if err := mapstructure.Decode(res.Result, &ret); err != nil {
		return nil, fmt.Errorf("error decoding result: %s", err)
	}
	return &ret, nil
}

// PrivateSubmitTransferToSubaccount makes a request to private/submit_transfer_to_subaccount
func (e *Exchange) PrivateSubmitTransferToSubaccount(params *private.SubmitTransferToSubaccountRequest) (*private.SubmitTransferToSubaccountResponse, error) {
	res, err := e.rpcRequest("private/submit_transfer_to_subaccount", params)
	if err != nil {
		return nil, err
	}
	var ret private.SubmitTransferToSubaccountResponse
	if err := mapstructure.Decode(res.Result, &ret); err != nil {
		return nil, fmt.Errorf("error decoding result: %s", err)
	}
	return &ret, nil
}

// PrivateSubmitTransferToUser makes a request to private/submit_transfer_to_user
func (e *Exchange) PrivateSubmitTransferToUser(params *private.SubmitTransferToUserRequest) (*private.SubmitTransferToUserResponse, error) {
	res, err := e.rpcRequest("private/submit_transfer_to_user", params)
	if err != nil {
		return nil, err
	}
	var ret private.SubmitTransferToUserResponse
	if err := mapstructure.Decode(res.Result, &ret); err != nil {
		return nil, fmt.Errorf("error decoding result: %s", err)
	}
	return &ret, nil
}

// PrivateSubscribe makes a request to private/subscribe
func (e *Exchange) PrivateSubscribe(params *private.SubscribeRequest) (*private.SubscribeResponse, error) {
	res, err := e.rpcRequest("private/subscribe", params)
	if err != nil {
		return nil, err
	}
	var ret private.SubscribeResponse
	if err := mapstructure.Decode(res.Result, &ret); err != nil {
		return nil, fmt.Errorf("error decoding result: %s", err)
	}
	return &ret, nil
}

// PrivateToggleNotificationsFromSubaccount makes a request to private/toggle_notifications_from_subaccount
func (e *Exchange) PrivateToggleNotificationsFromSubaccount(params *private.ToggleNotificationsFromSubaccountRequest) (*private.ToggleNotificationsFromSubaccountResponse, error) {
	res, err := e.rpcRequest("private/toggle_notifications_from_subaccount", params)
	if err != nil {
		return nil, err
	}
	var ret private.ToggleNotificationsFromSubaccountResponse
	if err := mapstructure.Decode(res.Result, &ret); err != nil {
		return nil, fmt.Errorf("error decoding result: %s", err)
	}
	return &ret, nil
}

// PrivateToggleSubaccountLogin makes a request to private/toggle_subaccount_login
func (e *Exchange) PrivateToggleSubaccountLogin(params *private.ToggleSubaccountLoginRequest) (*private.ToggleSubaccountLoginResponse, error) {
	res, err := e.rpcRequest("private/toggle_subaccount_login", params)
	if err != nil {
		return nil, err
	}
	var ret private.ToggleSubaccountLoginResponse
	if err := mapstructure.Decode(res.Result, &ret); err != nil {
		return nil, fmt.Errorf("error decoding result: %s", err)
	}
	return &ret, nil
}

// PrivateUnsubscribe makes a request to private/unsubscribe
func (e *Exchange) PrivateUnsubscribe(params *private.UnsubscribeRequest) (*private.UnsubscribeResponse, error) {
	res, err := e.rpcRequest("private/unsubscribe", params)
	if err != nil {
		return nil, err
	}
	var ret private.UnsubscribeResponse
	if err := mapstructure.Decode(res.Result, &ret); err != nil {
		return nil, fmt.Errorf("error decoding result: %s", err)
	}
	return &ret, nil
}

// PrivateWithdraw makes a request to private/withdraw
func (e *Exchange) PrivateWithdraw(params *private.WithdrawRequest) (*private.WithdrawResponse, error) {
	res, err := e.rpcRequest("private/withdraw", params)
	if err != nil {
		return nil, err
	}
	var ret private.WithdrawResponse
	if err := mapstructure.Decode(res.Result, &ret); err != nil {
		return nil, fmt.Errorf("error decoding result: %s", err)
	}
	return &ret, nil
}

// PublicAuth makes a request to public/auth
func (e *Exchange) PublicAuth(params *public.AuthRequest) (*public.AuthResponse, error) {
	res, err := e.rpcRequest("public/auth", params)
	if err != nil {
		return nil, err
	}
	var ret public.AuthResponse
	if err := mapstructure.Decode(res.Result, &ret); err != nil {
		return nil, fmt.Errorf("error decoding result: %s", err)
	}
	return &ret, nil
}

// PublicDisableHeartbeat makes a request to public/disable_heartbeat
func (e *Exchange) PublicDisableHeartbeat(params *public.DisableHeartbeatRequest) (*public.DisableHeartbeatResponse, error) {
	res, err := e.rpcRequest("public/disable_heartbeat", params)
	if err != nil {
		return nil, err
	}
	var ret public.DisableHeartbeatResponse
	if err := mapstructure.Decode(res.Result, &ret); err != nil {
		return nil, fmt.Errorf("error decoding result: %s", err)
	}
	return &ret, nil
}

// PublicGetAnnouncements makes a request to public/get_announcements
func (e *Exchange) PublicGetAnnouncements(params *public.GetAnnouncementsRequest) (*public.GetAnnouncementsResponse, error) {
	res, err := e.rpcRequest("public/get_announcements", params)
	if err != nil {
		return nil, err
	}
	var ret public.GetAnnouncementsResponse
	if err := mapstructure.Decode(res.Result, &ret); err != nil {
		return nil, fmt.Errorf("error decoding result: %s", err)
	}
	return &ret, nil
}

// PublicGetBookSummaryByCurrency makes a request to public/get_book_summary_by_currency
func (e *Exchange) PublicGetBookSummaryByCurrency(params *public.GetBookSummaryByCurrencyRequest) (*public.GetBookSummaryByCurrencyResponse, error) {
	res, err := e.rpcRequest("public/get_book_summary_by_currency", params)
	if err != nil {
		return nil, err
	}
	var ret public.GetBookSummaryByCurrencyResponse
	if err := mapstructure.Decode(res.Result, &ret); err != nil {
		return nil, fmt.Errorf("error decoding result: %s", err)
	}
	return &ret, nil
}

// PublicGetBookSummaryByInstrument makes a request to public/get_book_summary_by_instrument
func (e *Exchange) PublicGetBookSummaryByInstrument(params *public.GetBookSummaryByInstrumentRequest) (*public.GetBookSummaryByInstrumentResponse, error) {
	res, err := e.rpcRequest("public/get_book_summary_by_instrument", params)
	if err != nil {
		return nil, err
	}
	var ret public.GetBookSummaryByInstrumentResponse
	if err := mapstructure.Decode(res.Result, &ret); err != nil {
		return nil, fmt.Errorf("error decoding result: %s", err)
	}
	return &ret, nil
}

// PublicGetCurrencies makes a request to public/get_currencies
func (e *Exchange) PublicGetCurrencies(params *public.GetCurrenciesRequest) (*public.GetCurrenciesResponse, error) {
	res, err := e.rpcRequest("public/get_currencies", params)
	if err != nil {
		return nil, err
	}
	var ret public.GetCurrenciesResponse
	if err := mapstructure.Decode(res.Result, &ret); err != nil {
		return nil, fmt.Errorf("error decoding result: %s", err)
	}
	return &ret, nil
}

// PublicGetHistoricalVolatility makes a request to public/get_historical_volatility
func (e *Exchange) PublicGetHistoricalVolatility(params *public.GetHistoricalVolatilityRequest) (*public.GetHistoricalVolatilityResponse, error) {
	res, err := e.rpcRequest("public/get_historical_volatility", params)
	if err != nil {
		return nil, err
	}
	var ret public.GetHistoricalVolatilityResponse
	if err := mapstructure.Decode(res.Result, &ret); err != nil {
		return nil, fmt.Errorf("error decoding result: %s", err)
	}
	return &ret, nil
}

// PublicGetInstruments makes a request to public/get_instruments
func (e *Exchange) PublicGetInstruments(params *public.GetInstrumentsRequest) (*public.GetInstrumentsResponse, error) {
	res, err := e.rpcRequest("public/get_instruments", params)
	if err != nil {
		return nil, err
	}
	var ret public.GetInstrumentsResponse
	if err := mapstructure.Decode(res.Result, &ret); err != nil {
		return nil, fmt.Errorf("error decoding result: %s", err)
	}
	return &ret, nil
}

// PublicGetLastSettlementsByCurrency makes a request to public/get_last_settlements_by_currency
func (e *Exchange) PublicGetLastSettlementsByCurrency(params *public.GetLastSettlementsByCurrencyRequest) (*public.GetLastSettlementsByCurrencyResponse, error) {
	res, err := e.rpcRequest("public/get_last_settlements_by_currency", params)
	if err != nil {
		return nil, err
	}
	var ret public.GetLastSettlementsByCurrencyResponse
	if err := mapstructure.Decode(res.Result, &ret); err != nil {
		return nil, fmt.Errorf("error decoding result: %s", err)
	}
	return &ret, nil
}

// PublicGetLastSettlementsByInstrument makes a request to public/get_last_settlements_by_instrument
func (e *Exchange) PublicGetLastSettlementsByInstrument(params *public.GetLastSettlementsByInstrumentRequest) (*public.GetLastSettlementsByInstrumentResponse, error) {
	res, err := e.rpcRequest("public/get_last_settlements_by_instrument", params)
	if err != nil {
		return nil, err
	}
	var ret public.GetLastSettlementsByInstrumentResponse
	if err := mapstructure.Decode(res.Result, &ret); err != nil {
		return nil, fmt.Errorf("error decoding result: %s", err)
	}
	return &ret, nil
}

// PublicGetLastTradesByCurrency makes a request to public/get_last_trades_by_currency
func (e *Exchange) PublicGetLastTradesByCurrency(params *public.GetLastTradesByCurrencyRequest) (*public.GetLastTradesByCurrencyResponse, error) {
	res, err := e.rpcRequest("public/get_last_trades_by_currency", params)
	if err != nil {
		return nil, err
	}
	var ret public.GetLastTradesByCurrencyResponse
	if err := mapstructure.Decode(res.Result, &ret); err != nil {
		return nil, fmt.Errorf("error decoding result: %s", err)
	}
	return &ret, nil
}

// PublicGetLastTradesByCurrencyAndTime makes a request to public/get_last_trades_by_currency_and_time
func (e *Exchange) PublicGetLastTradesByCurrencyAndTime(params *public.GetLastTradesByCurrencyAndTimeRequest) (*public.GetLastTradesByCurrencyAndTimeResponse, error) {
	res, err := e.rpcRequest("public/get_last_trades_by_currency_and_time", params)
	if err != nil {
		return nil, err
	}
	var ret public.GetLastTradesByCurrencyAndTimeResponse
	if err := mapstructure.Decode(res.Result, &ret); err != nil {
		return nil, fmt.Errorf("error decoding result: %s", err)
	}
	return &ret, nil
}

// PublicGetLastTradesByInstrument makes a request to public/get_last_trades_by_instrument
func (e *Exchange) PublicGetLastTradesByInstrument(params *public.GetLastTradesByInstrumentRequest) (*public.GetLastTradesByInstrumentResponse, error) {
	res, err := e.rpcRequest("public/get_last_trades_by_instrument", params)
	if err != nil {
		return nil, err
	}
	var ret public.GetLastTradesByInstrumentResponse
	if err := mapstructure.Decode(res.Result, &ret); err != nil {
		return nil, fmt.Errorf("error decoding result: %s", err)
	}
	return &ret, nil
}

// PublicGetLastTradesByInstrumentAndTime makes a request to public/get_last_trades_by_instrument_and_time
func (e *Exchange) PublicGetLastTradesByInstrumentAndTime(params *public.GetLastTradesByInstrumentAndTimeRequest) (*public.GetLastTradesByInstrumentAndTimeResponse, error) {
	res, err := e.rpcRequest("public/get_last_trades_by_instrument_and_time", params)
	if err != nil {
		return nil, err
	}
	var ret public.GetLastTradesByInstrumentAndTimeResponse
	if err := mapstructure.Decode(res.Result, &ret); err != nil {
		return nil, fmt.Errorf("error decoding result: %s", err)
	}
	return &ret, nil
}

// PublicGetOptionMarkPrices makes a request to public/get_option_mark_prices
func (e *Exchange) PublicGetOptionMarkPrices(params *public.GetOptionMarkPricesRequest) (*public.GetOptionMarkPricesResponse, error) {
	res, err := e.rpcRequest("public/get_option_mark_prices", params)
	if err != nil {
		return nil, err
	}
	var ret public.GetOptionMarkPricesResponse
	if err := mapstructure.Decode(res.Result, &ret); err != nil {
		return nil, fmt.Errorf("error decoding result: %s", err)
	}
	return &ret, nil
}

// PublicGetOrderBook makes a request to public/get_order_book
func (e *Exchange) PublicGetOrderBook(params *public.GetOrderBookRequest) (*public.GetOrderBookResponse, error) {
	res, err := e.rpcRequest("public/get_order_book", params)
	if err != nil {
		return nil, err
	}
	var ret public.GetOrderBookResponse
	if err := mapstructure.Decode(res.Result, &ret); err != nil {
		return nil, fmt.Errorf("error decoding result: %s", err)
	}
	return &ret, nil
}

// PublicGetTime makes a request to public/get_time
func (e *Exchange) PublicGetTime(params *public.GetTimeRequest) (*public.GetTimeResponse, error) {
	res, err := e.rpcRequest("public/get_time", params)
	if err != nil {
		return nil, err
	}
	var ret public.GetTimeResponse
	if err := mapstructure.Decode(res.Result, &ret); err != nil {
		return nil, fmt.Errorf("error decoding result: %s", err)
	}
	return &ret, nil
}

// PublicGetTradeVolumes makes a request to public/get_trade_volumes
func (e *Exchange) PublicGetTradeVolumes(params *public.GetTradeVolumesRequest) (*public.GetTradeVolumesResponse, error) {
	res, err := e.rpcRequest("public/get_trade_volumes", params)
	if err != nil {
		return nil, err
	}
	var ret public.GetTradeVolumesResponse
	if err := mapstructure.Decode(res.Result, &ret); err != nil {
		return nil, fmt.Errorf("error decoding result: %s", err)
	}
	return &ret, nil
}

// PublicHello makes a request to public/hello
func (e *Exchange) PublicHello(params *public.HelloRequest) (*public.HelloResponse, error) {
	res, err := e.rpcRequest("public/hello", params)
	if err != nil {
		return nil, err
	}
	var ret public.HelloResponse
	if err := mapstructure.Decode(res.Result, &ret); err != nil {
		return nil, fmt.Errorf("error decoding result: %s", err)
	}
	return &ret, nil
}

// PublicSetHeartbeat makes a request to public/set_heartbeat
func (e *Exchange) PublicSetHeartbeat(params *public.SetHeartbeatRequest) (*public.SetHeartbeatResponse, error) {
	res, err := e.rpcRequest("public/set_heartbeat", params)
	if err != nil {
		return nil, err
	}
	var ret public.SetHeartbeatResponse
	if err := mapstructure.Decode(res.Result, &ret); err != nil {
		return nil, fmt.Errorf("error decoding result: %s", err)
	}
	return &ret, nil
}

// PublicSubscribe makes a request to public/subscribe
func (e *Exchange) PublicSubscribe(params *public.SubscribeRequest) (*public.SubscribeResponse, error) {
	res, err := e.rpcRequest("public/subscribe", params)
	if err != nil {
		return nil, err
	}
	var ret public.SubscribeResponse
	if err := mapstructure.Decode(res.Result, &ret); err != nil {
		return nil, fmt.Errorf("error decoding result: %s", err)
	}
	return &ret, nil
}

// PublicTest makes a request to public/test
func (e *Exchange) PublicTest(params *public.TestRequest) (*public.TestResponse, error) {
	res, err := e.rpcRequest("public/test", params)
	if err != nil {
		return nil, err
	}
	var ret public.TestResponse
	if err := mapstructure.Decode(res.Result, &ret); err != nil {
		return nil, fmt.Errorf("error decoding result: %s", err)
	}
	return &ret, nil
}

// PublicTicker makes a request to public/ticker
func (e *Exchange) PublicTicker(params *public.TickerRequest) (*public.TickerResponse, error) {
	res, err := e.rpcRequest("public/ticker", params)
	if err != nil {
		return nil, err
	}
	var ret public.TickerResponse
	if err := mapstructure.Decode(res.Result, &ret); err != nil {
		return nil, fmt.Errorf("error decoding result: %s", err)
	}
	return &ret, nil
}

// PublicUnsubscribe makes a request to public/unsubscribe
func (e *Exchange) PublicUnsubscribe(params *public.UnsubscribeRequest) (*public.UnsubscribeResponse, error) {
	res, err := e.rpcRequest("public/unsubscribe", params)
	if err != nil {
		return nil, err
	}
	var ret public.UnsubscribeResponse
	if err := mapstructure.Decode(res.Result, &ret); err != nil {
		return nil, fmt.Errorf("error decoding result: %s", err)
	}
	return &ret, nil
}