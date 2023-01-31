package draw_line

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"math"
	"os"
)

const (
	dx = 500
	dy = 300
)

type Putpixel func(x,y int)

func  drawline(x0, y0, x1, y1 int, brush Putpixel) {
	dx := math.Abs(float64(x1-x0))
	dy := math.Abs(float64(y1-y0))
	sx, sy := 1, 1
	if x0 >= x1 {
		sx = -1
	}
	if y0 >= y1 {
		sy = -1
	}
	err := dx - dy
	for {
		brush(x0, y0)
		if x0 == x1 && y0 == y1 {
			return
		}
		e2 := err * 2
		if e2 > -dy {
			err -= dy
			x0 += sx
		}
		if e2 < dx {
			err += dx
			y0 += sy
		}
	}
}

func DrawLine() {
	file, err := os.Create("test.jpg")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	nrgba := image.NewNRGBA(image.Rect(0, 0, dx, dy))
	drawline(1, 1, dx-2, dy-2, func(x, y int) {

		nrgba.Set(x, y, color.RGBA{uint8(x), uint8(y), 0, 255})
	})
	for y := 0; y < dy; y++ {
		nrgba.Set(1, y, color.White)
		nrgba.Set(dx-1, y, color.White)
	}
	err = jpeg.Encode(file, nrgba, &jpeg.Options{100})      //图像质量值为100，是最好的图像显示
	if err != nil {
		fmt.Println(err)
	}

}