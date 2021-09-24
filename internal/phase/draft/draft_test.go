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

	tester := DraftTestCreator{
		Player1: PlayerInfo{
			Name:      "player1",
			Resources: resource.GroupedResources{resource.Brown: 12},
			HandSetup: []CardInfo{
				{
					UseCost:     resource.GroupedResources{resource.Brown: 1},
					AcquireCost: resource.GroupedResources{resource.Brown: 1},
					Orientation: deck.Trade,
					Name:        "player1-Trade-1Brown",
				},
			},
		},
		Player2: PlayerInfo{
			Name:      "player2",
			Resources: resource.GroupedResources{resource.Brown: 1},
			HandSetup: []CardInfo{
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
	draftTest.ExecuteChooseCard(0)

	_, hasPlayer1Card := draftTest.Player1.RoundTableaCards["player1-Trade-1Brown"]
	assert.True(t, hasPlayer1Card)
}

func TestSecondPlayerDraftCard(t *testing.T) {
	ctrl := gomock.NewController(t)

	testCreator := DraftTestCreator{
		Player1: PlayerInfo{
			Name:      "player1",
			Resources: resource.GroupedResources{resource.Brown: 2},
			HandSetup: []CardInfo{
				{
					UseCost:     resource.GroupedResources{resource.Brown: 1},
					AcquireCost: resource.GroupedResources{resource.Brown: 1},
					Orientation: deck.Trade,
					Name:        "player1-Trade-1Brown",
				},
			},
		},
		Player2: PlayerInfo{
			Name:      "player2",
			Resources: resource.GroupedResources{resource.Brown: 1},
			HandSetup: []CardInfo{
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
	draftTest.ExecuteChooseCard(0)
	draftTest.ExecuteSkipUseResources()

	draftTest.ExecuteChooseCard(0)

	_, hasPlayer2Card := draftTest.Player2.RoundTableaCards["player2-Trade-1Brown"]
	assert.True(t, hasPlayer2Card)
}

type DraftTest struct {
	Player1 *player.Player
	Player2 *player.Player
	Draft   *draft.Draft
}

func (d DraftTest) ExecuteChooseCard(index int) {
	actions := d.Draft.GetActionsHandler()
	actions.Actions[phase.ChooseCard][index].Execute(&core.GameState{})
}

func (d DraftTest) ExecuteSkipUseResources() {
	actions := d.Draft.GetActionsHandler()
	actions.Actions[phase.SkipUseResources][0].Execute(&core.GameState{})
}

type DraftTestCreator struct {
	Player1 PlayerInfo
	Player2 PlayerInfo
}

func (d DraftTestCreator) Create(ctrl *gomock.Controller) DraftTest {
	p1HandSize := len(d.Player1.HandSetup)
	p2HandSize := len(d.Player2.HandSetup)

	if p1HandSize != p2HandSize {
		panic(fmt.Sprintf("player cards must be the same length, p1: %v, p2: %v", p1HandSize, p2HandSize))
	}

	player1, p1Cards := d.Player1.Create(ctrl)
	player2, p2Cards := d.Player2.Create(ctrl)

	manager := phase.NewPlayerManager([]*player.Player{player1, player2})
	phase := phase.Phase{PlayerManager: manager, GameState: &core.GameState{}}

	deck := mocks.NewMockDeckWithoutShuffle(ctrl, append(p2Cards, p1Cards...))
	draftPhase := draft.NewDraftPhase(phase, deck, len(p1Cards))

	return DraftTest{Player1: player1, Player2: player2, Draft: draftPhase}
}

type PlayerInfo struct {
	Name      string
	Resources resource.GroupedResources
	HandSetup []CardInfo
}

func (p *PlayerInfo) Create(ctrl *gomock.Controller) (*player.Player, []deck.Card) {
	playerInfo := player.NewPlayer(p.Name, rules.NewDefaultGameRules())
	playerInfo.ResourceHandler.AddResources(p.Resources)

	var cards []deck.Card
	for _, card := range p.HandSetup {
		cards = append(cards, card.Create(ctrl))
	}

	return playerInfo, cards
}

type CardInfo struct {
	UseCost     resource.GroupedResources
	AcquireCost resource.GroupedResources
	Orientation deck.CardOrientation
	Name        string
}

func (c CardInfo) Create(ctrl *gomock.Controller) *mocks.MockCard {
	validCardAction := mocks.NewMockOrientationActions(ctrl)
	validCardAction.EXPECT().PerformPlayToTableaAction(gomock.Any()).AnyTimes()
	validCardAction.EXPECT().UseCost().Return(c.UseCost).AnyTimes()
	validCardAction.EXPECT().AcquireCost().Return(c.AcquireCost).AnyTimes()
	validCardAction.EXPECT().IsOrientationValid(gomock.Any()).Return(true).AnyTimes()
	validCardAction.EXPECT().PerformUseResourceAction(gomock.Any()).Do(func(_ *core.GameState) {
		// TODO: reduce resources for the player here
	}).AnyTimes()

	invalidCardAction := mocks.NewMockOrientationActions(ctrl)
	invalidCardAction.EXPECT().PerformPlayToTableaAction(gomock.Any()).AnyTimes()
	invalidCardAction.EXPECT().UseCost().Return(c.UseCost).AnyTimes()
	invalidCardAction.EXPECT().AcquireCost().Return(c.AcquireCost).AnyTimes()
	invalidCardAction.EXPECT().IsOrientationValid(gomock.Any()).Return(false).AnyTimes()

	card := mocks.NewMockCard(ctrl)
	card.EXPECT().ID().Return(c.Name).AnyTimes()

	card.EXPECT().GetOrientationActions(gomock.Any()).DoAndReturn(func(orientation deck.CardOrientation) deck.OrientationActions {
		if orientation == c.Orientation {
			return validCardAction
		}
		return invalidCardAction
	}).AnyTimes()

	return card
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

	return mocks.NewMockDeckWithoutShuffle(ctrl, cards)
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
