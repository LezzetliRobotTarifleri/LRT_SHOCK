package config

import "github.com/tarm/serial"

// PortSetup function is used to define port settings
func PortSetup() *serial.Config {
	config := &serial.Config{
		Name: "COM10",
		Baud: 9600,
	}

	return config
}
