package shared

import (
	"log"
	"reflect"
	"strings"
)

// LeaklessOptions are the options that leakless should use
// for running child processes
type LeaklessOptions struct {
	// Windows-specific process options
	Windows *WindowsOptions
}

// WindowsOptions is a collection of Windows-specific process options
type WindowsOptions struct {
	// Whether to hide GUI window or not (default: `false`)
	//
	// Set this to `true` to avoid flashing console windows
	HideWindow bool
}

// BuildOptionsString builds a string of flags from *LeaklessOptions struct
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

// ParseOptionsString recovers *LeaklessOptions struct from a string of flags
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
