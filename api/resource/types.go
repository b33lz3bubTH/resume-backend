package handler

// ResourceType represents the type of resource being accessed
type ResourceType string

const (
	ResourceTypeBootcamp ResourceType = "bootcamps"
	ResourceTypeJournal  ResourceType = "journal"
	ResourceTypeMeme     ResourceType = "memes"
	ResourceTypeCategory ResourceType = "categories"
)

// isValidResourceType checks if the resource type is valid
func isValidResourceType(rt ResourceType) bool {
	return rt == ResourceTypeBootcamp || rt == ResourceTypeJournal ||
		rt == ResourceTypeMeme || rt == ResourceTypeCategory
}

