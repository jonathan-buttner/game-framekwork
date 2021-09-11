package phase_test

import (
	"testing"

	"github.com/jonathan-buttner/game-framework/internal/core"
	"github.com/jonathan-buttner/game-framework/internal/phase"
	"github.com/jonathan-buttner/game-framework/internal/player"
	"github.com/jonathan-buttner/game-framework/internal/resource"
	"github.com/jonathan-buttner/game-framework/internal/rules"
	"github.com/stretchr/testify/assert"
)

func TestNoActionsWhenNoResources(t *testing.T) {
	player := player.NewPlayer("player", rules.NewDefaultGameRules())
	resStep := phase.NewReduceResourcesStep(player)

	assert.Nil(t, resStep.GetActions())
}

func TestExecutingActionReducesYellow(t *testing.T) {
	player := player.NewPlayer("player", rules.NewDefaultGameRules())
	player.ResourceHandler.AddResources(resource.GroupedResources{resource.Yellow: 2})

	resStep := phase.NewReduceResourcesStep(player)

	assert.Len(t, resStep.GetActions(), 1)

	resStep.GetActions()[0].Execute(&core.GameState{})
	assert.Equal(t, player.ResourceHandler.Resources[resource.Yellow], 1)
}

func TestActionPerResourceType(t *testing.T) {
	player := player.NewPlayer("player", rules.NewDefaultGameRules())
	player.ResourceHandler.AddResources(resource.GroupedResources{resource.Yellow: 2, resource.Brown: 1, resource.Green: 1})

	resStep := phase.NewReduceResourcesStep(player)

	assert.Len(t, resStep.GetActions(), 3)
}
