package zerodoc

type MessageType uint8

const (
	MSG_USAGE MessageType = iota
	MSG_PERF
	MSG_GEO
	MSG_FLOW
	MSG_CONSOLE_LOG
	MSG_TYPE
	MSG_FPS
	MSG_LOG_USAGE

	MSG_VTAP_USAGE
)

const (
	MAX_STRING_LENGTH = 1024
)
