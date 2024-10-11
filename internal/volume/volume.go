package volume

import (
	"errors"
	"strings"
	"time"

	"pausefy/internal/helpers"
)

// Found URL: https://github.com/itchyny/volume-go on 2023-05-20, so renamed the package from "speaker" to "volume", but decided to
// keep this implementation instead of using this package

// NOTE: Using an int would be a performant option, but as the status should be "readible", a string has been used instead
type Status string

const (
	StatusOn      Status = "on"
	StatusOff     Status = "off"
	StatusUnknown Status = "unknown"
)

type MonitorFunc func(speakerStatus Status)

func Monitor(monitorFn MonitorFunc) error {
	prevStatus, err := getStatus()
	if err != nil {
		return err
	}

	for {
		time.Sleep(512 * time.Millisecond)

		currStatus, err := getStatus()
		if err == nil && prevStatus != currStatus {
			monitorFn(currStatus)
			prevStatus = currStatus
		}
	}
}

func getStatus() (Status, error) {
	out, err := helpers.ExecCmd([]string{"amixer", "get", "Master"})
	if err != nil {
		return StatusUnknown, err
	}

	for _, l := range strings.Split(out, "\n") {
		if !strings.Contains(l, "Playback") || !strings.Contains(l, "%") {
			continue
		}
		if strings.Contains(l, "[off]") || strings.Contains(l, "yes") {
			return StatusOff, nil
		}
		if strings.Contains(l, "[on]") || strings.Contains(l, "no") {
			return StatusOn, nil
		}
	}
	return StatusUnknown, errors.New(`unsupported "axmixer" output`)
}
