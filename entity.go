package gogame

/*
* Entity. See NetEntity for networked representation.
 */
type Entity struct {
	Id uint32

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

	return &NetEntity{
		Id:        ent.Id,
		Component: components,
	}
}
