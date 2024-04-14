package service

import (
	"fmt"
	"iso8583/processor"
)

type MessageService struct {
	MethodType string
}

func (s *MessageService) CreateMsg() {

	msgBuilded := s.findProcessor()
	FileServiceInst.CreateCommand(s.MethodType, msgBuilded, true)

	fmt.Printf("ISO Message: %s \n", msgBuilded)
}

func (s *MessageService) findProcessor() string {
	var msgBuilded string
	processors := []IProcessor{
		processor.RequestAuthProcessor{},
		processor.RequestCancelProcessor{},
		processor.RequestUnmkProcessor{},
		processor.RequestPreAuthProcessor{},
		processor.RequestConfirmPreAuthProcessor{},
	}

	i := New()
	msg := i.CreateISO(s.MethodType)

	for _, processor := range processors {
		if processor.CanProcess(s.MethodType) {
			msgBuilded = processor.ConsumerProcessor(msg)
		}
	}

	return msgBuilded
}
