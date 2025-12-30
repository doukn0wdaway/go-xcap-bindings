//go:build linux

package internal

import (
	_ "embed"
	"os"
)

//go:embed bin/xcap-cli
var binary []byte

func embeddedBinary() (data []byte, name string) {
	return binary, "xcap-cli"
}

func binaryFileMode() os.FileMode {
	return 0o755
}
