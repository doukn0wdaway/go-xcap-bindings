package xcap

import (
	"fmt"
	"os/exec"
	"strconv"

	"github.com/doukn0wdaway/go-xcap-bindings/internal"
)

func run(args ...string) (string, error) {
	binaryPath, err := internal.ExtractBinary()
	if err != nil {
		return "", err
	}

	out, err := exec.Command(binaryPath, args...).CombinedOutput()
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

func GetScreenshot(monitorId int64) (string, error) {

	out, err := run("screenshot", strconv.Itoa(int(monitorId)))
	if err != nil {
		return "", fmt.Errorf("xcap-cli failed: %w\n%s", err, out)
	}

	return string(out), nil
}
