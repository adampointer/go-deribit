package main

import (
	"fmt"
	"log"

	"github.com/adampointer/go-deribit/v3"
	"github.com/adampointer/go-deribit/v3/client/account_management"
	"github.com/adampointer/go-deribit/v3/client/private"
	"github.com/adampointer/go-deribit/v3/client/public"
	flag "github.com/spf13/pflag"
)

func main() {
	key := flag.String("access-key", "", "Access key")
	secret := flag.String("secret-key", "", "Secret access key")
	flag.Parse()
	errs := make(chan error)
	stop := make(chan bool)
	e, err := deribit.NewExchange(true, errs, stop)

	if err != nil {
		log.Fatalf("Error creating connection: %s", err)
	}
	if err := e.Connect(); err != nil {
		log.Fatalf("Error connecting to exchange: %s", err)
	}
	go func() {
		err := <-errs
		stop <- true
		log.Fatalf("RPC error: %s", err)
	}()
	client := e.Client()
	// Hit the test RPC endpoint
	res, err := client.Public.GetPublicTest(&public.GetPublicTestParams{})
	if err != nil {
		log.Fatalf("Error testing connection: %s", err)
	}
	log.Printf("Connected to Deribit API v%s", *res.Payload.Result.Version)
	if err := e.Authenticate(*key, *secret); err != nil {
		log.Fatalf("Error authenticating: %s", err)
	}

	// Account summary
	summary, err := client.AccountManagement.GetPrivateGetAccountSummary(&account_management.GetPrivateGetAccountSummaryParams{Currency: "BTC"})
	if err != nil {
		log.Fatalf("Error getting account summary: %s", err)
	}
	fmt.Printf("Available funds: %f\n", *summary.Payload.Result.AvailableFunds)

	// Buy
	buyParams := &private.GetPrivateBuyParams{
		Amount:         10,
		InstrumentName: "BTC-PERPETUAL",
		Type:           strPointer("market"),
	}
	buy, err := client.Private.GetPrivateBuy(buyParams)
	if err != nil {
		log.Fatalf("Error submitting buy order: %s", err)
	}
	fmt.Printf("Brought at %f\n", buy.Payload.Result.Order.AveragePrice)

	// Order book subscription
	book, err := e.SubscribeBookGroup("BTC-PERPETUAL", "none", "1", "100ms")
	if err != nil {
		log.Fatalf("Error subscribing to the book: %s", err)
	}
	for b := range book {
		fmt.Printf("Top bid: %f Top ask: %f\n", b.Bids[0][0], b.Asks[0][0])
	}
	e.Close()
}

func strPointer(str string) *string {
	return &str
}
