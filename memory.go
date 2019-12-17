package main

import (
	"log"

	linuxproc "github.com/c9s/goprocinfo/linux"
)

func sampleMemory() (t, a uint64) {
	d, err := linuxproc.ReadMemInfo("/proc/meminfo")
	if err != nil {
		log.Fatal("couldn't read from /proc/meminfo")
	}

	t, a = d.MemTotal, d.MemAvailable

	return t, a
}
