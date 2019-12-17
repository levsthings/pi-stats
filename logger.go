package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"time"
)

const (
	dir     = ".pi-stat"
	logPath = dir + "/"
	perms   = 0770
	maxLogs = 7
)

func logger() {
	h, m, s := time.Now().Clock()
	t := fmt.Sprintf("%d:%d:%d", h, m, s)

	log := fmt.Sprintf("%s, %s, %s, %s, %s\n", t, getUptime(), sampleCPU(), sampleTemp(), sampleMemory())

	write(log)
	rotate()
}

func write(d string) {
	t := time.Now().Format("02-01-2006")
	v := []byte(d)

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.Mkdir(dir, perms)
	}

	f, err := os.OpenFile(logPath+t, os.O_APPEND|os.O_CREATE|os.O_WRONLY, perms)
	if err != nil {
		log.Fatal(err)
	}

	if _, err := f.Write(v); err != nil {
		log.Fatal(err)
	}

	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}

func rotate() {
	logs, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal("couldn't rotate logs", err)
	}

	if len(logs) > maxLogs {
		sort.Slice(logs, func(i, j int) bool {
			t1, _ := time.Parse("02-01-2006", logs[i].Name())
			t2, _ := time.Parse("02-01-2006", logs[j].Name())
			return t1.Before(t2)
		})

		err := os.Remove(logPath + logs[0].Name())
		if err != nil {
			log.Println("error deleting oldest log", err)
		}
	}

}
