package util

type ConstantsUtils struct{}

// MTI CONSTANTS
func (s *ConstantsUtils) REQUEST_BRAND() string {
	return "0100"
}

func (s *ConstantsUtils) REQUEST_AUTHORIZATION() string {
	return "0200"
}

func (s *ConstantsUtils) REQUEST_UNMAKING() string {
	return "0420"
}

func (s *ConstantsUtils) REQUEST_CANCELING() string {
	return "0400"
}

func (s *ConstantsUtils) REQUEST_PRE_AUTHORIZATION() string {
	return "0220"
}

func (this *ConstantsUtils) GET_METHOD_DESCRIPTION(mti string) string {
	switch mti {
	case this.REQUEST_AUTHORIZATION():
		return "REQUEST_AUTHORIZATION"
	case this.REQUEST_BRAND():
		return "REQUEST_BRAND"
	case this.REQUEST_CANCELING():
		return "REQUEST_CANCELING"
	case this.REQUEST_UNMAKING():
		return "REQUEST_UNMAKING"
	case this.REQUEST_PRE_AUTHORIZATION():
		return "REQUEST_PRE_AUTHORIZATION"
	default:
		return "Not found Mti"
	}
}

// PROCESSING CODE CONSTANTS
func (s *ConstantsUtils) PROCESSING_CODE_CREDIT() string {
	return "003000"
}

func (s *ConstantsUtils) PROCESSING_CODE_DEBIT() string {
	return "002000"
}

var ConstantsUtilsInst = &ConstantsUtils{}
