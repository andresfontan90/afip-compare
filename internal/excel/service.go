package excel

import (
	"fmt"
	"strings"

	"github.com/andresfontan90/afip-compare/internal/utils"
	"github.com/sqweek/dialog"
	"github.com/xuri/excelize/v2"
)

func SelectExcel() (*ExcelData, error) {
	// Select file
	selectedFile, err := dialog.File().Title("selecciona un archivo excel").Filter("excel", "xls", "xlsx").Load()
	if err != nil {
		return nil, fmt.Errorf("error seleccionando excel: %s", err.Error())
	}

	if utils.IsEmptyString(selectedFile) {
		return nil, fmt.Errorf("no seleccionó ningón archivo")
	}

	// Load excel file
	excelFile, err := excelize.OpenFile(selectedFile)
	if err != nil {
		return nil, fmt.Errorf("error abriendo excel: %s", err.Error())
	}

	defer excelFile.Close()

	sheetList := excelFile.GetSheetList()

	// Validate that the Excel file has at least two sheets
	if len(sheetList) < 2 {
		return nil, fmt.Errorf("el archivo debe contener al menos 2 hojas para poder procesar")
	}

	// Select the first two sheets
	if utils.IsEmptyString(sheetList[0]) || utils.IsEmptyString(sheetList[1]) {
		return nil, fmt.Errorf("el archivo Excel debe contener nombres válidos en ambas hojas")
	}

	data := &ExcelData{
		FileName: selectedFile,
		FileData: excelFile,
		Sheet1:   sheetList[0],
		Sheet2:   sheetList[1],
	}

	return data, nil
}

func readExcel(excelFile *excelize.File, sheetName string) error {
	rows, err := excelFile.GetRows(sheetName)
	if err != nil {
		return fmt.Errorf("error leyendo filas: %s", err)
	}

	for rowIdx, row := range rows {
		for colIdx := range row {
			// Convertir índice de columna a letra (A, B, C, ...)
			colLetter, err := excelize.ColumnNumberToName(colIdx + 1)
			if err != nil {
				return fmt.Errorf("error obteniendo letra de columna: %s", err)
			}

			// Obtener valor de celda tipo "A1", "B2", etc.
			cellRef := fmt.Sprintf("%s%d", colLetter, rowIdx+1)
			value, err := excelFile.GetCellValue(sheetName, cellRef)
			if err != nil {
				return fmt.Errorf("error obteniendo valor en %s: %s", cellRef, err)
			}

			// Opcional: normalizar o parsear si es número
			value = normalizeDecimal(value)

			fmt.Printf("Valor en %s: %s\n", cellRef, value)
		}
	}

	return nil

}

func normalizeDecimal(input string) string {
	input = strings.TrimSpace(input)

	// Si tiene ',' como decimal y '.' como miles → lo convertimos
	if strings.Contains(input, ",") && strings.Contains(input, ".") {
		input = strings.ReplaceAll(input, ".", "")
		input = strings.ReplaceAll(input, ",", ".")
		return input
	}

	// Si solo tiene ',' asumimos que es decimal
	if strings.Contains(input, ",") && !strings.Contains(input, ".") {
		input = strings.ReplaceAll(input, ",", ".")
		return input
	}

	// Si ya tiene '.' como decimal, lo dejamos
	return input
}
