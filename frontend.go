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
	Init()
	Destroy()
}

type FrontendGameRules interface {
	Init()
	SetGameOperatingMode(opMode GameOperatingMode)
	Destroy()
}
