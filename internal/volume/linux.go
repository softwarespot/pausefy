package volume

import (
	"fmt"
	"strings"

	"github.com/softwarespot/pausefy/internal/helpers"
)

func getStatusLinux() (Status, error) {
	out, err := helpers.ExecCmd([]string{"amixer", "get", "Master"})
	if err != nil {
		return StatusUnknown, fmt.Errorf(`get volume: %w`, err)
	}

	for l := range strings.SplitSeq(out, "\n") {
		if !strings.Contains(l, "Playback") || !strings.Contains(l, "%") {
			continue
		}
		switch {
		case strings.Contains(l, "[off]"):
			return StatusOff, nil
		case strings.Contains(l, "[on]"):
			return StatusOn, nil
		}
	}
	return StatusUnknown, fmt.Errorf(`unsupported "amixer" output: %q`, out)
}
