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
	gameState   *core.GameState
	player      *player.Player
	cardID      string
	orientation deck.CardOrientation
}

func (d *DraftAction) Execute() error {
	return d.player.PlayCardFromHand(d.cardID, d.orientation, d.gameState)
}

type Action interface {
	Execute() error
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

	d.step = &DraftCardStep{d.playerManager, d.gameState}
}

func (d *Draft) GetActions() []Action {
	return d.step.GetActions()
}

func (d *Draft) PerformAction(action Action) {

}

type Step interface {
	GetActions() []Action
}

type DraftCardStep struct {
	playerManager *PlayerManager
	gameState     *core.GameState
}

func (d *DraftCardStep) GetActions() []Action {
	hand := d.playerManager.CurrentPlayer().GetHand()

	// validActions := getValidActions(hand) this determines based on the state of the player
	// what card orientations are valid
	var actions []Action
	for _, card := range hand {
		// TODO: need to finish the orientations
		actions = append(actions, &DraftAction{player: d.playerManager.CurrentPlayer(), cardID: card.ID(), orientation: deck.VictoryPoints, gameState: d.gameState})
	}

	return actions
}
