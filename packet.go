package main

type IdentifyPacket struct {
	Operation OperationCode      `json:"op"`
	Data      IdentifyPacketData `json:"d"`
}

type IdentifyPacketData struct {
	Token   string
	Intents uint64
	Os      string
	Browser string
	Device  string
}

type HeartbeatPacket struct {
	Operation OperationCode `json:"op"`
	Sequence  int           `json:"s"`
}
