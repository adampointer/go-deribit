package deribit

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/adampointer/go-deribit/v3/client/account_management"
	"github.com/adampointer/go-deribit/v3/client/market_data"
	"github.com/adampointer/go-deribit/v3/client/public"
	"github.com/adampointer/go-deribit/v3/client/supporting"

	"github.com/adampointer/go-deribit/v3/client"
	"github.com/adampointer/go-deribit/v3/client/websocket_only"
	"github.com/iancoleman/strcase"
	"github.com/stretchr/testify/assert"
)

type testCase struct {
	method string
	f      func(*client.Deribit, []byte) ([]byte, error)
}

var publicMethods = []*testCase{
	{
		"DisableHeartbeat",
		func(c *client.Deribit, req []byte) ([]byte, error) {
			var params = struct {
				Params websocket_only.GetPublicDisableHeartbeatParams
			}{}
			if err := json.Unmarshal(req, &params); err != nil {
				return nil, err
			}
			res, err := c.WebsocketOnly.GetPublicDisableHeartbeat(&params.Params)
			if err != nil {
				return nil, err
			}
			return json.Marshal(res.Payload)
		},
	},
	{
		"GetAnnouncements",
		func(c *client.Deribit, req []byte) ([]byte, error) {
			var params = struct {
				Params account_management.GetPublicGetAnnouncementsParams
			}{}
			if err := json.Unmarshal(req, &params); err != nil {
				return nil, err
			}
			res, err := c.AccountManagement.GetPublicGetAnnouncements(&params.Params)
			if err != nil {
				return nil, err
			}
			return json.Marshal(res.Payload)
		},
	},
	{
		"GetBookSummaryByCurrency",
		func(c *client.Deribit, req []byte) ([]byte, error) {
			var params = struct {
				Params market_data.GetPublicGetBookSummaryByCurrencyParams
			}{}
			if err := json.Unmarshal(req, &params); err != nil {
				return nil, err
			}
			res, err := c.MarketData.GetPublicGetBookSummaryByCurrency(&params.Params)
			if err != nil {
				return nil, err
			}
			return json.Marshal(res.Payload)
		},
	},
	{
		"GetBookSummaryByInstrument",
		func(c *client.Deribit, req []byte) ([]byte, error) {
			var params = struct {
				Params market_data.GetPublicGetBookSummaryByInstrumentParams
			}{}
			if err := json.Unmarshal(req, &params); err != nil {
				return nil, err
			}
			res, err := c.MarketData.GetPublicGetBookSummaryByInstrument(&params.Params)
			if err != nil {
				return nil, err
			}
			return json.Marshal(res.Payload)
		},
	},
	{
		"GetCurrencies",
		func(c *client.Deribit, req []byte) ([]byte, error) {
			var params = struct {
				Params public.GetPublicGetCurrenciesParams
			}{}
			if err := json.Unmarshal(req, &params); err != nil {
				return nil, err
			}
			res, err := c.Public.GetPublicGetCurrencies(&params.Params)
			if err != nil {
				return nil, err
			}
			return json.Marshal(res.Payload)
		},
	},
	{
		"GetHistoricalVolatility",
		func(c *client.Deribit, req []byte) ([]byte, error) {
			var params = struct {
				Params public.GetPublicGetHistoricalVolatilityParams
			}{}
			if err := json.Unmarshal(req, &params); err != nil {
				return nil, err
			}
			res, err := c.Public.GetPublicGetHistoricalVolatility(&params.Params)
			if err != nil {
				return nil, err
			}
			return json.Marshal(res.Payload)
		},
	},
	{
		"GetInstruments",
		func(c *client.Deribit, req []byte) ([]byte, error) {
			var params = struct {
				Params public.GetPublicGetInstrumentsParams
			}{}
			if err := json.Unmarshal(req, &params); err != nil {
				return nil, err
			}
			res, err := c.Public.GetPublicGetInstruments(&params.Params)
			if err != nil {
				return nil, err
			}
			return json.Marshal(res.Payload)
		},
	},
	{
		"GetLastSettlementsByCurrency",
		func(c *client.Deribit, req []byte) ([]byte, error) {
			var params = struct {
				Params market_data.GetPublicGetLastSettlementsByCurrencyParams
			}{}
			if err := json.Unmarshal(req, &params); err != nil {
				return nil, err
			}
			res, err := c.MarketData.GetPublicGetLastSettlementsByCurrency(&params.Params)
			if err != nil {
				return nil, err
			}
			return json.Marshal(res.Payload)
		},
	},
	{
		"GetLastSettlementsByInstrument",
		func(c *client.Deribit, req []byte) ([]byte, error) {
			var params = struct {
				Params market_data.GetPublicGetLastSettlementsByInstrumentParams
			}{}
			if err := json.Unmarshal(req, &params); err != nil {
				return nil, err
			}
			res, err := c.MarketData.GetPublicGetLastSettlementsByInstrument(&params.Params)
			if err != nil {
				return nil, err
			}
			return json.Marshal(res.Payload)
		},
	},
	{
		"GetLastTradesByCurrency",
		func(c *client.Deribit, req []byte) ([]byte, error) {
			var params = struct {
				Params market_data.GetPublicGetLastTradesByCurrencyParams
			}{}
			if err := json.Unmarshal(req, &params); err != nil {
				return nil, err
			}
			res, err := c.MarketData.GetPublicGetLastTradesByCurrency(&params.Params)
			if err != nil {
				return nil, err
			}
			return json.Marshal(res.Payload)
		},
	},
	{
		"GetLastTradesByCurrencyAndTime",
		func(c *client.Deribit, req []byte) ([]byte, error) {
			var params = struct {
				Params market_data.GetPublicGetLastTradesByCurrencyAndTimeParams
			}{}
			if err := json.Unmarshal(req, &params); err != nil {
				return nil, err
			}
			res, err := c.MarketData.GetPublicGetLastTradesByCurrencyAndTime(&params.Params)
			if err != nil {
				return nil, err
			}
			return json.Marshal(res.Payload)
		},
	},
	{
		"GetLastTradesByInstrument",
		func(c *client.Deribit, req []byte) ([]byte, error) {
			var params = struct {
				Params market_data.GetPublicGetLastTradesByInstrumentParams
			}{}
			if err := json.Unmarshal(req, &params); err != nil {
				return nil, err
			}
			res, err := c.MarketData.GetPublicGetLastTradesByInstrument(&params.Params)
			if err != nil {
				return nil, err
			}
			return json.Marshal(res.Payload)
		},
	},
	{
		"GetLastTradesByInstrumentAndTime",
		func(c *client.Deribit, req []byte) ([]byte, error) {
			var params = struct {
				Params market_data.GetPublicGetLastTradesByInstrumentAndTimeParams
			}{}
			if err := json.Unmarshal(req, &params); err != nil {
				return nil, err
			}
			res, err := c.MarketData.GetPublicGetLastTradesByInstrumentAndTime(&params.Params)
			if err != nil {
				return nil, err
			}
			return json.Marshal(res.Payload)
		},
	},
	{
		"GetOrderBook",
		func(c *client.Deribit, req []byte) ([]byte, error) {
			var params = struct {
				Params market_data.GetPublicGetOrderBookParams
			}{}
			if err := json.Unmarshal(req, &params); err != nil {
				return nil, err
			}
			res, err := c.MarketData.GetPublicGetOrderBook(&params.Params)
			if err != nil {
				return nil, err
			}
			return json.Marshal(res.Payload)
		},
	},
	{
		"GetTime",
		func(c *client.Deribit, req []byte) ([]byte, error) {
			var params = struct {
				Params supporting.GetPublicGetTimeParams
			}{}
			if err := json.Unmarshal(req, &params); err != nil {
				return nil, err
			}
			res, err := c.Supporting.GetPublicGetTime(&params.Params)
			if err != nil {
				return nil, err
			}
			return json.Marshal(res.Payload)
		},
	},
	{
		"GetTradeVolumes",
		func(c *client.Deribit, req []byte) ([]byte, error) {
			var params = struct {
				Params market_data.GetPublicGetTradeVolumesParams
			}{}
			if err := json.Unmarshal(req, &params); err != nil {
				return nil, err
			}
			res, err := c.MarketData.GetPublicGetTradeVolumes(&params.Params)
			if err != nil {
				return nil, err
			}
			return json.Marshal(res.Payload)
		},
	},
	{
		"GetTradingviewChartData",
		func(c *client.Deribit, req []byte) ([]byte, error) {
			var params = struct {
				Params public.GetPublicGetTradingviewChartDataParams
			}{}
			if err := json.Unmarshal(req, &params); err != nil {
				return nil, err
			}
			res, err := c.Public.GetPublicGetTradingviewChartData(&params.Params)
			if err != nil {
				return nil, err
			}
			return json.Marshal(res.Payload)
		},
	},
	{
		"Hello",
		func(c *client.Deribit, req []byte) ([]byte, error) {
			var params = struct {
				Params public.GetPublicHelloParams
			}{}
			if err := json.Unmarshal(req, &params); err != nil {
				return nil, err
			}
			res, err := c.Public.GetPublicHello(&params.Params)
			if err != nil {
				return nil, err
			}
			return json.Marshal(res.Payload)
		},
	},
	{
		"Test",
		func(c *client.Deribit, req []byte) ([]byte, error) {
			var params = struct {
				Params public.GetPublicTestParams
			}{}
			if err := json.Unmarshal(req, &params); err != nil {
				return nil, err
			}
			res, err := c.Public.GetPublicTest(&params.Params)
			if err != nil {
				return nil, err
			}
			return json.Marshal(res.Payload)
		},
	},
	{
		"Ticker",
		func(c *client.Deribit, req []byte) ([]byte, error) {
			var params = struct {
				Params public.GetPublicTickerParams
			}{}
			if err := json.Unmarshal(req, &params); err != nil {
				return nil, err
			}
			res, err := c.Public.GetPublicTicker(&params.Params)
			if err != nil {
				return nil, err
			}
			return json.Marshal(res.Payload)
		},
	},
	{
		"Unsubscribe",
		func(c *client.Deribit, req []byte) ([]byte, error) {
			var params = struct {
				Params websocket_only.GetPublicUnsubscribeParams
			}{}
			if err := json.Unmarshal(req, &params); err != nil {
				return nil, err
			}
			res, err := c.WebsocketOnly.GetPublicUnsubscribe(&params.Params)
			if err != nil {
				return nil, err
			}
			return json.Marshal(res.Payload)
		},
	},
}

func TestPublicMethods(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode")
	}

	errs := make(chan error)
	stop := make(chan bool)

	go func() {
		e := <-errs
		t.Fatalf("unexpected error: %s", e)
	}()

	conn, err := NewExchange(true, errs, stop)
	assert.Nil(t, err)
	assert.Nil(t, conn.Connect())

	for _, tt := range publicMethods {
		t.Run(tt.method, func(t *testing.T) {
			req, err := getExample(tt.method)
			assert.Nil(t, err)
			res, err := tt.f(conn.Client(), req)
			assert.Nil(t, err)
			t.Log(string(res))
		})
	}

	assert.Nil(t, conn.Close())
}

func getExample(method string) ([]byte, error) {
	file, err := os.Open(fmt.Sprintf("schema/examples/public/%s.request.json", strcase.ToSnake(method)))
	if err != nil {
		return nil, err
	}
	defer file.Close()
	b, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	return convertKeys(b), err
}

func convertKeys(j json.RawMessage) json.RawMessage {
	m := make(map[string]json.RawMessage)
	if err := json.Unmarshal([]byte(j), &m); err != nil {
		// Not a JSON object
		return j
	}

	for k, v := range m {
		fixed := fixKey(k)
		delete(m, k)
		m[fixed] = convertKeys(v)
	}

	b, err := json.Marshal(m)
	if err != nil {
		return j
	}

	return json.RawMessage(b)
}

func fixKey(key string) string {
	return strcase.ToCamel(key)
}
