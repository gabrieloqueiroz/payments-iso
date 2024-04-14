package service

import (
	"fmt"
	"iso8583/util"
	"os"
	"time"
)

type FileService struct{}

func (s *FileService) GetFileName(methodType string) string {
	return time.Now().Format("2006-01-02 15:04:05") + " - " + methodType
}

func (s *FileService) GetFilePath() string {
	basePath := "../commands/" + time.Now().Format("2006-01-02")
	_, err := os.Stat(basePath)

	if os.IsNotExist(err) {
		errDir := os.MkdirAll(basePath, 0755)

		if errDir != nil {
			fmt.Println("Erro ao criar diretório:", errDir)
		} else {
			fmt.Println("Diretório criado com sucesso.")
		}
	}

	return basePath + "/"
}

func (s *FileService) CreateCommand(methodType string, hexString string, isSource bool) {
	if isSource {
		methodType = util.ConstantsUtilsInst.GET_METHOD_DESCRIPTION_SOURCE(methodType)
	} else {
		methodType = util.ConstantsUtilsInst.GET_METHOD_DESCRIPTION_TARGET(methodType)
	}

	path := FileServiceInst.GetFilePath()
	fileName := FileServiceInst.GetFileName(methodType)

	file, err := os.OpenFile(path+fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println(err)
	}

	header := "echo '\\x02\\xe4"
	fotter := "' | nc localhost 7117"

	fullCommand := header + util.StringUtilInst.ConvertAsciiToHex(hexString) + fotter

	file.WriteString(fullCommand)

	file.Close()
}

var FileServiceInst = &FileService{}
