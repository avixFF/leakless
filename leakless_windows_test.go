// +build windows

package leakless_test

import (
	"testing"

	"github.com/ysmood/leakless"
	"github.com/ysmood/leakless/pkg/shared"
)

func TestHideWindow(t *testing.T) {
	launcher := leakless.New(
		&shared.LeaklessOptions{
			Windows: &shared.WindowsOptions{
				HideWindow: true,
			},
		},
	)

	// Run "test" program that spawns a "zombie" background process.
	cmd := launcher.Command(p("dist/test"), "on")

	cmd.Run()
}
