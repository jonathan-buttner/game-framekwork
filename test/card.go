package test

import (
	"github.com/golang/mock/gomock"
	"github.com/jonathan-buttner/game-framework/internal/deck"
	"github.com/jonathan-buttner/game-framework/test/mocks"
)

func NewPositionedMockCard(ctrl *gomock.Controller, orientation deck.CardOrientation, actions deck.OrientationActions) deck.PositionedCard {
	mockCard := mocks.NewMockCard(ctrl)
	mockCard.EXPECT().GetOrientationActions(gomock.Any()).Return(actions).AnyTimes()
	return deck.NewPositionedCard(mockCard, orientation)
}
