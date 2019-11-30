package deribit

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/cenkalti/backoff"
	"github.com/go-openapi/strfmt"

	"github.com/adampointer/go-deribit/v3/client/operations"
	"github.com/gorilla/websocket"
)

const (
	liveURL = "wss://www.deribit.com/ws/api/v2/"
	testURL = "wss://test.deribit.com/ws/api/v2/"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10
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
				pending:       make(map[uint64]*RPCCall, 1),
				subscriptions: make(map[string]*RPCSubscription),
				counter:       1,
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
	c.SetPongHandler(func(string) error { c.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	e.connMgr.conn = c
	// Start listening for responses
	go e.read()

	ctx, cancel := context.WithCancel(context.Background())
	e.connMgr.stopPinging = cancel
	go e.pinging(ctx)

	authed := false
	if clientID != "" && clientSecret != "" {
		// re-authenticate on reconnect
		if err := e.Authenticate(); err != nil {
			return fmt.Errorf("Error re-authenticating: %s", err)
		}
		authed = true
	}

	if f := e.OnConnect; f != nil {
		f()
	}
	e.resubscribe(authed)
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

// SetDisconnectHandler overrides the default disconnect handler
func (e *Exchange) SetDisconnectHandler(f func(*RPCCore)) {
	e.onDisconnect = f
}

// Reconnect reconnect is already built-in on OnDisconnect. Use this method only within OnDisconnect to override it
func (e *Exchange) Reconnect(core *RPCCore) {
	e.connMgr.stopPinging()
	if err := retry(e.Connect); err != nil {
		log.Printf("retry to reconnect failed %v", err)
		e.errors <- fmt.Errorf("retry to reconnect failed: %w", err)
	}
}

// Client returns an initialised API client
func (e *Exchange) Client() *operations.Client {
	if e.client == nil {
		e.client = operations.New(e, strfmt.Default)
	}
	return e.client
}

func (e *Exchange) pinging(ctx context.Context) {
	ticker := time.NewTicker(pingPeriod)
	for {
		select {
		case <-ctx.Done():
			return
		case <-e.stop:
			ticker.Stop()
			return
		case <-ticker.C:
			e.connMgr.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := e.connMgr.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

func (e *Exchange) resubscribe(authed bool) {
	if !authed {
		// We re-wire public subscriptions
		for chan0 := range e.calls.getSubscriptions() {
			log.Printf("Attempt at reconnecting subscription: %v", chan0)
			req := &operations.GetPublicSubscribeParams{Channels: []string{chan0}}

			if _, err := e.Client().GetPublicSubscribe(req); err != nil {
				log.Printf("Reconnection failed: %v", err)
				e.calls.deleteSubscription(chan0)
				continue
			}
			log.Printf("Subscription %v successfully re-wired", chan0)
		}
		return
	}

	// We re-wire private subscriptions
	for chan0 := range e.calls.getSubscriptions() {
		log.Printf("Attempt at reconnecting subscription: %v", chan0)
		req := &operations.GetPrivateSubscribeParams{Channels: []string{chan0}}

		if _, err := e.Client().GetPrivateSubscribe(req); err != nil {
			log.Printf("Reconnection failed: %v", err)
			e.calls.deleteSubscription(chan0)
			continue
		}
		log.Printf("Subscription %v successfully re-wired", chan0)
	}
}

func retry(operation func() error) error {
	b := backoff.NewExponentialBackOff()
	b.InitialInterval = time.Minute
	b.MaxInterval = 3 * time.Minute
	b.MaxElapsedTime = 15 * time.Minute

	notify := func(err error, t time.Duration) {
		log.Printf("%v retry in %s", err, t.Round(time.Second).String())
	}
	return backoff.RetryNotify(operation, b, notify)
}
