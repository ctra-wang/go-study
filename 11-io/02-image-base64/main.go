package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"github.com/fogleman/gg"
	"github.com/polds/imgbase64"
	"image/jpeg"
	"io"
	"log"
	"os"
	"strings"
)

func main() {

	//
	img := imgbase64.FromRemote("www.baidu.com")
	//fmt.Println(img)

	//imgUrl := ""
	//
	//h := md5.New()
	//h.Write([]byte(imgUrl))
	//key := hex.EncodeToString(h.Sum(nil))
	//fmt.Println(key)
	//// 查询redis
	//
	////获取远端图片
	//res, err := http.Get(imgUrl)
	//if err != nil {
	//	fmt.Println("A error occurred!")
	//	return
	//}
	//defer res.Body.Close()
	//fmt.Println("-------------")
	//
	//data, err := io.ReadAll(res.Body)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//imageBase64 := base64.StdEncoding.EncodeToString(data)
	i := strings.Index(img, ",")
	dec := base64.NewDecoder(base64.StdEncoding, strings.NewReader(img[i+1:]))
	//fmt.Println(dec)

	fmt.Println("-------===========")
	f, err := os.Create("ddd.jpg")
	if err != nil {
		log.Println(err)
	}
	defer os.RemoveAll(f.Name())
	_, err = io.Copy(f, dec)
	if err != nil {
		log.Println(err)
	}
	log.Println("---图片接收完成------名称是", "ddd.jpg")

	//var waterImage image.Image
	waterImage, err := gg.LoadImage("ddd.jpg")
	// 如果没有加载到图片，则创建图片
	if err != nil {
		fmt.Println("-------")
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

	buffer := bytes.NewBuffer(nil)
	//dc.EncodePNG(buffer)
	//b := buffer.Bytes()
	//
	//str := base64.StdEncoding.EncodeToString(b)
	//fmt.Println(str)
	//name := "/Users/ctra_wl/Desktop/go-test/test_pic_1.png"
	//
	//newfile, err := os.Create(name)
	//if err != nil {
	//	fmt.Println(err.Error())
	//	return
	//}
	////defer newfile.Close()
	//// 将文件保存输出，并设置压缩比
	err = jpeg.Encode(buffer, dc.Image(), &jpeg.Options{Quality: 60})
	b := buffer.Bytes()
	str := base64.StdEncoding.EncodeToString(b)
	fmt.Println(str)
	//if err != nil {
	//	fmt.Println(err.Error())
	//	return
	//}
	//return

	//f, err := os.Create("/imgs/ddd.png")
	//if err != nil {
	//	log.Println(err)
	//}
	////defer f.Close()
	//_, err = io.Copy(f, dec)
	//if err != nil {
	//	log.Println(err)
	//}
	//log.Println("---图片接收完成------名称是", "ddd.png")

	//var base64Encoding string
	//mimeType := http.DetectContentType(data)
	//
	//switch mimeType {
	//case "image/jpeg":
	//	base64Encoding += "data:image/jpeg;base64,"
	//case "image/png":
	//	base64Encoding += "data:image/png;base64,"
	//}
	//
	//imageBase64 := base64.StdEncoding.EncodeToString(data)
	//
	//// base64 string  存入redis
	//
	////base64.NewEncoding()
	//fmt.Println("base64", imageBase64)

}
