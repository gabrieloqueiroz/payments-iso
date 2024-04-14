package processor

import (
	"fmt"
	"iso8583/builder"
	"iso8583/util"
)

type RequestUnmkProcessor struct{}

func (s RequestUnmkProcessor) CanProcess(mti string) bool {
	return mti == util.ConstantsUtilsInst.REQUEST_UNMAKING()
}

func (s RequestUnmkProcessor) ConsumerProcessor(msg *builder.MessageBuilder) string {

	fieldBuilderInst.SetBit03Credit(msg)
	fieldBuilderInst.SetBit04(msg)
	fieldBuilderInst.SetBit07(msg)
	fieldBuilderInst.SetBit11(msg)
	fieldBuilderInst.SetBit12(msg)
	fieldBuilderInst.SetBit13(msg)
	fieldBuilderInst.SetBit22(msg)
	fieldBuilderInst.SetBit23(msg)
	fieldBuilderInst.SetBit42(msg)
	fieldBuilderInst.SetBit41(msg)
	fieldBuilderInst.SetBit42(msg)
	fieldBuilderInst.SetBit48(msg)
	fieldBuilderInst.SetBit49(msg)
	fieldBuilderInst.SetBit61(msg)
	fieldBuilderInst.SetBit62(msg)
	fieldBuilderInst.SetBit90(msg)
	fieldBuilderInst.SetBit126(msg)

	isoMessage, err := msg.Build()
	if err != nil {
		fmt.Printf("Error building ISO message: %v  \n", err)
	}

	return isoMessage
}
