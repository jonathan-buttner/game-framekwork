package phase

import (
	"github.com/jonathan-buttner/game-framework/internal/core"
	"github.com/jonathan-buttner/game-framework/internal/resource"
)

type ReduceResourcesStep struct {
	phase *PhaseWithTurnHandler
}

func NewReduceResourcesStep(phase *PhaseWithTurnHandler) ReduceResourcesStep {
	return ReduceResourcesStep{phase}
}

func (r ReduceResourcesStep) GetActions() []Action {
	var actions []Action

	for resType, count := range r.phase.PlayerManager.CurrentPlayer().ResourceHandler.Resources {
		if count > 0 {
			actions = append(actions, &reduceResourcesAction{phase: r.phase, resourceType: resType})
		}
	}

	return actions
}

type reduceResourcesAction struct {
	phase        *PhaseWithTurnHandler
	resourceType resource.ResourceType
}

func (r *reduceResourcesAction) Execute(gameState *core.GameState) error {
	err := r.phase.PlayerManager.CurrentPlayer().ResourceHandler.RemoveResources(resource.GroupedResources{r.resourceType: 1})

	if r.phase.PlayerManager.CurrentPlayer().ResourceCountExceedsLimit() {
		r.phase.Step = NewReduceResourcesStep(r.phase)
	} else {
		r.phase.PlayerTurnHandler.GoToNextPlayer()
	}

	return err
}
