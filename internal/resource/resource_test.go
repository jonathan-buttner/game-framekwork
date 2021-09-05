package resource_test

import (
	"testing"

	"github.com/jonathan-buttner/game-framework/internal/resource"
	"github.com/stretchr/testify/assert"
)

func TestAddingEmptyArrayResultsInZeroValue(t *testing.T) {
	resourceHandler := resource.NewResourceHandler()

	resourceHandler.AddResources(resource.GroupedResources{})
	assert.Equal(t, resourceHandler.Total, 0)
}

func TestTwoBrownsEquals8VAlue(t *testing.T) {
	resourceHandler := resource.NewResourceHandler()

	resourceHandler.AddResources(resource.GroupedResources{resource.Brown: 2})
	assert.Equal(t, resourceHandler.Total, 8)
}

func TestTwoGreensEquals6VAlue(t *testing.T) {
	resourceHandler := resource.NewResourceHandler()

	resourceHandler.AddResources(resource.GroupedResources{resource.Green: 2})
	assert.Equal(t, resourceHandler.Total, 6)
}

func TestCallingAddMultipleTimesIncreaseTotal(t *testing.T) {
	resourceHandler := resource.NewResourceHandler()

	resourceHandler.AddResources(resource.GroupedResources{resource.Green: 1})
	resourceHandler.AddResources(resource.GroupedResources{resource.Yellow: 1})

	assert.Equal(t, resourceHandler.Total, 4)
}

func TestResourcesAreAccumulatedByCategory(t *testing.T) {
	resourceHandler := resource.NewResourceHandler()

	resourceHandler.AddResources(resource.GroupedResources{resource.Green: 1, resource.Yellow: 1})
	resourceHandler.AddResources(resource.GroupedResources{resource.Yellow: 1})

	assert.Equal(t, len(resourceHandler.Resources), 2)
	assert.Equal(t, resourceHandler.Resources[resource.Yellow], 2)
	assert.Equal(t, resourceHandler.Resources[resource.Green], 1)
}

func TestCount(t *testing.T) {
	resourceHandler := resource.NewResourceHandler()

	resourceHandler.AddResources(resource.GroupedResources{resource.Green: 1, resource.Yellow: 1})
	resourceHandler.AddResources(resource.GroupedResources{resource.Yellow: 1})

	assert.Equal(t, resourceHandler.Count, 3)
}

func TestHasResourcesMultipleBrown(t *testing.T) {
	resourceHandler := resource.NewResourceHandler()

	resourceHandler.AddResources(resource.GroupedResources{resource.Green: 1, resource.Brown: 2})
	assert.True(t, resourceHandler.HasResources(map[resource.ResourceType]int{resource.Brown: 2}))
}

func TestDoesNotHaveResourcesMultipleBrown(t *testing.T) {
	resourceHandler := resource.NewResourceHandler()

	resourceHandler.AddResources(resource.GroupedResources{resource.Green: 1, resource.Brown: 2})
	assert.False(t, resourceHandler.HasResources(map[resource.ResourceType]int{resource.Brown: 3}))
}

func TestHasResourcesMultipleColors(t *testing.T) {
	resourceHandler := resource.NewResourceHandler()

	resourceHandler.AddResources(resource.GroupedResources{resource.Green: 1, resource.Brown: 2})
	assert.True(t, resourceHandler.HasResources(map[resource.ResourceType]int{resource.Brown: 2, resource.Green: 1}))
}

func TestRemoveResourcesDoesNotHaveRequired(t *testing.T) {
	resourceHandler := resource.NewResourceHandler()

	assert.NotNil(t, resourceHandler.RemoveResources(map[resource.ResourceType]int{resource.Brown: 1}))
}

func TestRemoveResources(t *testing.T) {
	resourceHandler := resource.NewResourceHandler()

	resourceHandler.AddResources(resource.GroupedResources{resource.Green: 2})
	assert.Nil(t, resourceHandler.RemoveResources(resource.GroupedResources{resource.Green: 1}))
	assert.True(t, resourceHandler.HasResources(resource.GroupedResources{resource.Green: 1}))
	assert.Equal(t, resourceHandler.Count, 1)
	assert.Equal(t, resourceHandler.Total, 3)
}

func TestRemoveResourcesKeyNoLongerExists(t *testing.T) {
	resourceHandler := resource.NewResourceHandler()

	resourceHandler.AddResources(resource.GroupedResources{resource.Green: 1})
	assert.Nil(t, resourceHandler.RemoveResources(resource.GroupedResources{resource.Green: 1}))

	_, hasGreen := resourceHandler.Resources[resource.Green]
	assert.False(t, hasGreen)
	assert.Equal(t, resourceHandler.Count, 0)
	assert.Equal(t, resourceHandler.Total, 0)
}
