// +build windows

package main

import (
	"os"
	"os/exec"
	"strconv"
	"syscall"

	"github.com/ysmood/leakless/pkg/shared"
)

func osSetupCmd(cmd *exec.Cmd, opts *shared.LeaklessOptions) error {
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: opts.Windows.HideWindow}
	return nil
}

func kill(p *os.Process) {
	_ = exec.Command("taskkill", "/t", "/f", "/pid", strconv.Itoa(p.Pid)).Run()
}
