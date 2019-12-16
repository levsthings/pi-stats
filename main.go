package main

import "fmt"

func main() {

	cpu := sampleCPU()
	fmt.Println(cpu)

	mem := sampleMemory()
	fmt.Println(mem)

	temp := sampleTemp()
	fmt.Println(temp)

	uptime := getUptime()
	fmt.Println(uptime)

}
