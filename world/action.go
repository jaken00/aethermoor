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

func (e *Entity) MoveEntity() {
	//Call both get lowest type and get nearest cell resource
}

func getNearestCellResource(current_position Vec2, worldMap World, resourceName ResourceType) Vec2 {
	directions := [8]Vec2{{1, 0}, {0, 1}, {-1, 0}, {0, -1}, {1, 1}, {1, -1}, {-1, 1}, {-1, -1}}

	for _, dir := range directions {
		checkPos := Vec2{
			XPos: current_position.XPos + dir.XPos,
			YPos: current_position.YPos + dir.YPos,
		}

		// bounds check
		if checkPos.XPos < 0 || checkPos.XPos >= worldMap.X_len ||
			checkPos.YPos < 0 || checkPos.YPos >= worldMap.Y_len {
			continue
		}

		entityNames, exists := worldMap.CellEntities[checkPos]
		if !exists {
			continue
		}

		for _, name := range entityNames {
			entity := worldMap.Entities[name]
			for _, prod := range entity.Produces {
				if prod.Type == resourceName {
					return checkPos
				}
			}
		}
	}

	return Vec2{-1, -1} // not found
}

//need make generalized function that takes in an entity current position (so maybe just a Vec2) and get the current location of the need of that cell. so goes through the 8 cells touching it and goes to
//find the need thjta the currentNeedType returns. -> consume until capacity move on
//first we do a check if current cell has grass or food or wahtevber -> then we move on!
//we need to also implement health attack defesne stats in a combat file as well
