package pistats

import (
	"errors"

	linuxproc "github.com/c9s/goprocinfo/linux"
)

// Memory contains total and available memory values in KBs
type Memory struct {
	Total     uint64
	Available uint64
}

func sampleMemory() (*Memory, error) {
	d, err := linuxproc.ReadMemInfo("/proc/meminfo")
	if err != nil {
		return nil, errors.New("couldn't read from /proc/meminfo")
	}

	mem := Memory{d.MemTotal, d.MemAvailable}

	return &mem, nil
}
