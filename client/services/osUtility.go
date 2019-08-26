package OsUtility

import (
	"serverInfo/client/models"

	"github.com/shirou/gopsutil/disk"
)

func GetStatus() SharedModel.UsageStat {
	// var Status SharedModel.UsageStat

	v, _ := disk.Usage("/")
	return v
}
