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
	"github.com/jonathan-buttner/game-framework/test"
	"github.com/jonathan-buttner/game-framework/test/mocks"
	"github.com/stretchr/testify/assert"
)

func TestSetupDeals5Cards(t *testing.T) {
	ctrl := gomock.NewController(t)

	_, players := newDraftPhase(ctrl, 5)

	assert.Len(t, players[0].GetHand(), 5)
	assert.Len(t, players[1].GetHand(), 5)
}

func TestGetActionsReturns20Actions(t *testing.T) {
	ctrl := gomock.NewController(t)

	phase, _ := newDraftPhase(ctrl, 5)

	assert.Len(t, phase.GetActions(), 5*len(deck.Orientations))
}

func TestRotateHands(t *testing.T) {
	ctrl := gomock.NewController(t)

	draftPhase, players := newDraftPhase(ctrl, 1)

	assert.Equal(t, "1", players[0].GetHand()[0].ID())
	assert.Equal(t, "0", players[1].GetHand()[0].ID())

	// this goes to player2
	draftPhase.NextPlayer()
	// this causes the hands to rotate and resets the players
	draftPhase.NextPlayer()

	assert.Equal(t, "0", players[0].GetHand()[0].ID())
	assert.Equal(t, "1", players[1].GetHand()[0].ID())

	// this goes to player1
	draftPhase.NextPlayer()
	// rotate hands
	draftPhase.NextPlayer()

	assert.Equal(t, "1", players[0].GetHand()[0].ID())
	assert.Equal(t, "0", players[1].GetHand()[0].ID())
}

func TestDraftASingleCard(t *testing.T) {
	ctrl := gomock.NewController(t)

	tester := test.DraftTestCreator{
		Player1: test.PlayerInfo{
			Name:      "player1",
			Resources: resource.GroupedResources{resource.Brown: 12},
			HandSetup: []test.CardInfo{
				{
					UseCost:     resource.GroupedResources{resource.Brown: 1},
					AcquireCost: resource.GroupedResources{resource.Brown: 1},
					Orientation: deck.Trade,
					Name:        "player1-Trade-1Brown",
				},
			},
		},
		Player2: test.PlayerInfo{
			Name:      "player2",
			Resources: resource.GroupedResources{resource.Brown: 1},
			HandSetup: []test.CardInfo{
				{
					UseCost:     resource.GroupedResources{resource.Brown: 1},
					AcquireCost: resource.GroupedResources{resource.Brown: 1},
					Orientation: deck.Trade,
					Name:        "player2-Trade-1Brown",
				},
			},
		},
	}

	draftTest := tester.Create(ctrl)
	draftTest.ExecuteChooseCardByID("player1-Trade-1Brown")

	_, hasPlayer1Card := draftTest.Player1.RoundTableaCards["player1-Trade-1Brown"]
	assert.True(t, hasPlayer1Card)
}

func TestSecondPlayerDraftCard(t *testing.T) {
	ctrl := gomock.NewController(t)

	testCreator := test.DraftTestCreator{
		Player1: test.PlayerInfo{
			Name:      "player1",
			Resources: resource.GroupedResources{resource.Brown: 2},
			HandSetup: []test.CardInfo{
				{
					UseCost:     resource.GroupedResources{resource.Brown: 1},
					AcquireCost: resource.GroupedResources{resource.Brown: 1},
					Orientation: deck.Trade,
					Name:        "player1-Trade-1Brown",
				},
			},
		},
		Player2: test.PlayerInfo{
			Name:      "player2",
			Resources: resource.GroupedResources{resource.Brown: 1},
			HandSetup: []test.CardInfo{
				{
					UseCost:     resource.GroupedResources{resource.Brown: 1},
					AcquireCost: resource.GroupedResources{resource.Brown: 1},
					Orientation: deck.Trade,
					Name:        "player2-Trade-1Brown",
				},
			},
		},
	}

	draftTest := testCreator.Create(ctrl)
	draftTest.ExecuteChooseCardByID("player1-Trade-1Brown")
	draftTest.ExecuteSkipUseResources()

	draftTest.ExecuteChooseCardByID("player2-Trade-1Brown")

	_, hasPlayer2Card := draftTest.Player2.RoundTableaCards["player2-Trade-1Brown"]
	assert.True(t, hasPlayer2Card)
}

func newDraftPhase(ctrl *gomock.Controller, numCardsToDeal int) (*draft.Draft, []*player.Player) {
	basePhase, players := createBasePhase()
	deck := createDeck(ctrl, numCardsToDeal, len(players))
	phase := draft.NewDraftPhase(basePhase, deck, numCardsToDeal)

	return phase, players
}

func createDeck(ctrl *gomock.Controller, numCardToDeal int, numberOfPlayers int) draft.Deck {
	cardAction := mocks.NewMockOrientationActions(ctrl)
	cardAction.EXPECT().PerformPlayToTableaAction(gomock.Any()).AnyTimes()
	cardAction.EXPECT().UseCost().Return(resource.GroupedResources{resource.Yellow: 1}).AnyTimes()
	cardAction.EXPECT().AcquireCost().Return(resource.GroupedResources{resource.Yellow: 1}).AnyTimes()
	cardAction.EXPECT().IsOrientationValid(gomock.Any()).Return(true).AnyTimes()

	var cards []deck.Card

	for i := 0; i < numCardToDeal*numberOfPlayers; i++ {
		card := mocks.NewMockCard(ctrl)
		card.EXPECT().ID().Return(fmt.Sprintf("%v", i)).AnyTimes()
		card.EXPECT().GetOrientationActions(gomock.Any()).Return(cardAction).AnyTimes()

		cards = append(cards, card)
	}

	return test.NewMockDeckWithoutShuffle(ctrl, cards)
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
