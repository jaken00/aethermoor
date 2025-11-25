package world

func (e *Entity) tickNeed(needType string) {
	need, exists := e.Needs[needType]
	if !exists {
		return
	}

	need.Current -= 1

}

func (e *Entity) needComparison(needType string) bool {
	need, exists := e.Needs[needType]
	if !exists {
		return false
	}
	return need.Current < need.Threshold
}

/*
	func (e *Entity) locateNeed(needType string, cell *Cell) {
		need, exists := e.Needs[needType]
		if !exists {
			return
		}

		//need to create a map[string]cell.CellType mapping. I need mappings and data connectionms

		switch cell.CellType {
		case "PLAINS":
			//go to!
			//Will fix this later
		}

}
*/
func (e *Entity) ActOrchestrator() {
	foodBool := e.needComparison("food")
	shelterBool := e.needComparison("shelter")

	if foodBool && shelterBool {
		// go get food
	} else if foodBool {
		//go get food
	} else {
		//go to shelter -> if no shelter close just walk to a random 8 square
		//to do this we need to create a for loop and break when we have the first choice
		//need a tick food as well
	}

}
