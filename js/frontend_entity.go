package js

import (
	"github.com/fuserobotics/gogame"
	"github.com/gopherjs/gopherjs/js"
)

type JsFrontendEntity struct {
	*js.Object
}

func (fe *JsFrontendEntity) Init() {
	if fe.Object.Get("init") != js.Undefined {
		fe.Object.Call("init")
	}
}

func (je *JsFrontendEntity) Call(fn string, args ...interface{}) interface{} {
	return je.Object.Call(fn, args)
}

func (fe *JsFrontendEntity) AddComponent(id uint32) gogame.FrontendComponent {
	if fe.Object.Get("addComponent") == js.Undefined {
		return nil
	}
	res := fe.Object.Call("addComponent", id)
	if res == nil || res == js.Undefined {
		return nil
	}
	return &JsFrontendComponent{Object: res}
}

func (fe *JsFrontendEntity) InitLate() {
	if fe.Object.Get("initLate") == js.Undefined {
		return
	}
	fe.Object.Call("initLate")
}

func (fe *JsFrontendEntity) Destroy() {
	if fe.Object.Get("destroy") == js.Undefined {
		return
	}
	fe.Object.Call("destroy")
}
