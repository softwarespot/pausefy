package cmd

import "fmt"

func cmdHelp() {
	helpText := `Usage: ./pausefy [OPTIONS]

Pause Spotify when the "mute" key on the keyboard is pressed for Darwin/Debian based OS'es

Options:
  -h, --help      Show this help text and exit.
  -v, --version   Display the version of the application and exit.
  -j, --json      Output the version as JSON.

Examples:
  ./pausefy`
	fmt.Println(helpText)
}
