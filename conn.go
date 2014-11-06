package main

import (
	"github.com/gorilla/websocket"
)

type connection struct {
	ws     *websocket.Conn
	respCh chan Response
	reqCh  chan Request
	doneCh chan bool
}

func (c *connection) reciever() {
	for {
		var req Request
		err := c.ws.ReadJSON(&req)
		if err != nil {
			break
		}
		c.reqCh <- req
	}
	c.ws.Close()
}

func (c *connection) sender() {
	for resp := range c.respCh {
		err := c.ws.WriteJSON(resp)
		if err != nil {
			break
		}
	}
	c.ws.Close()
}

var upgrader = &websocket.Upgrader{ReadBufferSize: 4096, WriteBufferSize: 4096}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil) //upgrader.Upgrade(w, r HTTP.HEADER) use later
	if err != nil {
		return
	}
	c := &connection{respCh: make(chan Response), reqCh: make(chan Request), doneCh: make(chan bool), ws: ws}
	h.register <- c
	defer func() {
		doneCh <- true
		h.unregister <- c
	}()
	go c.sender()
	go c.parseReq()
	c.receiver()
}
