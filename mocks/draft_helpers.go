package mocks

import (
	"github.com/golang/mock/gomock"
	"github.com/jonathan-buttner/game-framework/internal/deck"
	"github.com/jonathan-buttner/game-framework/internal/phase/draft"
)

func NewMockDeckWithoutShuffle(ctrl *gomock.Controller, cards []deck.Card) draft.Deck {
	mockDeck := NewMockDeck(ctrl)
	concreteDeck := deck.NewDeck(cards)

	mockDeck.EXPECT().Shuffle().AnyTimes()
	mockDeck.EXPECT().DealCards(gomock.Any(), gomock.Any()).Do(func(handSize int, player deck.Player) {
		concreteDeck.DealCards(handSize, player)
	}).AnyTimes()

	return mockDeck
}
