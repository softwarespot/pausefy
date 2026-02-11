package spotify

import (
	"fmt"
	"strings"

	"github.com/godbus/dbus/v5"
)

type appLinux struct {
	conn *dbus.Conn
	bus  dbus.BusObject
}

func newAppLinux() (*appLinux, error) {
	conn, err := dbus.SessionBus()
	if err != nil {
		return nil, fmt.Errorf("connecting to session bus: %w", err)
	}

	return &appLinux{
		conn: conn,
		bus: conn.Object(
			"org.mpris.MediaPlayer2.spotify",
			dbus.ObjectPath("/org/mpris/MediaPlayer2"),
		),
	}, nil
}

func (a *appLinux) play() error {
	if err := a.execMethod("Play"); err != nil {
		return fmt.Errorf("send play command: %w", err)
	}
	return nil
}

func (a *appLinux) pause() error {
	if err := a.execMethod("Pause"); err != nil {
		return fmt.Errorf("send pause command: %w", err)
	}
	return nil
}

func (a *appLinux) execMethod(name string) error {
	if res := a.bus.Call(fmt.Sprintf("org.mpris.MediaPlayer2.Player.%s", name), 0); res.Err != nil {
		return fmt.Errorf("execute method %q: %w", name, res.Err)
	}
	return nil
}

func (a *appLinux) status() (Status, error) {
	status, err := a.bus.GetProperty("org.mpris.MediaPlayer2.Player.PlaybackStatus")
	if err != nil {
		return StatusUnknown, fmt.Errorf("get status: %w", err)
	}

	switch strings.Trim(status.String(), `"`) {
	case "Playing":
		return StatusPlaying, nil
	case "Paused":
		return StatusPaused, nil
	default:
		return StatusUnknown, nil
	}
}

func (a *appLinux) isRunning() (bool, error) {
	obj := a.conn.Object("org.freedesktop.DBus", "/org/freedesktop/DBus")
	res := obj.Call("org.freedesktop.DBus.NameHasOwner", 0, a.bus.Destination())

	var isRunning bool
	if err := res.Store(&isRunning); err != nil {
		return false, fmt.Errorf("get running state: %w", err)
	}
	return isRunning, nil
}
