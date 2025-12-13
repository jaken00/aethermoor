package world

import (
	"math/rand"
)

func tickNeed(entity *Entity) bool {
	var die bool
	for _, need := range entity.Needs {

		need.Current -= 1
		if need.Current <= 0 {
			die = true
			return die
		}
	}
	die = false
	return die
}

// This gets the current lowest need type when deciding action | Returns: NEEDTYPE
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

func getRandomAdjacentPosition(pos Vec2, worldMap *World) Vec2 {
	directions := []struct{ dx, dy int }{
		{-1, 0}, {1, 0}, {0, -1}, {0, 1}, // cardinal
		{-1, -1}, {-1, 1}, {1, -1}, {1, 1}, // diagonal
	}

	validPositions := []Vec2{}

	for _, dir := range directions {
		newX := pos.XPos + dir.dx
		newY := pos.YPos + dir.dy

		if newX >= 0 && newX < len(worldMap.Grid) &&
			newY >= 0 && newY < len(worldMap.Grid[0]) {
			validPositions = append(validPositions, Vec2{XPos: newX, YPos: newY})
		}
	}

	if len(validPositions) == 0 {
		return pos
	}

	return validPositions[rand.Intn(len(validPositions))]
}

func (e *Entity) MoveEntity(worldMap *World) {

	lowestNeedType := getLowestNeedtype(e)
	prev_position := *e.Position
	positionToMove := getNearestCellResource(*e.Position, worldMap, ResourceType(lowestNeedType))

	// Move randomly if no resource found
	if positionToMove.XPos < 0 || positionToMove.YPos < 0 {
		positionToMove = getRandomAdjacentPosition(prev_position, worldMap)
	}

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
			if produces.Type == resourceNeeded && produces.Current > 0 {
				produces.Current--
				return true
			}
		}
	}

	return false

}
