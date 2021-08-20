package core

type player interface {
	setHand(cards []Card)
}

type stack interface {
	Size() int
	Pop() Card
}

type stackInternal struct {
	items []interface{}
}

func (s *stackInternal) Size() int {
	return len(s.items)
}

func (s *stackInternal) Pop() interface{} {
	if s.Size() <= 0 {
		panic("stack is empty during pop")
	}

	top := len(s.items) - 1
	item := s.items[top]

	// avoid a memory leak
	s.items[top] = nil

	// remove the from the stack
	s.items = s.items[:top]
	return item
}

func NewStack(items []interface{}) *stackInternal {
	return &stackInternal{
		items: items,
	}
}

type cardStack struct {
	*stackInternal
}

func (c *cardStack) Pop() Card {
	v := c.stackInternal.Pop().(Card)
	return v
}

func NewCardStack(cards []Card) stack {
	convertCards := make([]interface{}, len(cards))

	for i, v := range cards {
		convertCards[i] = v
	}

	return &cardStack{
		stackInternal: NewStack(convertCards),
	}
}

type Deck struct {
	cards stack
}

func (d Deck) dealCards(handSize int, players []player) {
	hands := make([][]Card, len(players))
	for i := 0; i < len(players); i++ {
		hands[i] = make([]Card, 0, handSize)
	}

	for i := 0; i < handSize; i++ {
		for playerIndex := 0; playerIndex < len(players); playerIndex++ {
			hands[playerIndex] = append(hands[playerIndex], d.cards.Pop())
		}
	}
}

func NewDeck(cards []Card) Deck {
	return Deck{
		cards: NewCardStack(cards),
	}
}
