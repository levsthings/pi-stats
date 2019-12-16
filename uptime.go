package main

import (
	"fmt"
	"log"

	linuxproc "github.com/c9s/goprocinfo/linux"
)

func getUptime() string {
	d, err := linuxproc.ReadUptime("/proc/uptime")
	if err != nil {
		log.Fatal("couldn't read from /proc/uptime")
	}

	t := d.GetTotalDuration()
	return fmt.Sprintf("Uptime:", t)
}
