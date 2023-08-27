package main

import (
	"log/slog"

	"github.com/captain-corgi/go-shorts/qrcode"
)

func main() {
	if err := qrcode.Generate("Hello, world", "hello.png"); err != nil {
		slog.Info("No content created")
	}
}
