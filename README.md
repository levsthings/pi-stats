# pi-stats


A tiny stat checker for the Raspberry Pi. It checks uptime, CPU activity, CPU temperature and memory stats.


### Usage


You can import `pi-stats` as a library and use the raw data to feed your own formatter and logger:

```go
package main

import pistats "github.com/levsthings/pi-stats"

func main() {
    data := pistats.GetData()
}
```

### Example Program


`pi-stats` comes with an example program that lives in the `example` folder which can be built and used for local logging or getting console outputs.

You can also download the latest binary from Github [releases](https://github.com/levsthings/pi-stats/releases). 

If ran without any flags, the program will write the data once to stdout and exit. If you supply the `--mode log` flag, the program will run
in it's intended background mode writing a log file every 5 minutes infinitely.

See usage:

```terminal
Usage of pi-stats:
  -mode string
    	Expected input: '--mode console' or '--mode log' (default "console")
```

You can run the binary manually or add it to your startup routine via `systemd` or `rc.local`.The log outputs will automatically go to a directory named 
`.pi-stats`, and will be wherever the binary is run. If you add it to your startup routine, pay attention to the execution context.


### Sample Output


```terminal
16:26:10, uptime: 8h27m14s, CPU 1: 16.98%, CPU2: 13.13%, CPU3: 18.00%, CPU4: 11.88%, temp: 36°C, memtotal: 15952MB, memused: 6103MB, memavailable: 9849MB
```



