package OsUtility

import (
	"encoding/json"
	"fmt"
	"log"
	"monitor-server/client/models"
	"time"

	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/docker"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
)

// func GetMem(manager *SharedModel.DetailedManager) {
// 	var virtual *mem.VirtualMemoryStat
// 	// var Status SharedModel.UsageStat
// 	virtual, _ = mem.VirtualMemory()
// 	manager.VirtualMemory <- virtual
//
// }
//
// func GetDisk(detailManager SharedModel.DetailedManager) {
// 	var Usage *disk.UsageStat
// 	// var Status SharedModel.UsageStat
// 	Usage, _ = disk.Usage("/")
// 	detailManager.DiskUsage <- Usage
//
// }
func GetCpu(connection chan []byte) {
	var masterSend SharedModel.MasterData
	var virtual []cpu.InfoStat
	// var Status SharedModel.UsageStat
	ticker := time.NewTicker(5 * time.Second)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				virtual, _ = cpu.Info()
				masterSend.InfoType = "CPU_DATA"
				masterSend.Cpu = virtual

				stat, err := json.Marshal(masterSend)
				failOnError(err, "failed while getting CPU details")
				connection <- stat
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()

}

func GetMem(connection chan []byte) {
	var virtual *mem.VirtualMemoryStat
	var masterSend SharedModel.MasterData

	// var Status SharedModel.UsageStat
	ticker := time.NewTicker(5 * time.Second)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				virtual, _ = mem.VirtualMemory()
				masterSend.InfoType = "MEM_DATA"
				masterSend.VirtualMemory = virtual
				stat, err := json.Marshal(masterSend)

				failOnError(err, "failed while getting Memory details")
				connection <- stat
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()

}

func GetDisk(connection chan []byte) {

	var masterSend SharedModel.MasterData

	// var Status SharedModel.UsageStat
	ticker := time.NewTicker(4 * time.Second)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				masterSend.DiskUsage, _ = disk.Usage("/")
				masterSend.InfoType = "DISK_DATA"
				stat, err := json.Marshal(masterSend)
				failOnError(err, "failed while getting disk space")
				connection <- stat
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()

}

func GetDocker(connection chan []byte) {
	var masterSend SharedModel.MasterData

	// var Status SharedModel.UsageStat
	ticker := time.NewTicker(4 * time.Second)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				masterSend.DockerState, _ = docker.GetDockerStat()
				masterSend.InfoType = "DOCKER_DATA"
				stat, err := json.Marshal(masterSend)
				failOnError(err, "failed while getting disk space")
				connection <- stat
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()

}
func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}
