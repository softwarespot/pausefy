package volume

import (
	"fmt"

	"github.com/softwarespot/pausefy/internal/helpers"
)

func getStatusDarwin() (Status, error) {
	out, err := helpers.ExecCmd([]string{"osascript", "-e", "output muted of (get volume settings)"})
	if err != nil {
		return StatusUnknown, fmt.Errorf(`get volume: %w`, err)
	}

	switch out {
	case "false":
		return StatusOn, nil
	case "true":
		return StatusOff, nil
	default:
		return StatusUnknown, fmt.Errorf(`unsupported "osascript" output: %q`, out)
	}
}
