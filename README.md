[![Build Status](https://travis-ci.com/adampointer/go-deribit.svg?branch=master)](https://travis-ci.com/adampointer/go-deribit)  [![Go Report Card](https://goreportcard.com/badge/github.com/adampointer/go-deribit)](https://goreportcard.com/report/github.com/adampointer/go-deribit)  [![codebeat badge](https://codebeat.co/badges/5bf32114-b7e1-4e70-91bf-fae2449fe2cb)](https://codebeat.co/projects/github-com-adampointer-go-deribit-master)

# go-deribit

## V3 

This project is now using go1.13 with Go Modules, but should remain compatible with `dep`. Also, as there are some breaking changes introduced by the latest schema changes from the remote API, I have decided to carry on development in the new `v3` namespace with the project root containing the code tagged `v2.x`.

`import "github.com/adampointer/go-deribit/v3"`

We now have the latest API methods which were recently released such as `public/get_tradingview_chart_data`.

I recommend using the `v3` project in your projects as all onward development will now be within this project.

[GoDoc API Documentation](https://godoc.org/github.com/adampointer/go-deribit/v3)

## Overview

Go library for using the [Deribit's](https://www.deribit.com/reg-3027.8327) **v2** Websocket API. 

Deribit is a modern, fast BitCoin derivatives exchange. 

This library is a port of the [official wrapper libraries](https://github.com/deribit) to Go.

If you wish to try it out, be kind and use my affiliate link [https://www.deribit.com/reg-3027.8327](https://www.deribit.com/reg-3027.8327)

Or tip me!

btc: 3HmLfHJvrJuM48zHFY6HstUCxbuwV3dvxd

eth: 0x9Dc9129185E79211534D0039Af1C6f1ff585F5e3

ltc: MEpFjCdR3uXd6QjuJTSu3coLtcSWY3S2Hg


*p.s.* If you want a good BitMEX client library then try [go-bitmex](https://github.com/adampointer/go-bitmex)

[V2 API Documentation](http://docs.deribit.com/v2/?javascript#deribit-api-v2-0-0)

## Example Usage

Look at `cmd/example/main.go`

```
make build
example --access-key XXX --secret-key YYYYYY
```

## Reconnect Behaviour

There is a heartbeat which triggers every 10 seconds to keep the websocket connection alive. In the event of a connection error, the library will automatically attempt to reconnect, re-authenticate and reestablish subscriptions.

This behaviour is overrideable with the `SetDisconnectHandler` method.

```
// Example reconnect code
exchange.SetDisconnectHandler(func (core *deribit.RPCCore) {
    exg := &deribit.NewExchangeFromCore(true, core)
	log.Warn("Disconnected from exchange. Attempting reconnection...")
	if err := exg.Connect(); err != nil {
		log.Fatalf("Error re-connecting to exchange: %s", err)
	}
	log.Info("Reconnected")
})
```

## Logging

The standard logger has been used within the library. You can plug this into your own application's logger by overriding the output io.Writer.

```
logger := logrus.New()
exchange.SetLogOutput(logger.Writer())
```

## Development

The `models` and `client` directories are where all the requests and responses are stored. The contents is automatically generated from the `schema` directory by `go-swagger`.

If you need to rebuild these use `make generate-models`.

The RPC subscriptions are also auto-generated. Use `make generate-methods` to rebuild these. They are in `rpc_subscriptions.go`.
