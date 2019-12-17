package main

import "time"

func main() {
	for {
		logger()
		time.Sleep(time.Minute * 5)
	}
}
