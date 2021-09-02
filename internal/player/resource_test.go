package player_test

import (
	"testing"

	"github.com/jonathan-buttner/game-framework/internal/player"
	"github.com/stretchr/testify/assert"
)

func TestAddingEmptyArrayResultsInZeroValue(t *testing.T) {
	resourceHandler := player.NewResourceHandler()

	resourceHandler.AddResources([]player.Resource{})
	assert.Equal(t, resourceHandler.Total, 0)
}

func TestTwoBrownsEquals8VAlue(t *testing.T) {
	resourceHandler := player.NewResourceHandler()

	resourceHandler.AddResources([]player.Resource{{ResourceType: player.Brown, Count: 2}})
	assert.Equal(t, resourceHandler.Total, 8)
}

func TestTwoGreensEquals6VAlue(t *testing.T) {
	resourceHandler := player.NewResourceHandler()

	resourceHandler.AddResources([]player.Resource{{ResourceType: player.Green, Count: 2}})
	assert.Equal(t, resourceHandler.Total, 6)
}

func TestCallingAddMultipleTimesIncreaseTotal(t *testing.T) {
	resourceHandler := player.NewResourceHandler()

	resourceHandler.AddResources([]player.Resource{{ResourceType: player.Green, Count: 1}})
	resourceHandler.AddResources([]player.Resource{{ResourceType: player.Yellow, Count: 1}})

	assert.Equal(t, resourceHandler.Total, 4)
}

func TestResourcesAreAccumulatedByCategory(t *testing.T) {
	resourceHandler := player.NewResourceHandler()

	resourceHandler.AddResources([]player.Resource{{ResourceType: player.Green, Count: 1}, {ResourceType: player.Yellow, Count: 1}})
	resourceHandler.AddResources([]player.Resource{{ResourceType: player.Yellow, Count: 1}})

	assert.Equal(t, len(resourceHandler.Resources), 2)
	assert.Equal(t, resourceHandler.Resources[player.Yellow], 2)
	assert.Equal(t, resourceHandler.Resources[player.Green], 1)
}
