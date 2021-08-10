package main

import (
	"image/color"

	"github.com/skip2/go-qrcode"
	"springmars.com/qrcode/pojo"
	"springmars.com/qrcode/qr"
)

func create() {
	qc := pojo.NewQrCode()
	qc.Content = "springmars.com"
	qc.Level = qrcode.High
	qc.Size = 256
	qc.BackgroundColor = color.RGBA{250, 205, 90, 255}
	qc.ForegroundColor = color.Black
	// 文件目录必须是已经存在的，不然无法生存
	qc.FileName = "./qr/springmars.png"

	qc.CreateQrCode()
}

func main() {
	qr.TestQrCode()
	qr.CustomQrCode()
	create()
}
