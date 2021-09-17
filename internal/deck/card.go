package deck

import "github.com/jonathan-buttner/game-framework/internal/resource"

//go:generate mockgen -destination=../../mocks/mock_card.go -package=mocks github.com/jonathan-buttner/game-framework/internal/deck Card,OrientationActions

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
	Cost() resource.GroupedResources
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
	case c == Payment:
		return "payment"
	default:
		return "invalid"
	}
}

var Orientations = [...]CardOrientation{
	VictoryPoints,
	Upgrade,
	Trade,
	Generate,
	Payment,
}

type Game interface{}
