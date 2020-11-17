package app

import (
	"github.com/tarm/serial"
)

// Connect opens and returns serial port as defined in config
func Connect(config *serial.Config) (*serial.Port, error) {
	conn, err := serial.OpenPort(config)
	return conn, err
}
