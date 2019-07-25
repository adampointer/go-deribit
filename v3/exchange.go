package deribit

import (
	"errors"
	"io"
	"log"
	"time"

	"github.com/go-openapi/strfmt"

	"github.com/adampointer/go-deribit/v3/client/operations"
	"github.com/gorilla/websocket"
)

const (
	liveURL = "wss://www.deribit.com/ws/api/v2/"
	testURL = "wss://test.deribit.com/ws/api/v2/"
)

// ErrTimeout - request timed out
var ErrTimeout = errors.New("timed out waiting for a response")

// Exchange is an API wrapper with the exchange
type Exchange struct {
	url    string
	test   bool
	client *operations.Client
	RPCCore
}

// NewExchange creates a new API wrapper
// key and secret can be ignored if you are only calling public endpoints
func NewExchange(test bool, errs chan error, stop chan bool) (*Exchange, error) {
	exc := &Exchange{
		RPCCore: RPCCore{
			calls: &callManager{
				pending: make(map[uint64]*RPCCall, 1),
				counter: 1,
			},
			subs: &subManager{
				subscriptions: make(map[string]*RPCSubscription),
			},
			connMgr: &connManager{},
			errors:  errs,
			stop:    stop,
		},
	}
	exc.onDisconnect = exc.Reconnect
	exc.url = liveURL
	if test {
		exc.test = true
		exc.url = testURL
	}
	return exc, nil
}

// NewExchangeFromCore creates a new exchange from an existing RPCCore
func NewExchangeFromCore(test bool, core RPCCore) *Exchange {
	exc := &Exchange{RPCCore: core}
	exc.onDisconnect = exc.Reconnect
	exc.url = liveURL
	if test {
		exc.test = true
		exc.url = testURL
	}
	return exc
}

// Connect to the websocket API
func (e *Exchange) Connect() error {
	c, _, err := websocket.DefaultDialer.Dial(e.url, nil)
	if err != nil {
		return err
	}
	e.connMgr.conn = c
	// Start listening for responses
	go e.read()
	go e.heartbeat()
	return nil
}

// Close the websocket connection
func (e *Exchange) Close() error {
	return e.close()
}

// SetLogOutput set log output
func (e *Exchange) SetLogOutput(w io.Writer) {
	log.SetOutput(w)
}

// Override the default disconnect handler
func (e *Exchange) SetDisconnectHandler(f func(*RPCCore)) {
	e.onDisconnect = f
}

// Reconnect reconnect is already built-in on OnDisconnect. Use this method only within OnDisconnect to override it
func (e *Exchange) Reconnect(core *RPCCore) {
	// Rebuild the connection and the subscriptions
	c, _, err := websocket.DefaultDialer.Dial(e.url, nil)
	if err != nil {
		log.Printf("Error in the default dialer %v", err)
	} else {
		// This seems to have worked
		log.Printf("Reconnected to the API...")
		e.connMgr.conn = c
		go e.read()

		// We re-authenticated
		if err := e.Authenticate(); err != nil {
			log.Fatalf("Error re-authenticating: %s", err)
		}

		// We re-wire the subscriptions
		e.rewireSubscriptions()
	}
}

// Client returns an initialised API client
func (e *Exchange) Client() *operations.Client {
	if e.client == nil {
		e.client = operations.New(e, strfmt.Default)
	}
	return e.client
}

func (e *Exchange) rewireSubscriptions() {
	for chan0 := range e.subs.getSubscriptions() {
		log.Printf("Attempt at reconnecting subscription: %v", chan0)
		if _, err := e.Client().GetPrivateSubscribe(&operations.GetPrivateSubscribeParams{Channels: []string{chan0}}); err != nil {
			log.Printf("Reconnection failed: %v", err)
			e.subs.deleteSubscription(chan0)
		} else {
			log.Printf("Subscription %v successfully re-wired", chan0)
		}
	}
}

func (e *Exchange) heartbeat() {
	ticker := time.NewTicker(10 * time.Second)
	go func() {
		for {
			select {
			case <-ticker.C:
				e.testConnection()
			case <-e.stop:
				ticker.Stop()
			}
		}
	}()
}

func (e *Exchange) testConnection() {
	if _, err := e.Client().GetPublicTest(&operations.GetPublicTestParams{}); err != nil {
		// We've got an error, so we reconnect
		e.onDisconnect(&e.RPCCore)
	}
}
