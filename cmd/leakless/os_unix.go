// +build !windows

package main

import (
	"os"
	"os/exec"
	"syscall"

	"github.com/ysmood/leakless/pkg/shared"
)

func osSetupCmd(cmd *exec.Cmd, opts *shared.LeaklessOptions) error {
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
	return nil
}

func kill(p *os.Process) {
	_ = syscall.Kill(-p.Pid, syscall.SIGKILL)
}
