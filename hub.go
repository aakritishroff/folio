package main

import (
	"code.google.com/p/go.net/websocket"
	"encoding/json"
	"os"
)

type hub struct {
	conns      map[*connection]bool
	register   chan *connection
	unregister chan *connection
}

var h = hub{
	conns:      make(map[*connection]bool),
	register:   make(chan *connection),
	unregister: make(chan *connection),
}

func (h *hub) init() {
	for {
		select {
		case c := <-h.register:
			h.conns[c] = true
		case c := <-h.unregister:
			if _, ok := h.conns[c]; ok {
				delete(h.conns, c)
				close(c.respCh)
				close(c.reqCh)
				close(c.doneCh)
			}
		}
	}
}
