//go:build windows

package internal

import (
	_ "embed"
	"os"
)

//go:embed bin/xcap-cli.exe
var binary []byte

func embeddedBinary() (data []byte, name string) {
	return binary, "xcap-cli.exe"
}

func binaryFileMode() os.FileMode {
	return 0o666
}
