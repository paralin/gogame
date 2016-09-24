package gogame

type EntityFactory interface {
	Spawn(id uint32) *Entity
}
