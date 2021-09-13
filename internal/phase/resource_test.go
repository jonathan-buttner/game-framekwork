package phase_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/jonathan-buttner/game-framework/internal/core"
	"github.com/jonathan-buttner/game-framework/internal/phase"
	"github.com/jonathan-buttner/game-framework/internal/player"
	"github.com/jonathan-buttner/game-framework/internal/resource"
	"github.com/jonathan-buttner/game-framework/internal/rules"
	"github.com/jonathan-buttner/game-framework/mocks"
	"github.com/stretchr/testify/assert"
)

func TestNoActionsWhenNoResources(t *testing.T) {
	ctrl := gomock.NewController(t)

	handler, _ := createPhaseHandlerAndPlayer(ctrl)

	resStep := phase.NewReduceResourcesStep(handler)

	assert.Nil(t, resStep.GetActions())
}

func createPhaseHandlerAndPlayer(ctrl *gomock.Controller) (*mocks.MockPhaseHandler, *player.Player) {
	player1 := player.NewPlayer("player", rules.NewDefaultGameRules())

	handler := mocks.NewMockPhaseHandler(ctrl)
	handler.EXPECT().CurrentPlayer().Return(player1).AnyTimes()
	handler.EXPECT().SetStep(gomock.Any()).AnyTimes()

	return handler, player1
}

func TestExecutingActionReducesYellow(t *testing.T) {
	ctrl := gomock.NewController(t)

	handler, player1 := createPhaseHandlerAndPlayer(ctrl)
	handler.EXPECT().NextPlayer().AnyTimes()
	player1.ResourceHandler.AddResources(resource.GroupedResources{resource.Yellow: 2})

	resStep := phase.NewReduceResourcesStep(handler)

	assert.Len(t, resStep.GetActions(), 1)

	resStep.GetActions()[0].Execute(&core.GameState{})
	assert.Equal(t, player1.ResourceHandler.Resources[resource.Yellow], 1)
}

func TestActionPerResourceType(t *testing.T) {
	ctrl := gomock.NewController(t)

	handler, player := createPhaseHandlerAndPlayer(ctrl)
	player.ResourceHandler.AddResources(resource.GroupedResources{resource.Yellow: 2, resource.Brown: 1, resource.Green: 1})

	resStep := phase.NewReduceResourcesStep(handler)

	assert.Len(t, resStep.GetActions(), 3)
}

func TestStaysInReduceStateWhenStillOverTheLimit(t *testing.T) {
	ctrl := gomock.NewController(t)

	handler, player := createPhaseHandlerAndPlayer(ctrl)
	handler.EXPECT().NextPlayer().Times(1)

	player.SetGameRules(rules.NewResourceLimit(1))
	player.ResourceHandler.AddResources(resource.GroupedResources{resource.Brown: 3})

	resStep := phase.NewReduceResourcesStep(handler)

	assert.Len(t, resStep.GetActions(), 1)
	resStep.GetActions()[0].Execute(&core.GameState{})

	assert.Equal(t, player.ResourceHandler.Resources[resource.Brown], 2)
	assert.Len(t, resStep.GetActions(), 1)
	resStep.GetActions()[0].Execute(&core.GameState{})
	assert.Equal(t, player.ResourceHandler.Resources[resource.Brown], 1)
}
