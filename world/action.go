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

func (e *Entity) MoveEntity(worldMap *World) {
	lowestNeedType := getLowestNeedtype(e)
	prev_position := *e.Position
	positionToMove := getNearestCellResource(*e.Position, worldMap, ResourceType(lowestNeedType))

	oldCell := &worldMap.Grid[prev_position.XPos][prev_position.YPos]

	for i, entity := range oldCell.CellEntities {
		if entity.Name == e.Name {
			oldCell.CellEntities = append(oldCell.CellEntities[:i], oldCell.CellEntities[i+1:]...)
			break
		}
	}
	oldList := worldMap.CellEntities[prev_position]

	for i, name := range oldList {
		if name == e.Name {
			worldMap.CellEntities[prev_position] = append(oldList[:i], oldList[i+1:]...)
			break
		}
	}

	e.Position = &positionToMove

	newCell := &worldMap.Grid[positionToMove.XPos][positionToMove.YPos]
	newCell.CellEntities = append(newCell.CellEntities, e)

	worldMap.CellEntities[positionToMove] = append(worldMap.CellEntities[positionToMove], e.Name)

}

func getNearestCellResource(current_position Vec2, worldMap *World, resourceName ResourceType) Vec2 {
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

func (e *Entity) CheckCurrentCell(worldMap *World, resourceNeeded ResourceType) bool {

	current_pos := e.Position
	current_cell := worldMap.Grid[current_pos.XPos][current_pos.YPos]

	for _, potential_entities := range current_cell.CellEntities {
		for _, produces := range potential_entities.Produces {
			if produces.Type == resourceNeeded {
				produces.Current--
				return true
			}
		}
	}

	return false

}
