package phase

import (
	"github.com/jonathan-buttner/game-framework/internal/deck"
	"github.com/jonathan-buttner/game-framework/internal/player"
)

type Phase struct {
	currentPlayer *player.Player
}

type DraftAction struct {
	player *player.Player
	cardID string
}

type Draft struct {
	Phase
	players []*player.Player
	deck    *deck.Deck
}

func (d *Draft) Setup() {
	d.deck.Shuffle()
	d.deck.DealCards(5, d.convertToDeckPlayer())
	d.currentPlayer = d.players[0]
}

func (d *Draft) convertToDeckPlayer() []deck.Player {
	deckPlayers := make([]deck.Player, len(d.players))

	for i := range d.players {
		deckPlayers[i] = d.players[i]
	}

	return deckPlayers
}

// func (d *Draft) GetActions() DraftAction {
// 	d.currentPlayer.GetHand()
// }

func (d *Draft) PerformAction(action DraftAction) {

}
