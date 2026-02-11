package spotify

import (
	"errors"
	"runtime"
)

var errUnsupportedOS = errors.New("unsupported OS for Spotify app")

// NOTE: Found the package URL: https://github.com/dawidd6/go-spotify-dbus, after this was initially created

// NOTE: Using an int would be a performant option, but as the status should be "readible", a string has been used instead
type Status string

const (
	StatusPlaying Status = "playing"
	StatusPaused  Status = "paused"
	StatusUnknown Status = "unknown"
)

type App struct {
	darwin *appDarwin
	linux  *appLinux
}

func New() (*App, error) {
	app := &App{
		darwin: nil,
		linux:  nil,
	}

	// Build constraints should be used, but this is simpler to maintain.
	// See URL: https://stackoverflow.com/a/19847868
	switch runtime.GOOS {
	case "darwin":
		var err error
		if app.darwin, err = newAppDarwin(); err != nil {
			return nil, err
		}
		return app, nil
	case "linux":
		var err error
		if app.linux, err = newAppLinux(); err != nil {
			return nil, err
		}
		return app, nil
	default:
		return nil, errUnsupportedOS
	}
}

func (a *App) Play() error {
	switch runtime.GOOS {
	case "darwin":
		return a.darwin.play()
	case "linux":
		return a.linux.play()
	default:
		return errUnsupportedOS
	}
}

func (a *App) Pause() error {
	switch runtime.GOOS {
	case "darwin":
		return a.darwin.pause()
	case "linux":
		return a.linux.pause()
	default:
		return errUnsupportedOS
	}
}

func (a *App) Status() (Status, error) {
	switch runtime.GOOS {
	case "darwin":
		return a.darwin.status()
	case "linux":
		return a.linux.status()
	default:
		return StatusUnknown, errUnsupportedOS
	}
}

func (a *App) IsRunning() (bool, error) {
	switch runtime.GOOS {
	case "darwin":
		return a.darwin.isRunning()
	case "linux":
		return a.linux.isRunning()
	default:
		return false, errUnsupportedOS
	}
}
