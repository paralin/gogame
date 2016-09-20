package gogame

import "fmt"

/*
 * Passed at startup, registers all the components.
 */
type ComponentTable map[uint32]ComponentFactory

func NewComponentTable() ComponentTable {
	return make(map[uint32]ComponentFactory)
}

/*
 * Validate a component table.
 */
func (t *ComponentTable) Validate() error {
	for id, factory := range *t {
		if factory.Meta().Id != id {
			return fmt.Errorf("Component table ID %d does not match metadata ID %d", id, factory.Meta().Id)
		}

		testInstance := factory.New()
		if testInstance == nil {
			return fmt.Errorf("Component table ID %d factory returned nil.", id)
		}
		if testInstance.Meta().Id != id {
			return fmt.Errorf("Component table ID %d does not match ID of built instance %d", id, testInstance.Meta().Id)
		}
	}

	return nil
}
