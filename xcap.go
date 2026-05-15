package xcap

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"image"
	"io"
	"os"
	"os/exec"
	"strconv"

	"github.com/doukn0wdaway/go-xcap-bindings/internal"
)

func run(args ...string) (string, error) {
	binaryPath, err := internal.ExtractBinary()
	if err != nil {
		return "", err
	}
	cmd := exec.Command(binaryPath, args...)

	configureCmd(cmd)

	out, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("xcap-cli failed: %w\n%s", err, out)
	}

	return string(out), nil
}

func GetMonitors() (string, error) {
	out, err := run("get-monitors")
	if err != nil {
		return "", fmt.Errorf("xcap-cli failed: %w\n%s", err, out)
	}

	return string(out), nil
}

func GetScreenshot(monitorId int64) (*image.RGBA, error) {
	binaryPath, err := internal.ExtractBinary()
	if err != nil {
		return nil, err
	}

	cmd := exec.Command(binaryPath, "screenshot", strconv.Itoa(int(monitorId)))

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}
	cmd.Stderr = os.Stderr

	if err := cmd.Start(); err != nil {
		return nil, err
	}

	data, err := io.ReadAll(stdout)
	if err != nil {
		return nil, err
	}

	if err := cmd.Wait(); err != nil {
		return nil, err
	}

	r := bytes.NewReader(data)

	var w, h uint32
	if err := binary.Read(r, binary.LittleEndian, &w); err != nil {
		return nil, err
	}
	if err := binary.Read(r, binary.LittleEndian, &h); err != nil {
		return nil, err
	}

	pix := make([]byte, w*h*4)

	if _, err := io.ReadFull(r, pix); err != nil {
		return nil, err
	}

	img := &image.RGBA{
		Pix:    pix,
		Stride: int(w * 4),
		Rect:   image.Rect(0, 0, int(w), int(h)),
	}
	return img, nil
}
