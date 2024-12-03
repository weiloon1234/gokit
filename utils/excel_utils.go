package utils

import (
	"github.com/xuri/excelize/v2"
)

func GenerateExcelFile(fileName string, headerColumns []string, data [][]interface{}) (*excelize.File, error) {
	f := excelize.NewFile()

	sheetName := "Sheet1"
	_, err := f.NewSheet(sheetName)
	if err != nil {
		return nil, err
	}

	// Write Header
	for i, header := range headerColumns {
		cell, _ := excelize.CoordinatesToCellName(i+1, 1)
		err := f.SetCellValue(sheetName, cell, header)
		if err != nil {
			return nil, err
		}
	}

	// Write Data Rows
	for rowIndex, row := range data {
		for colIndex, value := range row {
			cell, _ := excelize.CoordinatesToCellName(colIndex+1, rowIndex+2)
			err := f.SetCellValue(sheetName, cell, value)
			if err != nil {
				return nil, err
			}
		}
	}

	return f, nil
}
