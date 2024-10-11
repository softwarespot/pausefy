package cmd

import (
	"log"
	"time"

	"pausefy/internal/helpers"
	"pausefy/internal/spotify"
	"pausefy/internal/volume"

	"github.com/pkg/errors"
)

func cmdStart() {
	// DISCLAIMER: Given the nature of the application, errors which occur can just be logged to STDOUT, as the application
	// shouldn't exit and can recover eventually, for example when Spotify has started/re-started

	helpers.Retry(func(iter int) error {
		sfy, err := spotify.New()
		if err != nil {
			return errors.Wrapf(err, "Spotify not running. Try %d", iter)
		}

		monitorFn, err := getMonitorFunc(sfy)
		if err != nil {
			return errors.Wrapf(err, "creating monitor function. Try %d", iter)
		}

		log.Printf("start monitoring %s\n", helpers.ExecutableName())
		if err := volume.Monitor(monitorFn); err != nil {
			return errors.Wrapf(err, "start speaker monitoring. Try %d", iter)
		}
		return nil
	}, 5*time.Second)
}

func getMonitorFunc(sfy *spotify.App) (volume.MonitorFunc, error) {
	currSpotifyStatus, err := sfy.Status()
	if err != nil {
		return nil, errors.Wrap(err, "get Spotify status")
	}

	return func(speakerStatus volume.Status) {
		// Handle errors by logging to STDOUT,
		// as there is no reason to exit the application, as this is likely temporary

		isRunning, err := sfy.IsRunning()
		if err != nil {
			log.Printf("speaker status changed to %q, got an error whilst checking if Spotify is running\n", speakerStatus)
			log.Println(err)
			return
		}

		if !isRunning {
			log.Printf("speaker status changed to %q, Spotify is not running\n", speakerStatus)
			return
		}

		switch speakerStatus {
		case volume.StatusOn:
			if currSpotifyStatus == spotify.StatusPlaying {
				log.Printf("speaker status changed to %q, setting Spotify to play\n", speakerStatus)
				if err := sfy.Play(); err != nil {
					log.Println(err)
				}
			} else {
				log.Printf("speaker status changed to %q, not setting Spotify to play\n", speakerStatus)
			}
		case volume.StatusOff:
			log.Printf("speaker status changed to %q, setting Spotify to pause\n", speakerStatus)
			if err := sfy.Pause(); err != nil {
				log.Println(err)
			}
		case volume.StatusUnknown:
			log.Println("unable to determine the current speaker status")
		}

		if currSpotifyStatus, err = sfy.Status(); err != nil {
			log.Println(err)
		}
	}, nil
}
