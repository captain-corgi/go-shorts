package qrcode

import (
	"log/slog"

	"github.com/skip2/go-qrcode"
)

// Generate create qr code from string
func Generate(data string, outputFilePath string) error {
	err := qrcode.WriteFile(data, qrcode.Medium, 256, outputFilePath)
	if err != nil {
		slog.Error("Sorry couldn't create qrcode:,%v", err)
		return err
	}
	return nil
}
