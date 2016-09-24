package js

import (
	"github.com/fuserobotics/gogame"
	"github.com/gopherjs/gopherjs/js"
)

// Wraps a IFrontendGameRules in TypeScript
type JsFrontendGameRules struct {
	*js.Object
}

func (fe *JsFrontendGameRules) Init() {
	if fe.Object.Get("init") == js.Undefined {
		return
	}
	fe.Object.Call("init")
}

func (fe *JsFrontendGameRules) SetGameOperatingMode(opMode gogame.GameOperatingMode) {
	if fe.Object.Get("setGameOperatingMode") == js.Undefined {
		return
	}
	fe.Object.Call("setGameOperatingMode", opMode)
}

func (fe *JsFrontendGameRules) Destroy() {
	if fe.Object.Get("destroy") == js.Undefined {
		return
	}
	fe.Object.Call("destroy")
}
