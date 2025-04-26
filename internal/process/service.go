package process

import (
	"encoding/csv"
	"fmt"
	"math"
	"os"
	"strings"
	"time"

	"github.com/andresfontan90/afip-compare/internal/config"
	"github.com/andresfontan90/afip-compare/internal/excel"
	"github.com/andresfontan90/afip-compare/internal/utils"
	"github.com/schollz/progressbar/v3"
)

func Process() error {
	LoadColumnNames()

	excelData, err := excel.SelectExcel()
	if err != nil {
		return err
	}

	defer excelData.FileData.Close()

	fmt.Printf("se cruzarán los datos de la hoja '%s' y la hoja '%s'", excelData.Sheet1, excelData.Sheet2)

	fmt.Println("")
	fmt.Println("Esto puede tardar varios minutos...")
	fmt.Println("")

	// Read sheet 1
	fmt.Println("Procesando hoja 1...")

	cols1, err := excelData.FileData.GetCols(excelData.Sheet1)
	if err != nil {
		return fmt.Errorf("error leyendo la hoja: %s. %s", excelData.Sheet1, err.Error())
	}

	barCols1 := progressbar.NewOptions(len(cols1),
		progressbar.OptionSetDescription("Leyendo columnas"),
		progressbar.OptionShowCount(),
		progressbar.OptionSetWidth(30),
	)

	for _, col := range cols1 {
		barCols1.Add(1)

		loadColFile(cuitSheet1, col, 1)
		loadColFile(dateSheet1, col, 1)
		loadColFile(puntoSheet1, col, 1)
		loadColFile(compSheet1, col, 1)
		loadColFile(mountSheet1, col, 1)
		loadColFile(mountNetoSheet1, col, 1)
		loadColFile(taxSheet1, col, 1)
		loadColFile(sourceSheet1, col, 1)
	}

	barCols1.Finish()

	if len(fileMap1) != 8 {
		return fmt.Errorf("error en hoja 1. La cantidad de columnas leidas es incorrecta")
	}

	if len(fileMap1[columNames[cuitSheet1]]) != len(fileMap1[columNames[taxSheet1]]) ||
		len(fileMap1[columNames[cuitSheet1]]) != len(fileMap1[columNames[dateSheet1]]) ||
		len(fileMap1[columNames[cuitSheet1]]) != len(fileMap1[columNames[puntoSheet1]]) ||
		len(fileMap1[columNames[cuitSheet1]]) != len(fileMap1[columNames[compSheet1]]) ||
		len(fileMap1[columNames[cuitSheet1]]) != len(fileMap1[columNames[mountSheet1]]) ||
		len(fileMap1[columNames[cuitSheet1]]) != len(fileMap1[columNames[sourceSheet1]]) ||
		len(fileMap1[columNames[cuitSheet1]]) != len(fileMap1[columNames[mountNetoSheet1]]) {

		return fmt.Errorf("error en hoja 1. Hay columnas que tienen mas registros que otras")
	}

	// Read sheet 2
	fmt.Println("")
	fmt.Println("")
	fmt.Println("Procesando hoja 2...")

	cols2, err := excelData.FileData.GetCols(excelData.Sheet2)
	if err != nil {
		return fmt.Errorf("error leyendo la hoja: %s. %s", excelData.Sheet2, err.Error())
	}

	barCols2 := progressbar.NewOptions(len(cols2),
		progressbar.OptionSetDescription("Leyendo columnas"),
		progressbar.OptionShowCount(),
		progressbar.OptionSetWidth(30),
	)

	for _, col := range cols2 {
		barCols2.Add(1)

		loadColFile(cuitSheet2, col, 2)
		loadColFile(dateSheet2, col, 2)
		loadColFile(puntoSheet2, col, 2)
		loadColFile(compSheet2, col, 2)
		loadColFile(mountSheet2, col, 2)
		loadColFile(mountNetoSheet2, col, 2)
		loadColFile(taxSheet2, col, 2)
		loadColFile(sourceSheet2, col, 2)
	}

	barCols2.Finish()

	if len(fileMap2) != 8 {
		return fmt.Errorf("error en hoja 2. La cantidad de columnas leidas es incorrecta")
	}

	if len(fileMap2[columNames[cuitSheet2]]) != len(fileMap2[columNames[taxSheet2]]) ||
		len(fileMap2[columNames[cuitSheet2]]) != len(fileMap2[columNames[dateSheet2]]) ||
		len(fileMap2[columNames[cuitSheet2]]) != len(fileMap2[columNames[puntoSheet2]]) ||
		len(fileMap2[columNames[cuitSheet2]]) != len(fileMap2[columNames[compSheet2]]) ||
		len(fileMap2[columNames[cuitSheet2]]) != len(fileMap2[columNames[mountSheet2]]) ||
		len(fileMap2[columNames[cuitSheet2]]) != len(fileMap2[columNames[sourceSheet2]]) ||
		len(fileMap2[columNames[cuitSheet2]]) != len(fileMap2[columNames[mountNetoSheet2]]) {

		return fmt.Errorf("error en hoja 2. Hay columnas que tienen mas registros que otras")
	}

	equalRows := make([]Entity, 0)
	equalRows2 := make([]Entity, 0)
	notInFile2 := make([]Entity, 0)
	notInFile1 := make([]Entity, 0)

	// Compare files
	fmt.Println("")
	fmt.Println("")
	fmt.Println("Cruzando la hoja 1 contra la hoja 2...")

	cuit1 := fileMap1[columNames[cuitSheet1]]
	cuit2 := fileMap2[columNames[cuitSheet2]]

	date1 := fileMap1[columNames[dateSheet1]]
	date2 := fileMap2[columNames[dateSheet2]]

	punto1 := fileMap1[columNames[puntoSheet1]]
	punto2 := fileMap2[columNames[puntoSheet2]]

	comp1 := fileMap1[columNames[compSheet1]]
	comp2 := fileMap2[columNames[compSheet2]]

	mount1 := fileMap1[columNames[mountSheet1]]
	mount2 := fileMap2[columNames[mountSheet2]]

	mountNeto1 := fileMap1[columNames[mountNetoSheet1]]
	mountNeto2 := fileMap2[columNames[mountNetoSheet2]]

	tax1 := fileMap1[columNames[taxSheet1]]
	tax2 := fileMap2[columNames[taxSheet2]]

	source1 := fileMap1[columNames[sourceSheet1]]
	source2 := fileMap2[columNames[sourceSheet2]]

	barRows1 := progressbar.NewOptions(len(cuit1),
		progressbar.OptionSetDescription("Comparando"),
		progressbar.OptionShowCount(),
		progressbar.OptionSetWidth(30),
	)

	for index1, value1 := range cuit1 {
		barRows1.Add(1)

		value1Formated := utils.NormalizeString(value1)
		var founded bool

		mount1Aux, err := utils.StringToNumber(mount1[index1])
		if err != nil {
			return fmt.Errorf("error parseando el importe %s", mount1[index1])
		}

		mountNeto1Aux, err := utils.StringToNumber(mountNeto1[index1])
		if err != nil {
			return fmt.Errorf("error parseando el importe %s", mountNeto1[index1])
		}

		tax1Aux, err := utils.StringToNumber(tax1[index1])
		if err != nil {
			return fmt.Errorf("error parseando el importe %s", tax1[index1])
		}

		date1Aux, err := utils.StringToDate(date1[index1])
		if err != nil {
			return fmt.Errorf("error parseando fecha %s", date1[index1])
		}

		for index2, value2 := range cuit2 {
			mount2Aux, err := utils.StringToNumber(mount2[index2])
			if err != nil {
				return fmt.Errorf("error parseando el importe %s", mount2[index2])
			}

			tax2Aux, err := utils.StringToNumber(tax2[index2])
			if err != nil {
				return fmt.Errorf("error parseando el importe %s", tax2[index2])
			}

			date2Aux, err := utils.StringToDate(date2[index2])
			if err != nil {
				return fmt.Errorf("error parseando fecha %s", date2[index2])
			}

			if strings.EqualFold(value1Formated, utils.NormalizeString(value2)) &&
				//strings.EqualFold(utils.NormalizeString(punto1[index1]), utils.NormalizeString(punto2[index2])) &&
				strings.EqualFold(utils.NormalizeString(comp1[index1]), utils.NormalizeString(comp2[index2])) &&
				utils.Abs(utils.DateDifference(date1Aux.UTC().Local(), date2Aux.UTC().Local())) <= config.AppConfig.DateToleranceDays &&
				math.Abs(math.Abs(mount1Aux)-math.Abs(mount2Aux)) <= config.AppConfig.AmountTolerance &&
				//math.Abs(mountNeto1Aux)-math.Abs(mountNeto2Aux) <= config.AppConfig.AmountTolerance &&
				math.Abs(math.Abs(tax1Aux)-math.Abs(tax2Aux)) <= config.AppConfig.AmountTolerance {

				founded = true

				entity := Entity{
					Cuit:   utils.NormalizeString(cuit1[index1]),
					Date:   date1Aux,
					Punto:  utils.NormalizeString(punto1[index1]),
					Comp:   utils.NormalizeString(comp1[index1]),
					Mount:  mount1Aux,
					Neto:   mountNeto1Aux,
					Tax:    tax1Aux,
					Source: source1[index1],
				}

				equalRows = append(equalRows, entity)
				break
			}
		}

		if !founded {
			entity := Entity{
				Cuit:   utils.NormalizeString(cuit1[index1]),
				Date:   date1Aux,
				Punto:  utils.NormalizeString(punto1[index1]),
				Comp:   utils.NormalizeString(comp1[index1]),
				Mount:  mount1Aux,
				Neto:   mountNeto1Aux,
				Tax:    tax1Aux,
				Source: source1[index1],
			}

			notInFile2 = append(notInFile2, entity)
		}
	}

	barRows1.Finish()

	fmt.Println("")
	fmt.Println("")
	fmt.Println("Cruzando la hoja 2 contra la hoja 1...")

	barRows2 := progressbar.NewOptions(len(cuit2),
		progressbar.OptionSetDescription("Comparando"),
		progressbar.OptionShowCount(),
		progressbar.OptionSetWidth(30),
	)

	for index2, value2 := range cuit2 {
		barRows2.Add(1)

		value2Formated := utils.NormalizeString(value2)
		var founded bool

		mount2Aux, err := utils.StringToNumber(mount2[index2])
		if err != nil {
			return fmt.Errorf("error parseando el importe %s", mount2[index2])
		}

		mountNeto2Aux, err := utils.StringToNumber(mountNeto2[index2])
		if err != nil {
			return fmt.Errorf("error parseando el importe  %s", mountNeto2[index2])
		}

		tax2Aux, err := utils.StringToNumber(tax2[index2])
		if err != nil {
			return fmt.Errorf("error parseando el importe  %s", tax2[index2])
		}

		date2Aux, err := utils.StringToDate(date2[index2])
		if err != nil {
			return fmt.Errorf("error parseando fecha  %s", date2[index2])
		}

		for index1, value1 := range cuit1 {
			mount1Aux, err := utils.StringToNumber(mount1[index1])
			if err != nil {
				return fmt.Errorf("error parseando el importe  %s", mount1[index1])
			}

			tax1Aux, err := utils.StringToNumber(tax1[index1])
			if err != nil {
				return fmt.Errorf("error parseando el importe  %s", tax1[index1])
			}

			date1Aux, err := utils.StringToDate(date1[index1])
			if err != nil {
				return fmt.Errorf("error parseando fecha  %s", date1[index1])
			}

			if strings.EqualFold(utils.NormalizeString(value1), value2Formated) &&
				//strings.EqualFold(utils.NormalizeString(punto1[index1]), utils.NormalizeString(punto2[index2])) &&
				strings.EqualFold(utils.NormalizeString(comp1[index1]), utils.NormalizeString(comp2[index2])) &&
				utils.Abs(utils.DateDifference(date1Aux.UTC().Local(), date2Aux.UTC().Local())) <= config.AppConfig.DateToleranceDays &&
				math.Abs(math.Abs(mount1Aux)-math.Abs(mount2Aux)) <= config.AppConfig.AmountTolerance &&
				//math.Abs(mountNeto1Aux)-math.Abs(mountNeto2Aux) <= config.AppConfig.AmountTolerance &&
				math.Abs(math.Abs(tax1Aux)-math.Abs(tax2Aux)) <= config.AppConfig.AmountTolerance {

				founded = true

				entity := Entity{
					Cuit:   utils.NormalizeString(cuit2[index2]),
					Date:   date2Aux,
					Punto:  utils.NormalizeString(punto2[index2]),
					Comp:   utils.NormalizeString(comp2[index2]),
					Mount:  mount2Aux,
					Neto:   mountNeto2Aux,
					Tax:    tax2Aux,
					Source: source2[index2],
				}

				equalRows2 = append(equalRows2, entity)

				break
			}
		}

		if !founded {
			entity := Entity{
				Cuit:   utils.NormalizeString(cuit2[index2]),
				Date:   date2Aux,
				Punto:  utils.NormalizeString(punto2[index2]),
				Comp:   utils.NormalizeString(comp2[index2]),
				Mount:  mount2Aux,
				Neto:   mountNeto2Aux,
				Tax:    tax2Aux,
				Source: source2[index2],
			}

			notInFile1 = append(notInFile1, entity)
		}
	}

	barRows2.Finish()

	// Create output files
	if err := createOuputFileCSV(equalRows, "cruces.csv"); err != nil {
		return err
	}
	if err := createOuputFileCSV(equalRows2, "cruces2.csv"); err != nil {
		return err
	}
	if err := createOuputFileCSV(notInFile2, "notInFile2.csv"); err != nil {
		return err
	}
	if err := createOuputFileCSV(notInFile1, "notInFile1.csv"); err != nil {
		return err
	}

	// Validate results
	if len(cuit1) != len(equalRows)+len(notInFile2) || len(cuit2) != len(equalRows)+len(notInFile1) {
		return fmt.Errorf("inconsitencia en los datos procesados. Cantidades: \nfile1:%d\nfile2:%d\nprocesado1:%d\nprocesado2:%d",
			len(cuit1), len(cuit2), len(equalRows)+len(notInFile2), len(equalRows)+len(notInFile1))
	}

	return nil
}

func LoadColumnNames() {
	columNames[cuitSheet1] = "cuit"
	columNames[cuitSheet2] = "cuit"
	columNames[dateSheet1] = "fecha"
	columNames[dateSheet2] = "fecha"
	columNames[puntoSheet1] = "punto"
	columNames[puntoSheet2] = "punto"
	columNames[compSheet1] = "comprobante"
	columNames[compSheet2] = "comprobante"
	columNames[mountSheet1] = "importe"
	columNames[mountSheet2] = "importe"
	columNames[mountNetoSheet1] = "importeNeto"
	columNames[mountNetoSheet2] = "importeNeto"
	columNames[taxSheet1] = "impuesto"
	columNames[taxSheet2] = "impuesto"
	columNames[sourceSheet1] = "fuente"
	columNames[sourceSheet2] = "fuente"
}

func createOuputFileCSV(data []Entity, fileName string) error {
	csvFile, err := os.Create(fileName)
	if err != nil {
		return fmt.Errorf("error creando csv %s: %s", fileName, err.Error())
	}
	defer csvFile.Close()

	csvwriter := csv.NewWriter(csvFile)
	csvwriter.Comma = []rune(config.AppConfig.CsvSeparator)[0]

	headers := []string{"CUIT", "FECHA", "PUNTO", "COMPROBANTE", "IMPORTE", "IMPORTE NETO", "IMPUESTO", "FUENTE"}
	_ = csvwriter.Write(headers)

	for _, d := range data {
		_ = csvwriter.Write([]string{d.Cuit, d.Date.Format(time.DateOnly), d.Punto, d.Comp, fmt.Sprintf("%.2f", d.Mount), fmt.Sprintf("%.2f", d.Neto), fmt.Sprintf("%.2f", d.Tax), d.Source})
	}
	csvwriter.Flush()

	return nil
}

func loadColFile(colName string, colList []string, fileNumber int) {
	var flgFoundCol bool
	values := make([]string, 0)

	for _, cell := range colList {
		if !flgFoundCol {
			if strings.EqualFold(utils.NormalizeString(cell), colName) {
				flgFoundCol = true
			}
		} else {
			if !utils.IsEmptyString(cell) {
				values = append(values, cell)
			}
		}
	}

	if len(values) > 0 {
		if fileNumber == 1 {
			fileMap1[columNames[colName]] = values
		}
		if fileNumber == 2 {
			fileMap2[columNames[colName]] = values
		}
	}
}
