package internal

import (
	_ "embed"
	"runtime"
)

// Linux
//
//go:embed bin/xcap-cli
var linuxAmd64 []byte

// Windows
//
//go:embed bin/xcap-cli
var windowsAmd64 []byte

func EmbeddedBinary() (data []byte, name string) {
	switch runtime.GOOS + "/" + runtime.GOARCH {
	case "linux/amd64":
		return linuxAmd64, "xcap-cli"
	case "windows/amd64":
		return windowsAmd64, "xcap-cli.exe"
	default:
		panic("unsupported platform: " + runtime.GOOS + "/" + runtime.GOARCH)
	}
}
