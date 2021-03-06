package components

import (
	"github.com/paralin/gogame"
	"github.com/paralin/gogame/components/transform"
)

func RegisterComponents(table gogame.ComponentTable) {
	// Register transform component
	table[transform.TransformComponentMeta.Id] = transform.TransformComponentFactory
}

// Build a pre-initialized component table
func NewComponentTable() gogame.ComponentTable {
	table := gogame.NewComponentTable()
	RegisterComponents(table)
	return table
}
