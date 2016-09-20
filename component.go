package gogame

type Component interface {
	// Initialize on an entity. No guerantee of execution order.
	Init(ent *Entity)

	// Return the metadata for this component
	Meta() ComponentMeta

	// Build initial snapshot data for component creation
	InitData() []byte
}

type ComponentFactory interface {
	// Return the metadata for this component
	Meta() ComponentMeta

	// Instantiate a new component
	New() Component
}

/*
 * To initial network representation
 */
func ComponentNetworkInit(comp Component) *NetComponent {
	meta := comp.Meta()
	res := &NetComponent{
		Id:       meta.Id,
		InitData: comp.InitData(),
	}
	return res
}
