package deribit

import (
	"errors"
	"sync"

	"github.com/adampointer/go-deribit/v3/client/operations"
	"github.com/adampointer/go-deribit/v3/models"
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
	OnDisconnect func(*Exchange) // triggers on a failed read from connection

	url           string
	test          bool
	conn          *websocket.Conn
	mutex         sync.Mutex
	pending       map[uint64]*RPCCall
	subscriptions map[string]*RPCSubscription
	counter       uint64
	errors        chan error
	stop          chan bool
	auth          *models.PublicAuthResponse
	client        *operations.Client
}

// NewExchange creates a new API wrapper
// key and secret can be ignored if you are only calling public endpoints
func NewExchange(test bool, errs chan error, stop chan bool) (*Exchange, error) {
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
	if err := e.conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "")); err != nil {
		return err
	}
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
