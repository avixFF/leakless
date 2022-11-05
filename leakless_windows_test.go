// +build windows

package leakless_test

import (
	"testing"

	"github.com/ysmood/leakless"
	"github.com/ysmood/leakless/pkg/shared"
	"github.com/ysmood/leakless/pkg/utils"
)

func buildWindowsBinary() {
	utils.Exec("go", "dist",
		"build",
		"-ldflags", "-H=windowsgui",
		"-o", "test_windows.exe",
		"../cmd/test",
	)
}

func TestHideWindow(t *testing.T) {
	launcher := leakless.New(
		&shared.LeaklessOptions{
			Windows: &shared.WindowsOptions{
				HideWindow: true,
			},
		},
	)

	// Run "test" program that spawns a "zombie" background process.
	cmd := launcher.Command(p("dist/test_windows"), "on")

	// TODO: Figure out why `test_windows.exe`, compiled with `-H=windowsgui`
	// always makes `leakless.exe` visible, even when `HideWindow` is `true`.

	cmd.Run()
}
