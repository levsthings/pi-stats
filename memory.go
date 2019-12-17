package main

import (
	"log"

	linuxproc "github.com/c9s/goprocinfo/linux"
)

func sampleMemory() (uint64, uint64) {
	t, a := calcMem(readMemData())

	return t, a
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
