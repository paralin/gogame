package gogame

import "fmt"

type Component interface {
	// Initialize on an entity. No guerantee of execution order.
	Init(ent *Entity)

	// Initialize on an entity, with init data.
	InitWithData(ent *Entity, data []byte)

	// Initialize frontend component
	InitFrontend(fe FrontendComponent)

	// Called after all components have been Init() or InitWithData()
	InitLate()

	// Check if Update() tick will be required.
	ShouldUpdate() bool

	// Update tick func
	Update()

	// Return the metadata for this component
	Meta() ComponentMeta

	// Build initial snapshot data for component creation
	InitData() []byte

	// Destroy the component. Frontend component is automatically cleaned before this is called.
	Destroy()
}

/*
 * To initial network representation
 */
func ComponentToNetworkInit(comp Component) *NetComponent {
	meta := comp.Meta()
	res := &NetComponent{
		Id:       meta.Id,
		InitData: comp.InitData(),
	}
	return res
}

func (gr *Game) ComponentFromId(id uint32) (Component, error) {
	// Find a factory for this component
	fact, ok := gr.ComponentTable[id]
	if !ok {
		return nil, fmt.Errorf("Unable to find component factory for %d.", id)
	}
	return fact.New(), nil
}
