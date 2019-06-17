# go-deribit

## V2 API - Major Breaking Changes

Both the remote API and this implementation have changed a lot since v1. The deprecated, but still functioning v1 API has been tagged `v1.0.0`.

[V2 API Documentation](http://docs.deribit.com/v2/?javascript#deribit-api-v2-0-0)

## Overview

Go library for using the [Deribit's](https://www.deribit.com/reg-3027.8327) **v2** Websocket API. 

Deribit is a modern, fast BitCoin derivatives exchange. If you are using BitMex then you are doing it wrong! Deribit does not freeze up during higher than average load. Also, it is peer-to-peer, not run by market makers on lucrative contracts who want to liquidate you.

This library is a port of the [official wrapper libraries](https://github.com/deribit) to Go.

If you wish to try it out, be kind and use my affiliate link [https://www.deribit.com/reg-3027.8327](https://www.deribit.com/reg-3027.8327)

## Example Usage

Look at `cmd/example/main.go`

```
make build
example --access-key XXX --secret-key YYYYYY
```

## Development

The `models` and `client` directories are where all the requests and responses are stored. The contents is automatically generated from the `schema` directory by `go-swagger`.

If you need to rebuild these use `make generate-models`.

The RPC subscriptions are also auto-generated. Use `make generate-methods` to rebuild these. They are in `rpc_subscriptions.go`.
