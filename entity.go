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

	// Map of component ID to implementation instance
	Components map[uint32]Component

	// Frontend entity if set
	FrontendEntity FrontendEntity
}

/* Call after creating a fresh entity (without init data). */
func (ent *Entity) InitComponents() {
	for _, comp := range ent.Components {
		comp.Init(ent)
	}
}

/* Call after creating any type of entity. */
func (ent *Entity) LateInitComponents() {
	for _, comp := range ent.Components {
		comp.InitLate()
	}
}

func (ent *Entity) InitFrontendEntity() {
	if ent.FrontendEntity == nil {
		return
	}
	for _, comp := range ent.Components {
		fe := ent.FrontendEntity.AddComponent(comp.Meta().Id)
		if fe != nil {
			comp.InitFrontend(fe)
		}
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

/*
 * Instantiate from network representation.
 */
func (gr *Game) EntityFromNetInit(ent *NetEntity) (*Entity, error) {
	entTable := gr.EntityTable
	res := &Entity{
		Id:         ent.Id,
		Components: make(map[uint32]Component),
	}

	if ent.ParentId != 0 {
		if entTable == nil {
			return nil, errors.New("Entity has parent, but no entity table given.")
		}
		parent, ok := entTable[ent.ParentId]
		if !ok {
			return nil, fmt.Errorf("Entity %d has unknown parent ID %d", ent.Id, ent.ParentId)
		}
		res.Parent = parent
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
		if len(comp.InitData) > 0 {
			rcomp.InitWithData(res, comp.InitData)
		} else {
			rcomp.Init(res)
		}
	}

	res.LateInitComponents()

	return res, nil
}
