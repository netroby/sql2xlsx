package main

import (
	"fmt"
	"github.com/tealeg/xlsx"
)

func main() {
	var (
		file *xlsx.File
		sheet *xlsx.Sheet
		row *xlsx.Row
		cell *xlsx.Cell
		err error
	)
	file = xlsx.NewFile()
	sheet = file.AddSheet("Sheet1")
	row = sheet.AddRow()
	cell = row.AddCell()
	cell.Value = "I am a cell!"
	cell.Value = "中文测试"
	err = file.Save("demo.xlsx")
	if err != nil {
		fmt.Printf(err.Error())
	}
}
