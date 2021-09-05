package deck

import "github.com/jonathan-buttner/game-framework/internal/resource"

//go:generate mockgen -destination=../../mocks/mock_card.go -package=mocks github.com/jonathan-buttner/game-framework/internal/deck Card,CardAction

type Card interface {
	ID() string
	IsOrientationValid(orientation CardOrientation) bool
	GetOrientationAction(orientation CardOrientation) CardAction
	Cost() resource.GroupedResources
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

type CardAction interface {
	PerformEndRoundAction(game Game)
	PerformEndTurnAction(game Game)
	PerformPlayToTableaAction(game Game)
}

type PositionedCard struct {
	Card
	CardAction

	Orientation CardOrientation
}

func NewPositionedCard(card Card, orientation CardOrientation) PositionedCard {
	return PositionedCard{card, card.GetOrientationAction(orientation), orientation}
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
)

func (c CardOrientation) String() string {
	switch {
	case c == VictoryPoints:
		return "victory points"
	case c == Upgrade:
		return "upgrade"
	case c == Trade:
		return "trade"
	case c == Generate:
		return "generate"
	default:
		return "invalid"
	}
}

var Orientations = [...]CardOrientation{
	VictoryPoints,
	Upgrade,
	Trade,
	Generate,
}

type Game interface {
}
