package resource

import "fmt"

//go:generate stringer -type=ResourceType

type ResourceType int

const (
	Yellow ResourceType = iota
	Red
	Green
	Brown
)

func (r ResourceType) Value() int {
	switch {
	case r == Yellow:
		return 1
	case r == Red:
		return 2
	case r == Green:
		return 3
	case r == Brown:
		return 4
	default:
		return 0
	}
}

type Resource struct {
	ResourceType

	Count int
}

type ResourceHandler struct {
	Total     int
	Count     int
	Resources map[ResourceType]int
}

func NewResourceHandler() *ResourceHandler {
	return &ResourceHandler{Resources: make(map[ResourceType]int)}
}

func (r *ResourceHandler) RemoveResources(resources GroupedResources) error {
	if !r.HasResources(resources) {
		return fmt.Errorf("player does not have all of the requested resources %v", resources)
	}

	for resType, resCount := range resources {
		r.Total -= (resCount * resType.Value())
		r.Resources[resType] -= resCount
		r.Count -= resCount

		if r.Resources[resType] <= 0 {
			delete(r.Resources, resType)
		}
	}

	return nil
}

func (r *ResourceHandler) AddResources(resources GroupedResources) {
	for resType, count := range resources {
		r.Total += (count * resType.Value())
		r.Resources[resType] += count
		r.Count += count
	}
}

func (r *ResourceHandler) HasResources(neededResources GroupedResources) bool {
	for requiredResourceType, requiredNumberResources := range neededResources {
		numberOfType, hasType := r.Resources[requiredResourceType]
		if !hasType || requiredNumberResources > numberOfType {
			return false
		}
	}

	return true
}

type GroupedResources map[ResourceType]int
