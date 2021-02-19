package config

import "github.com/tarm/serial"

// PortSetup function is used to define port settings
func PortSetup(comPort string, baudRate int) *serial.Config {
	config := &serial.Config{
		Name: comPort,
		Baud: baudRate,
	}

	return config
}
