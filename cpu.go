package pistats

import (
	"log"
	"time"

	linuxproc "github.com/c9s/goprocinfo/linux"
)

// CPU contains float32 values for each CPU core on a Raspberry PI
type CPU [4]float32

func sampleCPU() CPU {
	p := readCPUdata()
	time.Sleep(time.Second * 1)
	c := readCPUdata()

	return calcAllCores(c, p)
}

func readCPUdata() *linuxproc.Stat {
	d, err := linuxproc.ReadStat("/proc/stat")
	if err != nil {
		log.Fatal("couldn't read from /proc/stat")
	}

	return d
}

func calcCore(c, p linuxproc.CPUStat) float32 {
	// htop formula for calculating CPU activity
	pIdle := p.Idle + p.IOWait
	cIdle := c.Idle + c.IOWait

	pNonIdle := p.User + p.Nice + p.System + p.IRQ + p.SoftIRQ + p.Steal
	cNonIdle := c.User + c.Nice + c.System + c.IRQ + c.SoftIRQ + c.Steal

	pTotal := pIdle + pNonIdle
	cTotal := cIdle + cNonIdle

	difTotal := cTotal - pTotal
	difIdle := cIdle - pIdle

	return (float32(difTotal) - float32(difIdle)) / float32(difTotal)
}

func calcAllCores(curr, prev *linuxproc.Stat) CPU {
	stats := CPU{}
	for i := range stats {
		stats[i] = calcCore(curr.CPUStats[i], prev.CPUStats[i]) * 100
	}

	return stats
}
