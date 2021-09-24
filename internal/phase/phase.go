package phase

import (
	"fmt"

	"github.com/jonathan-buttner/game-framework/internal/core"
	"github.com/jonathan-buttner/game-framework/internal/player"
)

//go:generate mockgen -destination=../../mocks/mock_phase.go -package=mocks github.com/jonathan-buttner/game-framework/internal/phase PhaseHandler
//go:generate stringer -type=ActionType

type Step interface {
	GetActions() []Action
}

type Action interface {
	fmt.Stringer
	Execute(gameState *core.GameState) error
	Type() ActionType
}

type ActionType int

const (
	ChooseCard ActionType = iota
	ReduceResources
	UseResources
	SkipUseResources
)

type ActionsHandler struct {
	DetailedActions map[ActionType]map[string]Action
	Actions         map[ActionType][]Action
}

func NewActionsHandler(actions []Action) ActionsHandler {
	actionsMap := make(map[ActionType][]Action)

	for _, action := range actions {
		actionsArray, ok := actionsMap[action.Type()]
		if ok {
			actionsMap[action.Type()] = append(actionsArray, action)
		} else {
			actionsMap[action.Type()] = []Action{action}
		}
	}

	detailedActions := make(map[ActionType]map[string]Action)
	for _, action := range actions {
		stringToAction, ok := detailedActions[action.Type()]

		if ok {
			stringToAction[action.String()] = action
		} else {
			detailedActions[action.Type()] = map[string]Action{action.String(): action}
		}
	}

	return ActionsHandler{DetailedActions: detailedActions, Actions: actionsMap}
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
