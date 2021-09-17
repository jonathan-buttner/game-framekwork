package mocks

import (
	"github.com/golang/mock/gomock"
	"github.com/jonathan-buttner/game-framework/internal/deck"
)

func NewPositionedMockCard(ctrl *gomock.Controller, orientation deck.CardOrientation, actions deck.OrientationActions) deck.PositionedCard {
	mockCard := NewMockCard(ctrl)
	mockCard.EXPECT().GetOrientationActions(gomock.Any()).Return(actions).AnyTimes()
	return deck.NewPositionedCard(mockCard, orientation)
}
