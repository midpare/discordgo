package main

type OperationCode int

const (
	Dispatch OperationCode = iota
	Heartbeat
	Identify
	PresenceUpdate
	VoiceStateUpdate
	Resume OperationCode = iota + 1
	Reconnect
	RequestGuildMembers
	InvalidSession
	Hello
	HeartbeatACK
)
