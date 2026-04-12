package main

import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
)

func getBasicDeviceInfo(dInfo *DeviceInfo) error {
	info, err := host.Info()
	if err != nil {
		return fmt.Errorf("host.Info: %w", err)
	}

	pcts, err := cpu.Percent(time.Second, false)
	if err != nil {
		return fmt.Errorf("cpu.Percent: %w", err)
	}

	v, err := mem.VirtualMemory()
	if err != nil {
		return fmt.Errorf("mem.VirtualMemory: %w", err)
	}

	d, err := disk.Usage("/")
	if err != nil {
		return fmt.Errorf("disk.Usage: %w", err)
	}

	infos, _ := cpu.Info()

	dInfo.OS = info.OS
	dInfo.Platform = info.Platform
	dInfo.PlatformVersion = info.PlatformVersion
	dInfo.CPUUsage = pcts[0]
	dInfo.DiskUsagePercent = d.UsedPercent
	dInfo.MemoryUsedPercent = v.UsedPercent

	if len(infos) > 0 {
		dInfo.CPU = infos[0].ModelName
		dInfo.CPUCores = infos[0].Cores
	}

	return nil

}
