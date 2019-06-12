package main

import (
	"fmt"
	"log"

	"github.com/adampointer/go-deribit/client/operations"
	flag "github.com/spf13/pflag"

	"github.com/adampointer/go-deribit"
)

type totals struct {
	btc, eth, total float64
}

func main() {
	test := flag.Bool("test", false, "Use test.deribit.com")
	key := flag.String("access-key", "", "Access key")
	secret := flag.String("secret-key", "", "Secret access key")
	flag.Parse()

	errs := make(chan error)
	stop := make(chan bool)
	e, err := deribit.NewExchange(*test, errs, stop)

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
	res, err := client.GetPublicTest(&operations.GetPublicTestParams{})
	if err != nil {
		log.Fatalf("Error testing connection: %s", err)
	}
	log.Printf("Connected to Deribit API v%s", *res.Payload.Result.Version)
	if err := e.Authenticate(*key, *secret); err != nil {
		log.Fatalf("Error authenticating: %s", err)
	}

	var total totals
	currencies := []string{"BTC", "ETH"}

	for _, sym := range currencies {
		summary, err := client.GetPrivateGetAccountSummary(&operations.GetPrivateGetAccountSummaryParams{Currency: string(sym)})
		if err != nil {
			log.Fatalf("Error getting account summary: %s", err)
		}
		equity := *summary.Payload.Result.Equity
		index, err := client.GetPublicGetIndex(&operations.GetPublicGetIndexParams{Currency: string(sym)})
		if err != nil {
			log.Fatalf("Error getting index price: %s", err)
		}
		var price float64
		switch sym {
		case "BTC":
			price = index.Payload.Result.BTC
			total.btc = price * equity
		case "ETH":
			price = index.Payload.Result.ETH
			total.eth = price * equity
		}
	}
	total.total = total.btc + total.eth
	fmt.Printf("\n Portfolio Value\n\n BTC: $%.2f\n ETH $%.2f\n Total: $%.2f\n\n", total.btc, total.eth, total.total)
}
