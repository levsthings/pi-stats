package main

import "time"

func main() {
	// TODO: Remove formatting logic from data collectors
	// TODO: Add formatting logic to logger
	// TODO: Add README
	for {
		logger()
		time.Sleep(time.Minute * 5)
	}
}
