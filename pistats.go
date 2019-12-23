package pistats

// Data describes the structure that GetData returns
type Data struct {
	Uptime      string
	CPU         cpu
	Temperature string
	Memory      memory
}

// GetData calls all data points and returns Data
func GetData() Data {
	data := Data{
		getUptime(),
		sampleCPU(),
		sampleTemp(),
		sampleMemory(),
	}

	return data
}
