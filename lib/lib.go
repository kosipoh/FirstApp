package lib

import (
	libEx "github.com/360EntSecGroup-Skylar/excelize"
)

func ReadEx() string {
	file := libEx.NewFile()
	file, _ = libEx.OpenFile("file/Foresma.xlsx")
	return file.GetRows("Лист1")[0][0]

}
