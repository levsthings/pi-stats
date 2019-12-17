package main

import (
	"fmt"
	"log"

	linuxproc "github.com/c9s/goprocinfo/linux"
)

const memFormat = `memtotal: %dKB memused: %dKB memavailable: %dKB`

func sampleMemory() string {
	t, a := calcMem(readMemData())

	return fmt.Sprintf(memFormat, t, t-a, a)
}

func calcMem(d *linuxproc.MemInfo) (t, a uint64) {
	totalMem := d.MemTotal
	availMem := d.MemAvailable

	return totalMem, availMem
}

func readMemData() *linuxproc.MemInfo {
	d, err := linuxproc.ReadMemInfo("/proc/meminfo")
	if err != nil {
		log.Fatal("couldn't read from /proc/meminfo")

	}

	return d
}
