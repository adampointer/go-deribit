package main

import (
	"log"

	flag "github.com/spf13/pflag"

	"github.com/adampointer/go-deribit"
)

func main() {
	key := flag.String("access-key", "", "Access key")
	secret := flag.String("secret-key", "", "Secret access key")
	flag.Parse()
	errs := make(chan error)
	stop := make(chan bool)
	e, err := deribit.NewExchange(*key, *secret, false, errs, stop)
	if err != nil {
		log.Fatalf("Error creating connection: %s", err)
	}
	if err := e.Connect(); err != nil {
		log.Fatalf("Error connecting to exchange: %s", err)
	}
	go func() {
		err := <- errs
		stop <- true
		log.Fatalf("RPC error: %s", err)
	}()
	// Hit the test RPC endpoint
	res, err := e.PublicTest(nil)
	if err != nil {
		log.Fatalf("Error testing connection: %s", err)
	}
	log.Printf("Connected to Deribit API v%s", res.Version)
}
