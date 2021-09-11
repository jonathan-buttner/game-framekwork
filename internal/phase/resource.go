package phase

import (
	"github.com/jonathan-buttner/game-framework/internal/core"
	"github.com/jonathan-buttner/game-framework/internal/player"
	"github.com/jonathan-buttner/game-framework/internal/resource"
)

type ReduceResourcesStep struct {
	player *player.Player
}

func NewReduceResourcesStep(player *player.Player) ReduceResourcesStep {
	return ReduceResourcesStep{player}
}

func (r ReduceResourcesStep) GetActions() []Action {
	var actions []Action

	for resType, count := range r.player.ResourceHandler.Resources {
		if count > 0 {
			actions = append(actions, &reduceResourcesAction{player: r.player, resourceType: resType})
		}
	}

	return actions
}

type reduceResourcesAction struct {
	player       *player.Player
	resourceType resource.ResourceType
}

func (r *reduceResourcesAction) Execute(gameState *core.GameState) error {
	return r.player.ResourceHandler.RemoveResources(resource.GroupedResources{r.resourceType: 1})
}
