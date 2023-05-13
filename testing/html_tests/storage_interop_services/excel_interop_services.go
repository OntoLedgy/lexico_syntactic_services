package storage_interop_services

import (
	"fmt"
	"github.com/xuri/excelize/v2"
)

//TODO : replace/merge with storage interop services

type TableData struct {
	Headers []string
	Rows    [][]string
}

func ReadExcelFile(filePath string) []string {
	f, err := excelize.OpenFile(filePath)
	if err != nil {
		fmt.Println("Error opening Excel file:", err)
		return nil
	}

	var classIDs []string
	rows, err := f.GetRows("Sheet1")
	if err != nil {
		fmt.Println("Error reading rows from Sheet1:", err)
		return nil
	}

	for _, row := range rows {
		for _, colCell := range row {
			classIDs = append(classIDs, colCell)
		}
	}

	return classIDs
}

func WriteTableDataToSheet(sheetName string, data TableData, outputPath string) {

	f := excelize.NewFile()

	// Create a new sheet
	sheetIndex, _ := f.NewSheet(sheetName)

	// Set the active sheet
	f.SetActiveSheet(sheetIndex)

	// Write header row
	for col, header := range data.Headers {
		colName, _ := excelize.ColumnNumberToName(col + 1)
		cell := fmt.Sprintf("%s1", colName)
		f.SetCellValue(sheetName, cell, header)
	}

	// Write data rows
	for row, rowData := range data.Rows {
		for col, cellData := range rowData {
			colName, _ := excelize.ColumnNumberToName(col + 1)
			cell := fmt.Sprintf("%s%d", colName, row+2)
			f.SetCellValue(sheetName, cell, cellData)
		}
	}

	// Save the Excel file
	f.SaveAs(outputPath)

}

func WriteMapToExcel(
	data map[string]interface{},
	outputPath string,
	sheetName string) error {

	f := excelize.NewFile()

	sheet := sheetName
	// Create a new sheet
	sheetIndex, _ := f.NewSheet(sheet)

	// Set the active sheet
	f.SetActiveSheet(sheetIndex)

	// Write headers
	colIndex := 1
	for header := range data {
		cell, _ := excelize.CoordinatesToCellName(colIndex, 1)
		err := f.SetCellValue(sheet, cell, header)
		if err != nil {
			return err
		}
		colIndex++
	}

	// Write data
	colIndex = 1
	for _, value := range data {
		rowIndex := 2
		values, ok := value.([]interface{})
		if !ok {
			return fmt.Errorf("value is not a slice of interface{}")
		}
		for _, v := range values {
			cell, _ := excelize.CoordinatesToCellName(colIndex, rowIndex)
			err := f.SetCellValue(sheet, cell, v)
			if err != nil {
				return err
			}
			rowIndex++
		}
		colIndex++
	}

	// Save the Excel file
	err := f.SaveAs(outputPath)
	if err != nil {
		return err
	}

	return nil
}
