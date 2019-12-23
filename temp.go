package pistats

import (
	"regexp"
)

func sampleTemp() string {
	t := parseTemp()

	return t
}

func parseTemp() string {
	d := readTempData()
	r := regexp.MustCompile("[^0-9.]+")

	return r.ReplaceAllString(d, "")
}

func readTempData() string {
	// out, err := exec.Command("vcgencmd", "measure_temp").Output()
	// if err != nil {
	// 	log.Fatal("couldn't read from vcgencmd")
	// }

	// return string(out)

	return "36.2"
}
