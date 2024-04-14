package processor

import (
	"iso8583/builder"
	"iso8583/domain"
	"iso8583/util"
)

type fieldBuilder struct{}

func (s *fieldBuilder) SetBit02(msg *builder.MessageBuilder) {
	// Primary Account Number
	msg.AddField(domain.BIT_002_PAN, "5413330089020011")
}

func (s *fieldBuilder) SetBit03Credit(msg *builder.MessageBuilder) {
	// Processing Code
	msg.AddField(domain.BIT_003_PROC_CODE, util.ConstantsUtilsInst.PROCESSING_CODE_CREDIT())
}

func (s *fieldBuilder) SetBit03Debit(msg *builder.MessageBuilder) {
	// Processing Code
	msg.AddField(domain.BIT_003_PROC_CODE, util.ConstantsUtilsInst.PROCESSING_CODE_CREDIT())
}

func (s *fieldBuilder) SetBit04(msg *builder.MessageBuilder) {
	// Amount Transaction
	msg.AddField(domain.BIT_004_TRAN_AMOUNT, "000000001000")
}

func (s *fieldBuilder) SetBit07(msg *builder.MessageBuilder) {
	// Transmission Date
	msg.AddField(domain.BIT_007_TRAN_DATE_TIME, "0410180927")
}

func (s *fieldBuilder) SetBit11(msg *builder.MessageBuilder) {
	// NSU - Number Sequencial Unique
	msg.AddField(domain.BIT_011_SYS_TRACE_AUDIT_NUM, "009194")
}

func (s *fieldBuilder) SetBit12(msg *builder.MessageBuilder) {
	// Time, Local Transaction
	msg.AddField(domain.BIT_012_LOCAL_TRAN_TIME, "150927")
}

func (s *fieldBuilder) SetBit13(msg *builder.MessageBuilder) {
	// Time, Local Transaction
	msg.AddField(domain.BIT_013_LOCAL_TRAN_DATE, "0410")
}

func (s *fieldBuilder) SetBit14(msg *builder.MessageBuilder) {
	// Expiration Date
	msg.AddField(domain.BIT_014_EXPIRATION_DATE, "0410")
}

func (s *fieldBuilder) SetBit17(msg *builder.MessageBuilder) {
	// Capture Date
	msg.AddField(domain.BIT_017_CAPTURE_DATE, "0410")
}

func (s *fieldBuilder) SetBit18(msg *builder.MessageBuilder) {
	// Merchant Type
	msg.AddField(domain.BIT_018_MERCHANT_TYPE, "1520")
}

func (s *fieldBuilder) SetBit19(msg *builder.MessageBuilder) {
	// Acquiring Institution Country Code
	msg.AddField(domain.BIT_019_ACQUIRER_COUNTRY_CODE, "076")
}

func (s *fieldBuilder) SetBit20(msg *builder.MessageBuilder) {
	// PAN Extended - Country Code
	msg.AddField(domain.BIT_020_PAN_EXTENDED_COUNTRY_CODE, "076")
}

func (s *fieldBuilder) SetBit21(msg *builder.MessageBuilder) {
	// Forwarding Institution Country Code
	msg.AddField(domain.BIT_020_PAN_EXTENDED_COUNTRY_CODE, "076")
}

func (s *fieldBuilder) SetBit22(msg *builder.MessageBuilder) {
	// Point of Service Entry Mode
	msg.AddField(domain.BIT_022_POS_ENTRY_MODE, "051")
}

func (s *fieldBuilder) SetBit23(msg *builder.MessageBuilder) {
	// Application PAN Sequence Number
	msg.AddField(domain.BIT_023_CARD_SEQUENCE_NUM, "003")
}

func (s *fieldBuilder) SetBit25(msg *builder.MessageBuilder) {
	// Point of Service Condition Code
	msg.AddField(domain.BIT_025_POS_CONDITION_CODE, "00")
}

func (s *fieldBuilder) SetBit26(msg *builder.MessageBuilder) {
	// Point of Service Capture Code
	msg.AddField(domain.BIT_026_POS_PIN_CAPTURE_CODE, "12")
}

func (s *fieldBuilder) SetBit27(msg *builder.MessageBuilder) {
	// Authorizing Identification Response Length
	msg.AddField(domain.BIT_027_AUTH_ID_RSP, "1")
}

func (s *fieldBuilder) SetBit28(msg *builder.MessageBuilder) {
	// Amount, Transaction Fee
	msg.AddField(domain.BIT_028_TRAN_FEE_AMOUNT, "C00000000")
}

func (s *fieldBuilder) SetBit29(msg *builder.MessageBuilder) {
	// Amount, Settlement Fee
	msg.AddField(domain.BIT_029_SETTLEMENT_FEE_AMOUNT, "D00000000")
}

func (s *fieldBuilder) SetBit30(msg *builder.MessageBuilder) {
	// Amount, Transaction Processing Fee
	msg.AddField(domain.BIT_030_TRAN_PROC_FEE_AMOUNT, "D00000000")
}

func (s *fieldBuilder) SetBit31(msg *builder.MessageBuilder) {
	// Amount, Settlement Processing Fee
	msg.AddField(domain.BIT_031_SETTLEMENT_PROC_FEE_AMOUNT, "C00000000")
}

func (s *fieldBuilder) SetBit32(msg *builder.MessageBuilder) {
	// Acquiring Institution Identification Code
	msg.AddField(domain.BIT_032_ACQUIRING_INST_ID_CODE, "014553")
}

func (s *fieldBuilder) SetBit33(msg *builder.MessageBuilder) {
	// Forward Institution Identification Code
	msg.AddField(domain.BIT_033_FORWARDING_INT_ID_CODE, "014553")
}

func (s *fieldBuilder) SetBit34(msg *builder.MessageBuilder) {
	// Primary Account Number, Extended
	msg.AddField(domain.BIT_034_PRIMARY_ACCOUNT_NUMBER_EXTENDED, "12345678901234567890")
}

func (s *fieldBuilder) SetBit35(msg *builder.MessageBuilder) {
	// Track 2
	msg.AddField(domain.BIT_035_TRACK_2_DATA, "5413330089020011D2512601079360805")
}

func (s *fieldBuilder) SetBit37(msg *builder.MessageBuilder) {
	// Retrievel Reference Number
	msg.AddField(domain.BIT_037_RETRIEVAL_REF_NUM, "000000088900")
}

func (s *fieldBuilder) SetBit38(msg *builder.MessageBuilder) {
	// Authorization Identification Response
	msg.AddField(domain.BIT_038_AUTH_ID_RESPONSE, "123456")

}

func (s *fieldBuilder) SetBit39(msg *builder.MessageBuilder) {
	// Response Code
	msg.AddField(domain.BIT_039_RESPONSE_CODE, "00")

}

func (s *fieldBuilder) SetBit41(msg *builder.MessageBuilder) {
	// Terminal Id
	msg.AddField(domain.BIT_041_CARD_ACCEPTOR_TERMINAL_ID, "GT000131")
}

func (s *fieldBuilder) SetBit42(msg *builder.MessageBuilder) {
	// Card Acceptor Id
	msg.AddField(domain.BIT_042_CARD_ACCEPTOR_ID_CODE, "012005532486001")
}

func (s *fieldBuilder) SetBit43(msg *builder.MessageBuilder) {
	// Card Acceptor Transaction
	msg.AddField(domain.BIT_043_CARD_ACCEPTOR_NAME_LOCATION, "ENTREPA*Homologa  o E  SAO PAULO     BRA")
}

func (s *fieldBuilder) SetBit48(msg *builder.MessageBuilder) {
	// Additional Data - Private
	msg.AddField(domain.BIT_048_ADDITIONAL_DATA, " 6105000017601U1401C??o E012014MTIP32 MCD 15A01301440402244000115")
	// msg.AddField(48, "009008ENTREPAY010013ASAPCARD12345")
}

func (s *fieldBuilder) SetBit49(msg *builder.MessageBuilder) {
	// Currency Code, Transaction
	msg.AddField(domain.BIT_049_TRAN_CURRENCY_CODE, "986")
}

func (s *fieldBuilder) SetBit50(msg *builder.MessageBuilder) {
	// Country Code
	msg.AddField(domain.BIT_050_SETTLEMENT_CURRENCY_CODE, "076")
}

func (s *fieldBuilder) SetBit51(msg *builder.MessageBuilder) {
	// Currency Code, Cardholder Billing
	msg.AddField(domain.BIT_051_CURRENCY_CODE_CARDHOLDER, "986")
}

func (s *fieldBuilder) SetBit52(msg *builder.MessageBuilder) {
	// Crypto Pass
	msg.AddField(domain.BIT_052_PIN_DATA, "D002E27D002E2727")
}

func (s *fieldBuilder) SetBit53(msg *builder.MessageBuilder) {
	// Security Related Control Information
	msg.AddField(domain.BIT_053_SECURITY_RELATED_CONTROL_INFORMATION, "2600000000000000")
}

func (s *fieldBuilder) SetBit54(msg *builder.MessageBuilder) {
	// Additional Amount Transaction
	msg.AddField(domain.BIT_054_ADDITIONAL_AMOUNTS, "0010120000000022000020011")
}

func (s *fieldBuilder) SetBit55(msg *builder.MessageBuilder) {
	// Emv Data
	msg.AddField(domain.BIT_055_INTEGRATED_CIRCUIT_CARD, "FF20c9820218009F2701809F26081DDDB9B651A2A8CA9F36020214950580402480009F34034203009F3704E065DFEF9F3303E0F0C85F280200769F10120610A000002A0000000000000000000000FF9A032404099F1A0200769F3501229C01008407A00000000410109F02060000000022005F2A0209865F24032008319F150255118E1200000000000000004201410342031E031F039F120948697065726361726450094869706572636172645F3401005F20184D45524341444F2F48204142455254555241202020202020")
}

func (s *fieldBuilder) SetBit57(msg *builder.MessageBuilder) {
	// Amount, Cash
	msg.AddField(domain.BIT_057_AUTHORISATION_LIFE_CYCLE, "2500")
}

func (s *fieldBuilder) SetBit58(msg *builder.MessageBuilder) {
	// Authorizing Agent Institution ID
	msg.AddField(domain.BIT_058_AUTHORISING_AGENT_INSTITUTION, "1234")
}

func (s *fieldBuilder) SetBit59(msg *builder.MessageBuilder) {
	// Echo Data
	msg.AddField(domain.BIT_059_ECHO_DATA, "XXXXXXYYYYYYYYYZ")
}

func (s *fieldBuilder) SetBit60(msg *builder.MessageBuilder) {

	msg.AddField(domain.BIT_060_ADVICE_REASON_CODE, "D002E27131784908")
}

func (s *fieldBuilder) SetBit61(msg *builder.MessageBuilder) {
	// Terminal Data
	msg.AddField(domain.BIT_061_POINT_OF_SERIVE_DATA, "001008GP0104AD0020101493572496003026000000000030007603015000  004016000001.63 230602005003PAX006004A9100070101493572496")
}

func (s *fieldBuilder) SetBit62(msg *builder.MessageBuilder) {
	// Additional Data
	msg.AddField(domain.BIT_062_ADDITIONAL_DATA, "001008LISTO IP002031Av. Brigadeiro Faria Lima, 1663003009Sao Paulo004002SP00500307600600805412001007004899900801412345678910111009011110307809140101650316056B2C5EFCE7C34D97DD2831FB40AA8765377052C19A78128C4CDF977F76A487F650CBEB4C8857780458AFF3E8868CEEE091202EEC91E6F57C850756A7F4EE83B6F41E8D5B5B488CA6252B6D427EF905A011042200001D002E27131784908FFFFFFFFFFF00A2000C5012014A0000000041010013020MASTERCARD          ")
}

func (s *fieldBuilder) SetBit64(msg *builder.MessageBuilder) {
	// Network Management Information Code
	msg.AddField(domain.BIT_064_NETWORK_MANAGEMENT_INF_CODE, "IFFFXPTO")
}

func (s *fieldBuilder) SetBit65(msg *builder.MessageBuilder) {
	// Network Management Information Code
	msg.AddField(domain.BIT_065_NETWORK_MANAGEMENT_INF_CODE, "IFFFXPTO")
}

func (s *fieldBuilder) SetBit66(msg *builder.MessageBuilder) {
	// Network Management Information Code
	msg.AddField(domain.BIT_066_SETTLEMENT_CODE, "IFFFXPTO")
}

func (s *fieldBuilder) SetBit67(msg *builder.MessageBuilder) {
	// Extended Payment Code
	msg.AddField(domain.BIT_067_EXTENDED_PAYMENT_CODE, "IFFFXPTO")
}

func (s *fieldBuilder) SetBit90(msg *builder.MessageBuilder) {
	// Extended Payment Code
	msg.AddField(domain.BIT_090_ORIGINAL_DATA_ELEMENTS, "010000919404101809270000000000000000000000")
}

func (s *fieldBuilder) SetBit120(msg *builder.MessageBuilder) {
	// Additional Data POS
	msg.AddField(domain.BIT_120_ADDITIONAL_POS_DATA, "9990580010320200007618013010342260000056451300300200038006374686")
}

func (s *fieldBuilder) SetBit121(msg *builder.MessageBuilder) {
	// Special Transaction Data
	msg.AddField(domain.BIT_121_SPECIAL_TRANSACTION_DATA, "002012000298200875003009v.1.0.0.0004021032019050815541075201")
}

func (s *fieldBuilder) SetBit126(msg *builder.MessageBuilder) {
	// Tlv Crypto Info
	msg.AddField(domain.BIT_126_SWITCH_PRIVATE_DATA, "010013002020FFFFFFFFFFF00A2000C50030012004020000000000000000000000050011006004    ")
}

var fieldBuilderInst = &fieldBuilder{}
