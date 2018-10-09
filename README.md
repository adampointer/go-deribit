# go-deribit

## Overview

Go library for using the [Deribit's](https://www.deribit.com/reg-3027.8327) Websocket API. 

Deribit is a modern, fast BitCoin derivatives exchange. If you are using BitMex then you are doing it wrong! Deribit does not freeze up during higher than average load and is peer-to-peer, not run by market makers on lucrative contracts.

This library is a limited port of the [official wrapper libraries](https://github.com/deribit) to Go.

If you wish to try it out, be kind and use my affiliate link [https://www.deribit.com/reg-3027.8327](https://www.deribit.com/reg-3027.8327)

## Current Limitations

This is a work in progress! As I am building this library as I develop my own trading bot, I am only implementing features as and when I need them. I have not gone and added every single RPC method. It is, however, trivial to do so and I absolutely accept pull requests.

Clean shutdown does not work properly. Calling `Close()` should close the websocket connection and all its associated channels, but doesn't as it is blocking trying to read from the socket.

## Example Usage

```
package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	api "github.com/adampointer/go-deribit"
	"github.com/mitchellh/mapstructure"
	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		panic(fmt.Sprintf("can't initialize zap logger: %v", err))
	}
	defer logger.Sync()

    // Trap ctrl+c
	terminate := make(chan os.Signal, 1)
	signal.Notify(terminate, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL)

    // Create channels to receive errors and to stop
	errs := make(chan error)
	stop := make(chan bool)
	deribit, err := api.NewExchange("YOUR-API-KEY", "YOUR-SECRET-KEY", true, errs, stop)
	if err != nil {
		logger.Fatal("Error creating connection", zap.String("message", err.Error()))
	}
	if err := deribit.Connect(); err != nil {
		logger.Fatal("Error connecting to exchange", zap.String("message", err.Error()))
	}

    // Hit the test RPC endpoint
	res, err := deribit.Test(true)
	if err != nil {
		logger.Fatal("Error testing connection", zap.String("message", err.Error()))
	}
	logger.Info("Connected to Deribit API", zap.String("version", res.APIBuild))

	// Test subs and notifications
	trades := make(chan []interface{})
	if err := deribit.Subscribe(api.EvtTrade, trades); err != nil {
		logger.Fatal("Error subscribing to trades", zap.String("message", err.Error()))
	}

    // Enter the main loop
Loop:
	for {
		select {
		case <-terminate:
            // On ctrl+c
			logger.Warn("Terminating")
			if err := deribit.Close(); err != nil {
				logger.Fatal("Error closing websocket", zap.String("message", err.Error()))
			}
			//stop <- true
			break Loop
		case trade := <-trades:
            // Log out Trade events as we receieve them
			for _, trd := range trade {
				var evt api.EvtTradeResponse
				if err := mapstructure.Decode(trd, &evt); err != nil {
					logger.Fatal("Unable to decode map", zap.String("message", err.Error()))
				}
				fmt.Printf("TRADE: price %f : direction %s\n", evt.Price, evt.Direction)
				logger.Info("Trade", zap.Float64("price", evt.Price), zap.String("direction", evt.Direction))
			}
		case err := <-errs:
			fmt.Printf("Error: %s\n", err)
			logger.Error("RPC error", zap.String("message", err.Error()))
		}
	}
}
```
