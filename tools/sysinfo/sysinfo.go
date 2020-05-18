package sysinfo

import (
	"github.com/elastic/go-sysinfo"
)

// GetHostInfo is function to read Host System Information
func GetHostInfo() (*HostInfo, error) {
	info, err := sysinfo.Host()
	if err != nil {
		return nil, err
	}

	hi := HostInfo{}
	hi.Hostname = info.Info().Hostname
	hi.Architecture = info.Info().Architecture
	hi.BootTime = info.Info().BootTime
	hi.Containerized = info.Info().Containerized
	hi.IPs = info.Info().IPs
	hi.KernelVersion = info.Info().KernelVersion
	hi.MACs = info.Info().MACs

	OS := OSInfo{}
	OS.Family = info.Info().OS.Family
	OS.Platform = info.Info().OS.Platform
	OS.Name = info.Info().OS.Name
	OS.Version = info.Info().OS.Version
	OS.Major = info.Info().OS.Major
	OS.Minor = info.Info().OS.Minor
	OS.Patch = info.Info().OS.Patch
	OS.Build = info.Info().OS.Build
	OS.Codename = info.Info().OS.Codename

	hi.OS = &OS

	hi.Timezone = info.Info().Timezone
	hi.TimezoneOffsetSec = info.Info().TimezoneOffsetSec
	hi.UniqueID = info.Info().UniqueID

	return &hi, err
}
