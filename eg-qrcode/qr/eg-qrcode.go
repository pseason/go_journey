package qr

import (
	"image/color"
	"log"

	"github.com/skip2/go-qrcode"
)

func TestQrCode() {
	// 测试qrcode 生成
	qrcode.WriteFile("www.baidu.com", qrcode.Medium, 256, "./qr/baidu.png")
}

// 自定义qrcode
func CustomQrCode() {
	qr, err := qrcode.New("www.hao123.com", qrcode.High)
	if err == nil {
		qr.BackgroundColor = color.RGBA{250, 205, 90, 255}
		qr.ForegroundColor = color.Black
		qr.WriteFile(255, "./qr/hao123.png")
	} else {
		log.Fatal("自定义qrcode error", err)
	}
}
