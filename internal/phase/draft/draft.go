package draft

import (
	"fmt"

	"github.com/jonathan-buttner/game-framework/internal/core"
	"github.com/jonathan-buttner/game-framework/internal/deck"
	"github.com/jonathan-buttner/game-framework/internal/phase"
	"github.com/jonathan-buttner/game-framework/internal/player"
)

/**
States:
Choose card
Trade/Complete goal/skip
Reduce resources

Next player (repeat choose card - reduce)

Switch hands
Repeat
*/

type Draft struct {
	phase.Phase
	deck             *deck.Deck
	startingHandSize int
}

func NewDraftPhase(phase phase.Phase, deck *deck.Deck, startingHandSize int) *Draft {
	return &Draft{Phase: phase, deck: deck, startingHandSize: startingHandSize}
}

func (d *Draft) Setup() {
	d.deck.Shuffle()

	// deal 5 cards to each player
	d.PlayerManager.ExecuteForPlayers(func(player *player.Player) error {
		d.deck.DealCards(d.startingHandSize, player)
		return nil
	})

	d.Step = chooseCardStep{d}
}

func (d *Draft) GetActions() []phase.Action {
	return d.Step.GetActions()
}

func (d *Draft) PerformAction(action phase.Action) {
	err := action.Execute(d.GameState)
	if err != nil {
		panic(fmt.Errorf("executing action failed err: %v", err))
	}

	// if d.PlayerManager.CurrentPlayer().ResourceCountExceedsLimit() {
	// 	d.step = phase.NewReduceResourcesStep(d.PlayerManager.CurrentPlayer())
	// } else {
	// 	d.PlayerManager.NextPlayer()
	// 	d.step = chooseCardStep{d}
	// }
}

func (d *Draft) GoToNextPlayer() {

}

func (d *Draft) createPhaseWithTurnHandler() *phase.PhaseWithTurnHandler {
	return &phase.PhaseWithTurnHandler{PlayerTurnHandler: d, Phase: d.Phase}
}

type chooseCardStep struct {
	draft *Draft
}

func (d chooseCardStep) GetActions() []phase.Action {
	cardsInHandWithPositions := d.draft.PlayerManager.CurrentPlayer().GetHand().AllPositionCombinations()
	validCards := d.draft.PlayerManager.CurrentPlayer().ValidOrientations(cardsInHandWithPositions)

	var actions []phase.Action
	for _, card := range validCards {
		actions = append(actions, &chooseCardAction{draft: d.draft, card: card})
	}

	return actions
}

type chooseCardAction struct {
	draft *Draft
	card  deck.PositionedCard
}

func (d *chooseCardAction) Execute(gameState *core.GameState) error {
	err := d.draft.PlayerManager.CurrentPlayer().PlayCardFromHand(d.card.ID(), d.card.Orientation, gameState)
	d.draft.Step = useResourcesStep{d.draft}

	return err
}

type useResourcesStep struct {
	draft *Draft
}

func (u useResourcesStep) GetActions() []phase.Action {
	return nil
}

type useResourcesAction struct {
	draft *Draft
	card  deck.PositionedCard
}

func (u *useResourcesAction) Execute(gameState *core.GameState) error {
	return nil
}

type skipUseResourcesAction struct {
	draft *Draft
}

func (s *skipUseResourcesAction) Execute(gameState *core.GameState) error {
	if s.draft.PlayerManager.CurrentPlayer().ResourceCountExceedsLimit() {
		s.draft.Step = phase.NewReduceResourcesStep(s.draft.createPhaseWithTurnHandler())
	} else {
		// switch to next player and go back to choose card step
	}

	return nil
}
