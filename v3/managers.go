package deribit

import (
	"sync"

	"github.com/gorilla/websocket"
)

type callManager struct {
	reqMutex   sync.Mutex
	pending    map[uint64]*RPCCall
	countMutex sync.Mutex
	counter    uint64
}

func (c *callManager) nextID() uint64 {
	c.countMutex.Lock()
	defer c.countMutex.Unlock()
	id := c.counter
	c.counter++
	return id
}

func (c *callManager) addCall(call *RPCCall) {
	// Create a new request ID
	c.reqMutex.Lock()
	call.Req.ID = c.nextID()
	c.pending[call.Req.ID] = call
	c.reqMutex.Unlock()
}

func (c *callManager) deleteCall(call *RPCCall) {
	c.reqMutex.Lock()
	delete(c.pending, call.Req.ID)
	c.reqMutex.Unlock()
}

func (c *callManager) getCall(id uint64) *RPCCall {
	c.reqMutex.Lock()
	defer c.reqMutex.Unlock()
	return c.pending[id]
}

func (c *callManager) closeAllWithError(err error) {
	c.reqMutex.Lock()
	for _, call := range c.pending {
		call.Error = err
		call.Done <- true
	}
	c.reqMutex.Unlock()
}

type subManager struct {
	subMutex      sync.Mutex
	subscriptions map[string]*RPCSubscription
}

func (s *subManager) getSubscription(channel string) *RPCSubscription {
	s.subMutex.Lock()
	defer s.subMutex.Unlock()
	return s.subscriptions[channel]
}

func (s *subManager) getSubscriptions() map[string]*RPCSubscription {
	s.subMutex.Lock()
	defer s.subMutex.Unlock()
	return s.subscriptions
}

func (s *subManager) deleteSubscription(channel string) {
	s.subMutex.Lock()
	delete(s.subscriptions, channel)
	s.subMutex.Unlock()
}

type connManager struct {
	closedMutex sync.Mutex
	isClosed    bool
	conn        *websocket.Conn
}

func (c *connManager) closed() bool {
	c.closedMutex.Lock()
	defer c.closedMutex.Unlock()
	return c.isClosed
}

func (c *connManager) close() {
	c.closedMutex.Lock()
	c.isClosed = true
	c.closedMutex.Unlock()
}
