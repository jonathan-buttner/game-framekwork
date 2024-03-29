package deck

import (
	"fmt"

	"github.com/jonathan-buttner/game-framework/internal/resource"
)

//go:generate mockgen -destination=../../mocks/mock_card.go -package=mocks github.com/jonathan-buttner/game-framework/internal/deck Card,OrientationActions
//go:generate stringer -type=CardOrientation

// TODO: rename
type Card interface {
	ID() string
	GetOrientationActions(orientation CardOrientation) OrientationActions
}

type Cards []Card

func (c Cards) AllPositionCombinations() []PositionedCard {
	var allPositions []PositionedCard
	for _, card := range c {
		for _, orientation := range Orientations {
			allPositions = append(allPositions, NewPositionedCard(card, orientation))
		}
	}

	return allPositions
}

// TODO: rename
type OrientationActions interface {
	// TODO: should the perform action calls return errors?

	// TODO: this should check that the user performing the action can pay a resource if this is the first
	// time doing a convert resource card
	PerformUseResourceAction(game Game)
	PerformPlayToTableaAction(game Game)
	// Cost to execute the card's action
	UseCost() resource.GroupedResources
	// Cost to put the card in your tablea
	AcquireCost() resource.GroupedResources
	IsOrientationValid(game Game) bool
}

type PositionedCard struct {
	Card
	OrientationActions

	Orientation CardOrientation
}

func NewPositionedCard(card Card, orientation CardOrientation) PositionedCard {
	return PositionedCard{card, card.GetOrientationActions(orientation), orientation}
}

func (p PositionedCard) String() string {
	return fmt.Sprintf("%v[%v]", p.ID(), p.Orientation.String())
}

type NamedCard struct {
	Name string
}

func (c NamedCard) ID() string {
	return c.Name
}

type CardOrientation int

const (
	VictoryPoints CardOrientation = iota
	Upgrade
	Trade
	Generate
	Payment
)

var Orientations = [...]CardOrientation{
	VictoryPoints,
	Upgrade,
	Trade,
	Generate,
	Payment,
}

type Game interface{}
