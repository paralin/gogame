package gogame

type ComponentFactory interface {
	// Return the metadata for this component
	Meta() ComponentMeta

	// Instantiate a new component
	New() Component
}
