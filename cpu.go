package main

import (
	"fmt"
	"log"
	"time"

	linuxproc "github.com/c9s/goprocinfo/linux"
)

type cpu [4]float32

const format = `CPU 1: %.2f, CPU2: %.2f, CPU3: %.2f, CPU4: %.2f \n`

func sampleCPUdata() string {
	// take sample CPU activity
	p := readCPUdata()

	time.Sleep(time.Second * 1)
	c := readCPUdata()

	stats := calcAllCores(c, p)

	return fmt.Sprintf(format, stats[0], stats[1], stats[2], stats[3])
}

func readCPUdata() *linuxproc.Stat {
	// read data from /proc/stat
	d, err := linuxproc.ReadStat("/proc/stat")

	if err != nil {
		log.Fatal("couldn't read from /proc/stat")
	}

	return d
}

func calcAllCores(curr, prev *linuxproc.Stat) cpu {
	// calculate CPU activity for all cores
	stats := cpu{}
	for i := range stats {
		stats[i] = calcCore(curr.CPUStats[i], prev.CPUStats[i]) * 100
	}
	return stats
}

func calcCore(c, p linuxproc.CPUStat) float32 {
	// htop formula for calculating CPU activity
	pIdle := p.Idle + p.IOWait
	cIdle := c.Idle + c.IOWait

	pNidle := p.User + p.Nice + p.System + p.IRQ + p.SoftIRQ + p.Steal
	cNidle := c.User + c.Nice + c.System + c.IRQ + c.SoftIRQ + c.Steal

	pTotal := pIdle + pNidle
	cTotal := cIdle + cNidle

	difTotal := cTotal - pTotal
	difIdle := cIdle - pIdle

	return (float32(difTotal) - float32(difIdle)) / float32(difTotal)
}
