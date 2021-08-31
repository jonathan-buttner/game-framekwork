package phase

import (
	"github.com/jonathan-buttner/game-framework/internal/deck"
	"github.com/jonathan-buttner/game-framework/internal/player"
)

type Phase struct {
	playerManager *PlayerManager
}

type DraftAction struct {
	player *player.Player
	cardID string
}

type Draft struct {
	Phase
	deck *deck.Deck
}

func (d *Draft) Setup() {
	d.deck.Shuffle()

	// deal 5 cards to each player
	d.playerManager.ExecuteForPlayers(func(player *player.Player) error {
		d.deck.DealCards(5, d.playerManager.CurrentPlayer())
		return nil
	})
}

func (d *Draft) GetActions() []DraftAction {
	d.currentPlayer.GetHand()
}

func (d *Draft) PerformAction(action DraftAction) {

}
