package js

import (
	"github.com/fuserobotics/gogame"
	"github.com/gopherjs/gopherjs/js"
)

// Wraps a IFrontend in TypeScript
type JsFrontend struct {
	*js.Object
}

func (fe *JsFrontend) Init() gogame.FrontendGameRules {
	if fe.Object.Get("init") == nil {
		return nil
	}
	fegr := fe.Object.Call("init")
	if fegr != nil {
		return &JsFrontendGameRules{Object: fegr}
	}
	return nil
}

func (fe *JsFrontend) AddEntity(entity *gogame.Entity) gogame.FrontendEntity {
	if fe.Object.Get("addEntity") == nil {
		return nil
	}
	res := fe.Object.Call("addEntity", entity.ToNetworkInit())
	if res == nil || res == js.Undefined {
		return nil
	}
	return &JsFrontendEntity{Object: res}
}

func (fe *JsFrontend) Destroy() {
	if fe.Object.Get("destroy") == js.Undefined {
		return
	}
	fe.Object.Call("destroy")
}
