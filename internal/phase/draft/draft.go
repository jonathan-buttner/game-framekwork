package draft

import (
	"fmt"

	"github.com/jonathan-buttner/game-framework/internal/core"
	"github.com/jonathan-buttner/game-framework/internal/deck"
	"github.com/jonathan-buttner/game-framework/internal/phase"
)

//go:generate mockgen -destination=../../../mocks/mock_draft.go -package=mocks github.com/jonathan-buttner/game-framework/internal/phase/draft Deck

/**
States:
Choose card
Trade/Complete goal/skip potentially pull trade out into its own step because the player must pay at least a yellow?
Reduce resources

Next player (repeat choose card - reduce)

Switch hands
Repeat
*/

type Deck interface {
	Shuffle()
	DealCards(handSize int, player deck.Player)
}

type Draft struct {
	phase.Phase
	deck             Deck
	startingHandSize int
}

func NewDraftPhase(phase phase.Phase, deck Deck, startingHandSize int) *Draft {
	draftPhase := &Draft{Phase: phase, deck: deck, startingHandSize: startingHandSize}
	draftPhase.setup()

	return draftPhase
}

func (d *Draft) setup() {
	d.deck.Shuffle()

	// deal 5 cards to each player
	d.PlayerManager.ExecuteForPlayers(func(state phase.ExecutionState) error {
		d.deck.DealCards(d.startingHandSize, state.CurrentPlayer)
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
}

func (d *Draft) NextPlayer() {
	if d.PlayerManager.HasMorePlayers() {
		d.PlayerManager.NextPlayer()
	} else {
		d.rotateHands()
		d.PlayerManager.RotateStartPlayer()
	}

	d.Step = chooseCardStep{d}
}

func (d *Draft) rotateHands() {
	var hand []deck.Card

	d.PlayerManager.ExecuteForPlayers(func(state phase.ExecutionState) error {
		if hand != nil {
			tempHand := state.CurrentPlayer.GetHand()
			state.CurrentPlayer.SetHand(hand)
			hand = tempHand
		} else {
			hand = state.CurrentPlayer.GetHand()
			state.CurrentPlayer.SetHand(state.PrevPlayer.GetHand())
		}
		return nil
	})
}

type chooseCardStep struct {
	draft *Draft
}

func (c chooseCardStep) GetActions() []phase.Action {
	cardsInHandWithPositions := c.draft.PlayerManager.CurrentPlayer().GetHand().AllPositionCombinations()

	var actions []phase.Action
	for _, cardWithPosition := range cardsInHandWithPositions {
		if c.isChoosableCard(cardWithPosition) {
			actions = append(actions, &chooseCardAction{draft: c.draft, card: cardWithPosition})
		}
	}

	return actions
}

func (c chooseCardStep) isChoosableCard(card deck.PositionedCard) bool {
	return c.draft.PlayerManager.CurrentPlayer().ResourceHandler.HasResources(card.Cost()) &&
		card.IsOrientationValid(c.draft.GameState)
}

type chooseCardAction struct {
	draft *Draft
	card  deck.PositionedCard
}

func (c *chooseCardAction) Execute(gameState *core.GameState) error {
	err := c.draft.PlayerManager.CurrentPlayer().PlayCardFromHand(c.card.ID(), c.card.Orientation, gameState)
	c.draft.SetStep(phase.UseResourcesStep{Phase: c.draft})

	return err
}
