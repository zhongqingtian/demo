package Excelize

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
)

func CreateExcelize() {
	// 创建一个默认模版的文件
	f := excelize.NewFile()
	// 创建工作表
	index := f.NewSheet("Sheet2")
	// 设置单元格的值
	f.SetCellValue("Sheet2", "A1", "hello world")
	f.SetCellValue("Sheet2", "A2", "hello world")
	f.SetCellValue("Sheet2", "A3", "hello world")
	f.SetCellValue("Sheet2", "A4", "hello world")
	f.SetCellValue("Sheet2", "A5", "hello world")
	f.SetCellValue("Sheet1", "B1", 100)
	f.SetCellValue("Sheet1", "B2", 100)
	f.SetCellValue("Sheet1", "B3", 100)
	f.SetCellValue("Sheet1", "B3", 100)

	// 设置工作簿的默认工作表
	f.SetActiveSheet(index)
	// 根据指定路径保存文件
	err := f.SaveAs("./Book1.xlsx")
	if err != nil {
		fmt.Println(err)
	}
}
