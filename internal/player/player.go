package player

import (
	"fmt"
	"log"

	"github.com/jonathan-buttner/game-framework/internal/deck"
)

type Player struct {
	Hand             map[string]Card
	TableaCards      map[string]Card
	RoundTableaCards map[string]Card
}

func (p *Player) SetHand(cards []Card) {
	for _, card := range cards {
		p.Hand[card.ID()] = card
	}
}

func (p *Player) PlayCardFromHand(cardID string, orientation deck.CardOrientation) error {
	cardFromHand, ok := p.Hand[cardID]
	if !ok {
		log.Fatalf("requested card from hand: %v does not exist to play", cardID)
	}

	if !cardFromHand.IsOrientationValid(orientation) {
		return fmt.Errorf("requested card orientation: %v is not valid", orientation.String())
	}

	p.RoundTableaCards[cardFromHand.ID()] = cardFromHand
	delete(p.Hand, cardFromHand.ID())

	// cardFromHand.performAction(p)
	return nil
}

type Card interface {
	PerformAction(player Player, orientation deck.CardOrientation)
	IsOrientationValid(orientation deck.CardOrientation) bool
	ID() string
}
