package volume

import (
	"errors"
	"runtime"
	"time"
)

// Found URL: https://github.com/itchyny/volume-go on 2023-05-20, so renamed the package from "speaker" to "volume", but decided to
// keep this implementation instead of using this package

// NOTE: Using an int would be a performant option, but as the status should be "readable", a string has been used instead
type Status string

const (
	StatusOn      Status = "on"
	StatusOff     Status = "off"
	StatusUnknown Status = "unknown"
)

type MonitorFunc func(status Status, err error)

func Monitor(monitorFn MonitorFunc) error {
	prevStatus, err := getStatus()
	if err != nil {
		return err
	}

	ticker := time.NewTicker(512 * time.Millisecond)
	defer ticker.Stop()

	for range ticker.C {
		currStatus, err := getStatus()
		if err != nil {
			monitorFn(StatusUnknown, err)
			continue
		}

		if currStatus != prevStatus {
			monitorFn(currStatus, nil)
			prevStatus = currStatus
		}
	}
	return nil
}

func getStatus() (Status, error) {
	// Build constraints should be used, but this is simpler to maintain.
	// See URL: https://stackoverflow.com/a/19847868
	switch runtime.GOOS {
	case "darwin":
		return getStatusDarwin()
	case "linux":
		return getStatusLinux()
	default:
		return StatusUnknown, errors.New("unsupported OS for volume status")
	}
}
