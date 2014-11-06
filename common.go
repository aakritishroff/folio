package main

import (
	"encoding/json"
)

//Uppercase first letter for all fields!
type Request struct {
	Name string
	Body json.RawMessage
}

type Response struct {
	Body json.RawMessage
	Err  string //maybe?
}

type CreateArgs struct {
	Prefix   string
	Slug     string
	CopyFrom *PageData
}

type CreateReply struct {
	State *PageState
	Err   string
}

type ReadArgs struct {
	Page           *PageState
	WaitForChanges time.Duration
}
