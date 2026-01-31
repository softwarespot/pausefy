package spotify

import (
	"fmt"

	"github.com/softwarespot/pausefy/internal/helpers"
)

// /Applications/Spotify.app/Contents/Resources/Spotify.sdef for AppleScript reference

type appDarwin struct{}

func newAppDarwin() (*appDarwin, error) {
	return &appDarwin{}, nil
}

func (a *appDarwin) play() error {
	// Doesn't return any output, even if Spotify is not running
	if _, err := execScript(`tell application "Spotify" to play`); err != nil {
		return fmt.Errorf("send play command: %w", err)
	}
	return nil
}

func (a *appDarwin) pause() error {
	// Doesn't return any output, even if Spotify is not running
	if _, err := execScript(`tell application "Spotify" to pause`); err != nil {
		return fmt.Errorf("send pause command: %w", err)
	}
	return nil
}

func (a *appDarwin) status() (Status, error) {
	out, err := execScript(`tell application "Spotify" to return player state`)
	if err != nil {
		return StatusUnknown, fmt.Errorf("get status: %w", err)
	}

	switch out {
	case "playing":
		return StatusPlaying, nil
	case "paused", "stopped":
		return StatusPaused, nil
	default:
		return StatusUnknown, nil
	}
}

func (a *appDarwin) isRunning() (bool, error) {
	out, err := execScript(`tell application "System Events" to return (exists process "Spotify")`)
	if err != nil {
		return false, err
	}

	switch out {
	case "true":
		return true, nil
	case "false":
		return false, nil
	default:
		return false, fmt.Errorf(`unsupported "osascript" output: %q`, out)
	}
}

func execScript(script string) (string, error) {
	out, err := helpers.ExecCmd([]string{"osascript", "-e", script})
	if err != nil {
		return "", fmt.Errorf(`execute "osascript" %q: %w`, script, err)
	}
	return out, nil
}
