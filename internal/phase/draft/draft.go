package draft

import (
	"fmt"

	"github.com/jonathan-buttner/game-framework/internal/core"
	"github.com/jonathan-buttner/game-framework/internal/deck"
	"github.com/jonathan-buttner/game-framework/internal/phase"
	"github.com/jonathan-buttner/game-framework/internal/player"
)

const startingHandSize = 5

type draftAction struct {
	player *player.Player
	card   deck.PositionedCard
}

func (d *draftAction) Execute(gameState *core.GameState) error {
	return d.player.PlayCardFromHand(d.card.ID(), d.card.Orientation, gameState)
}

type Draft struct {
	phase.Phase
	deck *deck.Deck
	step phase.Step
}

func NewDraftPhase(phase phase.Phase, deck *deck.Deck) *Draft {
	return &Draft{Phase: phase, deck: deck}
}

func (d *Draft) Setup() {
	d.deck.Shuffle()

	// deal 5 cards to each player
	d.PlayerManager.ExecuteForPlayers(func(player *player.Player) error {
		d.deck.DealCards(startingHandSize, player)
		return nil
	})

	d.step = draftCardStep{d.PlayerManager.CurrentPlayer()}
}

func (d *Draft) GetActions() []phase.Action {
	return d.step.GetActions()
}

func (d *Draft) PerformAction(action phase.Action) {
	err := action.Execute(d.GameState)
	if err != nil {
		panic(fmt.Errorf("executing action failed err: %v", err))
	}

	if d.PlayerManager.CurrentPlayer().ResourceCountExceedsLimit() {
		d.step = phase.NewReduceResourcesStep(d.PlayerManager.CurrentPlayer())
	} else {
		d.PlayerManager.NextPlayer()
		d.step = draftCardStep{d.PlayerManager.CurrentPlayer()}
	}
}

type draftCardStep struct {
	Player *player.Player
}

func (d draftCardStep) GetActions() []phase.Action {
	cardsInHandWithPositions := d.Player.GetHand().AllPositionCombinations()
	validCards := d.Player.ValidOrientations(cardsInHandWithPositions)

	var actions []phase.Action
	for _, card := range validCards {
		actions = append(actions, &draftAction{player: d.Player, card: card})
	}

	return actions
}
