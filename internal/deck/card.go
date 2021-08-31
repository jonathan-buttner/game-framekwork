package deck

import "github.com/jonathan-buttner/game-framework/internal/core"

//go:generate mockgen -destination=../../mocks/mock_card.go -package=mocks github.com/jonathan-buttner/game-framework/internal/deck Card

type Card interface {
	ID() string
	IsOrientationValid(orientation CardOrientation) bool
	GetOrientationAction(orientation CardOrientation) CardAction
}

type CardAction interface {
	PerformEndRoundAction(game core.Game)
	PerformEndTurnAction(game core.Game)
	PerformPlayToTableaAction(game core.Game)
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
