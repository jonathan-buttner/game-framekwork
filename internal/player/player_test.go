package player_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/jonathan-buttner/game-framework/internal/deck"
	"github.com/jonathan-buttner/game-framework/internal/player"
	"github.com/jonathan-buttner/game-framework/internal/resource"
	"github.com/jonathan-buttner/game-framework/internal/rules"
	"github.com/jonathan-buttner/game-framework/mocks"
	"github.com/stretchr/testify/assert"
)

func TestValidOrientationsFiltersOutCard(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	needs2BrownsCard := mocks.NewMockCard(ctrl)
	needs2BrownsCard.EXPECT().GetOrientationAction(gomock.Any()).Return(nil)
	needs2BrownsCard.EXPECT().Cost().Return(map[resource.ResourceType]int{resource.Brown: 2})

	card := deck.NewPositionedCard(needs2BrownsCard, deck.VictoryPoints)

	player := player.NewPlayer("player1", rules.NewDefaultGameRules())
	assert.Len(t, player.ValidOrientations([]deck.PositionedCard{card}), 0)
}

func TestValidOrientationsIncludesCard(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	needs2BrownsCard := mocks.NewMockCard(ctrl)
	needs2BrownsCard.EXPECT().GetOrientationAction(gomock.Any()).Return(nil)
	needs2BrownsCard.EXPECT().Cost().Return(map[resource.ResourceType]int{resource.Brown: 2})
	needs2BrownsCard.EXPECT().ID().Return("needs2Browns")

	card := deck.NewPositionedCard(needs2BrownsCard, deck.VictoryPoints)

	player := player.NewPlayer("player1", rules.NewDefaultGameRules())
	player.ResourceHandler.AddResources(resource.GroupedResources{resource.Brown: 3})

	validCards := player.ValidOrientations([]deck.PositionedCard{card})
	assert.Len(t, validCards, 1)
	assert.Equal(t, validCards[0].ID(), "needs2Browns")
}
