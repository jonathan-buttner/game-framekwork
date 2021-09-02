package core

import "github.com/jonathan-buttner/game-framework/internal/player"

// TODO: rename
// This should be used internally by player and cards etc
type GameState struct {
	players []*player.Player
}
