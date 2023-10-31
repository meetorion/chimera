package sseKit

type messageType uint

const (
	// MessageTypeRaw 对于data，不做任何处理
	MessageTypeRaw messageType = iota + 1

	// MessageTypeEncode 对于data，编码一下（前端需对应处理）
	MessageTypeEncode

	// MessageTypeBase64 对于data，base64编码一下（前端需对应处理）
	MessageTypeBase64
)