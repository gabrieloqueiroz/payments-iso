package processor

import (
	"fmt"
	"iso8583/builder"
	"iso8583/util"
)

type RequestCancelProcessor struct{}

func (s RequestCancelProcessor) CanProcess(mti string) bool {
	return mti == util.ConstantsUtilsInst.REQUEST_CANCELING()
}

func (s RequestCancelProcessor) ConsumerProcessor(msg *builder.MessageBuilder) string {

	fieldBuilderInst.SetBit02(msg)
	fieldBuilderInst.SetBit03Credit(msg)
	fieldBuilderInst.SetBit04(msg)
	fieldBuilderInst.SetBit07(msg)
	fieldBuilderInst.SetBit12(msg)
	fieldBuilderInst.SetBit13(msg)
	fieldBuilderInst.SetBit14(msg)
	fieldBuilderInst.SetBit22(msg)
	fieldBuilderInst.SetBit23(msg)
	fieldBuilderInst.SetBit35(msg)
	fieldBuilderInst.SetBit41(msg)
	fieldBuilderInst.SetBit42(msg)
	fieldBuilderInst.SetBit48(msg)
	fieldBuilderInst.SetBit49(msg)
	fieldBuilderInst.SetBit54(msg)
	fieldBuilderInst.SetBit61(msg)
	fieldBuilderInst.SetBit90(msg)
	fieldBuilderInst.SetBit120(msg)
	fieldBuilderInst.SetBit126(msg)

	isoMessage, err := msg.Build()
	if err != nil {
		fmt.Printf("Error building ISO message: %v  \n", err)
	}

	return isoMessage
}
