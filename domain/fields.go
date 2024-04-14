package domain

type Fields struct{}

const (
	BIT_002_PAN                                  = 2
	BIT_003_PROC_CODE                            = 3
	BIT_004_TRAN_AMOUNT                          = 4
	BIT_005_SETTLE_AMOUNT                        = 5
	BIT_006_CARDHOLDER_BILLING                   = 6
	BIT_007_TRAN_DATE_TIME                       = 7
	BIT_009_CONVERSION_RATE_SETTLEMENT           = 9
	BIT_010_CONVERSION_RATE_CARDHOLDER           = 10
	BIT_011_SYS_TRACE_AUDIT_NUM                  = 11
	BIT_012_LOCAL_TRAN_TIME                      = 12
	BIT_013_LOCAL_TRAN_DATE                      = 13
	BIT_014_EXPIRATION_DATE                      = 14
	BIT_015_SELLTLEMENT_DATE                     = 15
	BIT_016_CONVERSION_DATE                      = 16
	BIT_017_CAPTURE_DATE                         = 17
	BIT_018_MERCHANT_TYPE                        = 18
	BIT_019_ACQUIRER_COUNTRY_CODE                = 19
	BIT_020_PAN_EXTENDED_COUNTRY_CODE            = 20
	BIT_021_FORWARDING_INST_COUNTRY_CODE         = 20
	BIT_022_POS_ENTRY_MODE                       = 22
	BIT_023_CARD_SEQUENCE_NUM                    = 23
	BIT_025_POS_CONDITION_CODE                   = 25
	BIT_026_POS_PIN_CAPTURE_CODE                 = 26
	BIT_027_AUTH_ID_RSP                          = 27
	BIT_028_TRAN_FEE_AMOUNT                      = 28
	BIT_029_SETTLEMENT_FEE_AMOUNT                = 29
	BIT_030_TRAN_PROC_FEE_AMOUNT                 = 30
	BIT_031_SETTLEMENT_PROC_FEE_AMOUNT           = 31
	BIT_032_ACQUIRING_INST_ID_CODE               = 32
	BIT_033_FORWARDING_INT_ID_CODE               = 33
	BIT_034_PRIMARY_ACCOUNT_NUMBER_EXTENDED      = 34
	BIT_035_TRACK_2_DATA                         = 35
	BIT_037_RETRIEVAL_REF_NUM                    = 37
	BIT_038_AUTH_ID_RESPONSE                     = 38
	BIT_039_RESPONSE_CODE                        = 39
	BIT_040_SERVICE_RESTRICTION_CODE             = 40
	BIT_041_CARD_ACCEPTOR_TERMINAL_ID            = 41
	BIT_042_CARD_ACCEPTOR_ID_CODE                = 42
	BIT_043_CARD_ACCEPTOR_NAME_LOCATION          = 43
	BIT_044_ADDITIONAL_RESPONSE_DATA             = 44
	BIT_045_TRACK_1_DATA                         = 45
	BIT_048_ADDITIONAL_DATA                      = 48
	BIT_049_TRAN_CURRENCY_CODE                   = 49
	BIT_050_SETTLEMENT_CURRENCY_CODE             = 50
	BIT_051_CURRENCY_CODE_CARDHOLDER             = 51
	BIT_052_PIN_DATA                             = 52
	BIT_053_SECURITY_RELATED_CONTROL_INFORMATION = 53
	BIT_054_ADDITIONAL_AMOUNTS                   = 54
	BIT_055_INTEGRATED_CIRCUIT_CARD              = 55
	BIT_056_MESSAGE_REASON_CODE                  = 56
	BIT_057_AUTHORISATION_LIFE_CYCLE             = 57
	BIT_058_AUTHORISING_AGENT_INSTITUTION        = 58
	BIT_059_ECHO_DATA                            = 59
	BIT_060_ADVICE_REASON_CODE                   = 60
	BIT_061_POINT_OF_SERIVE_DATA                 = 61
	BIT_062_ADDITIONAL_DATA                      = 62
	BIT_063_NETWORK_DATA                         = 63
	BIT_064_NETWORK_MANAGEMENT_INF_CODE          = 64
	BIT_065_NETWORK_MANAGEMENT_INF_CODE          = 65
	BIT_066_SETTLEMENT_CODE                      = 66
	BIT_067_EXTENDED_PAYMENT_CODE                = 67
	BIT_070_NETWORK_MANAGEMENT_INFORMATION_CODE  = 70
	BIT_073_DATE_ACTION                          = 73
	BIT_074_CREDITS_NUMBER                       = 74
	BIT_075_CREDITS_REVERSAL_NUMBER              = 75
	BIT_076_DEBITS_NUMBER                        = 76
	BIT_077_DEBITS_REVERSAL_NUMBER               = 77
	BIT_078_TRANSFER_NUMBER                      = 78
	BIT_079_TRANSFER_REVERSAL_NUMBER             = 79
	BIT_080_INQUIRIES_NUMBER                     = 80
	BIT_081_AUTHORISATIONS_NUMBER                = 81
	BIT_082_CREDITS_PROCESSING_FEE_AMOUNT        = 82
	BIT_083_CREDITS_TRANSACTION_FEE_AMOUNT       = 83
	BIT_084_DEBITS_PROCESSING_FEE_AMOUNT         = 84
	BIT_085_DEBITS_TRANSACTION_FEE_AMOUNT        = 85
	BIT_086_CREDITS_AMOUNT                       = 86
	BIT_087_CREDITS_REVERSAL_AMOUNT              = 87
	BIT_088_DEBITS_AMOUNT                        = 88
	BIT_089_DEBITS_REVERSAL_AMOUNT               = 89
	BIT_090_ORIGINAL_DATA_ELEMENTS               = 90
	BIT_091_FILE_UPDATE_CODE                     = 91
	BIT_095_REPLACEMENT_AMOUNTS                  = 95
	BIT_097_AMOUNT_NET_SETTLEMENT                = 97
	BIT_098_PAYEE                                = 98
	BIT_100_RECEIVING_INST_ID_CODE               = 100
	BIT_101_FILE_NAME                            = 101
	BIT_102_ACCOUNT_ID_1                         = 102
	BIT_103_ACCOUNT_ID_2                         = 103
	BIT_118_PAYMENTS_NUMBER                      = 118
	BIT_110_ADDTIONAL_DATA_2                     = 110
	BIT_119_PAYMENTS_REVERSAL_NUMBER             = 119
	BIT_120_ADDITIONAL_POS_DATA                  = 120
	BIT_121_SPECIAL_TRANSACTION_DATA             = 121
	BIT_123_                                     = 123
	BIT_126_SWITCH_PRIVATE_DATA                  = 126
	BIT_127_                                     = 127
)

var FieldsInst = &Fields{}