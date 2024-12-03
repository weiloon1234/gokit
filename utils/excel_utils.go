package utils

import (
	"fmt"
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
			switch v := value.(type) {
			case string:
				// Set as string value
				err := f.SetCellValue(sheetName, cell, v)
				if err != nil {
					return nil, err
				}
			case float64, float32, int, int64:
				// Convert numbers to string to prevent scientific notation issues
				err := f.SetCellValue(sheetName, cell, fmt.Sprintf("'%v", v)) // Prepend apostrophe to treat as text
				if err != nil {
					return nil, err
				}
			case bool:
				// Set boolean values directly
				err := f.SetCellValue(sheetName, cell, v)
				if err != nil {
					return nil, err
				}
			case nil:
				// Handle nil values (set cell as empty)
				err := f.SetCellValue(sheetName, cell, "")
				if err != nil {
					return nil, err
				}
			default:
				// Default case to treat value as string
				err := f.SetCellValue(sheetName, cell, fmt.Sprintf("'%v", v)) // Prepend apostrophe to treat as text
				if err != nil {
					return nil, err
				}
			}
		}
	}

	return f, nil
}
