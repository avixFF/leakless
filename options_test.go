package leakless_test

import (
	"strings"
	"testing"

	"github.com/ysmood/leakless/pkg/shared"
)

func TestOptionsBuilder(t *testing.T) {
	if len(shared.BuildOptionsString(nil)) != 0 {
		t.Fatalf(
			"Options string must be empty when options are empty, got: \"%s\"",
			shared.BuildOptionsString(nil),
		)
	}

	options := &shared.LeaklessOptions{}

	// Test --hide-window flag

	if strings.Contains(shared.BuildOptionsString(options), "--hide-window") {
		t.Fatalf(
			"Options string must NOT contain \"--hide-window\" flag, got: \"%s\"",
			shared.BuildOptionsString(nil),
		)
	}

	options = &shared.LeaklessOptions{
		Windows: &shared.WindowsOptions{
			HideWindow: true,
		},
	}

	if !strings.Contains(shared.BuildOptionsString(options), "--hide-window") {
		t.Fatalf(
			"Options string must contain \"--hide-window\" flag, got: \"%s\"",
			shared.BuildOptionsString(nil),
		)
	}
}

func TestOptionsParser(t *testing.T) {
	options := shared.ParseOptionsString("")

	if options == nil {
		t.Fatal("Options must not be nil after parsing")
	}

	// Test --hide-window flag

	if options.Windows.HideWindow != false {
		t.Fatal("LeaklessOptions.Windows.HideWindow must be false by default")
	}

	options = shared.ParseOptionsString("--hide-window")

	if options.Windows.HideWindow != true {
		t.Fatal("LeaklessOptions.Windows.HideWindow must be true when the \"--hide-window\" flag is specified")
	}
}
