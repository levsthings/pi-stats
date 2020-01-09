package pistats

// Data describes the structure that GetData returns
type Data struct {
	Uptime
	*CPU
	Temperature
	*Memory
}

// GetData calls all data points and returns Data
func GetData() (*Data, error) {

	u, err := getUptime()
	if err != nil {
		return nil, err
	}

	c, err := sampleCPU()
	if err != nil {
		return nil, err
	}

	t, err := sampleTemp()
	if err != nil {
		return nil, err
	}

	m, err := sampleMemory()
	if err != nil {
		return nil, err
	}

	data := Data{
		u,
		c,
		t,
		m,
	}

	return &data, nil
}
