package phase_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/jonathan-buttner/game-framework/internal/core"
	"github.com/jonathan-buttner/game-framework/internal/deck"
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

func TestUseResourcesStepGetActionIncludesVictoryPointCards(t *testing.T) {
	ctrl := gomock.NewController(t)

	handler, player := createPhaseHandlerAndPlayer(ctrl)
	setupPlayerWithBrownAndOneVictoryCard(ctrl, player)

	step := phase.UseResourcesStep{Phase: handler}
	actions := step.GetActions()
	actions[0].Execute(&core.GameState{})

	assert.Len(t, actions, 2)
}

func TestUseResourcesStepGetActionIncludesTradeCards(t *testing.T) {
	ctrl := gomock.NewController(t)

	handler, player := createPhaseHandlerAndPlayer(ctrl)
	setupPlayerWithBrownAndOneTradeCard(ctrl, player)

	step := phase.UseResourcesStep{Phase: handler}
	actions := step.GetActions()
	actions[0].Execute(&core.GameState{})

	assert.Len(t, actions, 2)
}

func TestUseResourcesStepGetActionIncludesSkip(t *testing.T) {
	ctrl := gomock.NewController(t)

	handler, _ := createPhaseHandlerAndPlayer(ctrl)
	handler.EXPECT().NextPlayer().Times(1)

	step := phase.UseResourcesStep{Phase: handler}
	actions := step.GetActions()
	actions[0].Execute(&core.GameState{})

	assert.Len(t, actions, 1)
}

func createPhaseHandlerAndPlayer(ctrl *gomock.Controller) (*mocks.MockPhaseHandler, *player.Player) {
	player1 := player.NewPlayer("player", rules.NewDefaultGameRules())

	handler := mocks.NewMockPhaseHandler(ctrl)
	handler.EXPECT().CurrentPlayer().Return(player1).AnyTimes()
	handler.EXPECT().SetStep(gomock.Any()).AnyTimes()

	return handler, player1
}

func setupPlayerWithBrownAndOneVictoryCard(ctrl *gomock.Controller, player *player.Player) {
	player.ResourceHandler.AddResources(resource.GroupedResources{resource.Brown: 1})

	actions := mocks.NewMockOrientationActions(ctrl)
	actions.EXPECT().UseCost().Return(resource.GroupedResources{resource.Brown: 1}).AnyTimes()
	actions.EXPECT().PerformUseResourceAction(gomock.Any()).Times(1)
	aCard := mocks.NewPositionedMockCard(ctrl, deck.VictoryPoints, actions)

	player.CardsByOrientation[deck.VictoryPoints] = []deck.PositionedCard{aCard}
}

func setupPlayerWithBrownAndOneTradeCard(ctrl *gomock.Controller, player *player.Player) {
	player.ResourceHandler.AddResources(resource.GroupedResources{resource.Brown: 1})

	actions := mocks.NewMockOrientationActions(ctrl)
	actions.EXPECT().UseCost().Return(resource.GroupedResources{resource.Brown: 1}).AnyTimes()
	actions.EXPECT().PerformUseResourceAction(gomock.Any()).Times(1)
	aCard := mocks.NewPositionedMockCard(ctrl, deck.Trade, actions)

	player.CardsByOrientation[deck.VictoryPoints] = []deck.PositionedCard{aCard}
}
