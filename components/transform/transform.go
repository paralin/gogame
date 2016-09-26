package transform

import (
	"github.com/fuserobotics/gogame"
	"github.com/golang/protobuf/proto"
)

var TransformComponentMeta gogame.ComponentMeta = gogame.ComponentMeta{
	Id: 1,
}

type TransformComponent struct {
	Entity   *gogame.Entity
	Data     TransformData
	Frontend gogame.FrontendComponent
}

// Initialize a brand new transform component
func (tc *TransformComponent) Init(ent *gogame.Entity) {
	tc.Entity = ent
	tc.Data.Position = &TransformPosition{
		X: 0,
		Y: 0,
	}
	tc.Data.Scale = &TransformScale{
		XScale: 0.2,
		YScale: 0.2,
	}
}

// Initialize a remotely created transform component, over the network
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
	tc.Frontend = fe
	tc.SyncPosition()
}

// Tells the frontend to update position.
func (tc *TransformComponent) SyncPosition() {
	if tc.Frontend == nil {
		return
	}
	tc.Frontend.Call("setPosition", &tc.Data)
}

func (tc *TransformComponent) Destroy() {
}

// Factory to spawn transform components
type transformComponentFactory struct {
}

func (tff *transformComponentFactory) Meta() gogame.ComponentMeta {
	return TransformComponentMeta
}

func (tff *transformComponentFactory) New() gogame.Component {
	return &TransformComponent{}
}

// Assert at compile time the component is valid
// This line will fail otherwise.
var TransformComponentFactory gogame.ComponentFactory = &transformComponentFactory{}
