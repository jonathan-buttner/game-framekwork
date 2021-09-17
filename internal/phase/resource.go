package phase

import (
	"github.com/jonathan-buttner/game-framework/internal/core"
	"github.com/jonathan-buttner/game-framework/internal/deck"
	"github.com/jonathan-buttner/game-framework/internal/resource"
)

type ReduceResourcesStep struct {
	Phase PhaseHandler
}

func NewReduceResourcesStep(phase PhaseHandler) ReduceResourcesStep {
	return ReduceResourcesStep{phase}
}

func (r ReduceResourcesStep) GetActions() []Action {
	var actions []Action

	for resType, count := range r.Phase.CurrentPlayer().ResourceHandler.Resources {
		if count > 0 {
			actions = append(actions, &reduceResourcesAction{phase: r.Phase, resourceType: resType})
		}
	}

	return actions
}

type reduceResourcesAction struct {
	phase        PhaseHandler
	resourceType resource.ResourceType
}

func (r *reduceResourcesAction) Execute(gameState *core.GameState) error {
	err := r.phase.CurrentPlayer().ResourceHandler.RemoveResources(resource.GroupedResources{r.resourceType: 1})

	if r.phase.CurrentPlayer().ResourceCountExceedsLimit() {
		r.phase.SetStep(NewReduceResourcesStep(r.phase))
	} else {
		r.phase.NextPlayer()
	}

	return err
}

type UseResourcesStep struct {
	Phase PhaseHandler
}

func (u UseResourcesStep) GetActions() []Action {
	victoryCards := u.Phase.CurrentPlayer().CardsByOrientation[deck.VictoryPoints]
	tradeCards := u.Phase.CurrentPlayer().CardsByOrientation[deck.Trade]

	actions := u.createActionsFromPlayableCards(victoryCards)

	// TODO: check if the user has enough resources to pay the initial trade cost
	actions = append(actions, u.createActionsFromPlayableCards(tradeCards)...)
	actions = append(actions, &skipUseResourcesAction{u.Phase})

	return actions
}

func (u UseResourcesStep) createActionsFromPlayableCards(cards []deck.PositionedCard) []Action {
	var actions []Action
	for _, card := range cards {
		if u.Phase.CurrentPlayer().ResourceHandler.HasResources(card.Cost()) {
			actions = append(actions, &useResourcesAction{phase: u.Phase, card: card})
		}
	}

	return actions
}

type useResourcesAction struct {
	phase PhaseHandler
	card  deck.PositionedCard
}

func (u *useResourcesAction) Execute(gameState *core.GameState) error {
	err := u.phase.CurrentPlayer().ResourceHandler.RemoveResources(u.card.Cost())
	if err != nil {
		return err
	}

	u.card.PerformUseResourceAction(gameState)

	// stay in the use resource step in case the player wants to use more of their resource to complete goals
	// or perform trades
	return nil
}

type skipUseResourcesAction struct {
	phase PhaseHandler
}

func (s *skipUseResourcesAction) Execute(gameState *core.GameState) error {
	if s.phase.CurrentPlayer().ResourceCountExceedsLimit() {
		s.phase.SetStep(NewReduceResourcesStep(s.phase))
	} else {
		s.phase.NextPlayer()
	}

	return nil
}
