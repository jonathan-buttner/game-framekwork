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
Trade/Complete goal/skip potentially pull trade out into its own step because the player must pay at least a yellow?
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

func (d *Draft) NextPlayer() {

}

type chooseCardStep struct {
	draft *Draft
}

func (d chooseCardStep) GetActions() []phase.Action {
	cardsInHandWithPositions := d.draft.PlayerManager.CurrentPlayer().GetHand().AllPositionCombinations()

	var actions []phase.Action
	for _, cardWithPosition := range cardsInHandWithPositions {
		if d.draft.PlayerManager.CurrentPlayer().ResourceHandler.HasResources(cardWithPosition.Cost()) {
			actions = append(actions, &chooseCardAction{draft: d.draft, card: cardWithPosition})
		}
	}

	return actions
}

type chooseCardAction struct {
	draft *Draft
	card  deck.PositionedCard
}

func (d *chooseCardAction) Execute(gameState *core.GameState) error {
	err := d.draft.PlayerManager.CurrentPlayer().PlayCardFromHand(d.card.ID(), d.card.Orientation, gameState)
	d.draft.SetStep(phase.UseResourcesStep{Phase: d.draft})

	return err
}
