package main

import (
	"fmt"
	"github.com/fogleman/gg"
	"image/jpeg"
	"os"
)

func main() {

	// 加载图片
	waterImage, err := gg.LoadJPG("11-io/01-image/1672197385562.jpg")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	// 通过图片实例化gg
	dc := gg.NewContextForImage(waterImage)
	dc.SetRGBA(1, 1, 1, 0)

	dc.SetRGB(0, 0, 0)
	// 加载字体，设置大小
	if err := dc.LoadFontFace("/Users/ctra_wl/Downloads/longshuhongheicuti.ttf", 16); err != nil {
		fmt.Println(err.Error())
		return
	}
	// 给图片添加文字，位置在（x, y） 处
	s := "张三"
	dc.DrawStringWrapped(s, 220, 390, 0, 0, 140, 1, gg.AlignCenter)

	s1 := "1200fenmi"
	dc.DrawStringWrapped(s1, 220, 460, 0, 0, 140, 1, gg.AlignCenter)
	name := "/Users/ctra_wl/Desktop/go-test/test_pic_1.png"

	newfile, err := os.Create(name)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	//defer newfile.Close()
	// 将文件保存输出，并设置压缩比
	err = jpeg.Encode(newfile, dc.Image(), &jpeg.Options{Quality: 60})
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	return

}
