package shared

import (
	"log"
	"reflect"
	"strings"
)

// Options that leakless should use for running child processes
type LeaklessOptions struct {
	// Windows-specific process options
	Windows *WindowsOptions
}

// Windows-specific process options
type WindowsOptions struct {
	// Whether to hide GUI window or not (default: `false`)
	//
	// Set this to `true` to avoid flashing console windows
	HideWindow bool
}

// Build options string from *LeaklessOptions struct to pass to leakless
func BuildOptionsString(options *LeaklessOptions) string {
	if options == nil {
		return ""
	}

	args := []string{}

	field, ok := reflect.TypeOf(options.Windows).Elem().FieldByName("HideWindow")
	if !ok {
		panic("Field not found")
	}
	log.Println(field.Tag.Get("flag"))

	if options.Windows != nil {
		if options.Windows.HideWindow {
			args = append(args, "--hide-window")
		}
	}

	return strings.Join(args, " ")
}

// Parse options string and recover *LeaklessOptions struct to use inside
// leakless
func ParseOptionsString(input string) *LeaklessOptions {
	options := &LeaklessOptions{
		Windows: &WindowsOptions{
			HideWindow: false,
		},
	}

	args := strings.Split(input, " ")

	for len(args) > 0 && strings.HasPrefix(args[0], "--") {
		arg := args[0]

		if arg == "--hide-window" {
			options.Windows.HideWindow = true
		}

		args = args[1:]
	}

	return options
}
