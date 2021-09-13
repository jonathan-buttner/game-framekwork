package phase

import (
	"github.com/jonathan-buttner/game-framework/internal/core"
	"github.com/jonathan-buttner/game-framework/internal/player"
)

//go:generate mockgen -destination=../../mocks/mock_phase.go -package=mocks github.com/jonathan-buttner/game-framework/internal/phase PhaseHandler

type Step interface {
	GetActions() []Action
}

type Action interface {
	Execute(gameState *core.GameState) error
}

type Phase struct {
	PlayerManager *PlayerManager
	GameState     *core.GameState
	Step          Step
}

func (p *Phase) SetStep(step Step) {
	p.Step = step
}

func (p *Phase) CurrentPlayer() *player.Player {
	return p.PlayerManager.CurrentPlayer()
}

type PhaseHandler interface {
	SetStep(step Step)
	CurrentPlayer() *player.Player
	NextPlayer()
}
