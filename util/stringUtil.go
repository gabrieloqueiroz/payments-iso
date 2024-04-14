package util

import (
	"fmt"
)

type StringUtil struct{}

func (s *StringUtil) ConvertAsciiToHex(msg string) string {
	var hexString string
	for _, char := range msg {
		hexString += fmt.Sprintf("\\x%02x", char)
	}

	return hexString
}

var StringUtilInst = &StringUtil{}
