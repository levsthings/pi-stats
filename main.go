package main

import "time"

func main() {
	// TODO: Clean up if it has more than 6 entries
	// TODO: Run as background process on system startup
	for {
		logger()
		time.Sleep(time.Minute * 5)
	}
}
