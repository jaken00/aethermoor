package world

import "fmt"

func tickNeed(entity *Entity) {

	for _, need := range entity.Needs {

		need.Current -= 1
		if need.Current >= 0 {
			fmt.Println("Need kill entity")
		}
	}
}

//This gets the current lowest need type when deciding action | Returns: NEEDTYPE
func getLowestNeedtype(e *Entity) NeedType {
	lowest := 50.0
	var current float64

	var currentLowestNeedType NeedType

	for _, need := range e.Needs {
		current = need.Current
		if current < lowest {
			currentLowestNeedType = need.Kind
			lowest = current

		}
	}

	return currentLowestNeedType
}

//need make generalized function that takes in an entity current position (so maybe just a Vec2) and get the current location of the need of that cell. so goes through the 8 cells touching it and goes to
//find the need thjta the currentNeedType returns. -> consume until capacity move on
//first we do a check if current cell has grass or food or wahtevber -> then we move on!
//we need to also implement health attack defesne stats in a combat file as well
