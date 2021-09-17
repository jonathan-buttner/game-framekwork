package player

import (
	"fmt"
	"log"

	"github.com/jonathan-buttner/game-framework/internal/deck"
	"github.com/jonathan-buttner/game-framework/internal/resource"
	"github.com/jonathan-buttner/game-framework/internal/rules"
)

type Player struct {
	hand map[string]deck.Card

	TableaCards        map[string]deck.PositionedCard
	RoundTableaCards   map[string]deck.PositionedCard
	GameRules          rules.GameRules
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

func (p *Player) SetGameRules(gameRules rules.GameRules) {
	p.GameRules = gameRules
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

func (p *Player) PlayCardFromHand(cardID string, orientation deck.CardOrientation, game deck.Game) error {
	cardFromHand, ok := p.hand[cardID]
	if !ok {
		log.Fatalf("requested card from hand: %v does not exist to play", cardID)
	}

	cardWithOrientation := deck.NewPositionedCard(cardFromHand, orientation)

	// TODO: this should return an error if you try to play a level 2 resource without having a level 1 resource
	if !cardWithOrientation.IsOrientationValid(game) {
		return fmt.Errorf("requested card orientation: %v is not valid", orientation.String())
	}

	delete(p.hand, cardFromHand.ID())

	p.CardsByOrientation.addCard(cardWithOrientation)
	cardWithOrientation.PerformPlayToTableaAction(game)
	return nil
}

type cardTypes map[deck.CardOrientation][]deck.PositionedCard

func newCardTypes() cardTypes {
	return make(cardTypes)
}

func (c cardTypes) addCard(card deck.PositionedCard) {
	c[card.Orientation] = append(c[card.Orientation], card)
}
