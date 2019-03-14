package deribit

import (
	"errors"
	"github.com/gorilla/websocket"
	"sync"
)

const (
	liveURL = "wss://www.deribit.com/ws/api/v2/"
	testURL = "wss://test.deribit.com/ws/api/v2/"
)

// ErrTimeout - request timed out
var ErrTimeout = errors.New("timed out waiting for a response")

// Exchange is an API wrapper with the exchange
type Exchange struct {
	key           string
	secret        string
	url           string
	test          bool
	conn          *websocket.Conn
	mutex         sync.Mutex
	pending       map[uint64]*RPCCall
	subscriptions map[string]*RPCSubscription
	counter       uint64
	errors        chan error
	stop          chan bool
}

// NewExchange creates a new API wrapper
// key and secret can be ignored if you are only calling public endpoints
func NewExchange(key, secret string, test bool, errs chan error, stop chan bool) (*Exchange, error) {
	exc := &Exchange{
		pending:       make(map[uint64]*RPCCall, 1),
		subscriptions: make(map[string]*RPCSubscription),
		counter:       1,
		errors:        errs,
		stop:          stop,
	}

	exc.url = liveURL
	if test {
		exc.test = true
		exc.url = testURL
	}
	exc.key = key
	exc.secret = secret
	return exc, nil
}

// Connect to the websocket API
func (e *Exchange) Connect() error {
	c, _, err := websocket.DefaultDialer.Dial(e.url, nil)
	if err != nil {
		return err
	}
	e.conn = c
	// Start listening for responses
	go e.read()
	//go e.heartbeat()
	return nil
}

// Close the websocket connection
func (e *Exchange) Close() error {
	// Ignore error as we will close the conn anyway
	e.conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	return e.conn.Close()
}

/*func (e *Exchange) heartbeat() {
	ticker := time.NewTicker(10 * time.Second)
	go func() {
		for {
			select {
			case <-ticker.C:
				if _, err := e.Ping(); err != nil {
					e.stop <- true
				}
			case <-e.stop:
				ticker.Stop()
				return
			}
		}
	}()
}*/
