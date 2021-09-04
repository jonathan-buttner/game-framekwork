package resource_test

import (
	"testing"

	"github.com/jonathan-buttner/game-framework/internal/resource"
	"github.com/stretchr/testify/assert"
)

func TestAddingEmptyArrayResultsInZeroValue(t *testing.T) {
	resourceHandler := resource.NewResourceHandler()

	resourceHandler.AddResources([]resource.Resource{})
	assert.Equal(t, resourceHandler.Total, 0)
}

func TestTwoBrownsEquals8VAlue(t *testing.T) {
	resourceHandler := resource.NewResourceHandler()

	resourceHandler.AddResources([]resource.Resource{{ResourceType: resource.Brown, Count: 2}})
	assert.Equal(t, resourceHandler.Total, 8)
}

func TestTwoGreensEquals6VAlue(t *testing.T) {
	resourceHandler := resource.NewResourceHandler()

	resourceHandler.AddResources([]resource.Resource{{ResourceType: resource.Green, Count: 2}})
	assert.Equal(t, resourceHandler.Total, 6)
}

func TestCallingAddMultipleTimesIncreaseTotal(t *testing.T) {
	resourceHandler := resource.NewResourceHandler()

	resourceHandler.AddResources([]resource.Resource{{ResourceType: resource.Green, Count: 1}})
	resourceHandler.AddResources([]resource.Resource{{ResourceType: resource.Yellow, Count: 1}})

	assert.Equal(t, resourceHandler.Total, 4)
}

func TestResourcesAreAccumulatedByCategory(t *testing.T) {
	resourceHandler := resource.NewResourceHandler()

	resourceHandler.AddResources([]resource.Resource{{ResourceType: resource.Green, Count: 1}, {ResourceType: resource.Yellow, Count: 1}})
	resourceHandler.AddResources([]resource.Resource{{ResourceType: resource.Yellow, Count: 1}})

	assert.Equal(t, len(resourceHandler.Resources), 2)
	assert.Equal(t, resourceHandler.Resources[resource.Yellow], 2)
	assert.Equal(t, resourceHandler.Resources[resource.Green], 1)
}

func TestCount(t *testing.T) {
	resourceHandler := resource.NewResourceHandler()

	resourceHandler.AddResources([]resource.Resource{{ResourceType: resource.Green, Count: 1}, {ResourceType: resource.Yellow, Count: 1}})
	resourceHandler.AddResources([]resource.Resource{{ResourceType: resource.Yellow, Count: 1}})

	assert.Equal(t, resourceHandler.Count, 3)
}

func TestHasResourcesMultipleBrown(t *testing.T) {
	resourceHandler := resource.NewResourceHandler()

	resourceHandler.AddResources([]resource.Resource{{ResourceType: resource.Green, Count: 1}, {ResourceType: resource.Brown, Count: 2}})
	assert.True(t, resourceHandler.HasResources(map[resource.ResourceType]int{resource.Brown: 2}))
}

func TestDoesNotHaveResourcesMultipleBrown(t *testing.T) {
	resourceHandler := resource.NewResourceHandler()

	resourceHandler.AddResources([]resource.Resource{{ResourceType: resource.Green, Count: 1}, {ResourceType: resource.Brown, Count: 2}})
	assert.False(t, resourceHandler.HasResources(map[resource.ResourceType]int{resource.Brown: 3}))
}

func TestHasResourcesMultipleColors(t *testing.T) {
	resourceHandler := resource.NewResourceHandler()

	resourceHandler.AddResources([]resource.Resource{{ResourceType: resource.Green, Count: 1}, {ResourceType: resource.Brown, Count: 2}})
	assert.True(t, resourceHandler.HasResources(map[resource.ResourceType]int{resource.Brown: 2, resource.Green: 1}))
}
