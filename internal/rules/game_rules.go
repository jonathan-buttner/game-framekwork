package rules

import "github.com/jonathan-buttner/game-framework/internal/resource"

const (
	resourceLimit = 10
)

// TODO: determine if this should be up'ed to a yellow
var initialTradeCost = resource.GroupedResources{}

type GameRules struct {
	ResourceLimit    int
	InitialTradeCost resource.GroupedResources
}

// TODO: allow trade cost to be passed in here
func NewResourceLimit(limit int) GameRules {
	return GameRules{ResourceLimit: limit}
}

func NewDefaultGameRules() GameRules {
	return GameRules{ResourceLimit: resourceLimit, InitialTradeCost: initialTradeCost}
}
