package config

import "github.com/dank/go-csgsi"

// GameSetup creates and returns a Game object
func GameSetup() *csgsi.Game {
	return csgsi.New(0)
}
