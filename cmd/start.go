package cmd

import (
	"fmt"
	"log"
	"time"

	"github.com/softwarespot/pausefy/internal/helpers"
	"github.com/softwarespot/pausefy/internal/spotify"
	"github.com/softwarespot/pausefy/internal/volume"
)

func cmdStart() {
	// DISCLAIMER: Given the nature of the application, errors which occur can just be logged to STDOUT, as the application
	// shouldn't exit and can recover eventually, for example when Spotify has started/re-started

	helpers.Retry(func(iter int) error {
		s, err := spotify.New()
		if err != nil {
			return fmt.Errorf("application Spotify not running. Try %d: %w", iter, err)
		}

		monitorFn, err := getMonitorFunc(s)
		if err != nil {
			return fmt.Errorf("creating monitor function. Try %d: %w", iter, err)
		}

		log.Printf("start monitoring %s\n", helpers.ExecutableName())
		if err := volume.Monitor(monitorFn); err != nil {
			return fmt.Errorf("start speaker monitoring. Try %d: %w", iter, err)
		}
		return nil
	}, 5*time.Second)
}

func getMonitorFunc(s *spotify.App) (volume.MonitorFunc, error) {
	currSpotifyStatus, err := s.Status()
	if err != nil {
		return nil, fmt.Errorf("get Spotify status: %w", err)
	}

	return func(speakerStatus volume.Status, speakerStatusErr error) {
		// Handle errors by logging to STDOUT,
		// as there is no reason to exit the application, as this is likely a temporary issue

		if speakerStatusErr != nil {
			log.Println("got an error of whilst checking speaker status")
			log.Println(speakerStatusErr)
			return
		}

		isRunning, err := s.IsRunning()
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
			// Restore playback only if Spotify was playing before speakers went off
			switch currSpotifyStatus {
			case spotify.StatusPlaying:
				log.Printf("speaker status changed to %q, setting Spotify to play\n", speakerStatus)
				if err := s.Play(); err != nil {
					log.Println(err)
				}
				currSpotifyStatus = spotify.StatusPlaying
			default:
				log.Printf("speaker status changed to %q, not setting Spotify to play\n", speakerStatus)
			}
		case volume.StatusOff:
			// Save the current Spotify status before pausing so we know whether to resume later
			var err error
			if currSpotifyStatus, err = s.Status(); err != nil {
				log.Println(err)
			}

			log.Printf("speaker status changed to %q, setting Spotify to pause\n", speakerStatus)
			if err := s.Pause(); err != nil {
				log.Println(err)
			}
		default:
			log.Printf("unable to determine the current speaker status. Current Spotify status is %q", currSpotifyStatus)
		}
	}, nil
}
