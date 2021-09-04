package phase

import (
	"github.com/jonathan-buttner/game-framework/internal/core"
	"github.com/jonathan-buttner/game-framework/internal/deck"
	"github.com/jonathan-buttner/game-framework/internal/player"
)

type Phase struct {
	playerManager *PlayerManager
	gameState     *core.GameState
}

type DraftAction struct {
	player *player.Player
	card   deck.PositionedCard
}

func (d *DraftAction) Execute(gameState *core.GameState) error {
	return d.player.PlayCardFromHand(d.card.ID(), d.card.Orientation, gameState)
}

type Action interface {
	Execute(gameState *core.GameState) error
}

type Draft struct {
	Phase
	deck *deck.Deck
	step Step
}

func (d *Draft) Setup() {
	d.deck.Shuffle()

	// deal 5 cards to each player
	d.playerManager.ExecuteForPlayers(func(player *player.Player) error {
		d.deck.DealCards(5, d.playerManager.CurrentPlayer())
		return nil
	})

	d.step = DraftCardStep{d.playerManager, d.gameState}
}

func (d *Draft) GetActions() []Action {
	return d.step.GetActions()
}

func (d *Draft) PerformAction(action Action) {
	action.Execute(d.gameState)
	// TODO: handle error

	if d.playerManager.CurrentPlayer().ResourceCountExceedsLimit() {
		d.step = ReduceResourcesStep{d.playerManager, d.gameState}
	}
	// Check if the player's resources are not valid (greater than 10)
	// if they are invalid them don't change players yet

	// if they are valid then change to new player
}

type Step interface {
	GetActions() []Action
}

type ReduceResourcesStep struct {
	// TODO: maybe switch these to current player?
	playerManager *PlayerManager

	// TODO: might not need this?
	gameState *core.GameState
}

func (ReduceResourcesStep) GetActions() []Action {
	// TODO: get the resources that have a count > 0 so as options to give back one
	return nil
}

type DraftCardStep struct {
	// TODO: maybe switch these to current player?
	playerManager *PlayerManager
	gameState     *core.GameState
}

func (d DraftCardStep) GetActions() []Action {
	cardsInHandWithPositions := d.playerManager.CurrentPlayer().GetHand().AllPositionCombinations()
	validCards := d.playerManager.CurrentPlayer().ValidOrientations(cardsInHandWithPositions)

	var actions []Action
	for _, card := range validCards {
		actions = append(actions, &DraftAction{player: d.playerManager.CurrentPlayer(), card: card})
	}

	return actions
}
