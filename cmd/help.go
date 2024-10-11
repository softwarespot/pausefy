package cmd

import "fmt"

func cmdHelp() {
	helpText := `Usage: ./pausefy-linux [OPTIONS]

Pause Spotify when the "mute" key on the keyboard is pressed for Debian based OS'es

Options:
  -h, --help      Show this help text and exit.

Examples:
  ./pausefy-linux`
	fmt.Println(helpText)
}
