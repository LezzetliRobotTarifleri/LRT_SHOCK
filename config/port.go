package config

import "github.com/tarm/serial"

// PortSetup function is used to define port settings
func PortSetup() *serial.Config {
	config := &serial.Config{
		Name: "COM6",
		Baud: 9600,
	}

	return config
}
