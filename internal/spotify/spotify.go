package spotify

import (
	"fmt"
	"strings"

	"github.com/godbus/dbus/v5"
)

// NOTE: Found the package URL: https://github.com/dawidd6/go-spotify-dbus, after this was initially created

// NOTE: Using an int would be a performant option, but as the status should be "readible", a string has been used instead
type Status string

const (
	StatusPlaying Status = "playing"
	StatusPaused  Status = "paused"
	StatusUnknown Status = "unknown"
)

type App struct {
	conn *dbus.Conn
	bus  dbus.BusObject
	dest string
}

func New() (*App, error) {
	conn, err := dbus.SessionBus()
	if err != nil {
		return nil, fmt.Errorf("connecting to session bus: %w", err)
	}

	dest := "org.mpris.MediaPlayer2.spotify"
	path := dbus.ObjectPath("/org/mpris/MediaPlayer2")

	return &App{
		conn: conn,
		bus:  conn.Object(dest, path),
		dest: dest,
	}, nil
}

func (a *App) Play() error {
	return a.runMethod("Play")
}

func (a *App) Pause() error {
	return a.runMethod("Pause")
}

func (a *App) Status() (Status, error) {
	v, err := a.bus.GetProperty("org.mpris.MediaPlayer2.Player.PlaybackStatus")
	if err != nil {
		return StatusUnknown, fmt.Errorf("getting playback status: %w", err)
	}

	switch strings.Trim(v.String(), `"`) {
	case "Playing":
		return StatusPlaying, nil
	case "Paused":
		return StatusPaused, nil
	default:
		return StatusUnknown, nil
	}
}

func (a *App) IsRunning() (bool, error) {
	obj := a.conn.Object("org.freedesktop.DBus", "/org/freedesktop/DBus")
	call := obj.Call("org.freedesktop.DBus.NameHasOwner", 0, a.dest)
	var isRunning bool
	return isRunning, call.Store(&isRunning)
}

func (a *App) runMethod(name string) error {
	call := a.bus.Call(fmt.Sprintf("org.mpris.MediaPlayer2.Player.%s", name), 0)
	if call.Err != nil {
		return fmt.Errorf("calling method %q: %w", name, call.Err)
	}
	return nil
}
