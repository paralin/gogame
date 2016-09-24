package gogame

type Frontend interface {
	Init() FrontendGameRules
	AddEntity(entity *Entity) FrontendEntity
	Destroy()
}

type FrontendEntity interface {
	Init()
	AddComponent(componentId uint32) FrontendComponent
	InitLate()
	Destroy()
}

type FrontendComponent interface {
	Call(name string, args ...interface{}) interface{}
	// Return true if Update() calls are needed
	Init() bool
	// Called after all of the InitLate() is called on FrontendEntity
	InitLate()
	// Update tick
	Update()
	Destroy()
}

type FrontendGameRules interface {
	Init()
	SetGameOperatingMode(opMode GameOperatingMode)
	Destroy()
}
