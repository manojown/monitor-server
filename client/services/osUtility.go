package OsUtility

import (
	"fmt"

	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
)

func GetMem() *mem.VirtualMemoryStat {
	var virtual *mem.VirtualMemoryStat
	// var Status SharedModel.UsageStat
	virtual, _ = mem.VirtualMemory()
	return virtual
}

func GetDisk() {
	// var virtual []disk.PartitionStat
	// var Status SharedModel.UsageStat
	virtual, _ := disk.Usage(, "/")
	fmt.Printf("Asdasd data is", virtual)
	return virtual
}
