package main

import "encoding/json"

type BaseEvent struct {
	Operation OperationCode `json:"op"`
}

type BasicEvent struct {
	Operation OperationCode   `json:"op"`
	Sequence  int             `json:"s"`
	Type      string          `json:"t"`
	Data      json.RawMessage `json:"d"`
}

type HelloEvent struct {
	Heartbeat int `json:"heartbeat_interval"`
}
