package pistats

import (
	"log"

	linuxproc "github.com/c9s/goprocinfo/linux"
)

// Memory contains total and available memory values in KBs
type Memory struct {
	Total     uint64
	Available uint64
}

func sampleMemory() Memory {
	d, err := linuxproc.ReadMemInfo("/proc/meminfo")
	if err != nil {
		log.Fatal("couldn't read from /proc/meminfo")
	}

	mem := Memory{d.MemTotal, d.MemAvailable}

	return mem
}
