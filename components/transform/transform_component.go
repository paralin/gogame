package transform

import (
	"github.com/fuserobotics/gogame"
	"github.com/golang/protobuf/proto"
)

var TransformComponentMeta gogame.ComponentMeta = gogame.ComponentMeta{
	Id: 1,
}

type TransformComponent struct {
	Entity *gogame.Entity
	Data   TransformData
}

func (tc *TransformComponent) Init(ent *gogame.Entity) {
	tc.Entity = ent
}

func (tc *TransformComponent) InitWithData(ent *gogame.Entity, data []byte) {
	tc.Entity = ent

	// parse data, handle error here somehow
	proto.Unmarshal(data, &tc.Data)
}

func (tc *TransformComponent) InitLate() {
}

func (tc *TransformComponent) Meta() gogame.ComponentMeta {
	return TransformComponentMeta
}

func (tc *TransformComponent) InitData() []byte {
	data, _ := proto.Marshal(&tc.Data)
	return data
}

// Assert at compile time the component is valid
// This line will fail otherwise.
var componentAssertion gogame.Component = &TransformComponent{}
