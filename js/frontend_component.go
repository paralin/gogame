package js

import (
	"github.com/gopherjs/gopherjs/js"
)

type JsFrontendComponent struct {
	*js.Object
	hasUpdateFunc bool
}

func (fc *JsFrontendComponent) Init() bool {
	defer func() {
		if fc.Object.Get("init") != js.Undefined {
			fc.Object.Call("init")
		}
	}()
	fc.hasUpdateFunc = fc.Object.Get("update") != js.Undefined
	return fc.hasUpdateFunc
}

func (fc *JsFrontendComponent) Call(fn string, args ...interface{}) interface{} {
	return fc.Object.Call(fn, args)
}

func (fc *JsFrontendComponent) InitLate() {
	if fc.Object.Get("initLate") != js.Undefined {
		fc.Object.Call("initLate")
	}
}

func (fc *JsFrontendComponent) Update() {
	if !fc.hasUpdateFunc {
		return
	}
	fc.Object.Call("update")
}

func (fc *JsFrontendComponent) Destroy() {
	if fc.Object.Get("destroy") == js.Undefined {
		return
	}
	fc.Object.Call("destroy")
}
