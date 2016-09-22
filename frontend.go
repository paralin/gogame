package gogame

type Frontend interface {
	Init()
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
