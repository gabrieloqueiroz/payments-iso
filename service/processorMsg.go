package service

import "iso8583/builder"

type IProcessor interface {
	CanProcess(mti string) bool
	ConsumerProcessor(msg *builder.MessageBuilder) string
}
