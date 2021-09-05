package player

import (
	"fmt"
	"log"

	"github.com/jonathan-buttner/game-framework/internal/deck"
	"github.com/jonathan-buttner/game-framework/internal/resource"
	"github.com/jonathan-buttner/game-framework/internal/rules"
)

type Player struct {
	rules.GameRules

	hand               map[string]deck.Card
	TableaCards        map[string]deck.PositionedCard
	RoundTableaCards   map[string]deck.PositionedCard
	CardsByOrientation cardTypes
	ResourceHandler    *resource.ResourceHandler
	Name               string
}

func NewPlayer(name string, gameRules rules.GameRules) *Player {
	return &Player{
		ResourceHandler:    resource.NewResourceHandler(),
		hand:               make(map[string]deck.Card),
		TableaCards:        make(map[string]deck.PositionedCard),
		RoundTableaCards:   make(map[string]deck.PositionedCard),
		CardsByOrientation: newCardTypes(),
		GameRules:          gameRules,
		Name:               name,
	}
}

func (p *Player) ResourceCountExceedsLimit() bool {
	return p.ResourceHandler.Count > p.GameRules.ResourceLimit
}

func (p *Player) SetHand(cards []deck.Card) {
	newHand := make(map[string]deck.Card)
	for _, card := range cards {
		newHand[card.ID()] = card
	}
	p.hand = newHand
}

func (p *Player) GetHand() deck.Cards {
	var hand []deck.Card
	for _, card := range p.hand {
		hand = append(hand, card)
	}

	return hand
}

func (p *Player) ValidOrientations(positionCards []deck.PositionedCard) []deck.PositionedCard {
	var validCards []deck.PositionedCard

	for _, cardWithPosition := range positionCards {
		if p.ResourceHandler.HasResources(cardWithPosition.Cost()) {
			validCards = append(validCards, cardWithPosition)
		}
	}

	return validCards
}

func (p *Player) PlayCardFromHand(cardID string, orientation deck.CardOrientation, game deck.Game) error {
	cardFromHand, ok := p.hand[cardID]
	if !ok {
		log.Fatalf("requested card from hand: %v does not exist to play", cardID)
	}

	if !cardFromHand.IsOrientationValid(orientation) {
		return fmt.Errorf("requested card orientation: %v is not valid", orientation.String())
	}

	cardWithOrientation := deck.NewPositionedCard(cardFromHand, orientation)
	delete(p.hand, cardFromHand.ID())

	p.CardsByOrientation.addCard(cardWithOrientation)
	cardWithOrientation.PerformPlayToTableaAction(game)
	return nil
}

func (p *Player) PerformEndRoundAction(game deck.Game, cardID string) error {
	card, ok := p.RoundTableaCards[cardID]
	if !ok {
		return fmt.Errorf("requested card id: %v is not valid", cardID)
	}

	card.PerformEndRoundAction(game)
	return nil
}

type cardTypes map[deck.CardOrientation][]deck.PositionedCard

func newCardTypes() cardTypes {
	return make(map[deck.CardOrientation][]deck.PositionedCard)
}

func (c cardTypes) addCard(card deck.PositionedCard) {
	c[card.Orientation] = append(c[card.Orientation], card)
}
