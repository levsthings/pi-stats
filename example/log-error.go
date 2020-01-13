package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

type errorOutput struct {
	Error   error
	Message string
}

// Log error builds an error using errorOutput and attempts to log it
// to disk, and exits with error code.

func logError(e errorOutput) {
	filename := "pi-stats.error.log"

	t := time.Now().Format("02-01-2006 15:04:05")
	s := fmt.Sprintf("time=%q, level=fatal, error=%q, msg=%q\n", t, e.Error, e.Message)
	b := []byte(s)

	f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	if _, err := f.Write(b); err != nil {
		log.Fatal(err)
	}

	if err := f.Close(); err != nil {
		log.Fatal(err)
	}

	os.Exit(1)
}
