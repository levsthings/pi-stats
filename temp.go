package main

import (
	"fmt"
	"log"
	"os/exec"
	"regexp"
	"strconv"
)

const tempFormat = `temp: %.0f°C`

func sampleTemp() string {
	d := parseTemp()

	f, err := strconv.ParseFloat(d, 64)
	if err != nil {
		log.Print("couldn't parse from vgencmd")
	}

	return fmt.Sprintf(tempFormat, f)
}

func parseTemp() string {
	d := readTempData()
	r := regexp.MustCompile("[^0-9.]+")

	return string(r.ReplaceAllString(d, ""))
}

func readTempData() string {
	out, err := exec.Command("vgencmd", "measure_temp").Output()
	if err != nil {
		log.Fatal("couldn't read from vgencmd")
	}

	s := string(out)

	return s
}
