package opencv

import (
	"fmt"
	"gocv.io/x/gocv"
	"image"
	"image/color"
)

const MinimumArea = 3000

func OutLinePic() {
	deviceID := 0

	webcam, err := gocv.OpenVideoCapture(deviceID)
	if err != nil {
		fmt.Printf("无法打开摄像头: %v\n", deviceID)
		return
	}
	defer webcam.Close()

	window := gocv.NewWindow("Main")
	defer window.Close()

	// 多开两窗口展示图片变化过程
	Delta := gocv.NewWindow("Delta")
	defer Delta.Close()
	Thresh := gocv.NewWindow("Thresh")
	defer Thresh.Close()

	img := gocv.NewMat()
	defer img.Close()

	imgDelta := gocv.NewMat()
	defer imgDelta.Close()

	imgThresh := gocv.NewMat()
	defer imgThresh.Close()

	mog2 := gocv.NewBackgroundSubtractorMOG2WithParams(500, 25, false)
	defer mog2.Close()

	status := "ready!!"                     //初始提示词
	statusColor := color.RGBA{0, 255, 0, 0} //提示词颜色
	fmt.Printf("开始读取设备: %v\n", deviceID)
	for {
		if ok := webcam.Read(&img); !ok {
			fmt.Printf("设备已关闭: %v\n", deviceID)
			return
		}
		if img.Empty() {
			continue
		}

		status = "ready!!"                     // 默认为ready
		statusColor = color.RGBA{0, 255, 0, 0} // 默认绿色

		// 获取前景
		mog2.Apply(img, &imgDelta)
		// 展示前景
		Delta.IMShow(imgDelta)

		// 二值化
		gocv.Threshold(imgDelta, &imgThresh, 25, 255, gocv.ThresholdBinary)

		// 膨胀
		gocv.Dilate(imgThresh, &imgThresh, gocv.NewMat())

		// 轮廓
		contours := gocv.FindContours(imgThresh, gocv.RetrievalExternal, gocv.ChainApproxSimple)
		for i, c := range contours {
			area := gocv.ContourArea(c)
			// 过滤
			if area < MinimumArea {
				continue
			}

			status = "move!!"
			statusColor = color.RGBA{0, 0, 255, 0}
			// 画轮廓
			gocv.DrawContours(&img, contours, i, statusColor, 2)
			// 画边框
			rect := gocv.BoundingRect(c)
			gocv.Rectangle(&img, rect, color.RGBA{255, 0, 0, 0}, 2)
		}

		// 画提示词
		gocv.PutText(&img, status, image.Pt(10, 50), gocv.FontHersheyPlain, 2.4, statusColor, 2)
		// 展示膨胀后图像
		Thresh.IMShow(imgThresh)
		// 展示源图
		window.IMShow(img)
		if window.WaitKey(1) == 27 {
			break
		}
	}
}
