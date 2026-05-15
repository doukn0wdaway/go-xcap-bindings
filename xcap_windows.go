//go:build windows

package xcap

import (
	"os/exec"
	"syscall"
)

func configureCmd(cmd *exec.Cmd) {
	cmd.SysProcAttr = &syscall.SysProcAttr{
		CreationFlags: syscall.SW_HIDE,
	}
}
