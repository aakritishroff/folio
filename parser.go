package main

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"time"
)

var requests = map[string]func() interface{}{
	"create": func() interface{} { return new(CreateArgs) },
	"read":   func() interface{} { return new(ReadArgs) },
	// "replace":      func() interface{} { return new(replaceArgs) },
	// "update":       func() interface{} { return new(updateArgs) },
	//"delete": func() interface{} { return new(DeleteArgs) },
	// "deletePrefix": func() interface{} { return new(deletePrefixArgs) },
	// "renamePrefix": func() interface{} { return new(renamePrefixArgs) },
}

func (c *connection) parseReq() {
	for {
		select {
		case r := <-c.reqCh:
			newReq := requests[r.Name]
			if newReq == nil {
				errmsg := "unknown command" + r.Name
				errCmd := &Response{Err: errmsg, Body: nil}
				c.respCh <- errCmd
			}
			reqType := newReq()
			if err := json.Unmarshal([]byte(r.Body), reqType); err != nil {
				errmsg = "malformed arguments" + r.Name
				errArgs := &Response{Err: errmsg, Body: nil}
				c.respCh <- errArgs
			}
			switch reqArgs := reqType.(type) {
			case *CreateArgs:
				r1 := create(reqArgs)
				reply, err := json.Marshal(r1)
				if err != nil {
					errReply := &Response{Err: "Error while parsing function reply", Body: nil}
					c.respCh <- errArgs
				}
				response := &Response{Err: "OK", Body: reply}
				c.respCh <- response
			case *ReadArgs:
				read(reqArgs)
			}
		case done := <-c.doneCh:
			return
		}
	}
}

func create(args *CreateArgs) *CreateReply {

}

func read(args *ReadArgs) {

}
