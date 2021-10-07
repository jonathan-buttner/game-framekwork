package test

import (
	"fmt"

	"github.com/golang/mock/gomock"
	"github.com/jonathan-buttner/game-framework/internal/core"
	"github.com/jonathan-buttner/game-framework/internal/deck"
	"github.com/jonathan-buttner/game-framework/internal/phase"
	"github.com/jonathan-buttner/game-framework/internal/phase/draft"
	"github.com/jonathan-buttner/game-framework/internal/player"
	"github.com/jonathan-buttner/game-framework/internal/resource"
	"github.com/jonathan-buttner/game-framework/internal/rules"
	"github.com/jonathan-buttner/game-framework/test/mocks"
)

func NewMockDeckWithoutShuffle(ctrl *gomock.Controller, cards []deck.Card) draft.Deck {
	mockDeck := mocks.NewMockDeck(ctrl)
	concreteDeck := deck.NewDeck(cards)

	mockDeck.EXPECT().Shuffle().AnyTimes()
	mockDeck.EXPECT().DealCards(gomock.Any(), gomock.Any()).Do(func(handSize int, player deck.Player) {
		concreteDeck.DealCards(handSize, player)
	}).AnyTimes()

	return mockDeck
}

type DraftTest struct {
	Player1 *player.Player
	Player2 *player.Player
	Draft   *draft.Draft
}

func (d DraftTest) ExecuteChooseCardByIndex(index int) {
	actions := d.Draft.GetActionsHandler()
	actions.Actions[phase.ChooseCard][index].Execute(&core.GameState{})
}

func (d DraftTest) ExecuteChooseCardByID(id string) {
	actions := d.Draft.GetActionsHandler()
	actions.DetailedActions[phase.ChooseCard][id].Execute(&core.GameState{})
}

func (d DraftTest) ExecuteSkipUseResources() {
	actions := d.Draft.GetActionsHandler()
	actions.Actions[phase.SkipUseResources][0].Execute(&core.GameState{})
}

type DraftTestSetupInfo struct {
	Player1 PlayerInfo
	Player2 PlayerInfo
}

func NewDraftTest(ctrl *gomock.Controller, setupInfo DraftTestSetupInfo) DraftTest {
	p1HandSize := len(setupInfo.Player1.HandSetup)
	p2HandSize := len(setupInfo.Player2.HandSetup)

	if p1HandSize != p2HandSize {
		panic(fmt.Sprintf("player cards must be the same length, p1: %v, p2: %v", p1HandSize, p2HandSize))
	}

	player1, p1Cards := setupInfo.Player1.Create(ctrl)
	player2, p2Cards := setupInfo.Player2.Create(ctrl)

	manager := phase.NewPlayerManager([]*player.Player{player1, player2})
	phase := phase.Phase{PlayerManager: manager, GameState: &core.GameState{}}

	deck := NewMockDeckWithoutShuffle(ctrl, append(p2Cards, p1Cards...))
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
