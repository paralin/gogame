package gogame

/*
* Entity. See NetEntity for networked representation.
 */
type Entity struct {
	Id uint32

	// Parent entity.
	Parent *Entity

	// Map of component ID to implementation instance
	Components map[uint32]Component
}

/*
 * Convert to networked representation.
 */
func (ent *Entity) ToNetworkInit() *NetEntity {
	components := make([]*NetComponent, len(ent.Components))
	i := 0
	for _, comp := range ent.Components {
		components[i] = ComponentNetworkInit(comp)
		i++
	}

	var parentId uint32
	if ent.Parent != nil && ent.Parent != ent {
		parentId = ent.Parent.Id
	}

	return &NetEntity{
		Id:        ent.Id,
		ParentId:  parentId,
		Component: components,
	}
}
