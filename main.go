package main

import (
	"fmt"
	"os"

	"go-who-ate-my-cpu/cli"
	"go-who-ate-my-cpu/display"

	"github.com/alecthomas/kong"
)

func main() {
	options := cli.CommandLineOptions{}
	kong.Parse(&options,
		kong.Name("gowhoatemy"),
		kong.Description("Find what is slowing your computer down."),
	)

	deviceInfo := display.DeviceInfo{}
	err := display.GetBasicDeviceInfo(&deviceInfo)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Could not get device info: ", err)
		os.Exit(1)
	}

	if cli.RunSelectedMode(options) {
		return
	}

	cli.RunInteractiveMenu(deviceInfo)
}
