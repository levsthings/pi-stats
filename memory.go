package pistats

import (
	"log"

	linuxproc "github.com/c9s/goprocinfo/linux"
)

type memory struct {
	total     uint64
	available uint64
}

func sampleMemory() memory {
	d, err := linuxproc.ReadMemInfo("/proc/meminfo")
	if err != nil {
		log.Fatal("couldn't read from /proc/meminfo")
	}

	mem := memory{d.MemTotal, d.MemAvailable}

	return mem
}
