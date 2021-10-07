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
	ID() string
}

type ActionType int

const (
	ChooseCard ActionType = iota
	ReduceResources
	UseResources
	SkipUseResources
)

type (
	ActionsByType  map[ActionType][]Action
	ActionsHandler struct {
		DetailedActions map[ActionType]map[string]Action
		Actions         ActionsByType
	}
)

func NewActionsHandler(actions []Action) ActionsHandler {
	actionsMap := getActionsByType(actions)

	detailedActions := make(map[ActionType]map[string]Action)
	for _, action := range actions {
		stringToAction, ok := detailedActions[action.Type()]

		if ok {
			stringToAction[action.ID()] = action
		} else {
			detailedActions[action.Type()] = map[string]Action{action.ID(): action}
		}
	}

	return ActionsHandler{DetailedActions: detailedActions, Actions: actionsMap}
}

func getActionsByType(actions []Action) ActionsByType {
	actionsMap := make(ActionsByType)

	for _, action := range actions {
		actionsArray, ok := actionsMap[action.Type()]
		if ok {
			actionsMap[action.Type()] = append(actionsArray, action)
		} else {
			actionsMap[action.Type()] = []Action{action}
		}
	}

	return actionsMap
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
