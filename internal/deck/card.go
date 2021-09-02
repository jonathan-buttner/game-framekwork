package deck

//go:generate mockgen -destination=../../mocks/mock_card.go -package=mocks github.com/jonathan-buttner/game-framework/internal/deck Card

type Card interface {
	ID() string
	IsOrientationValid(orientation CardOrientation) bool
	GetOrientationAction(orientation CardOrientation) CardAction
}

type CardAction interface {
	PerformEndRoundAction(game Game)
	PerformEndTurnAction(game Game)
	PerformPlayToTableaAction(game Game)
}

// TODO: move to a central location?
type Game interface {
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
