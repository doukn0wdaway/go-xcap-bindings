package internal

import (
	"os"
	"path/filepath"
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

		data, name := embeddedBinary() // TODO: ADD CACHING
		path := filepath.Join(dir, name)

		e = os.WriteFile(path, data, binaryFileMode())
		if e != nil {
			err = e
			return
		}

		extractedPath = path
	})

	return extractedPath, err
}
