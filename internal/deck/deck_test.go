package deck_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/jonathan-buttner/game-framework/internal/deck"
	"github.com/jonathan-buttner/game-framework/test/mocks"
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
	cards := []deck.Card{mocks.NewMockCard(ctrl)}

	d := deck.NewDeck(cards)
	d.DealCards(1, p)
	assert.Equal(t, cards, hand)
}

func TestEmptiesDeck(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	p := mocks.NewMockPlayer(ctrl)

	// hand := make([]deck.Card, 1)
	// p.EXPECT().SetHand(gomock.Any()).DoAndReturn(func(cards []deck.Card) {
	// 	copy(hand, cards)
	// })

	p.EXPECT().SetHand(gomock.Any())
	cards := []deck.Card{mocks.NewMockCard(ctrl)}

	d := deck.NewDeck(cards)
	d.DealCards(1, p)
	assert.Equal(t, d.Size(), 0)
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

	cardHello := mocks.NewMockCard(ctrl)

	cardHi := mocks.NewMockCard(ctrl)

	d := deck.NewDeck([]deck.Card{cardHello, cardHi})
	d.DealCards(1, player1)
	d.DealCards(1, player2)
	assert.Equal(t, []deck.Card{cardHi}, player1Hand)
	assert.Equal(t, []deck.Card{cardHello}, player2Hand)
}
