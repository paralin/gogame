package gogame

import (
	"errors"
	"fmt"
)

type EntityTable map[uint32]*Entity

func NewEntityTable() EntityTable {
	return make(map[uint32]*Entity)
}

/*
* Entity. See NetEntity for networked representation.
 */
type Entity struct {
	Id uint32

	// Parent entity.
	Parent *Entity

	// Children entities
	Children map[uint32]*Entity

	// Map of component ID to implementation instance
	Components map[uint32]Component

	// Tick components
	TickComponents map[uint32]Component

	// Frontend components
	FrontendComponents map[uint32]FrontendComponent

	// Frontend components that have an update function
	TickFrontendComponents map[uint32]FrontendComponent

	// Frontend entity if set
	FrontendEntity FrontendEntity

	// True if Update() does anything
	HasUpdateTick bool
}

// Construct a new entity
func NewEntity(id uint32) *Entity {
	return &Entity{
		Id:                     id,
		Children:               make(map[uint32]*Entity),
		Components:             make(map[uint32]Component),
		TickComponents:         make(map[uint32]Component),
		FrontendComponents:     make(map[uint32]FrontendComponent),
		TickFrontendComponents: make(map[uint32]FrontendComponent),
	}
}

func (ent *Entity) AddComponent(comp Component) {
	meta := comp.Meta()
	ent.Components[meta.Id] = comp
	if comp.ShouldUpdate() {
		ent.TickComponents[meta.Id] = comp
		ent.HasUpdateTick = true
	}
}

/* Calls Init() on all components. */
func (ent *Entity) InitComponents() {
	for _, comp := range ent.Components {
		comp.Init(ent)
	}
}

/* Calls InitLate() on all components and frontend components. */
func (ent *Entity) LateInitComponents() {
	for _, comp := range ent.Components {
		comp.InitLate()
	}
	for _, comp := range ent.FrontendComponents {
		comp.InitLate()
	}
}

/* Initializes the frontend entity, if one has been set. */
func (ent *Entity) InitFrontendEntity() {
	if ent.FrontendEntity == nil {
		return
	}
	for id, comp := range ent.Components {
		fe := ent.FrontendEntity.AddComponent(comp.Meta().Id)
		if fe != nil {
			fe.Init()
			ent.FrontendComponents[id] = fe
			comp.InitFrontend(fe)
		}
	}
}

func (ent *Entity) Update() {
	for _, comp := range ent.TickComponents {
		comp.Update()
	}
	for _, comp := range ent.TickFrontendComponents {
		comp.Update()
	}
}

// Destroy the entity
func (ent *Entity) Destroy() {
	for cid, comp := range ent.Components {
		fe, ok := ent.FrontendComponents[cid]
		if ok && fe != nil {
			fe.Destroy()
		}
		comp.Destroy()
	}
	if ent.FrontendEntity != nil {
		ent.FrontendEntity.Destroy()
	}
}

/*
 * Convert to networked representation.
 */
func (ent *Entity) ToNetworkInit() *NetEntity {
	components := make([]*NetComponent, len(ent.Components))
	i := 0
	for _, comp := range ent.Components {
		components[i] = ComponentToNetworkInit(comp)
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

// Adds child, before child is initialized
func (pent *Entity) AddChild(ent *Entity) {
	pent.Children[ent.Id] = ent
}

/*
 * Instantiate from network representation.
 */
func (gr *Game) EntityFromNetInit(ent *NetEntity) (*Entity, error) {
	entTable := gr.EntityTable
	res := NewEntity(ent.Id)

	if ent.ParentId != 0 {
		if entTable == nil {
			return nil, errors.New("Entity has parent, but no entity table given.")
		}
		parent, ok := entTable[ent.ParentId]
		if !ok {
			return nil, fmt.Errorf("Entity %d has unknown parent ID %d", ent.Id, ent.ParentId)
		}
		res.Parent = parent
		parent.AddChild(res)
	}

	for _, comp := range ent.Component {
		if _, ok := res.Components[comp.Id]; ok {
			return nil, fmt.Errorf("Entity %d has duplicate component %d.", res.Id, comp.Id)
		}
		rcomp, err := gr.ComponentFromId(comp.Id)
		if err != nil {
			return nil, err
		}
		res.Components[comp.Id] = rcomp
		rcomp.InitWithData(res, comp.InitData)
	}

	return res, nil
}
