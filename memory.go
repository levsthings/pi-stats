package main

import (
	"fmt"
	"log"

	linuxproc "github.com/c9s/goprocinfo/linux"
)

const memFormat = `Total Memory: %d Used Memory: %d Available Memory: %d`

func sampleMemory() string {
	// sample CPU activity
	t, a := calcMem(readMemData())

	return fmt.Sprintf(memFormat, t, t-a, a)
}

func calcMem(d *linuxproc.MemInfo) (t, a uint64) {
	totalMem := d.MemTotal
	availMem := d.MemAvailable

	return totalMem, availMem

}

func readMemData() *linuxproc.MemInfo {
	// read data from /proc/mem
	d, err := linuxproc.ReadMemInfo("/proc/meminfo")
	if err != nil {
		log.Fatal("couldn't read from /proc/meminfo")

	}
	return d
}
