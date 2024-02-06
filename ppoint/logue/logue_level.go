package logue

type Level int

const (
	FATAL Level = iota + 1
	ERROR
	WARNING
	INFO
	DEBUG
)
