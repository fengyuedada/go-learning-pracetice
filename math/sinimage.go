package mathImage

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
	"os"
)

//图片大小
const size = 300

//根据给定大小创建灰度图
func Pic() {
	resPic := image.NewGray(image.Rect(0, 0, size, size))

	//画布填充
	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			//填充颜色
			resPic.SetGray(x, y, color.Gray{Y: 255})
		}
	}

	//画图
	for x := 0; x < size; x++ {
		s := float64(x) * 2 * math.Pi / size

		y := size/2 - math.Sin(s)*size/2

		resPic.SetGray(x, int(y), color.Gray{})
	}

	//创建文件
	file, err := os.Create("sin.png")
	if err != nil {
		log.Fatal(err)
	}

	png.Encode(file, resPic)

	file.Close()
}
