package pistats

import (
	"errors"
	"os/exec"
	"regexp"
)

type Temperature string

func sampleTemp() (Temperature, error) {
	t, err := parseTemp()
	if err != nil {
		return "", err
	}
	return Temperature(t), nil
}

func parseTemp() (string, error) {
	d, err := readTempData()
	if err != nil {
		return "", err
	}
	r := regexp.MustCompile("[^0-9.]+")

	return r.ReplaceAllString(d, ""), nil
}

func readTempData() (string, error) {
	out, err := exec.Command("vcgencmd", "measure_temp").Output()
	if err != nil {
		return "", errors.New("couldn't read from vcgencmd")
	}

	return string(out), nil
}
