package deribit

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/adampointer/go-deribit/v3/client/operations"
	"github.com/iancoleman/strcase"
	"github.com/stretchr/testify/assert"
)

type testCase struct {
	method string
	f      func(*operations.Client, []byte) ([]byte, error)
}

var publicMethods = []*testCase{
	{
		"DisableHeartbeat",
		func(c *operations.Client, req []byte) ([]byte, error) {
			var params = struct {
				Params operations.GetPublicDisableHeartbeatParams
			}{}
			if err := json.Unmarshal(req, &params); err != nil {
				return nil, err
			}
			res, err := c.GetPublicDisableHeartbeat(&params.Params)
			if err != nil {
				return nil, err
			}
			return json.Marshal(res.Payload)
		},
	},
	{
		"GetAnnouncements",
		func(c *operations.Client, req []byte) ([]byte, error) {
			var params = struct {
				Params operations.GetPublicGetAnnouncementsParams
			}{}
			if err := json.Unmarshal(req, &params); err != nil {
				return nil, err
			}
			res, err := c.GetPublicGetAnnouncements(&params.Params)
			if err != nil {
				return nil, err
			}
			return json.Marshal(res.Payload)
		},
	},
	{
		"GetBookSummaryByCurrency",
		func(c *operations.Client, req []byte) ([]byte, error) {
			var params = struct {
				Params operations.GetPublicGetBookSummaryByCurrencyParams
			}{}
			if err := json.Unmarshal(req, &params); err != nil {
				return nil, err
			}
			res, err := c.GetPublicGetBookSummaryByCurrency(&params.Params)
			if err != nil {
				return nil, err
			}
			return json.Marshal(res.Payload)
		},
	},
	{
		"GetBookSummaryByInstrument",
		func(c *operations.Client, req []byte) ([]byte, error) {
			var params = struct {
				Params operations.GetPublicGetBookSummaryByInstrumentParams
			}{}
			if err := json.Unmarshal(req, &params); err != nil {
				return nil, err
			}
			res, err := c.GetPublicGetBookSummaryByInstrument(&params.Params)
			if err != nil {
				return nil, err
			}
			return json.Marshal(res.Payload)
		},
	},
	{
		"GetCurrencies",
		func(c *operations.Client, req []byte) ([]byte, error) {
			var params = struct {
				Params operations.GetPublicGetCurrenciesParams
			}{}
			if err := json.Unmarshal(req, &params); err != nil {
				return nil, err
			}
			res, err := c.GetPublicGetCurrencies(&params.Params)
			if err != nil {
				return nil, err
			}
			return json.Marshal(res.Payload)
		},
	},
	{
		"GetHistoricalVolatility",
		func(c *operations.Client, req []byte) ([]byte, error) {
			var params = struct {
				Params operations.GetPublicGetHistoricalVolatilityParams
			}{}
			if err := json.Unmarshal(req, &params); err != nil {
				return nil, err
			}
			res, err := c.GetPublicGetHistoricalVolatility(&params.Params)
			if err != nil {
				return nil, err
			}
			return json.Marshal(res.Payload)
		},
	},
	{
		"GetInstruments",
		func(c *operations.Client, req []byte) ([]byte, error) {
			var params = struct {
				Params operations.GetPublicGetInstrumentsParams
			}{}
			if err := json.Unmarshal(req, &params); err != nil {
				return nil, err
			}
			res, err := c.GetPublicGetInstruments(&params.Params)
			if err != nil {
				return nil, err
			}
			return json.Marshal(res.Payload)
		},
	},
	{
		"GetLastSettlementsByCurrency",
		func(c *operations.Client, req []byte) ([]byte, error) {
			var params = struct {
				Params operations.GetPublicGetLastSettlementsByCurrencyParams
			}{}
			if err := json.Unmarshal(req, &params); err != nil {
				return nil, err
			}
			res, err := c.GetPublicGetLastSettlementsByCurrency(&params.Params)
			if err != nil {
				return nil, err
			}
			return json.Marshal(res.Payload)
		},
	},
	{
		"GetLastSettlementsByInstrument",
		func(c *operations.Client, req []byte) ([]byte, error) {
			var params = struct {
				Params operations.GetPublicGetLastSettlementsByInstrumentParams
			}{}
			if err := json.Unmarshal(req, &params); err != nil {
				return nil, err
			}
			res, err := c.GetPublicGetLastSettlementsByInstrument(&params.Params)
			if err != nil {
				return nil, err
			}
			return json.Marshal(res.Payload)
		},
	},
	{
		"GetLastTradesByCurrency",
		func(c *operations.Client, req []byte) ([]byte, error) {
			var params = struct {
				Params operations.GetPublicGetLastTradesByCurrencyParams
			}{}
			if err := json.Unmarshal(req, &params); err != nil {
				return nil, err
			}
			res, err := c.GetPublicGetLastTradesByCurrency(&params.Params)
			if err != nil {
				return nil, err
			}
			return json.Marshal(res.Payload)
		},
	},
	{
		"GetLastTradesByCurrencyAndTime",
		func(c *operations.Client, req []byte) ([]byte, error) {
			var params = struct {
				Params operations.GetPublicGetLastTradesByCurrencyAndTimeParams
			}{}
			if err := json.Unmarshal(req, &params); err != nil {
				return nil, err
			}
			res, err := c.GetPublicGetLastTradesByCurrencyAndTime(&params.Params)
			if err != nil {
				return nil, err
			}
			return json.Marshal(res.Payload)
		},
	},
	{
		"GetLastTradesByInstrument",
		func(c *operations.Client, req []byte) ([]byte, error) {
			var params = struct {
				Params operations.GetPublicGetLastTradesByInstrumentParams
			}{}
			if err := json.Unmarshal(req, &params); err != nil {
				return nil, err
			}
			res, err := c.GetPublicGetLastTradesByInstrument(&params.Params)
			if err != nil {
				return nil, err
			}
			return json.Marshal(res.Payload)
		},
	},
	{
		"GetLastTradesByInstrumentAndTime",
		func(c *operations.Client, req []byte) ([]byte, error) {
			var params = struct {
				Params operations.GetPublicGetLastTradesByInstrumentAndTimeParams
			}{}
			if err := json.Unmarshal(req, &params); err != nil {
				return nil, err
			}
			res, err := c.GetPublicGetLastTradesByInstrumentAndTime(&params.Params)
			if err != nil {
				return nil, err
			}
			return json.Marshal(res.Payload)
		},
	},
	{
		"GetOrderBook",
		func(c *operations.Client, req []byte) ([]byte, error) {
			var params = struct {
				Params operations.GetPublicGetOrderBookParams
			}{}
			if err := json.Unmarshal(req, &params); err != nil {
				return nil, err
			}
			res, err := c.GetPublicGetOrderBook(&params.Params)
			if err != nil {
				return nil, err
			}
			return json.Marshal(res.Payload)
		},
	},
	{
		"GetTime",
		func(c *operations.Client, req []byte) ([]byte, error) {
			var params = struct {
				Params operations.GetPublicGetTimeParams
			}{}
			if err := json.Unmarshal(req, &params); err != nil {
				return nil, err
			}
			res, err := c.GetPublicGetTime(&params.Params)
			if err != nil {
				return nil, err
			}
			return json.Marshal(res.Payload)
		},
	},
	{
		"GetTradeVolumes",
		func(c *operations.Client, req []byte) ([]byte, error) {
			var params = struct {
				Params operations.GetPublicGetTradeVolumesParams
			}{}
			if err := json.Unmarshal(req, &params); err != nil {
				return nil, err
			}
			res, err := c.GetPublicGetTradeVolumes(&params.Params)
			if err != nil {
				return nil, err
			}
			return json.Marshal(res.Payload)
		},
	},
	{
		"GetTradingviewChartData",
		func(c *operations.Client, req []byte) ([]byte, error) {
			var params = struct {
				Params operations.GetPublicGetTradingviewChartDataParams
			}{}
			if err := json.Unmarshal(req, &params); err != nil {
				return nil, err
			}
			res, err := c.GetPublicGetTradingviewChartData(&params.Params)
			if err != nil {
				return nil, err
			}
			return json.Marshal(res.Payload)
		},
	},
	{
		"Hello",
		func(c *operations.Client, req []byte) ([]byte, error) {
			var params = struct {
				Params operations.GetPublicHelloParams
			}{}
			if err := json.Unmarshal(req, &params); err != nil {
				return nil, err
			}
			res, err := c.GetPublicHello(&params.Params)
			if err != nil {
				return nil, err
			}
			return json.Marshal(res.Payload)
		},
	},
	// This fails with invalid parameters but it is not clear why
	//{
	//	"SetHeartbeat",
	//	func(c *operations.Client, req []byte) ([]byte, error) {
	//		var params = struct {
	//			Params operations.GetPublicSetHeartbeatParams
	//		}{}
	//		if err := json.Unmarshal(req, &params); err != nil {
	//			return nil, err
	//		}
	//		spew.Dump(params)
	//		res, err := c.GetPublicSetHeartbeat(&params.Params)
	//		if err != nil {
	//			return nil, err
	//		}
	//		return json.Marshal(res.Payload)
	//	},
	//},
	{
		"Test",
		func(c *operations.Client, req []byte) ([]byte, error) {
			var params = struct {
				Params operations.GetPublicTestParams
			}{}
			if err := json.Unmarshal(req, &params); err != nil {
				return nil, err
			}
			res, err := c.GetPublicTest(&params.Params)
			if err != nil {
				return nil, err
			}
			return json.Marshal(res.Payload)
		},
	},
	{
		"Ticker",
		func(c *operations.Client, req []byte) ([]byte, error) {
			var params = struct {
				Params operations.GetPublicTickerParams
			}{}
			if err := json.Unmarshal(req, &params); err != nil {
				return nil, err
			}
			res, err := c.GetPublicTicker(&params.Params)
			if err != nil {
				return nil, err
			}
			return json.Marshal(res.Payload)
		},
	},
	{
		"Unsubscribe",
		func(c *operations.Client, req []byte) ([]byte, error) {
			var params = struct {
				Params operations.GetPublicUnsubscribeParams
			}{}
			if err := json.Unmarshal(req, &params); err != nil {
				return nil, err
			}
			res, err := c.GetPublicUnsubscribe(&params.Params)
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
