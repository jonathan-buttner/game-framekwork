package rules

const resourceLimit = 10

type GameRules struct {
	ResourceLimit int
}

func NewResourceLimit(limit int) GameRules {
	return GameRules{ResourceLimit: limit}
}

func NewDefaultGameRules() GameRules {
	return GameRules{ResourceLimit: resourceLimit}
}
