package excel

import "github.com/xuri/excelize/v2"

type ExcelData struct {
	FileName string
	FileData *excelize.File
	Sheet1   string
	Sheet2   string
}
