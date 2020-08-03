package Excelize

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
)

func ReadExcelize() {
	f, err := excelize.OpenFile("./Book1.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	// 获取工作表中指定单元格的值
	cell, err := f.GetCellValue("Sheet1", "B2")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(cell)
	// 获取Sheet1 上所有的单元格
	rows, err := f.GetRows("Sheet2")
	for _, row := range rows {
		for _, colCell := range row {
			fmt.Println(colCell, "\t")
		}
		fmt.Println()
	}
}
