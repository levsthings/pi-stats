package pistats

import (
	"errors"

	linuxproc "github.com/c9s/goprocinfo/linux"
)

type Uptime string

func getUptime() (Uptime, error) {
	d, err := linuxproc.ReadUptime("/proc/uptime")
	if err != nil {
		return "", errors.New("couldn't read from /proc/uptime")
	}
	s := d.GetTotalDuration().String()
	return Uptime(s), nil
}
