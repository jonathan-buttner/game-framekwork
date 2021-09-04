package resource

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

func (r *ResourceHandler) AddResources(resources []Resource) {
	for _, resource := range resources {
		r.Total += (resource.Count * resource.Value())
		r.Resources[resource.ResourceType] += resource.Count
		r.Count += resource.Count
	}
}

func (r *ResourceHandler) HasResources(neededResources ResourceRequirement) bool {
	for requiredResourceType, requiredNumberResources := range neededResources {
		numberOfType, hasType := r.Resources[requiredResourceType]
		if !hasType || requiredNumberResources > numberOfType {
			return false
		}
	}

	return true
}

type ResourceRequirement map[ResourceType]int
