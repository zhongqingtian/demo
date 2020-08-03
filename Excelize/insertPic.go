package Excelize

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
)

func InsertPic() {
	f, err := excelize.OpenFile("./Book3.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	// 插入图片
	/*err = f.AddPicture("Sheet1", "A2", "./image1.jpg", "")
	if err != nil {
		fmt.Println(err)
	}*/
	// 在工作表中插入图片，并设置图片的缩放比例
	err = f.AddPicture("Sheet1", "D2", "./image2.jpg", `{"x_scale": 0.5, "y_scale": 0.5}`)
	if err != nil {
		fmt.Println(err)
	}
	// 在工作表中插入图片，并设置图片的打印属性
	err = f.AddPicture("Sheet1", "H2", "./image3.gif", `{"x_offset": 15, "y_offset": 10, "print_obj": true, "lock_aspect_ratio": false, "locked": false}`)
	if err != nil {
		fmt.Println(err)
	}
	// 保存文件
	err = f.Save()
	if err != nil {
		fmt.Println(err)
	}
}
