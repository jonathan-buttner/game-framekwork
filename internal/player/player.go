package player

import (
	"fmt"
	"log"

	"github.com/jonathan-buttner/game-framework/internal/core"
	"github.com/jonathan-buttner/game-framework/internal/deck"
)

type Player struct {
	Hand               map[string]deck.Card
	TableaCards        map[string]playedCard
	RoundTableaCards   map[string]playedCard
	CardsByOrientation cardTypes
	Name               string
}

func NewPlayer(name string) *Player {
	return &Player{
		Hand:               make(map[string]deck.Card),
		TableaCards:        make(map[string]playedCard),
		RoundTableaCards:   make(map[string]playedCard),
		CardsByOrientation: newCardTypes(),
		Name:               name,
	}
}

func (p *Player) SetHand(cards []deck.Card) {
	newHand := make(map[string]deck.Card)
	for _, card := range cards {
		newHand[card.ID()] = card
	}
	p.Hand = newHand
}

func (p *Player) GetHand() []deck.Card {
	var hand []deck.Card
	for _, card := range p.Hand {
		hand = append(hand, card)
	}

	return hand
}

func (p *Player) PlayCardFromHand(cardID string, orientation deck.CardOrientation, game core.Game) error {
	cardFromHand, ok := p.Hand[cardID]
	if !ok {
		log.Fatalf("requested card from hand: %v does not exist to play", cardID)
	}

	if !cardFromHand.IsOrientationValid(orientation) {
		return fmt.Errorf("requested card orientation: %v is not valid", orientation.String())
	}

	cardWithOrientation := newPlayedCard(cardFromHand, orientation)
	delete(p.Hand, cardFromHand.ID())

	p.CardsByOrientation.addCard(cardWithOrientation)
	cardWithOrientation.PerformPlayToTableaAction(game)
	return nil
}

func (p *Player) PerformEndTurn(game core.Game) {
	// p.TurnCard.PerformEndTurnAction(game)
	// p.RoundTableaCards[p.TurnCard.ID()] = *p.TurnCard

}

func (p *Player) PerformEndRoundAction(game core.Game, cardID string) error {
	card, ok := p.RoundTableaCards[cardID]
	if !ok {
		return fmt.Errorf("requested card id: %v is not valid", cardID)
	}

	card.PerformEndRoundAction(game)
	return nil
}

type playedCard struct {
	deck.Card
	deck.CardAction

	orientation deck.CardOrientation
}

func newPlayedCard(card deck.Card, orientation deck.CardOrientation) playedCard {
	return playedCard{card, card.GetOrientationAction(orientation), orientation}
}

type cardTypes map[deck.CardOrientation][]playedCard

func newCardTypes() cardTypes {
	return make(map[deck.CardOrientation][]playedCard)
}

func (c cardTypes) addCard(card playedCard) {
	c[card.orientation] = append(c[card.orientation], card)
}

type Game interface {
}
