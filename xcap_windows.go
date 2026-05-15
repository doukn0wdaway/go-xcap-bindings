//go:build windows

package xcap

import (
	"os/exec"
	"syscall"

	"golang.org/x/sys/windows"
)

func configureCmd(cmd *exec.Cmd) {
	cmd.SysProcAttr = &syscall.SysProcAttr{
		HideWindow:    true,
		CreationFlags: windows.CREATE_NO_WINDOW,
	}
}
