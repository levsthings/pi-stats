package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

const (
	dir     = "out"
	logPath = dir + "/"
	perms   = 0770
)

func logger() {
	h, m, s := time.Now().Clock()
	t := fmt.Sprintf("%d:%d:%d", h, m, s)

	log := fmt.Sprintf("%s, %s, %s, %s, %s\n", t, getUptime(), sampleCPU(), sampleTemp(), sampleMemory())

	write(log)

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
