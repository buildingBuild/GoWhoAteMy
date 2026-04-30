package main

import (
	"fmt"
	"os"

	"github.com/alecthomas/kong"
)

type DeviceInfo struct {
	OS                string
	Platform          string
	PlatformVersion   string
	CPU               string
	CPUUsage          float64
	CPUCores          int32
	DiskUsagePercent  float64
	MemoryUsedPercent float64
}

type CommandLineOptions struct {
	CPU             bool `help:"Run CPU monitoring mode."`
	Memory          bool `help:"Run memory monitoring mode."`
	Network         bool `help:"Run network monitoring mode."`
	Computer        bool `help:"Run full computer monitoring mode with notifications."`
	IntervalSeconds int  `name:"interval" default:"10" help:"Seconds between computer checks."`
}

var deviceInfo DeviceInfo

func main() {
	cli := CommandLineOptions{}
	kong.Parse(&cli,
		kong.Name("gowhoatemy"),
		kong.Description("Find what is slowing your computer down."),
	)

	err := getBasicDeviceInfo(&deviceInfo)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Could not get device info: ", err)
		os.Exit(1)
	}

	if runSelectedMode(cli) {
		return
	}

	runInteractiveMenu()
}
