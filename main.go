package main

import (
	"fmt"
)

func main() {

	cpu := sampleCPU()
	fmt.Println(cpu)

	mem := sampleMemory()
	fmt.Println(mem)

	uptime := getUptime()
	fmt.Println(uptime)
}
