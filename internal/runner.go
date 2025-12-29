package internal

import (
	"os"
	"path/filepath"
	"runtime"
	"sync"
)

var (
	extractedPath string
	once          sync.Once
)

func ExtractBinary() (string, error) {
	var err error

	once.Do(func() {
		dir, e := os.UserCacheDir()
		if e != nil {
			err = e
			return
		}

		dir = filepath.Join(dir, "xcap")
		_ = os.MkdirAll(dir, 0o755)

		data, name := EmbeddedBinary()
		path := filepath.Join(dir, name)

		if _, e := os.Stat(path); e == nil {
			extractedPath = path
			return
		}

		mode := os.FileMode(0o755)
		if runtime.GOOS == "windows" {
			mode = 0o666
		}

		e = os.WriteFile(path, data, mode)
		if e != nil {
			err = e
			return
		}

		extractedPath = path
	})

	return extractedPath, err
}
