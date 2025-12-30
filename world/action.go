package world

import (
	"fmt"
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

	//Using this to force entity home so they dont get in a loop of eating -> shelter -> eating
	if e.EntitySettings.Activity != ShelterActivity {
		switch currentLowestNeedType {
		case NeedFood:
			e.EntitySettings.Activity = HuntingActivity
		case NeedShelter:
			e.EntitySettings.Activity = ShelterActivity
		}
	}

	return currentLowestNeedType
}

// This gets the resource needed for the lowest need | Returns: ResourceType
func getLowestNeedResource(e *Entity) ResourceType {
	lowest := 50.0
	var current float64
	var lowestNeedResource ResourceType

	for _, need := range e.Needs {
		current = need.Current
		if current < lowest {
			lowestNeedResource = need.Resource
			lowest = current
		}
	}

	return lowestNeedResource
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

	var positionToMove Vec2
	prev_position := *e.Position
	if e.EntitySettings.Activity == ShelterActivity {
		dx := e.Position.XPos - e.Home.XPos
		dy := e.Position.YPos - e.Home.YPos

		moveX := 0
		if dx > 1 {
			moveX = -1
		} else if dx < -1 {
			moveX = 1
		}

		moveY := 0
		if dy > 1 {
			moveY = -1
		} else if dy < -1 {
			moveY = 1
		}

		positionToMove = Vec2{XPos: e.Position.XPos + moveX, YPos: e.Position.YPos + moveY}
	} else {
		getLowestNeedtype(e) // Update activity based on lowest need

		positionToMove = getNearestCellResource(*e.Position, worldMap, getLowestNeedResource(e))

		// Move randomly if no resource found
		if positionToMove.XPos < -1 || positionToMove.YPos < -1 { //THese are at both -1 as resource not found returns a -2, -2 Vec2
			positionToMove = getRandomAdjacentPosition(prev_position, worldMap)
		}
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

	return Vec2{-2, -2} // not found
}

func (e *Entity) CheckCurrentCell(worldMap *World, resourceNeeded ResourceType) bool {
	current_pos := e.Position
	current_cell := worldMap.Grid[current_pos.XPos][current_pos.YPos]

	if e.EntitySettings.Activity == ShelterActivity {
		if e.Position.XPos == e.Home.XPos && e.Position.YPos == e.Home.YPos {
			shelterNeed := e.Needs[NeedShelter]
			fmt.Printf("** GOT SHELTER ** ") // **TODO** Shelter Check Goes here
			shelterNeed.Current += 3         // add in shelter reproduction

		}
	}

	for _, potential_entities := range current_cell.CellEntities {
		if potential_entities.Name == e.Name {
			continue
		}

		for i := range potential_entities.Produces {
			if potential_entities.Produces[i].Type == resourceNeeded && potential_entities.Produces[i].Current > 0 {
				potential_entities.Produces[i].Current--

				if potential_entities.Produces[i].Current <= 0 {
					potential_entities.Alive = false
				}

				for _, need := range e.Needs {
					if need.Resource == resourceNeeded {
						// TODO: Need do bounds checking (make sure it stays below threshold)

						// here we need to also need to create a corpse on death
						if e.Type == WolfEntity {
							killedEnemy := attack(e, potential_entities)
							if killedEnemy {
								need.Current += need.ConsumeRate
								fmt.Printf("** %s ATE %s ** (Current: %.1f -> %.1f)\n",
									e.Name, resourceNeeded, need.Current, need.Current+need.ConsumeRate)
							}
						} else {
							need.Current += need.ConsumeRate
						}

						break
					}
				}
				return true
			}
		}
	}
	return false
}
