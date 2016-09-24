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

func (tc *TransformComponent) ShouldUpdate() bool {
	return false
}

func (tc *TransformComponent) Update() {
}

func (tc *TransformComponent) Meta() gogame.ComponentMeta {
	return TransformComponentMeta
}

func (tc *TransformComponent) InitData() []byte {
	data, _ := proto.Marshal(&tc.Data)
	return data
}

func (tc *TransformComponent) InitFrontend(fe gogame.FrontendComponent) {
}

func (tc *TransformComponent) Destroy() {
}

type TransformComponentFactory struct {
}

func (tff *TransformComponentFactory) Meta() gogame.ComponentMeta {
	return TransformComponentMeta
}

func (tff *TransformComponentFactory) New() gogame.Component {
	return &TransformComponent{}
}

// Assert at compile time the component is valid
// This line will fail otherwise.
var componentAssertion gogame.Component = &TransformComponent{}
