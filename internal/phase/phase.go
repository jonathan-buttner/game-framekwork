package phase

import "github.com/jonathan-buttner/game-framework/internal/core"

type Step interface {
	GetActions() []Action
}

type Action interface {
	Execute(gameState *core.GameState) error
}

type Phase struct {
	PlayerManager *PlayerManager
	GameState     *core.GameState
}
