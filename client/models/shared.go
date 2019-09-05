package SharedModel

import (
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/docker"
	"github.com/shirou/gopsutil/mem"
)

type MasterData struct {
	InfoType      string                    `json:"infoType"`
	VirtualMemory *mem.VirtualMemoryStat    `json:"virtualMemory"`
	DiskUsage     *disk.UsageStat           `json:"diskUsage"`
	Cpu           []cpu.InfoStat            `json:"cpu"`
	CpuPercentage []cpu.InfoStat            `json:"cpuPercentage"`
	DockerState   []docker.CgroupDockerStat `json:"dockerState"`
}

type DetailedManager struct {
	VirtualMemory chan *mem.VirtualMemoryStat
	DiskUsage     chan *disk.UsageStat
	Cpu           chan []cpu.InfoStat
	CpuPercentage chan []cpu.InfoStat
}
