package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"time"

	pistats "github.com/levsthings/pi-stats"
)

const (
	dir     = ".pi-stats"
	logPath = dir + "/"
	maxLogs = 7

	consoleMode = "console"
	logMode     = "log"
)

func main() {
	mode := flag.String("mode", consoleMode, "Expected input: '--mode console' or '--mode log'")
	flag.Parse()

	if *mode == logMode {
		for {
			d, err := pistats.GetData()
			if err != nil {
				logError(errorOutput{
					err,
					"couldn't get temp data from pi-stats",
				})
			}
			write(format(d))
			rotate()

			time.Sleep(time.Minute * 5)
		}
	}

	d, err := pistats.GetData()
	if err != nil {
		logError(errorOutput{
			err,
			"couldn't get temp data from pi-stats",
		})
	}
	fmt.Print(format(d))
}

func format(d *pistats.Data) string {
	var (
		uptimeFormat = `uptime: %s`
		cpuFormat    = `CPU 1: %.2f%%, CPU2: %.2f%%, CPU3: %.2f%%, CPU4: %.2f%%`
		tempFormat   = `temp: %s°C`
		memFormat    = `memtotal: %dMB, memused: %dMB, memavailable: %dMB`
		logFormat    = "%s, %s, %s, %s, %s\n"
	)

	uptime := fmt.Sprintf(uptimeFormat, d.Uptime)

	cpuData := d.CPU
	cpu := fmt.Sprintf(cpuFormat, cpuData[0], cpuData[1], cpuData[2], cpuData[3])

	temp := fmt.Sprintf(tempFormat, d.Temperature)

	tmb, amb := d.Memory.Total/1024, d.Memory.Available/1024
	mem := fmt.Sprintf(memFormat, tmb, tmb-amb, amb)

	t := time.Now().Format("15:04:05")

	log := fmt.Sprintf(logFormat, t, uptime, cpu, temp, mem)

	return log
}

func write(d string) {
	t := time.Now().Format("02-01-2006")
	v := []byte(d)

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.Mkdir(dir, 0755)
		if err != nil {
			logError(errorOutput{
				Error:   err,
				Message: "couldn't create log folder",
			})
		}
	}

	f, err := os.OpenFile(logPath+t, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		logError(errorOutput{
			Error:   err,
			Message: "couldn't open stats log folder",
		})
	}

	if _, err := f.Write(v); err != nil {
		logError(errorOutput{
			err,
			"couldn't write to stats log file",
		})
	}

	if err := f.Close(); err != nil {
		logError(errorOutput{
			err,
			"couldn't close stats log file",
		})
	}
}

func rotate() {
	logs, err := ioutil.ReadDir(dir)
	if err != nil {
		logError(errorOutput{
			err,
			"couldn't rotate logs",
		})
	}

	if len(logs) > maxLogs {
		sort.Slice(logs, func(i, j int) bool {
			t1, _ := time.Parse("02-01-2006", logs[i].Name())
			t2, _ := time.Parse("02-01-2006", logs[j].Name())
			return t1.Before(t2)
		})

		err := os.Remove(logPath + logs[0].Name())
		if err != nil {
			logError(errorOutput{
				err,
				"couldn't delete oldest log file",
			})
		}
	}

}
