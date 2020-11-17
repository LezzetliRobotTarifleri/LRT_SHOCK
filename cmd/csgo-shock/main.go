// Copyright 2020 Lezzetli Robot Tarifleri
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"log"
	"time"

	"github.com/LezzetliRobotTarifleri/csgo-shock/app"
	"github.com/LezzetliRobotTarifleri/csgo-shock/config"
)

func main() {

	// Define connection settings
	conf := config.PortSetup()

	// Open and connect port with defined config parameters
	serialPort, err := app.Connect(conf)
	if err != nil {
		log.Fatal(err)
	}

	// We need at least 1 second after openning port
	// to send data via serial connection
	<-time.After(time.Second)

	// If device connected successfully, then we can
	// start listening game and send data
	game := config.GameSetup()

	// Send game data (player health) to serial port
	go app.Game(game, serialPort)

	// Starts listening to address provided.
	// If a POST request is received, it sends it through
	// the Game.Channel channel.
	game.Listen(":3000")
}
