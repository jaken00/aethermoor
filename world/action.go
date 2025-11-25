package world

func (e *Entity) getHunger() float64 {
	if need, exists := e.Needs["food"]; exists {
		return need.Current
	}
	return 0
}

func (e *Entity) getShelter() float64 {
	if need, exists := e.Needs["shelter"]; exists {
		return need.Current
	}
	return 0
}

func (e *Entity) needComparison(needType string) bool {
	need, exists := e.Needs[needType]
	if !exists {
		return false
	}
	return need.Current < need.Threshold
}

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
