package player

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
	Resources map[ResourceType]int
}

func NewResourceHandler() *ResourceHandler {
	return &ResourceHandler{Resources: make(map[ResourceType]int)}
}

func (r *ResourceHandler) AddResources(resources []Resource) {
	for _, resource := range resources {
		r.Total += (resource.Count * resource.Value())
		r.Resources[resource.ResourceType] += resource.Count
	}
}
