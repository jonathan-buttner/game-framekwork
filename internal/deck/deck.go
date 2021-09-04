package deck

import (
	"log"
	"math/rand"
	"time"
)

//go:generate mockgen -destination=../../mocks/mock_deck.go -package=mocks github.com/jonathan-buttner/game-framework/internal/deck Player

type Player interface {
	SetHand(cards []Card)
	GetHand() Cards
}

type stackInternal struct {
	items []interface{}
}

func (s *stackInternal) Size() int {
	return len(s.items)
}

func (s *stackInternal) Pop(numItems int) []interface{} {
	if s.Size() < numItems {
		log.Fatalf("requested more items: %v to be removed than were available: %v", numItems, s.Size())
	}

	var removedItems []interface{}
	for i := 0; i < numItems; i++ {
		top := len(s.items) - 1
		removedItems = append(removedItems, s.items[top])

		// avoid a memory leak
		s.items[top] = nil

		// remove the from the stack
		s.items = s.items[:top]
	}

	return removedItems
}

func newStack(items []interface{}) *stackInternal {
	return &stackInternal{
		items: items,
	}
}

type cardStack struct {
	*stackInternal
}

func (c *cardStack) Pop(numItems int) []Card {
	convertCards := make([]Card, numItems)
	items := c.stackInternal.Pop(numItems)

	for i, v := range items {
		convertCards[i] = v.(Card)
	}
	return convertCards
}

func (c *cardStack) Shuffle() {
	rand.Shuffle(c.Size(), func(i int, j int) {
		c.items[i], c.items[j] = c.items[j], c.items[i]
	})
}

func newCardStack(cards []Card) *cardStack {
	// maybe move this out?
	rand.Seed(time.Now().UnixNano())

	convertCards := make([]interface{}, len(cards))

	for i, v := range cards {
		convertCards[i] = v
	}

	return &cardStack{
		stackInternal: newStack(convertCards),
	}
}

type Deck struct {
	cards *cardStack
}

func (d *Deck) Shuffle() {
	d.cards.Shuffle()
}

func (d *Deck) DealCards(handSize int, player Player) {
	player.SetHand(d.cards.Pop(handSize))
}

func (d *Deck) Size() int {
	return d.cards.Size()
}

func NewDeck(cards []Card) *Deck {
	return &Deck{
		cards: newCardStack(cards),
	}
}
