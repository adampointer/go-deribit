# go-deribit

## Overview

Go library for using the [Deribit's](https://www.deribit.com/reg-3027.8327) Websocket API. 

Deribit is a modern, fast BitCoin derivatives exchange. If you are using BitMex then you are doing it wrong! Deribit does not freeze up during higher than average load. Also, it is peer-to-peer, not run by market makers on lucrative contracts who want to liquidate you.

This library is a limited port of the [official wrapper libraries](https://github.com/deribit) to Go.

If you wish to try it out, be kind and use my affiliate link [https://www.deribit.com/reg-3027.8327](https://www.deribit.com/reg-3027.8327)

## Current Limitations

This is a work in progress! As I am building this library as I develop my own trading bot, I am only implementing features as and when I need them. I have not gone and added every single RPC method. It is, however, trivial to do so and I absolutely accept pull requests.

## Example Usage

```
package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// Trap ctrl+c
	terminate := make(chan os.Signal, 1)
	signal.Notify(terminate, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL)

	// Create channels to receive errors and to stop
	errs := make(chan error)
	stop := make(chan bool)
	deribit, err := api.NewExchange("YOUR-API-KEY", "YOUR-SECRET-KEY", true, errs, stop)
	if err != nil {
		log.Fatal("Error creating connection")
	}
	if err := deribit.Connect(); err != nil {
		log.Fatalf("Error connecting to exchange: %s", err)
	}

	// Hit the test RPC endpoint
	res, err := deribit.Test(true)
	if err != nil {
		log.Fatalf("Error testing connection: %s", err)
	}
	log.Printf("Connected to Deribit API v%s", res.APIBuild)

	// Test subs and notifications
	trades, err := deribit.SubscribeTrades("BTC-PERPETUAL")
	if err != nil {
		log.Fatalf("Error subscribing to trades: %s", err)
	}
	log.Print("Subscribed to trades")

	// Enter the main loop and wait for things to pop out of channels
Loop:
	for {
		select {
		case <-terminate:
			// On ctrl+c
			logger.Warn("Terminating")
			if err := deribit.Close(); err != nil {
				log.Fatalf("Error closing websocket: %s", err)
			}
			//stop <- true
			break Loop
		case notification := <-trades:
            		trds, err := notification.DecodeTrades()
            		if err != nil {
			    	log.Fatalf("Decode error: %s", err)
            		}
			// Log out Trade events as we get them
			for _, trd := range trds {
				log.Printf("Trade %f %s", evt.Price, evt.Direction)
			}
		case err := <-errs:
			log.Fatalf("RPC error: %s", err)
		}
	}
}
```
