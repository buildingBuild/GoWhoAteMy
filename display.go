package main

import "fmt"

func displayDeviceInfo(dInfo DeviceInfo) {
	fmt.Printf(
		Red+"Device Info:\n"+Reset+
			"OS: %s | Platform: %s\n"+
			"Platform Version: %s | CPU: %s\n"+
			"CPU Usage: %.2f%% | CPU Cores: %d\n"+
			"Memory Usage: %.2f%% | Disk Usage: %.2f%%\n",
		dInfo.OS,
		dInfo.Platform,
		dInfo.PlatformVersion,
		dInfo.CPU,
		dInfo.CPUUsage,
		dInfo.CPUCores,
		dInfo.MemoryUsedPercent,
		dInfo.DiskUsagePercent,
	)
}
