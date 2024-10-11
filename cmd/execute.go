package cmd

import "flag"

func Execute() {
	var showHelp bool
	flagBoolVarP(&showHelp, "help", "h", false, "Display the help text and exit")

	flag.Parse()

	if showHelp {
		cmdHelp()
	}
	cmdStart()
}
