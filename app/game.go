package app

import (
	"log"
	"strconv"

	"github.com/dank/go-csgsi"
	"github.com/tarm/serial"
)

// Game function gets data from channel and sends
// to device connected to serial port
func Game(game *csgsi.Game, serialPort *serial.Port) {
	for state := range game.Channel {
		str := strconv.Itoa(state.Player.State.Health)
		_, err := serialPort.Write([]byte(str))
		if err != nil {
			log.Fatal(err)
		}
	}
}
