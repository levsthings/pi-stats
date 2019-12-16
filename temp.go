package main

import (
	"fmt"
	"log"
	"os/exec"
	"strconv"
	"strings"
)

const tempFormat = `temp: %.0f°C`

func sampleTemp() string {

	d := parseTemp()

	f, err := strconv.ParseFloat(d, 64)
	if err != nil {
		log.Fatal("couldn't parse from lm-sensors")
	}

	return fmt.Sprintf(tempFormat, f)
}

func parseTemp() string {
	d := readTempData()
	var temp string

	lines := strings.Split(d, "\n")
	for _, line := range lines {
		if strings.Contains(line, "temp1_input") {
			parts := strings.Split(line, ":")
			temp = strings.TrimSpace(parts[1])
		}
	}

	return temp
}

func readTempData() string {
	out, err := exec.Command("sensors", "-u").Output()
	if err != nil {
		log.Fatal("couldn't read from lm-sensors")
	}

	s := string(out)

	return s
}
