package service

import (
	"iso8583/builder"
)

// Iso8583 provides methods to parse and build ISO 8583 messages.
type Iso8583 struct {
	parser *Parser
	build  *builder.MessageBuilder
}

// New initializes a new ISO 8583 message parser and builder.
func New() *Iso8583 {
	return &Iso8583{
		parser: NewParser(),
		build:  builder.NewISO(),
	}
}

// Parse decodes an ISO 8583 message.
func (i *Iso8583) Parse(raw string) (*Parser, error) {
	msg, err := i.parser.Parse(raw)
	if err != nil {
		return msg, err
	}

	return msg, nil
}

// CreateISO initializes a new ISO 8583 message builder with an MTI.
func (i *Iso8583) CreateISO(mti string) *builder.MessageBuilder {
	// check MTI is filled, if true, reset it
	if len(i.build.MTI) > 0 {
		i.build.MTI = ""
	}

	// check if fields is filled, if true, reset it
	if i.build.Fields != nil {
		i.build.Fields = make(map[int]string)
	}

	// set MTI
	i.build.MTI = mti

	return i.build
}

// LogFields prints the fields of the ISO 8583 message.
func (i Iso8583) LogFields() {
	i.parser.LogFields()
}
