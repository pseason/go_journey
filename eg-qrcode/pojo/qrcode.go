package pojo

import (
	"fmt"
	"image/color"

	"github.com/skip2/go-qrcode"
)

/*自定义QrCode*/
type QrCode struct {
	Content         string
	Level           qrcode.RecoveryLevel
	Size            int
	ForegroundColor color.Color
	BackgroundColor color.Color
	FileName        string
}

func NewQrCode() *QrCode {
	fmt.Println("")
	return new(QrCode)
}

/** 创建qrcode */
func (qc *QrCode) CreateQrCode() error {
	qr, err := qrcode.New(qc.Content, qc.Level)
	if err == nil {
		qr.BackgroundColor = qc.BackgroundColor
		qr.ForegroundColor = qc.ForegroundColor
		qr.WriteFile(qc.Size, qc.FileName)
		return nil
	} else {
		return err
	}
}
