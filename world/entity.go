package world

func Die(e *Entity, worldMap *World) {
	println("ENTITY DEAD")
	entity_pos := *e.Position

	oldCell := &worldMap.Grid[entity_pos.XPos][entity_pos.YPos]

	for i, entity := range oldCell.CellEntities {
		if entity.Name == e.Name {
			oldCell.CellEntities = append(oldCell.CellEntities[:i], oldCell.CellEntities[i+1:]...)
			break
		}
	}
	oldList := worldMap.CellEntities[entity_pos]

	for i, name := range oldList {
		if name == e.Name {
			worldMap.CellEntities[entity_pos] = append(oldList[:i], oldList[i+1:]...)
			break
		}
	}
}
