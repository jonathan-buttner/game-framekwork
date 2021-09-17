package draft_test

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/jonathan-buttner/game-framework/internal/core"
	"github.com/jonathan-buttner/game-framework/internal/deck"
	"github.com/jonathan-buttner/game-framework/internal/phase"
	"github.com/jonathan-buttner/game-framework/internal/phase/draft"
	"github.com/jonathan-buttner/game-framework/internal/player"
	"github.com/jonathan-buttner/game-framework/internal/resource"
	"github.com/jonathan-buttner/game-framework/internal/rules"
	"github.com/jonathan-buttner/game-framework/mocks"
	"github.com/stretchr/testify/assert"
)

func TestSetupDeals5Cards(t *testing.T) {
	ctrl := gomock.NewController(t)

	phase, players := newDraftPhase(ctrl)
	phase.Setup()

	assert.Len(t, players[0].GetHand(), 5)
	assert.Len(t, players[1].GetHand(), 5)
}

func newDraftPhase(ctrl *gomock.Controller) (*draft.Draft, []*player.Player) {
	deck := createDeck(ctrl)
	basePhase, players := createBasePhase()
	phase := draft.NewDraftPhase(basePhase, deck, 5)

	return phase, players
}

func createDeck(ctrl *gomock.Controller) *deck.Deck {
	cardAction := mocks.NewMockOrientationActions(ctrl)
	cardAction.EXPECT().PerformPlayToTableaAction(gomock.Any()).AnyTimes()
	cardAction.EXPECT().Cost().Return(resource.GroupedResources{resource.Yellow: 1}).AnyTimes()
	cardAction.EXPECT().IsOrientationValid(gomock.Any()).Return(true).AnyTimes()

	var cards []deck.Card

	for i := 0; i < 10; i++ {
		card := mocks.NewMockCard(ctrl)
		card.EXPECT().ID().Return(fmt.Sprintf("%v", i)).AnyTimes()
		card.EXPECT().GetOrientationActions(gomock.Any()).Return(cardAction).AnyTimes()

		cards = append(cards, card)
	}

	return deck.NewDeck(cards)
}

func createBasePhase() (phase.Phase, []*player.Player) {
	player1 := player.NewPlayer("player1", rules.NewDefaultGameRules())
	player1.ResourceHandler.AddResources(resource.GroupedResources{resource.Yellow: 1})

	player2 := player.NewPlayer("player2", rules.NewDefaultGameRules())
	player2.ResourceHandler.AddResources(resource.GroupedResources{resource.Yellow: 1})

	players := []*player.Player{player1, player2}
	manager := phase.NewPlayerManager(players)

	return phase.Phase{PlayerManager: manager, GameState: &core.GameState{}}, players
}

func TestGetActionsReturns20Actions(t *testing.T) {
	ctrl := gomock.NewController(t)

	phase, _ := newDraftPhase(ctrl)

	phase.Setup()

	assert.Len(t, phase.GetActions(), 5*len(deck.Orientations))
}

func TestGetActionsForSecondPlayerReturns20Actions(t *testing.T) {
	ctrl := gomock.NewController(t)

	phase, _ := newDraftPhase(ctrl)

	phase.Setup()

	phase.PerformAction(phase.GetActions()[0])
	assert.Len(t, phase.GetActions(), 5*len(deck.Orientations))
}
