package deck_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/jonathan-buttner/game-framework/internal/deck"
	"github.com/jonathan-buttner/game-framework/mocks"
	"github.com/stretchr/testify/assert"
)

func TestSetsSinglePlayerWithOneCard(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	p := mocks.NewMockPlayer(ctrl)

	hand := make([]deck.Card, 1)
	p.EXPECT().SetHand(gomock.Any()).DoAndReturn(func(cards []deck.Card) {
		copy(hand, cards)
	})
	d := deck.NewDeck([]deck.Card{{Name: "hello"}})
	d.DealCards(1, []deck.Player{p})
	assert.Equal(t, []deck.Card{{Name: "hello"}}, hand)
}

func TestSetsMultiplePlayerWithOneCard(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	player1 := mocks.NewMockPlayer(ctrl)
	player1Hand := make([]deck.Card, 1)

	player1.EXPECT().SetHand(gomock.Any()).DoAndReturn(func(cards []deck.Card) {
		copy(player1Hand, cards)
	})

	player2 := mocks.NewMockPlayer(ctrl)
	player2Hand := make([]deck.Card, 1)

	player2.EXPECT().SetHand(gomock.Any()).DoAndReturn(func(cards []deck.Card) {
		copy(player2Hand, cards)
	})

	d := deck.NewDeck([]deck.Card{{Name: "hello"}, {Name: "card 2"}})
	d.DealCards(1, []deck.Player{player1, player2})
	assert.Equal(t, []deck.Card{{Name: "card 2"}}, player1Hand)
	assert.Equal(t, []deck.Card{{Name: "hello"}}, player2Hand)
}
